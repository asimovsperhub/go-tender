package controller

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"tender/internal/app/desk/dao"
	"tender/internal/app/desk/library/wxV3"
	"tender/internal/app/desk/model"
	"tender/internal/app/desk/model/do"
	"tender/internal/app/desk/model/entity"
	sysdao "tender/internal/app/system/dao"
	system_entity "tender/internal/app/system/model/entity"
	"tender/library/liberr"
	"time"
)

var (
	SingleSubscriptionOrderNotify = cSingleSubscriptionNotifyController{}
)

type cSingleSubscriptionNotifyController struct {
	BaseController
}

func (*cSingleSubscriptionNotifyController) WxNotify(r *ghttp.Request) {
	notifyReq, err := wechat.V3ParseNotify(r.Request)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 获取微信平台证书
	certMap := wxV3.GetClient(r.Context()).WxPublicKeyMap()
	// 验证异步通知的签名
	err = notifyReq.VerifySignByPKMap(certMap)
	if err != nil {
		xlog.Error(err)
		return
	}
	var settings *system_entity.PaySettings
	err = sysdao.PaySettings.Ctx(r.Context()).Scan(&settings)
	if err != nil {
		xlog.Error(err)
		return
	}
	result, err := notifyReq.DecryptCipherText(settings.WeixinApikey)
	fmt.Println("notify result: ", result)

	orderNo := result.OutTradeNo

	order := (*entity.SingleSubscriptionOrder)(nil)
	err = dao.SingleSubscriptionOrder.Ctx(r.Context()).Where("order_no", orderNo).Limit(1).Scan(&order)

	if order == nil {
		xlog.Error("order not found")
		return
	}

	if order.Status != "create" {
		r.Response.WriteJson(&wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
		return
	}

	err = g.DB().Transaction(r.Context(), func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SingleSubscriptionOrder.Ctx(r.Context()).Where("order_no", orderNo).Update(do.SingleSubscriptionOrder{
				PayTime:   time.Now().Unix(),
				UpdatedAt: gtime.Now(),
				Status:    "paid",
				PayType:   "wx",
			})

			liberr.ErrIsNil(ctx, err, "更新订单失败")
			if err != nil {
				return
			}

			user := &model.LoginUserRes{}
			_ = sysdao.MemberUser.Ctx(ctx).Where("id", order.UserId).Scan(user)
			_, err = dao.PurchaseLog.Ctx(r.Context()).Insert(do.PurchaseLog{
				UserId:    user.Id,
				NickName:  user.UserNickname,
				UpdatedAt: gtime.Now(),
				UserName:  user.UserName,
				Mobile:    user.Mobile,
				Memo:      getSingleSubscriptionOrderDesc(order.Type),
				Amount:    order.PayAmount,
				CreatedAt: gtime.Now(),
				PayType:   "wx",
				OrderNo:   order.OrderNo,
			})

			liberr.ErrIsNil(ctx, err, "增加明细失败")
			if err != nil {
				return
			}
		})

		return err
	})

	if err != nil {
		xlog.Error("order update failed")
		return
	}
	err = SubUpdateUser(r, order)
	if err != nil {
		xlog.Error("user single subscription  update failed")
		return
	}
	r.Response.WriteJson(&wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
}

func (*cSingleSubscriptionNotifyController) AlipayNotify(r *ghttp.Request) {

	notifyReq, err := alipay.ParseNotifyToBodyMap(r.Request)
	if err != nil {
		xlog.Error(err)
		return
	}
	var settings *system_entity.PaySettings
	err = sysdao.PaySettings.Ctx(r.Context()).Scan(&settings)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 支付宝异步通知验签（公钥证书模式）
	ok, err := alipay.VerifySignWithCert([]byte(settings.AlipayPublicCert), notifyReq)
	if !ok {
		r.Response.Write("failed")
		return
	}
	// 如果需要，可将 BodyMap 内数据，Unmarshal 到指定结构体指针 ptr
	aliResp := AlipayResult{}
	err = notifyReq.Unmarshal(&aliResp)

	// 此写法是 echo 框架返回支付宝的写法
	fmt.Println("notify result: ", aliResp)

	orderNo := aliResp.OutTradeNo

	order := (*entity.SingleSubscriptionOrder)(nil)
	err = dao.SingleSubscriptionOrder.Ctx(r.Context()).Where("order_no", orderNo).Limit(1).Scan(&order)

	if order == nil {
		xlog.Error("order not found")
		return
	}

	if order.Status != "create" {
		r.Response.WriteJson(&wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
		return
	}

	err = g.DB().Transaction(r.Context(), func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SingleSubscriptionOrder.Ctx(r.Context()).Where("order_no", orderNo).Update(do.SingleSubscriptionOrder{
				PayTime:   time.Now().Unix(),
				UpdatedAt: gtime.Now(),
				Status:    "paid",
				PayType:   "alipay",
			})

			liberr.ErrIsNil(ctx, err, "更新订单失败")
			if err != nil {
				return
			}

			user := &model.LoginUserRes{}
			_ = sysdao.MemberUser.Ctx(ctx).Where("id", order.UserId).Scan(user)
			_, err = dao.PurchaseLog.Ctx(r.Context()).Insert(do.PurchaseLog{
				UserId:    user.Id,
				NickName:  user.UserNickname,
				UpdatedAt: gtime.Now(),
				UserName:  user.UserName,
				Mobile:    user.Mobile,
				Memo:      getSingleSubscriptionOrderDesc(order.Type),
				Amount:    order.PayAmount,
				CreatedAt: gtime.Now(),
				PayType:   "alipay",
				OrderNo:   order.OrderNo,
			})

			if err != nil {
				return
			}

			//TODO: 增加订阅次数
		})

		return err
	})

	if err != nil {
		xlog.Error("order update failed")
		return
	}
	err = SubUpdateUser(r, order)
	if err != nil {
		xlog.Error("user single subscription  update failed")
		return
	}
	r.Response.Write("success")
}

func getSingleSubscriptionOrderDesc(orderType string) string {
	if orderType == "monthly" {
		return "新增月卡会员订阅"
	} else if orderType == "quarter" {
		return "新增季卡会员订阅"
	} else if orderType == "annual" {
		return "新增年卡会员订阅"
	}

	return ""
}
