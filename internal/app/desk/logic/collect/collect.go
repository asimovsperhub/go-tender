package collect

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"tender/api/v1/desk"
	"tender/internal/app/desk/dao"
	"tender/internal/app/desk/model/do"
	"tender/internal/app/desk/service"
	"tender/internal/app/system/consts"
	system_dao "tender/internal/app/system/dao"
	system_do "tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/library/liberr"
)

func init() {
	service.RegisterCollect(New())
}

type sCollect struct {
}

func New() *sCollect {
	return &sCollect{}
}

func (s *sCollect) CollectAdd(ctx context.Context, req *desk.CollectAddReq) (res *desk.CollectAddRes, err error) {
	result := (*entity.MemberCollect)(nil)
	user_id := service.Context().GetUserId(ctx)
	m := dao.MemberCollect.Ctx(ctx)
	m = m.Where("type = ?", req.Type)
	m = m.Where("user_id = ?", user_id)
	m = m.Where("article_id = ?", req.ArticleId)
	err = m.Scan(&result)
	if err != nil {
		return
	}
	if result != nil {
		err = errors.New("该文章已收藏")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.MemberCollect.Ctx(ctx).TX(tx).Insert(do.MemberCollect{
				Title:     req.Title,
				Type:      req.Type,
				Location:  req.Location,
				Industry:  req.Industry,
				ArticleId: req.ArticleId,
				Url:       req.Url,
				UserId:    user_id,
				CreatedAt: gtime.Now(),
			})

			if e != nil {
				err = errors.New("添加收藏失败")
				return
			}
		})
		return err
	})
	return
}

// CollectDel

func (s *sCollect) CollectDel(ctx context.Context, req *desk.CollectDelReq) (res *desk.CollectDelRes, err error) {
	m := dao.MemberCollect.Ctx(ctx)
	m = m.Where("type = ?", req.Type)
	m = m.Where("user_id = ?", req.UserId)
	m = m.Where("article_id = ?", req.ArticleId)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Delete()
		liberr.ErrIsNil(ctx, err, "删除数据失败")
	})
	return
}

//CollectGetList
func (s *sCollect) CollectGetList(ctx context.Context, req *desk.CollectGetListReq) (res *desk.CollectGetListRes, err error) {
	res = new(desk.CollectGetListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.MemberCollect.Ctx(ctx)
	order := "id"
	m = m.Where("user_id = ?", req.UserId)
	m = m.Where("type = ?", req.Type)
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取我的收藏列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.MemberCollectList)
		liberr.ErrIsNil(ctx, err, "获取我的收藏列表数据失败")
	})
	return
}

//CollectGet
func (s *sCollect) CollectGet(ctx context.Context, req *desk.CollectGetReq) (res *desk.CollectGetRes, err error) {
	res = new(desk.CollectGetRes)
	result := (*entity.MemberCollect)(nil)
	m := dao.MemberCollect.Ctx(ctx)
	m = m.Where("type = ?", req.Type)
	m = m.Where("user_id = ?", req.UserId)
	m = m.Where("article_id = ?", req.ArticleId)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&result)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	if result != nil {
		res.IsCollect = true
	} else {
		res.IsCollect = false
	}
	return
}

//CollectCountGet
func (s *sCollect) CollectCountGet(ctx context.Context, req *desk.CollectCountGetReq) (res *desk.CollectCountGetRes, err error) {
	res = new(desk.CollectCountGetRes)
	var count int
	m := dao.MemberCollect.Ctx(ctx)
	m = m.Where("type = ?", req.Type)
	m = m.Where("article_id = ?", req.ArticleId)
	err = g.Try(ctx, func(ctx context.Context) {
		count, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	res.Count = count
	return
}

func (s *sCollect) SubscribeAdd(ctx context.Context, req *desk.SubscribeAddReq) (res *desk.SubscribeAddRes, err error) {
	user_id := service.Context().GetUserId(ctx)
	user_info := (*entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = system_dao.MemberUser.Ctx(ctx).Where("id = ?", user_id).Scan(&user_info)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
		if err != nil {
			return
		}
	})
	sub, bysub, sub_count := 0, 0, 0

	if user_info != nil {
		sub = user_info.Subscribe
		bysub = user_info.BuySubscribe
		// subed = user_info.Subscribed
	}
	if sub+bysub == 0 {
		err = errors.New("可订阅数量为0")
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		// 已订阅数
		sub_count, err = dao.MemberSubscribe.Ctx(ctx).Where("user_id = ?", user_id).Count()

		liberr.ErrIsNil(ctx, err, "获取订阅信息失败")
		if err != nil {
			return
		}
	})
	if sub_count >= sub+bysub {
		err = errors.New("可订阅数量为0")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.MemberSubscribe.Ctx(ctx).TX(tx).Insert(do.MemberSubscribe{
				Name:      req.Name,
				Type:      req.Type,
				Industry:  req.Industry,
				Location:  req.Location,
				Keywords:  req.Keywords,
				UserId:    service.Context().GetUserId(ctx),
				CreatedAt: gtime.Now(),
			})

			if e != nil {
				err = errors.New("添加订阅失败")
				return
			}
		})
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := system_dao.MemberUser.Ctx(ctx).TX(tx).Where("id = ?", user_id).Update(
				system_do.MemberUser{Subscribed: sub_count + 1})
			if e != nil {
				err = errors.New("更新已订阅数失败")
				return
			}
		})
		return err
	})
	return
}

// CollectDel

func (s *sCollect) SubscribeDel(ctx context.Context, req *desk.SubscribeDelReq) (res *desk.SubscribeDelRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// WherePri 默认是主键  where 需要带字段
		_, err = dao.MemberSubscribe.Ctx(ctx).Where("id = ?", req.SubscribeId).Delete()
		if err != nil {
			err = errors.New("订阅删除失败")
			return
		}
	})
	return
}

//SubscribeEdit
func (s *sCollect) SubscribeEdit(ctx context.Context, req *desk.SubscribeEditReq) (res *desk.SubscribeEditRes, err error) {
	data := do.MemberSubscribe{}
	if req.Name != "" {
		data.Name = req.Name
	}
	if req.Type != "" {
		data.Type = req.Type
	}
	if req.Industry != "" {
		data.Industry = req.Industry
	}
	if req.Location != "" {
		data.Location = req.Location
	}
	if req.Keywords != "" {
		data.Keywords = req.Keywords
	}
	data.UpdatedAt = gtime.Now()
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.MemberSubscribe.Ctx(ctx).TX(tx).Where("id = ?", req.SubscribeId).Update(data)

			if e != nil {
				err = errors.New("更新订阅失败")
				return
			}
		})
		return err
	})
	return
}

//SubscribeList
func (s *sCollect) SubscribeList(ctx context.Context, req *desk.SubscribeListReq) (res *desk.SubscribeListRes, err error) {
	res = new(desk.SubscribeListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.MemberSubscribe.Ctx(ctx)
	order := "id"
	m = m.Where("user_id = ?", req.UserId)
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取我的订阅列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.MemberSubscribeList)
		liberr.ErrIsNil(ctx, err, "获取我的订阅列表数据失败")
	})
	return
}

//SubscribeGet

func (s *sCollect) SubscribeGet(ctx context.Context, req *desk.SubscribeGetReq) (res *desk.SubscribeGetRes, err error) {
	res = new(desk.SubscribeGetRes)
	m := dao.MemberSubscribe.Ctx(ctx)
	m = m.Where("id = ?", req.SubscribeId)
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err, "获取订阅失败")
		err = m.Scan(&res.MemberSubscribe)
		liberr.ErrIsNil(ctx, err, "获取订阅数据失败")
	})
	return
}

//SubscribeCan

func (s *sCollect) SubscribeCan(ctx context.Context, req *desk.SubscribeCanReq) (res *desk.SubscribeCanRes, err error) {
	res = new(desk.SubscribeCanRes)
	user_id := service.Context().GetUserId(ctx)
	user_info := (*entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = system_dao.MemberUser.Ctx(ctx).Where("id = ?", user_id).Scan(&user_info)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
		if err != nil {
			return
		}
	})
	sub, sub_count := 0, 0

	if user_info != nil {
		sub = user_info.Subscribe
	}
	err = g.Try(ctx, func(ctx context.Context) {
		sub_count, err = dao.MemberSubscribe.Ctx(ctx).Where("user_id = ?", user_id).Count()

		liberr.ErrIsNil(ctx, err, "获取订阅信息失败")
		if err != nil {
			return
		}
	})
	res.Count = sub - sub_count
	return
}
