package viporderpackage

import (
	"context"
	"errors"
	"fmt"
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
	system_dao "tender/internal/app/system/dao"
	system_entity "tender/internal/app/system/model/entity"
	"tender/library/libUtils"
	"tender/library/liberr"
	"time"
)

type sVipOrder struct {
}

func (s sVipOrder) QueryVipOrder(ctx context.Context, req *desk.QueryVipOrderInfoReq) (res *desk.QueryVipOrderInfoRes, err error) {
	res = (*desk.QueryVipOrderInfoRes)(nil)
	err = dao.VipOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&res)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("order not found")
	}

	return
}

func (s sVipOrder) QueryPaymentInfo(ctx context.Context, req *desk.QueryVipOrderPaymentReq) (res *desk.QueryVipOrderPaymentRes, err error) {
	order := (*entity.VipOrder)(nil)
	err = dao.VipOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&order)

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
	err = system_dao.PaySettings.Ctx(ctx).Scan(&settings)
	if err != nil {
		xlog.Error(err)
		return
	}
	bm.Set("appid", settings.WeixinAppid).
		Set("mchid", settings.WeixinMchid).
		Set("description", "会员订阅").
		Set("out_trade_no", order.OrderNo).
		Set("time_expire", expire).
		Set("notify_url", g.Cfg().MustGet(ctx, "pay.notifyUrlHost").String()+"/api/v1/desk/pay/wxnotify").
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
		res = &desk.QueryVipOrderPaymentRes{CodeUrl: wxRsp.Response.CodeUrl}
		return res, nil
	}

	return nil, errors.New(wxRsp.Error)
}

func (s sVipOrder) QueryPaymentInfoForAlipay(ctx context.Context, req *desk.QueryVipOrderPaymentAlipayReq) (res *desk.QueryVipOrderPaymentAlipayRes, err error) {
	order := (*entity.VipOrder)(nil)
	err = dao.VipOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&order)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	expire := "10m"
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", "会员订阅").
		Set("out_trade_no", req.OrderNo).
		Set("product_code", "FACE_TO_FACE_PAYMENT").
		Set("total_amount", float64(order.PayAmount)/100).
		Set("timeout_express", expire)

	aliRsp, err := ali.GetClient(ctx).SetNotifyUrl(g.Cfg().MustGet(ctx, "pay.notifyUrlHost").String()+"/api/v1/desk/pay/alinotify").TradePrecreate(ctx, bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		xlog.Errorf("client.TradePay(%+v),err:%+v", bm, err)
		return
	}
	res = &desk.QueryVipOrderPaymentAlipayRes{CodeUrl: aliRsp.Response.QrCode}

	return res, nil
}

func (s sVipOrder) CreateVipOrder(ctx context.Context, req *desk.CreateVipOrderReq) (res *desk.CreateVipOrderRes, err error) {

	memberFee := entity.MemberFee{}
	_ = dao.MemberFee.Ctx(ctx).Scan(&memberFee)
	user_id := service.Context().GetUserId(ctx)
	finduser := (*system_entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		m := system_dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%d'",
			system_dao.MemberUser.Columns().Id,
			user_id))
		err = m.Limit(1).Scan(&finduser)
		liberr.ErrIsNil(ctx, err, "用户信息不存在")
	})
	if err != nil {
		return
	}
	user_level := finduser.MemberLevel
	originalAmount := 0.0
	payAmount := 0.0
	if req.Type == "monthly" {
		if user_level > 1 {
			return nil, errors.New("不能购买低等级的会员")
		}
		originalAmount, _ = strconv.ParseFloat(memberFee.MonthlycardOriginal, 64)
		payAmount, _ = strconv.ParseFloat(memberFee.MonthlycardCurrent, 64)
	} else if req.Type == "quarter" {
		if user_level > 2 {
			return nil, errors.New("不能购买低等级的会员")
		}
		originalAmount, _ = strconv.ParseFloat(memberFee.QuartercardOriginal, 64)
		payAmount, _ = strconv.ParseFloat(memberFee.QuartercardCurrent, 64)
	} else if req.Type == "annual" {
		originalAmount, _ = strconv.ParseFloat(memberFee.AnnualcardOriginal, 64)
		payAmount, _ = strconv.ParseFloat(memberFee.AnnualcardCurrent, 64)
	} else {
		return nil, errors.New("type is wrong")
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			res = &desk.CreateVipOrderRes{}

			orderSn := libUtils.GenerateOrderSnWithPrefix("DY")
			id, e := dao.VipOrder.Ctx(ctx).TX(tx).InsertAndGetId(do.VipOrder{
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

func New() *sVipOrder {
	return &sVipOrder{}
}

func init() {
	// 当您尝试将具体类型分配或传递（或转换）为接口类型时，会出现此编译时错误；并且类型本身不实现接口，只是指向类型的指针
	service.RegisterVipOrder(New())
}
