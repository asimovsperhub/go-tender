package controller

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"tender/internal/app/desk/dao"
	"tender/internal/app/desk/model/do"
	"tender/internal/app/desk/model/entity"
	sysdao "tender/internal/app/system/dao"
	sysdo "tender/internal/app/system/model/do"
	sysentity "tender/internal/app/system/model/entity"
	"tender/library/liberr"
)

func VipUpdateUser(r *ghttp.Request, order *entity.VipOrder) (err error) {
	integral := (*sysentity.MemberIntegral)(nil)
	subscription := (*entity.MemberSubscription)(nil)
	user := (*sysentity.MemberUser)(nil)
	data := sysdo.MemberUser{}
	// 积分设置
	err = g.Try(r.Context(), func(ctx context.Context) {
		err = sysdao.MemberIntegral.Ctx(ctx).Scan(&integral)
		liberr.ErrIsNil(ctx, err, "获取积分信息失败")
		if err != nil {
			return
		}
	})
	// 订阅设置
	err = g.Try(r.Context(), func(ctx context.Context) {
		err = dao.MemberSubscription.Ctx(ctx).Scan(&subscription)
		liberr.ErrIsNil(ctx, err, "获取订阅信息失败")
		if err != nil {
			return
		}
	})
	err = g.Try(r.Context(), func(ctx context.Context) {
		err = sysdao.MemberUser.Ctx(ctx).Where("id", order.UserId).Scan(&user)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
		if err != nil {
			return
		}
	})
	//uercardsub := 0
	//// 升级会员
	//if user.MemberLevel == 1 {
	//	uercardsub = subscription.MonthlycardSubscription
	//} else if user.MemberLevel == 2 {
	//	uercardsub = subscription.QuartercardSubscription
	//} else if user.MemberLevel == 3 {
	//	uercardsub = subscription.AnnualcardSubscription
	//}
	//monthly quarter annual
	if order.Type == "monthly" {
		data.MemberLevel = 1
		data.Integral = user.Integral + integral.MonthlycardIntegral
		// data.Subscribe = user.Subscribe + subscription.MonthlycardSubscription - uercardsub
		data.Subscribe = subscription.MonthlycardSubscription
		if user.MaturityAt != nil {
			data.MaturityAt = user.MaturityAt.AddDate(0, 1, 0)
		} else {
			data.MaturityAt = gtime.Now().AddDate(0, 1, 0)
		}
	} else if order.Type == "quarter" {
		data.MemberLevel = 2
		data.Integral = user.Integral + integral.QuartercardIntegral
		// data.Subscribe = user.Subscribe + subscription.QuartercardSubscription - uercardsub
		data.Subscribe = subscription.QuartercardSubscription
		if user.MaturityAt != nil {
			data.MaturityAt = user.MaturityAt.AddDate(0, 3, 0)
		} else {
			data.MaturityAt = gtime.Now().AddDate(0, 3, 0)
		}
	} else if order.Type == "annual" {
		data.MemberLevel = 3
		data.Integral = user.Integral + integral.AnnualcardIntegral
		// data.Subscribe = user.Subscribe + subscription.AnnualcardSubscription - uercardsub
		data.Subscribe = subscription.AnnualcardSubscription
		if user.MaturityAt != nil {
			data.MaturityAt = user.MaturityAt.AddDate(1, 0, 0)
		} else {
			data.MaturityAt = gtime.Now().AddDate(1, 0, 0)
		}
	}
	err = g.DB().Transaction(r.Context(), func(ctx context.Context, tx gdb.TX) error {
		// 更新用户信息
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = sysdao.MemberUser.Ctx(r.Context()).Where("id", user.Id).Update(data)
			liberr.ErrIsNil(ctx, err, "更新用户失败")
			if err != nil {
				return
			}
		})

		return err
	})
	return
}

// 单次购买订阅数
func SubUpdateUser(r *ghttp.Request, order *entity.SingleSubscriptionOrder) (err error) {
	user := (*sysentity.MemberUser)(nil)
	data := sysdo.MemberUser{}
	err = g.Try(r.Context(), func(ctx context.Context) {
		err = sysdao.MemberUser.Ctx(ctx).Where("id", order.UserId).Scan(&user)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
		if err != nil {
			return
		}
	})
	data.BuySubscribe = user.BuySubscribe + 1
	err = g.DB().Transaction(r.Context(), func(ctx context.Context, tx gdb.TX) error {
		// 更新用户信息
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = sysdao.MemberUser.Ctx(r.Context()).Where("id", user.Id).Update(data)
			liberr.ErrIsNil(ctx, err, "更新用户失败")
			if err != nil {
				return
			}
		})

		return err
	})
	return
}

func IntegralUpdateUser(r *ghttp.Request, order *entity.RechargeOrder) (err error) {
	user := (*sysentity.MemberUser)(nil)
	data := sysdo.MemberUser{}
	err = g.Try(r.Context(), func(ctx context.Context) {
		err = sysdao.MemberUser.Ctx(ctx).Where("id", order.UserId).Scan(&user)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
		if err != nil {
			return
		}
	})
	integral := (*sysentity.MemberIntegral)(nil)
	// 积分设置
	err = g.Try(r.Context(), func(ctx context.Context) {
		err = sysdao.MemberIntegral.Ctx(ctx).Scan(&integral)
		liberr.ErrIsNil(ctx, err, "获取积分信息失败")
		if err != nil {
			return
		}
	})
	ratio := 0
	// 升级会员
	if user.MemberLevel == 1 {
		ratio = integral.MonthlycardRatio
	} else if user.MemberLevel == 2 {
		ratio = integral.QuartercardRatio
	} else if user.MemberLevel == 3 {
		ratio = integral.AnnualcardRatio
	} else {
		ratio = integral.Ratio
	}
	// 分
	in := int(float64(order.PayAmount/100) * float64(ratio))
	data.Integral = user.Integral + in
	err = g.DB().Transaction(r.Context(), func(ctx context.Context, tx gdb.TX) error {
		// 更新用户信息
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = sysdao.MemberUser.Ctx(r.Context()).Where("id", user.Id).Update(data)
			liberr.ErrIsNil(ctx, err, "更新用户失败")
			if err != nil {
				return
			}
		})

		return err
	})
	return
}

// 单次购买 更新知识记录
func KnowledgeUpdateHis(r *ghttp.Request, order *entity.KnowledgeOrder) (err error) {
	data := do.HisKnowledge{}
	data.UserId = order.UserId
	data.KnowledgeId = order.KnowledgeId
	err = g.DB().Transaction(r.Context(), func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.HisKnowledge.Ctx(r.Context()).Insert(data)
			liberr.ErrIsNil(ctx, err, "插入数据his知识失败")
			if err != nil {
				return
			}
		})
		return err
	})
	return
}
