package sys_membermanager

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"log"
	"strconv"
	"tender/api/v1/system"
	"tender/internal/app/system/consts"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
	"tender/library/liberr"
	"time"
)

func init() {
	service.RegisterMemberManger(New())
}

func New() *sSysMemberManger {
	return &sSysMemberManger{}
}

type sSysMemberManger struct {
}

func (s *sSysMemberManger) List(ctx context.Context, req *system.MemberUserSearchReq) (res *system.MemberUserSearchRes, err error) {
	res = new(system.MemberUserSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.MemberUser.Ctx(ctx)
	order := "id"
	if req.Name != "" {
		m = m.Where("user_nickname like ? ", "%"+req.Name+"%")
	}
	if req.Contact != "" {
		m = m.Where("mobile like ?", "%"+req.Contact+"%")
	}
	if req.Level != nil {
		m = m.Where("member_level = ?", *req.Level)
	}
	if req.Status != nil {
		m = m.Where("user_status = ?", *req.Status)
	}
	if len(req.DateRange) > 0 {
		m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0]+" 00:00:00", req.DateRange[1]+" 23:59:59")
	} else {
		if req.Start != "" {
			m = m.Where("created_at >=?", req.Start+" 00:00:00")
		}
		if req.End != "" {
			m = m.Where("created_at <=?", req.End+" 23:59:59")
		}
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取会员列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.MemberUserList)
		log.Println("获取会员列表-------->", res.MemberUserList)
		liberr.ErrIsNil(ctx, err, "获取会员列表数据失败")
	})
	return
}

//MemberUserEdit
func (s *sSysMemberManger) MemberUserEdit(ctx context.Context, req *system.MemberUserEditReq) (res *system.MemberUserEditRes, err error) {
	data := do.MemberUser{}
	if req.Integral != nil {
		data.Integral = *req.Integral
	}
	if req.Level != nil {
		data.MemberLevel = *req.Level
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.MemberUser.Ctx(ctx).Where("id = ?", req.Id).Update(data)
		liberr.ErrIsNil(ctx, err, "更新会员积分数据失败")
	})
	return
}

// Disable
func (s *sSysMemberManger) Disable(ctx context.Context, req *system.DisableMemberUserReq) (res *system.DisableMemberUserRes, err error) {
	dd, _ := time.ParseDuration(strconv.Itoa(req.Day*24) + "h")
	userStatus, releaseAt := 0, gtime.Now().Add(dd)
	if req.Day == 0 {
		userStatus = 1
		releaseAt = gtime.Now()
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.MemberUser.Ctx(ctx).WherePri(req.MemberUserId).Update(do.MemberUser{
			// 禁用的时间
			DisabledAt: gtime.Now(),
			// 释放时间
			ReleaseAt: releaseAt,
			// 用户状态
			UserStatus: userStatus,
			// 操作管理员id
			OperationId: service.Context().GetUserId(ctx),
			DisabledDay: req.Day,
		})
		liberr.ErrIsNil(ctx, err, "会员禁用失败")
	})
	return
}
func (s *sSysMemberManger) EditFee(ctx context.Context, req *system.MemberFeeReq) (res *system.MemberFeeRes, err error) {
	var res_id *entity.MemberFee
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.MemberFee.Ctx(ctx).Scan(&res_id)
		liberr.ErrIsNil(ctx, err, "获取会费信息失败")
	})
	data := do.MemberFee{}
	if req.MonthlycardOriginal != "" {
		MonthlycardOriginal, _ := strconv.ParseFloat(req.MonthlycardOriginal, 64)
		if MonthlycardOriginal <= 0 {
			err = errors.New("月卡原价 价格不能小于0")
			return
		}
		data.MonthlycardOriginal = req.MonthlycardOriginal
	}
	if req.MonthlycardCurrent != "" {
		MonthlycardCurrent, _ := strconv.ParseFloat(req.MonthlycardCurrent, 64)
		if MonthlycardCurrent <= 0 {
			err = errors.New("月卡现价 价格不能小于0")
			return
		}
		data.MonthlycardCurrent = req.MonthlycardCurrent
	}
	if req.QuartercardOriginal != "" {
		QuartercardOriginal, _ := strconv.ParseFloat(req.QuartercardOriginal, 64)
		if QuartercardOriginal <= 0 {
			err = errors.New("季卡原价 价格不能小于0")
			return
		}
		data.QuartercardOriginal = req.QuartercardOriginal
	}
	if req.QuartercardCurrent != "" {
		QuartercardCurrent, _ := strconv.ParseFloat(req.QuartercardCurrent, 64)
		if QuartercardCurrent <= 0 {
			err = errors.New("季卡现价 价格不能小于0")
			return
		}
		data.QuartercardCurrent = req.QuartercardCurrent
	}
	if req.AnnualcardOriginal != "" {
		AnnualcardOriginal, _ := strconv.ParseFloat(req.AnnualcardOriginal, 64)
		if AnnualcardOriginal <= 0 {
			err = errors.New("年卡原价 价格不能小于0")
			return
		}
		data.AnnualcardOriginal = req.AnnualcardOriginal
	}
	if req.AnnualcardCurrent != "" {
		AnnualcardCurrent, _ := strconv.ParseFloat(req.AnnualcardCurrent, 64)
		if AnnualcardCurrent <= 0 {
			err = errors.New("年卡现价 价格不能小于0")
			return
		}
		data.AnnualcardCurrent = req.AnnualcardCurrent
	}
	if req.DownloadKnowledge != "" {
		DownloadKnowledge, _ := strconv.ParseFloat(req.DownloadKnowledge, 64)
		if DownloadKnowledge <= 0 {
			err = errors.New("下载知识 价格不能小于0")
			return
		}
		data.DownloadKnowledge = req.DownloadKnowledge
	}
	if req.DownloadVideo != "" {
		DownloadVideo, _ := strconv.ParseFloat(req.DownloadVideo, 64)
		if DownloadVideo <= 0 {
			err = errors.New("下载视频知识 价格不能小于0")
			return
		}
		data.DownloadVideo = req.DownloadVideo
	}
	if res_id != nil {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.MemberFee.Ctx(ctx).WherePri(res_id.Id).Update(data)
			liberr.ErrIsNil(ctx, err, "修改会费设置失败")
		})
		return
	} else {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.MemberFee.Ctx(ctx).Insert(do.MemberFee{
				MonthlycardOriginal: req.MonthlycardOriginal,
				MonthlycardCurrent:  req.MonthlycardCurrent,
				QuartercardOriginal: req.QuartercardOriginal,
				QuartercardCurrent:  req.QuartercardCurrent,
				AnnualcardOriginal:  req.AnnualcardOriginal,
				AnnualcardCurrent:   req.AnnualcardCurrent,
				DownloadKnowledge:   req.DownloadKnowledge,
				DownloadVideo:       req.DownloadVideo,
			})
			liberr.ErrIsNil(ctx, err, "修改会费设置失败")
		})
	}
	return
}

func (s *sSysMemberManger) EditIn(ctx context.Context, req *system.MemberIntegralReq) (res *system.MemberIntegralRes, err error) {
	var res_id *entity.MemberIntegral
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.MemberIntegral.Ctx(ctx).Scan(&res_id)
		liberr.ErrIsNil(ctx, err, "获取积分信息失败")
	})
	data := do.MemberIntegral{}
	if req.MonthlycardIntegral != nil {
		data.MonthlycardIntegral = *req.MonthlycardIntegral
	}
	if req.QuartercardIntegral != nil {
		data.QuartercardIntegral = *req.QuartercardIntegral
	}
	if req.AnnualcardIntegral != nil {
		data.AnnualcardIntegral = *req.AnnualcardIntegral
	}
	if req.KnowledgeIntegral != nil {
		data.KnowledgeIntegral = *req.KnowledgeIntegral
	}
	if req.VideoIntegral != nil {
		data.VideoIntegral = *req.VideoIntegral
	}
	if req.IssueIntegral != nil {
		data.IssueIntegral = *req.IssueIntegral
	}
	if req.Ordinary != nil {
		data.Ordinary = *req.Ordinary
	}
	if req.Select != nil {
		data.Select = *req.Select
	}
	if req.Ratio != nil {
		data.Ratio = *req.Ratio
	}
	if req.MonthlycardRatio != nil {
		data.MonthlycardRatio = *req.MonthlycardRatio
	}
	if req.QuartercardRatio != nil {
		data.QuartercardRatio = *req.QuartercardRatio
	}
	if req.AnnualcardRatio != nil {
		data.AnnualcardRatio = *req.AnnualcardRatio
	}
	if res_id != nil {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.MemberIntegral.Ctx(ctx).WherePri(res_id.Id).Update(data)
			liberr.ErrIsNil(ctx, err, "修改积分设置失败")
		})
		return
	} else {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.MemberIntegral.Ctx(ctx).Insert(data)
			liberr.ErrIsNil(ctx, err, "添加积分设置失败")
		})
	}
	return
}

func (s *sSysMemberManger) EditSu(ctx context.Context, req *system.MemberSubscriptionReq) (res *system.MemberSubscriptionRes, err error) {
	var res_id *entity.MemberSubscription
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.MemberSubscription.Ctx(ctx).Scan(&res_id)
		liberr.ErrIsNil(ctx, err, "获取订阅信息失败")
	})
	data := do.MemberSubscription{}
	if req.MonthlycardSubscriptionPrice != "" {
		MonthlycardSubscriptionPrice, _ := strconv.ParseFloat(req.MonthlycardSubscriptionPrice, 64)
		if MonthlycardSubscriptionPrice <= 0 {
			err = errors.New("月卡新增加订购单价 价格不能小于0")
			return
		}
		data.MonthlycardSubscriptionPrice = req.MonthlycardSubscriptionPrice
	}
	if req.QuartercardSubscriptionPrice != "" {
		QuartercardSubscriptionPrice, _ := strconv.ParseFloat(req.QuartercardSubscriptionPrice, 64)
		if QuartercardSubscriptionPrice <= 0 {
			err = errors.New("季卡新增加订购单价 价格不能小于0")
			return
		}
		data.QuartercardSubscriptionPrice = req.QuartercardSubscriptionPrice
	}
	if req.AnnualcardSubscriptionPrice != "" {
		AnnualcardSubscriptionPrice, _ := strconv.ParseFloat(req.AnnualcardSubscriptionPrice, 64)
		if AnnualcardSubscriptionPrice <= 0 {
			err = errors.New("年卡新增加订购单价 价格不能小于0")
			return
		}
		data.AnnualcardSubscriptionPrice = req.AnnualcardSubscriptionPrice
	}
	if res_id != nil {
		if req.MonthlycardSubscription != nil {
			data.MonthlycardSubscription = *req.MonthlycardSubscription
		}
		if req.QuartercardSubscription != nil {
			data.QuartercardSubscription = *req.QuartercardSubscription
		}
		if req.AnnualcardSubscription != nil {
			data.AnnualcardSubscription = *req.AnnualcardSubscription
		}

		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.MemberSubscription.Ctx(ctx).WherePri(res_id.Id).Update(data)
			liberr.ErrIsNil(ctx, err, "修改订阅设置失败")
		})
		return
	} else {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.MemberSubscription.Ctx(ctx).Insert(do.MemberSubscription{
				MonthlycardSubscription:      req.MonthlycardSubscription,
				QuartercardSubscription:      req.QuartercardSubscription,
				AnnualcardSubscription:       req.AnnualcardSubscription,
				MonthlycardSubscriptionPrice: req.MonthlycardSubscriptionPrice,
				QuartercardSubscriptionPrice: req.QuartercardSubscriptionPrice,
				AnnualcardSubscriptionPrice:  req.AnnualcardSubscriptionPrice,
			})
			liberr.ErrIsNil(ctx, err, "添加订阅设置失败")
		})
		return
	}
}

func (s *sSysMemberManger) FindFee(ctx context.Context, req *system.MemberFeeFindReq) (res *system.MemberFeeFindRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.MemberFee.Ctx(ctx).Scan(&res)
		log.Println("获取会费信息失败-------------->", res)
		liberr.ErrIsNil(ctx, err, "获取会费信息失败")
	})
	return
}

func (s *sSysMemberManger) FindIn(ctx context.Context, req *system.MemberInFindReq) (res *system.MemberInFindRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.MemberIntegral.Ctx(ctx).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取积分信息失败")
	})
	return
}

func (s *sSysMemberManger) FindSu(ctx context.Context, req *system.MemberSuFindReq) (res *system.MemberSuFindRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.MemberSubscription.Ctx(ctx).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取订阅信息失败")
	})
	return
}
