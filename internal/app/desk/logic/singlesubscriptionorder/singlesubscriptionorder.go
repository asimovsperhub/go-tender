package viporderpackage

import (
	"context"
	"errors"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"strconv"
	"tender/api/v1/desk"
	"tender/internal/app/desk/dao"
	ali "tender/internal/app/desk/library/alipay"
	"tender/internal/app/desk/library/wxV3"
	"tender/internal/app/desk/model/do"
	"tender/internal/app/desk/model/entity"
	"tender/internal/app/desk/service"
	sysdao "tender/internal/app/system/dao"
	system_entity "tender/internal/app/system/model/entity"
	"tender/library/libUtils"
	"tender/library/liberr"
	"time"
)

type sSingleSubscriptionOrder struct {
}

func (s sSingleSubscriptionOrder) QuerySingleSubscriptionOrder(ctx context.Context, req *desk.QuerySingleSubscriptionOrderInfoReq) (res *desk.QuerySingleSubscriptionOrderInfoRes, err error) {
	res = (*desk.QuerySingleSubscriptionOrderInfoRes)(nil)
	err = dao.SingleSubscriptionOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&res)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("order not found")
	}

	return
}

func (s sSingleSubscriptionOrder) QueryPaymentInfo(ctx context.Context, req *desk.QuerySingleSubscriptionOrderPaymentReq) (res *desk.QuerySingleSubscriptionOrderPaymentRes, err error) {
	order := (*entity.SingleSubscriptionOrder)(nil)
	err = dao.SingleSubscriptionOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&order)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	var settings *system_entity.PaySettings
	err = sysdao.PaySettings.Ctx(ctx).Scan(&settings)
	if err != nil {
		xlog.Error(err)
		return
	}
	bm.Set("appid", settings.WeixinAppid).
		Set("mchid", settings.WeixinMchid).
		Set("description", "新增会员订阅").
		Set("out_trade_no", order.OrderNo).
		Set("time_expire", expire).
		Set("notify_url", g.Cfg().MustGet(ctx, "pay.notifyUrlHost").String()+"/api/v1/desk/pay/singlesubscription/wxnotify").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", order.PayAmount).
				Set("currency", "CNY")
		})

	client := wxV3.GetClient(ctx)
	wxRsp, err := client.V3TransactionNative(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	if wxRsp.Code == wechat.Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		res = &desk.QuerySingleSubscriptionOrderPaymentRes{CodeUrl: wxRsp.Response.CodeUrl}
		return res, nil
	}

	return nil, errors.New(wxRsp.Error)
}

func (s sSingleSubscriptionOrder) QueryPaymentInfoForAlipay(ctx context.Context, req *desk.QuerySingleSubscriptionOrderPaymentAlipayReq) (res *desk.QuerySingleSubscriptionOrderPaymentAlipayRes, err error) {
	order := (*entity.SingleSubscriptionOrder)(nil)
	err = dao.SingleSubscriptionOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&order)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	expire := "10m"
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", "新增会员订阅").
		Set("out_trade_no", req.OrderNo).
		Set("product_code", "FACE_TO_FACE_PAYMENT").
		Set("total_amount", float64(order.PayAmount)/100).
		Set("timeout_express", expire)

	aliRsp, err := ali.GetClient(ctx).SetNotifyUrl(g.Cfg().MustGet(ctx, "pay.notifyUrlHost").String()+"/api/v1/desk/pay/singlesubscription/alinotify").TradePrecreate(ctx, bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		xlog.Errorf("client.TradePay(%+v),err:%+v", bm, err)
		return
	}
	res = &desk.QuerySingleSubscriptionOrderPaymentAlipayRes{CodeUrl: aliRsp.Response.QrCode}

	return res, nil
}

func (s sSingleSubscriptionOrder) CreateSingleSubscriptionOrder(ctx context.Context, req *desk.CreateSingleSubscriptionOrderReq) (res *desk.CreateSingleSubscriptionOrderRes, err error) {

	memberSubscription := entity.MemberSubscription{}
	_ = dao.MemberSubscription.Ctx(ctx).Scan(&memberSubscription)

	originalAmount := 0.0
	payAmount := 0.0
	if req.Type == "monthly" {
		originalAmount, _ = strconv.ParseFloat(memberSubscription.MonthlycardSubscriptionPrice, 64)
		payAmount = originalAmount
	} else if req.Type == "quarter" {
		originalAmount, _ = strconv.ParseFloat(memberSubscription.QuartercardSubscriptionPrice, 64)
		payAmount = originalAmount
	} else if req.Type == "annual" {
		originalAmount, _ = strconv.ParseFloat(memberSubscription.AnnualcardSubscriptionPrice, 64)
		payAmount = originalAmount
	} else {
		return nil, errors.New("type is wrong")
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			res = &desk.CreateSingleSubscriptionOrderRes{}

			orderSn := libUtils.GenerateOrderSnWithPrefix("SS")
			id, e := dao.SingleSubscriptionOrder.Ctx(ctx).TX(tx).InsertAndGetId(do.SingleSubscriptionOrder{
				Memo:           req.Memo,
				PayType:        req.PayType,
				UserId:         service.Context().GetUserId(ctx),
				OrderNo:        orderSn,
				CreatedAt:      gtime.Now(),
				Status:         "create",
				OriginalAmount: int(originalAmount * 100), //单位: 分
				PayAmount:      int(payAmount * 100),      //单位: 分
				Type:           req.Type,
			})
			liberr.ErrIsNil(ctx, e, "创建订单失败")

			res.OrderId = id
			res.OrderNo = orderSn
			return
		})
		return err
	})
	return
}

func New() *sSingleSubscriptionOrder {
	return &sSingleSubscriptionOrder{}
}

func init() {
	// 当您尝试将具体类型分配或传递（或转换）为接口类型时，会出现此编译时错误；并且类型本身不实现接口，只是指向类型的指针
	service.RegisterSingleSubscriptionOrder(New())
}
