package bbs

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"log"
	"tender/api/v1/desk"
	"tender/internal/app/desk/consts"
	"tender/internal/app/desk/dao"
	"tender/internal/app/desk/model/do"
	"tender/internal/app/desk/model/entity"
	"tender/internal/app/desk/service"
	system_dao "tender/internal/app/system/dao"
	system_entity "tender/internal/app/system/model/entity"
	"tender/library/liberr"
)

func init() {
	service.RegisterBbs(New())
}

type sBbs struct {
}

func New() *sBbs {
	return &sBbs{}
}

func (s *sBbs) BbsPublish(ctx context.Context, req *desk.BbsPublishReq) (res *desk.BbsPublishRes, err error) {
	if len(req.Title) > 500 {
		err = errors.New("标题长度不能超过500")
		return
	}
	if len(req.Abstract) > 500 {
		err = errors.New("摘要长度不能超过500")
		return
	}
	if len(req.ReviewMessage) > 500 {
		err = errors.New("审核留言长度不能超过500")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.Bbs.Ctx(ctx).TX(tx).InsertAndGetId(do.Bbs{
				Title:          req.Title,
				Abstract:       req.Abstract,
				Classification: req.Classification,
				ReviewMessage:  req.ReviewMessage,
				ReviewStatus:   0,
				UserId:         service.Context().GetUserId(ctx),
				Content:        req.Content,
				CreatedAt:      gtime.Now(),
			})
			liberr.ErrIsNil(ctx, e, "发布论坛失败")
			if e != nil {
				return
			}
			//_, e = dao.BbsContent.Ctx(ctx).TX(tx).Insert(do.BbsContent{
			//	BbsId:   bbs_id,
			//	Content: req.Content,
			//})
			//liberr.ErrIsNil(ctx, e, "发布论坛内容失败")
			//if e != nil {
			//	return
			//}
		})
		return err
	})
	return
}

// BbsGetList

func (s *sBbs) BbsGetList(ctx context.Context, req *desk.BbsGetListReq) (res *desk.BbsGetListRes, err error) {
	res = new(desk.BbsGetListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.Bbs.Ctx(ctx)
	m = m.Where("user_id = ?", req.UserId)
	order := "created_at desc"
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取我的论坛列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.BbsList)
		liberr.ErrIsNil(ctx, err, "获取我的论坛列表数据失败")
	})
	for i := 0; i < len(res.BbsList); i++ {
		col := 0
		err = g.Try(ctx, func(ctx context.Context) {
			col, err = dao.MemberCollect.Ctx(ctx).Where("type = 'forum' and article_id = ? ", res.BbsList[i].Id).Count()
			liberr.ErrIsNil(ctx, err, "获取论坛内容失败")
		})
		res.BbsList[i].CollectCount = col
	}
	return
}

//BbsGetListAll

func (s *sBbs) BbsGetListAll(ctx context.Context, req *desk.BbsGetListAllReq) (res *desk.BbsGetListAllRes, err error) {
	bbs := [](*entity.Bbs)(nil)
	res = new(desk.BbsGetListAllRes)
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
		// m = m.Where("title like ?", "%"+req.KeyWords+"%")
	}
	orderby := "rank desc,created_at desc"
	if req.Sort != nil {
		switch *req.Sort {
		case 0:
			orderby = "rank desc,views desc"
		case 1:
			orderby = "rank desc,created_at asc"
		case 2:
			orderby = "rank desc,created_at desc"
		}
	}

	m = m.Where("review_status = 1")
	m = m.Where("status = 1")
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取论坛列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(orderby).Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛列表数据失败")
	})
	for i := 0; i < len(bbs); i++ {
		user_info := (*system_entity.MemberUser)(nil)
		err = g.Try(ctx, func(ctx context.Context) {
			user_m := system_dao.MemberUser.Ctx(ctx)
			user_m = user_m.Where("id = ?", bbs[i].UserId)
			err = user_m.Limit(1).Scan(&user_info)
			if user_info != nil {
				user_info.UserPassword = ""
			}
			liberr.ErrIsNil(ctx, err, "获取用户数据失败")
		})
		//bbs_content := (*entity.BbsContent)(nil)
		//m_ct := dao.BbsContent.Ctx(ctx)
		//m_ct = m_ct.Where("bbs_id = ?", bbs[i].Id)
		//err = g.Try(ctx, func(ctx context.Context) {
		//	err = m_ct.Scan(&bbs_content)
		//	liberr.ErrIsNil(ctx, err, "获取论坛内容失败")
		//})
		res.BbsList = append(res.BbsList, desk.BbsList{Bbs: bbs[i], UserInfo: user_info})
	}
	return
}

//BbsGet
func (s *sBbs) BbsGet(ctx context.Context, req *desk.BbsGetReq) (res *desk.BbsGetRes, err error) {
	res = new(desk.BbsGetRes)
	m := dao.Bbs.Ctx(ctx)
	m = m.Where("id = ?", req.BbsId)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.Bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛数据失败")
	})
	if err != nil {
		return
	}
	if res != nil {
		if res.Bbs.ReviewStatus == 0 {
			err = errors.New("帖子正在审核中")
			return
		}
		user_id := res.Bbs.UserId
		err = g.Try(ctx, func(ctx context.Context) {
			user_m := system_dao.MemberUser.Ctx(ctx)
			user_m = user_m.Where("id = ?", user_id)
			err = user_m.Limit(1).Scan(&res.UserInfo)
			if res.UserInfo != nil {
				res.UserInfo.UserPassword = ""
			}
			liberr.ErrIsNil(ctx, err, "获取用户数据失败")
		})
		//m_ct := dao.BbsContent.Ctx(ctx)
		//m_ct = m_ct.Where("bbs_id = ?", res.Bbs.Id)
		//err = g.Try(ctx, func(ctx context.Context) {
		//	err = m_ct.Scan(&res.BbsContent)
		//	liberr.ErrIsNil(ctx, err, "获取论坛内容失败")
		//})
		reply := ([]*entity.BbsReply)(nil)
		m_rp := dao.BbsReply.Ctx(ctx)
		m_rp = m_rp.Where("bbs_id = ?", res.Bbs.Id)
		err = g.Try(ctx, func(ctx context.Context) {
			err = m_rp.Scan(&reply)
			liberr.ErrIsNil(ctx, err, "获取论坛回复失败")
		})
		for i := 0; i < len(reply); i++ {
			rp_userinfo := (*system_entity.MemberUser)(nil)
			user_id = reply[i].UserId
			user_m := system_dao.MemberUser.Ctx(ctx)
			user_m = user_m.Where("id = ?", user_id)
			err = user_m.Limit(1).Scan(&rp_userinfo)
			if rp_userinfo != nil {
				rp_userinfo.UserPassword = ""
			}
			liberr.ErrIsNil(ctx, err, "获取用户数据失败")
			res.BbsReply = append(res.BbsReply, desk.BbsReply{UserInfo: rp_userinfo, Reply: reply[i]})
		}
	}
	return
}

// BbsEdit
func (s *sBbs) BbsEdit(ctx context.Context, req *desk.BbsEditReq) (res *desk.BbsEditRes, err error) {
	if len(req.Title) > 500 {
		err = errors.New("标题长度不能超过500")
		return
	}
	if len(req.Abstract) > 500 {
		err = errors.New("摘要长度不能超过500")
		return
	}
	if len(req.ReviewMessage) > 500 {
		err = errors.New("审核留言长度不能超过500")
		return
	}
	data := do.Bbs{}
	// content := do.BbsContent{}
	if req.Content != "" {
		data.Content = req.Content
	}
	if req.Abstract != "" {
		data.Abstract = req.Abstract
	}
	if req.Title != "" {
		data.Title = req.Title
	}
	if req.ReviewMessage != "" {
		data.ReviewMessage = req.ReviewMessage
	}
	data.ReviewStatus = 0
	data.UpdatedAt = gtime.Now()
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Bbs.Ctx(ctx).Where("id = ?", req.BbsId).Update(data)
		if err != nil {
			err = errors.New("知识编辑失败")
			return
		}
	})
	//err = g.Try(ctx, func(ctx context.Context) {
	//	_, err = dao.BbsContent.Ctx(ctx).Where("bbs_id= ?", req.BbsId).Update(content)
	//	if err != nil {
	//		err = errors.New("知识内容编辑失败")
	//		return
	//	}
	//})
	// 一般编辑比较少 所以在这修改而不是每次获取的时候查新的标题 修改收藏标题
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.MemberCollect.Ctx(ctx)
		m = m.Where("type = 'forum' and article_id = ?", req.BbsId)
		count, err := m.Count()
		log.Println(count)
		if err != nil {
			err = errors.New("获取收藏失败")
			return
		}
		if count > 0 {
			_, err = m.Update(&do.MemberCollect{Title: req.Title})
			err = errors.New("更新收藏失败")
			return
		}
	})
	return
}

//BbsDel
func (s *sBbs) BbsDel(ctx context.Context, req *desk.BbsDelReq) (res *desk.BbsDelRes, err error) {
	bbs := (*entity.Bbs)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Bbs.Ctx(ctx).Where("id = ?", req.BbsId).Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛数据失败")
	})

	login_id := service.Context().GetUserId(ctx)
	if uint64(bbs.UserId) == login_id {
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = dao.Bbs.Ctx(ctx).Where("id = ?", req.BbsId).Delete()
				if err != nil {
					err = errors.New("论坛删除失败")
					return
				}
			})
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = dao.BbsContent.Ctx(ctx).Where("bbs_id = ?", req.BbsId).Delete()
				if err != nil {
					err = errors.New("论坛内容删除失败")
					return
				}
			})
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = dao.BbsReply.Ctx(ctx).Where("bbs_id = ?", req.BbsId).Delete()
				if err != nil {
					err = errors.New("论坛回复删除失败")
					return
				}
			})
			return err
		})
	} else {
		err = errors.New("论坛文章不属于当前用户")
		return
	}
	return
}

//BbsBrowse
func (s *sBbs) BbsBrowse(ctx context.Context, req *desk.BbsBrowseReq) (res *desk.BbsBrowseRes, err error) {
	bbs := (*entity.Bbs)(nil)
	m := dao.Bbs.Ctx(ctx)
	m = m.Where("id = ?", req.BbsId)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛数据失败")
	})
	if err != nil {
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Update(do.Bbs{
			Views: bbs.Views + 1,
		})
		liberr.ErrIsNil(ctx, err, "更新论坛浏览量失败")
	})
	return
}

//BbsLike
func (s *sBbs) BbsLike(ctx context.Context, req *desk.BbsLikeReq) (res *desk.BbsLikeRes, err error) {
	bbs := (*entity.Bbs)(nil)
	user_id := service.Context().GetUserId(ctx)
	m := dao.Bbs.Ctx(ctx)
	m = m.Where("id = ?", req.BbsId)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛数据失败")
	})
	if err != nil {
		return
	}
	if req.Type == 1 {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = m.Update(do.Bbs{
				LikeCount: bbs.LikeCount + 1,
			})
			liberr.ErrIsNil(ctx, err, "更新论坛点赞量失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.BbsLike.Ctx(ctx).Insert(do.BbsLike{
				BbsId:     req.BbsId,
				UserId:    user_id,
				CreatedAt: gtime.Now(),
			})
			liberr.ErrIsNil(ctx, err, "插入论坛点赞失败")
		})
	}
	if req.Type == 0 {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = m.Update(do.Bbs{
				LikeCount: bbs.LikeCount - 1,
			})
			liberr.ErrIsNil(ctx, err, "更新论坛点赞量失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.BbsLike.Ctx(ctx).Where("bbs_id = ? and user_id = ?", req.BbsId, user_id).Delete()
			liberr.ErrIsNil(ctx, err, "删除论坛点赞失败")
		})
	}
	return
}

//BbsComment
func (s *sBbs) BbsComment(ctx context.Context, req *desk.BbsCommentReq) (res *desk.BbsCommentRes, err error) {
	bbs := (*entity.Bbs)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Bbs.Ctx(ctx).Where("id = ?", req.BbsId).Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛数据失败")
	})
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.Bbs.Ctx(ctx).Where("id = ? ", req.BbsId).Update(do.Bbs{
				ReplyCount: bbs.ReplyCount + 1,
			})
			liberr.ErrIsNil(ctx, err, "更新论坛评论量失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.BbsReply.Ctx(ctx).TX(tx).Insert(do.BbsReply{
				BbsId:     req.BbsId,
				Content:   req.Content,
				UserId:    service.Context().GetUserId(ctx),
				CreatedAt: gtime.Now(),
			})
			liberr.ErrIsNil(ctx, e, "论坛评论失败")
			if e != nil {
				return
			}
		})
		return err
	})
	return
}

//BbsCommentLike
func (s *sBbs) BbsCommentLike(ctx context.Context, req *desk.BbsCommentLikeReq) (res *desk.BbsCommentLikeRes, err error) {
	reply := (*entity.BbsReply)(nil)
	user_id := service.Context().GetUserId(ctx)
	m := dao.BbsReply.Ctx(ctx)
	m = m.Where("id = ?", req.CommentId)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&reply)
		liberr.ErrIsNil(ctx, err, "获取评论数据失败")
	})
	if err != nil {
		return
	}
	if req.Type == 1 {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = m.Update(do.BbsReply{
				LikeCount: reply.LikeCount + 1,
			})
			liberr.ErrIsNil(ctx, err, "更新评论点赞量失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.BbsLike.Ctx(ctx).Insert(do.BbsLike{
				BbsId:     req.BbsId,
				ReplyId:   req.CommentId,
				UserId:    user_id,
				CreatedAt: gtime.Now(),
			})
			liberr.ErrIsNil(ctx, err, "插入评论点赞失败")
		})
	}
	if req.Type == 0 {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = m.Update(do.BbsReply{
				LikeCount: reply.LikeCount - 1,
			})
			liberr.ErrIsNil(ctx, err, "更新评论点赞量失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.BbsLike.Ctx(ctx).Where("reply_id = ? and user_id = ?", req.CommentId, user_id).Delete()
			liberr.ErrIsNil(ctx, err, "删除评论点赞失败")
		})
	}
	return
}

//BbsCommentDel
func (s *sBbs) BbsCommentDel(ctx context.Context, req *desk.BbsCommentDelReq) (res *desk.BbsCommentDelRes, err error) {
	comment := (*entity.BbsReply)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.BbsReply.Ctx(ctx).Where("id = ?", req.CommentId).Scan(&comment)
		liberr.ErrIsNil(ctx, err, "获取评论数据失败")
	})
	bbs := (*entity.Bbs)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Bbs.Ctx(ctx).Where("id = ?", req.BbsId).Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛数据失败")
	})
	if err != nil {
		return
	}
	login_user := service.Context().GetLoginUser(ctx)
	if login_user != nil {
		if uint64(comment.UserId) == login_user.Id {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = g.Try(ctx, func(ctx context.Context) {
					_, err = dao.BbsReply.Ctx(ctx).Where("id = ?", req.CommentId).Delete()
					if err != nil {
						err = errors.New("评论删除失败")
						return
					}
				})
				err = g.Try(ctx, func(ctx context.Context) {
					_, err = dao.Bbs.Ctx(ctx).Where("id = ? ", req.BbsId).Update(do.Bbs{
						ReplyCount: bbs.LikeCount - 1,
					})
					liberr.ErrIsNil(ctx, err, "更新论坛评论量失败")
				})
				return err
			})
		} else {
			err = errors.New("评论不属于当前用户")
			return
		}

	}
	return
}

//BbsReply
func (s *sBbs) BbsReply(ctx context.Context, req *desk.BbsReplyReq) (res *desk.BbsReplyRes, err error) {
	bbs := (*entity.Bbs)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Bbs.Ctx(ctx).Where("id = ?", req.BbsId).Scan(&bbs)
		liberr.ErrIsNil(ctx, err, "获取论坛数据失败")
	})
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.Bbs.Ctx(ctx).Where("id = ? ", req.BbsId).Update(do.Bbs{
				ReplyCount: bbs.ReplyCount + 1,
			})
			liberr.ErrIsNil(ctx, err, "更新论坛评论量失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.BbsReply.Ctx(ctx).TX(tx).Insert(do.BbsReply{
				BbsId:     req.BbsId,
				Content:   req.Content,
				ReplyId:   req.ReplyId,
				UserId:    service.Context().GetUserId(ctx),
				CreatedAt: gtime.Now(),
			})
			liberr.ErrIsNil(ctx, e, "论坛回复失败")
			if e != nil {
				return
			}
		})
		return err
	})
	return
}

//BbsGetLike

func (s *sBbs) BbsGetLike(ctx context.Context, req *desk.BbsGetLikeReq) (res *desk.BbsGetLikeRes, err error) {
	res = new(desk.BbsGetLikeRes)
	bbslike := [](*entity.BbsLike)(nil)
	m := dao.BbsLike.Ctx(ctx)
	user_id := service.Context().GetUserId(ctx)
	if req.UserId != nil {
		user_id = uint64(*req.UserId)
	}
	if req.BbsId != nil {
		m = m.Where("bbs_id = ?", *req.BbsId)
	}
	m = m.Where("user_id = ?", user_id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&bbslike)
		liberr.ErrIsNil(ctx, err, "获取点赞列表数据失败")
	})
	if bbslike != nil {
		for i := 0; i < len(bbslike); i++ {
			if bbslike[i].ReplyId == 0 {
				res.LikeBbs = true
			} else {
				res.ReplyLike = append(res.ReplyLike, *bbslike[i])
			}
		}
	} else {
		res.LikeBbs = false
		res.ReplyLike = nil
	}
	return
}

//FeedbackPublish

func (s *sBbs) FeedbackPublish(ctx context.Context, req *desk.FeedbackPublishReq) (res *desk.FeedbackPublishRes, err error) {
	userinfo := (*system_entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		m := system_dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%d'", system_dao.MemberUser.Columns().Id, service.Context().GetUserId(ctx)))
		err = m.Limit(1).Scan(&userinfo)
	})
	if userinfo.MemberLevel < 1 {
		err = errors.New("非会员卡不能反馈")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.Feedback.Ctx(ctx).TX(tx).Insert(do.Feedback{
				BulletinId:         req.BulletinId,
				Company:            req.Company,
				ContactPerson:      req.ContactPerson,
				ContactInformation: req.ContactInformation,
				Remarks:            req.Remarks,
				Attachment:         req.Attachment,
				UserId:             userinfo.Id,
				UserName:           userinfo.UserNickname,
				CreatedAt:          gtime.Now(),
			})
			liberr.ErrIsNil(ctx, e, "发布反馈失败")
		})
		return err
	})
	return
}
