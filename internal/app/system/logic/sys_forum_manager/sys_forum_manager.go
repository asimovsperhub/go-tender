package sys_forum_manager

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	g_udid "github.com/google/uuid"
	"tender/api/v1/system"
	"tender/internal/app/desk/dao"
	desk_model "tender/internal/app/desk/model"
	"tender/internal/app/desk/model/do"
	desk_entity "tender/internal/app/desk/model/entity"
	"tender/internal/app/system/consts"
	system_dao "tender/internal/app/system/dao"
	system_do "tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
	"tender/internal/packed/websocket"
	"tender/library/liberr"
)

func init() {
	service.RegisterForumManger(New())
}

func New() *sSysForumManger {
	return &sSysForumManger{}
}

type sSysForumManger struct {
}

func (s *sSysForumManger) BbsGetListAll(ctx context.Context, req *system.BbsGetListAllReq) (res *system.BbsGetListAllRes, err error) {
	res = new(system.BbsGetListAllRes)
	bbs := [](*desk_model.Bbs)(nil)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.Bbs.Ctx(ctx)
	if req.KeyWords != "" {
		m = m.Where(fmt.Sprintf("MATCH(title,content,classification) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
	}
	order := "rank desc,created_at desc"
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取论坛列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛列表数据失败")
	})
	for i := 0; i < len(bbs); i++ {
		user_info := (*entity.MemberUser)(nil)
		err = g.Try(ctx, func(ctx context.Context) {
			user_m := system_dao.MemberUser.Ctx(ctx)
			user_m = user_m.Where("id = ?", bbs[i].UserId)
			err = user_m.Limit(1).Scan(&user_info)
			if user_info != nil {
				user_info.UserPassword = ""
			}
			liberr.ErrIsNil(ctx, err, "获取用户数据失败")
		})
		bbs_content := (*desk_entity.BbsContent)(nil)
		m_ct := dao.BbsContent.Ctx(ctx)
		m_ct = m_ct.Where("bbs_id = ?", bbs[i].Id)
		err = g.Try(ctx, func(ctx context.Context) {
			err = m_ct.Scan(&bbs_content)
			liberr.ErrIsNil(ctx, err, "获取论坛内容失败")
		})
		col := 0
		err = g.Try(ctx, func(ctx context.Context) {
			col, err = dao.MemberCollect.Ctx(ctx).Where("type = 'forum' and article_id = ? ", bbs[i].Id).Count()
			liberr.ErrIsNil(ctx, err, "获取论坛内容失败")
		})
		bbs[i].CollectCount = col
		bs := 0
		err = g.Try(ctx, func(ctx context.Context) {
			bs, err = dao.Bbs.Ctx(ctx).Where("user_id = ? ", bbs[i].UserId).Count()
			liberr.ErrIsNil(ctx, err, "获取论坛内容失败")
		})
		bbs[i].ReleaseCount = bs
		res.BbsList = append(res.BbsList, system.BbsList{Bbs: bbs[i], UserInfo: user_info})
	}
	return
}

//BbsGet
func (s *sSysForumManger) BbsGet(ctx context.Context, req *system.BbsGetReq) (res *system.BbsGetRes, err error) {
	res = new(system.BbsGetRes)
	m := dao.Bbs.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.Bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛数据失败")
	})
	if err != nil {
		return
	}
	//if res.Bbs != nil {
	//	m_ct := dao.BbsContent.Ctx(ctx)
	//	m_ct = m_ct.Where("bbs_id = ?", res.Bbs.Id)
	//	err = g.Try(ctx, func(ctx context.Context) {
	//		err = m_ct.Scan(&res.BbsContent)
	//		liberr.ErrIsNil(ctx, err, "获取论坛内容失败")
	//	})
	//}
	return
}

//BbsDel

func (s *sSysForumManger) BbsDel(ctx context.Context, req *system.BbsDelReq) (res *system.BbsDelRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Bbs.Ctx(ctx).Where("id = ?", req.Id).Update(do.Bbs{Status: 0})
		liberr.ErrIsNil(ctx, err, "删除论坛失败")

	})
	return
}

//BbsRestore

func (s *sSysForumManger) BbsRestore(ctx context.Context, req *system.BbsRestoreReq) (res *system.BbsRestoreRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Bbs.Ctx(ctx).Where("id = ?", req.Id).Update(do.Bbs{Status: 1})
		liberr.ErrIsNil(ctx, err, "恢复论坛失败")

	})
	return
}

//BbsForever

func (s *sSysForumManger) BbsForever(ctx context.Context, req *system.BbsForeverReq) (res *system.BbsForeverRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Bbs.Ctx(ctx).Where("id = ?", req.Id).Delete()
		liberr.ErrIsNil(ctx, err, "永久删除论坛失败")

	})
	return
}

//BbsGetReviewList
func (s *sSysForumManger) BbsGetReviewList(ctx context.Context, req *system.BbsGetReviewListReq) (res *system.BbsGetReviewListRes, err error) {
	res = new(system.BbsGetReviewListRes)
	bbs := [](*desk_model.Bbs)(nil)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.Bbs.Ctx(ctx)
	if req.KeyWords != "" {
		m = m.Where(fmt.Sprintf("MATCH(title,content,classification) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
	}
	order := "created_at desc"
	m = m.Where("review_status <> 1")
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取论坛列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛列表数据失败")
	})
	for i := 0; i < len(bbs); i++ {
		user_info := (*entity.MemberUser)(nil)
		err = g.Try(ctx, func(ctx context.Context) {
			user_m := system_dao.MemberUser.Ctx(ctx)
			user_m = user_m.Where("id = ?", bbs[i].UserId)
			err = user_m.Limit(1).Scan(&user_info)
			if user_info != nil {
				user_info.UserPassword = ""
			}
			liberr.ErrIsNil(ctx, err, "获取用户数据失败")
		})
		//bbs_content := (*desk_entity.BbsContent)(nil)
		//m_ct := dao.BbsContent.Ctx(ctx)
		//m_ct = m_ct.Where("bbs_id = ?", bbs[i].Id)
		//err = g.Try(ctx, func(ctx context.Context) {
		//	err = m_ct.Scan(&bbs_content)
		//	liberr.ErrIsNil(ctx, err, "获取论坛内容失败")
		//})
		res.BbsList = append(res.BbsList, system.BbsList{Bbs: bbs[i], UserInfo: user_info})
	}
	return
}

//BbsReview

func (s *sSysForumManger) BbsReview(ctx context.Context, req *system.BbsReviewReq) (res *system.BbsReviewRes, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.Bbs.Ctx(ctx).TX(tx).WherePri(req.Id).Update(do.Bbs{
				ReviewStatus:  req.Status,
				UpdatedAt:     gtime.Now(),
				ReviewMessage: req.OpReviewMessage,
				// 操作管理员id
				// OperationId: service.Context().GetUserId(ctx),
			})
			liberr.ErrIsNil(ctx, err, "审核论坛失败")
		})
		return err
	})
	bbs := (*desk_entity.Bbs)(nil)
	m := dao.Bbs.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	switch req.Status {
	case 1:
		req.OpReviewMessage = "帖子审核通过: " + req.OpReviewMessage
		userinfo := (*entity.MemberUser)(nil)
		user_m := system_dao.MemberUser.Ctx(ctx)
		err = g.Try(ctx, func(ctx context.Context) {
			user_m = user_m.Where(fmt.Sprintf("%s='%d'", system_dao.MemberUser.Columns().Id, bbs.UserId))
			err = user_m.Limit(1).Scan(&userinfo)
		})
		setting := (*entity.MemberIntegral)(nil)
		err = g.Try(ctx, func(ctx context.Context) {
			err = system_dao.MemberIntegral.Ctx(ctx).Scan(&setting)
			liberr.ErrIsNil(ctx, err, "获取积分信息失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			user_m = user_m.Where(fmt.Sprintf("%s='%d'", system_dao.MemberUser.Columns().Id, bbs.UserId))
			_, err = user_m.Update(system_do.MemberUser{Integral: userinfo.Integral + setting.IssueIntegral})
			liberr.ErrIsNil(ctx, err, "更新用户积分失败")
		})
	case 2:
		req.OpReviewMessage = "帖子审核未通过: " + req.OpReviewMessage
	}
	if req.Status != 0 {
		msgId := g_udid.New().String()
		_, err = system_dao.SysWsMsg.Ctx(ctx).Insert(system_do.SysWsMsg{
			MessageId: msgId,
			UserId:    bbs.UserId,
			Content:   req.OpReviewMessage,
			IsRead:    0,
			IsDel:     0,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		})
		if err != nil {
			return nil, fmt.Errorf("插入消息失败: %w", err)
		}

		websocket.SendToUser(uint64(bbs.UserId), &websocket.WResponse{
			Event: "context",
			Data: map[string]string{
				"messageId": msgId,
				"content":   req.OpReviewMessage,
			},
		})
	}
	return
}

//BbsTop

func (s *sSysForumManger) BbsTop(ctx context.Context, req *system.BbsTopReq) (res *system.BbsTopRes, err error) {
	res = new(system.BbsTopRes)
	rank := ""
	if req.Type == 1 {
		for _, id := range req.Ids {
			rank = "m" + id
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = dao.Bbs.Ctx(ctx).Where("id = ?", id).Update(do.Bbs{Rank: rank})
				liberr.ErrIsNil(ctx, err, "置顶论坛失败")
			})
		}
	} else {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.Bbs.Ctx(ctx).Where("id in (?)", req.Ids).Update("rank = ''")
			liberr.ErrIsNil(ctx, err, "取消置顶论坛失败")
		})
	}
	return
}
