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

type sKnowledgeOrder struct {
}

func (s sKnowledgeOrder) QueryOrder(ctx context.Context, req *desk.QueryKnowledgeOrderInfoReq) (res *desk.QueryKnowledgeOrderInfoRes, err error) {
	res = (*desk.QueryKnowledgeOrderInfoRes)(nil)
	err = dao.KnowledgeOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&res)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("order not found")
	}

	return
}

func (s sKnowledgeOrder) QueryPaymentInfo(ctx context.Context, req *desk.QueryKnowledgeOrderPaymentReq) (res *desk.QueryKnowledgeOrderPaymentRes, err error) {
	order := (*entity.KnowledgeOrder)(nil)
	err = dao.KnowledgeOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&order)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	//1645242806
	var settings *system_entity.PaySettings
	err = sysdao.PaySettings.Ctx(ctx).Scan(&settings)
	if err != nil {
		xlog.Error(err)
		return
	}
	bm.Set("appid", settings.WeixinAppid).
		Set("mchid", settings.WeixinMchid).
		Set("description", "知识库购买").
		Set("out_trade_no", order.OrderNo).
		Set("time_expire", expire).
		Set("notify_url", g.Cfg().MustGet(ctx, "pay.notifyUrlHost").String()+"/api/v1/desk/pay/knowledgeorder/wxnotify").
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
		res = &desk.QueryKnowledgeOrderPaymentRes{CodeUrl: wxRsp.Response.CodeUrl}
		return res, nil
	}

	return nil, errors.New(wxRsp.Error)
}

func (s sKnowledgeOrder) QueryPaymentInfoForAlipay(ctx context.Context, req *desk.QueryKnowledgeOrderPaymentAlipayReq) (res *desk.QueryKnowledgeOrderPaymentAlipayRes, err error) {
	order := (*entity.KnowledgeOrder)(nil)
	err = dao.KnowledgeOrder.Ctx(ctx).Where("order_no", req.OrderNo).Limit(1).Scan(&order)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	expire := "10m"
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", "知识库购买").
		Set("out_trade_no", req.OrderNo).
		Set("product_code", "FACE_TO_FACE_PAYMENT").
		Set("total_amount", float64(order.PayAmount)/100).
		Set("timeout_express", expire)

	aliRsp, err := ali.GetClient(ctx).SetNotifyUrl(g.Cfg().MustGet(ctx, "pay.notifyUrlHost").String()+"/api/v1/desk/pay/knowledgeorder/alinotify").TradePrecreate(ctx, bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		xlog.Errorf("client.TradePay(%+v),err:%+v", bm, err)
		return
	}
	res = &desk.QueryKnowledgeOrderPaymentAlipayRes{CodeUrl: aliRsp.Response.QrCode}

	return res, nil
}

func (s sKnowledgeOrder) CreateOrder(ctx context.Context, req *desk.CreateKnowledgeOrderReq) (res *desk.CreateKnowledgeOrderRes, err error) {

	memberFee := entity.MemberFee{}
	_ = dao.MemberFee.Ctx(ctx).Scan(&memberFee)

	downloadKnowledgeAmount, _ := strconv.ParseFloat(memberFee.DownloadKnowledge, 64)
	downloadVideoAmount, _ := strconv.ParseFloat(memberFee.DownloadVideo, 64)

	knowledge := entity.MemberKnowledge{}
	err = dao.MemberKnowledge.Ctx(ctx).Where("id", req.KnowledgeId).Limit(1).Scan(&knowledge)
	if err != nil {
		return nil, err
	}

	if knowledge.Id <= 0 {
		return nil, errors.New("knowledge not found")
	}

	originalAmount := 0.0
	payAmount := 0.0
	orderType := "knowledge"
	if knowledge.Type == 0 {
		orderType = "knowledge"
		originalAmount = downloadKnowledgeAmount
		payAmount = downloadKnowledgeAmount
	} else if knowledge.Type == 1 {
		orderType = "video"
		originalAmount = downloadVideoAmount
		payAmount = downloadVideoAmount
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			res = &desk.CreateKnowledgeOrderRes{}

			orderSn := libUtils.GenerateOrderSnWithPrefix("KN")
			id, e := dao.KnowledgeOrder.Ctx(ctx).TX(tx).InsertAndGetId(do.KnowledgeOrder{
				Memo:           req.Memo,
				PayType:        req.PayType,
				UserId:         service.Context().GetUserId(ctx),
				OrderNo:        orderSn,
				CreatedAt:      gtime.Now(),
				Status:         "create",
				OriginalAmount: int(originalAmount * 100), //单位: 分
				PayAmount:      int(payAmount * 100),      //单位: 分
				Type:           orderType,
				KnowledgeId:    req.KnowledgeId,
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

func New() *sKnowledgeOrder {
	return &sKnowledgeOrder{}
}

func init() {
	// 当您尝试将具体类型分配或传递（或转换）为接口类型时，会出现此编译时错误；并且类型本身不实现接口，只是指向类型的指针
	service.RegisterKnowledgeOrder(New())
}
