package publish

import (
	"context"
	"errors"
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
	service.RegisterPublish(New())
}

type sPublish struct {
}

func New() *sPublish {
	return &sPublish{}
}
func GetIntegral(ctx context.Context, intype string, integral int) (err error) {
	res := (*system_entity.MemberIntegral)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = system_dao.MemberIntegral.Ctx(ctx).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取积分信息失败")
	})
	if err != nil {
		return
	}
	if res != nil {
		if intype == "普通" {
			if integral > res.Ordinary {
				err = errors.New("积分超过上限")
				return
			}
		} else {
			if integral > res.Select {
				err = errors.New("积分超过上限")
				return
			}
		}
	}
	return
}

func (s *sPublish) KnowledgePublish(ctx context.Context, req *desk.KnowledgeReq) (res *desk.KnowledgeRes, err error) {
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
	err = GetIntegral(ctx, req.KnowledgeType, req.IntegralSetting)
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.MemberKnowledge.Ctx(ctx).TX(tx).Insert(do.MemberKnowledge{
				Title:                   req.Title,
				KnowledgeType:           req.KnowledgeType,
				Authority:               req.Authority,
				PrimaryClassification:   req.PrimaryClassification,
				SecondaryClassification: req.SecondaryClassification,
				IntegralSetting:         req.IntegralSetting,
				Content:                 req.Content,
				ReviewMessage:           req.ReviewMessage,
				ReviewStatus:            0,
				Type:                    req.Type,
				Abstract:                req.Abstract,
				CoverUrl:                req.CoverUrl,
				DetailsUrl:              req.DetailsUrl,
				VideoUrl:                req.VideoUrl,
				VideoIntroduction:       req.VideoIntroduction,
				AttachmentUrl:           req.AttachmentUrl,
				OpreviewMessage:         "",
				UserId:                  service.Context().GetUserId(ctx),
				UserName:                service.Context().GetLoginUser(ctx).UserName,
				CreatedAt:               gtime.Now(),
				Display:                 1,
				Crawler:                 0,
			})
			log.Println(e)
			liberr.ErrIsNil(ctx, e, "发布知识失败")
			if e != nil {
				return
			}
		})
		return err
	})
	return
}

func (s *sPublish) KnowledgePublishListGet(ctx context.Context, req *desk.KnowledgeGetListReq) (res *desk.KnowledgeGetListRes, err error) {
	res = new(desk.KnowledgeGetListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.MemberKnowledge.Ctx(ctx)
	order := "created_at desc"
	m = m.Where("user_id = ?", req.UserId)
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取我的知识列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.MemberKnowledgeList)
		liberr.ErrIsNil(ctx, err, "获取我的知识列表数据失败")
	})
	return
}

func (s *sPublish) KnowledgePublishGet(ctx context.Context, req *desk.KnowledgeGetReq) (res *desk.KnowledgeGetRes, err error) {
	res = new(desk.KnowledgeGetRes)
	m := dao.MemberKnowledge.Ctx(ctx)
	m = m.Where("id = ?", req.KnowledgeId)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.MemberKnowledge.MemberKnowledge)
		liberr.ErrIsNil(ctx, err, "获取知识数据失败")
	})
	if err != nil {
		return
	}
	if res.MemberKnowledge.MemberKnowledge != nil {
		if res.MemberKnowledge.MemberKnowledge.ReviewStatus == 0 {
			err = errors.New("知识正在审核中")
			return
		}
		user_id := res.MemberKnowledge.MemberKnowledge.UserId
		err = g.Try(ctx, func(ctx context.Context) {
			user_m := system_dao.MemberUser.Ctx(ctx)
			user_m = user_m.Where("id = ?", user_id)
			err = user_m.Limit(1).Scan(&res.MemberKnowledge.UserInfo)
			if res.MemberKnowledge.UserInfo != nil {
				res.MemberKnowledge.UserInfo.UserPassword = ""
			}
			liberr.ErrIsNil(ctx, err, "获取用户数据失败")
		})
		loginUser := service.Context().GetUserId(ctx)
		if loginUser != uint64(user_id) {
			if res.MemberKnowledge.MemberKnowledge.Authority == 1 {
				loginuserinfo := (*system_entity.MemberUser)(nil)
				err = g.Try(ctx, func(ctx context.Context) {
					loginuser_m := system_dao.MemberUser.Ctx(ctx)
					loginuser_m = loginuser_m.Where("id = ?", loginUser)
					err = loginuser_m.Limit(1).Scan(&loginuserinfo)
				})
				if loginuserinfo != nil {
					if loginuserinfo.MemberLevel == 0 {
						// 是否购买
						count := 0
						err = g.Try(ctx, func(ctx context.Context) {
							his_m := dao.HisKnowledge.Ctx(ctx)
							his_m = his_m.Where("user_id = ?", loginuserinfo.Id)
							his_m = his_m.Where("knowledge_id = ?", req.KnowledgeId)
							count, err = his_m.Count()
							liberr.ErrIsNil(ctx, err, "获取his知识失败")
							if err != nil {
								return
							}
						})
						if res.MemberKnowledge.MemberKnowledge.Authority == 1 {
							if count == 0 {
								if len(res.MemberKnowledge.MemberKnowledge.Content) > 3000 {
									res.MemberKnowledge.MemberKnowledge.Content = res.MemberKnowledge.MemberKnowledge.Content[:3000]
								} else {
									res.MemberKnowledge.MemberKnowledge.Content = res.MemberKnowledge.MemberKnowledge.Content
								}
								res.MemberKnowledge.MemberKnowledge.VideoUrl = ""
							}
						}
					}
				} else {
					if len(res.MemberKnowledge.MemberKnowledge.Content) > 3000 {
						res.MemberKnowledge.MemberKnowledge.Content = res.MemberKnowledge.MemberKnowledge.Content[:3000]
					} else {
						res.MemberKnowledge.MemberKnowledge.Content = res.MemberKnowledge.MemberKnowledge.Content
					}
				}
			}
		}
	}
	return
}

//KnowledgePublishEdit
//KnowledgeId             int    `p:"knowledgeId" v:"required#知识id不能为空"`
//	Title                   string `p:"title" `                  //标题
//	KnowledgeType           string `p:"knowledgeType"`           //知识类型
//	Authority               int    `p:"authority"`               //阅读下载权限
//	PrimaryClassification   string `p:"primaryClassification"`   //一级分类
//	SecondaryClassification string `p:"secondaryClassification"` // 二级分类
//	IntegralSetting         int    `p:"integralSetting"`         // 下载积分设置
//	Content                 string `p:"content"`                 // 正文
//	ReviewMessage           string `p:"reviewMessage"`           // 审核留言
func (s *sPublish) KnowledgePublishEdit(ctx context.Context, req *desk.KnowledgeEditReq) (res *desk.KnowledgeEditRes, err error) {
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
	data := do.MemberKnowledge{}
	if req.KnowledgeType != "" {
		data.Title = req.Title
	}
	if req.KnowledgeType != "" {
		data.KnowledgeType = req.KnowledgeType
	}
	if req.Authority != nil {
		data.Authority = &req.Authority
	}
	if req.PrimaryClassification != "" {
		data.PrimaryClassification = req.PrimaryClassification
	}
	if req.SecondaryClassification != "" {
		data.SecondaryClassification = req.SecondaryClassification
	}
	if req.IntegralSetting != nil {
		data.IntegralSetting = req.IntegralSetting
	}
	if req.Content != "" {
		data.Content = req.Content
	}
	if req.ReviewMessage != "" {
		data.ReviewMessage = req.ReviewMessage
	}
	if req.Abstract != "" {
		data.Abstract = req.Abstract
	}
	if req.Type != nil {
		data.Type = req.Type
		err = GetIntegral(ctx, req.KnowledgeType, *req.IntegralSetting)
		if err != nil {
			return
		}
	}
	if req.CoverUrl != "" {
		data.CoverUrl = req.CoverUrl
	}
	if req.DetailsUrl != "" {
		data.DetailsUrl = req.DetailsUrl
	}
	if req.VideoUrl != "" {
		data.VideoUrl = req.VideoUrl
	}
	if req.VideoIntroduction != "" {
		data.VideoIntroduction = req.VideoIntroduction
	}
	if req.AttachmentUrl != "" {
		data.AttachmentUrl = req.AttachmentUrl
	}
	data.UpdatedAt = gtime.Now()
	data.Display = 1
	m := dao.MemberKnowledge.Ctx(ctx)
	m = m.Where("id = ?", req.KnowledgeId)
	kl := (*entity.MemberKnowledge)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		// WherePri 默认是主键  where 需要带字段
		err = m.Scan(&kl)
		if err != nil {
			err = errors.New("查询知识失败")
			return
		}
	})
	if kl.IsAdmin != 1 {
		data.ReviewStatus = 0
	}
	err = g.Try(ctx, func(ctx context.Context) {
		// WherePri 默认是主键  where 需要带字段
		_, err = m.Update(data)
		if err != nil {
			err = errors.New("知识编辑失败")
			return
		}
	})
	if err != nil {
		err = errors.New("知识编辑失败")
		return
	}
	// 一般编辑比较少 所以在这修改而不是每次获取的时候查新的标题 修改收藏标题
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.MemberCollect.Ctx(ctx)
		m = m.Where("type = 'onlineKnowledgeBase' and article_id = ?", req.KnowledgeId)
		count, err := m.Count()
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

// KnowledgePublishDel
func (s *sPublish) KnowledgePublishDel(ctx context.Context, req *desk.KnowledgeDelReq) (res *desk.KnowledgeDelRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// WherePri 默认是主键  where 需要带字段
		_, err = dao.MemberKnowledge.Ctx(ctx).Where("id = ?", req.KnowledgeId).Delete()
		if err != nil {
			err = errors.New("知识删除失败")
			return
		}
	})
	if err != nil {
		err = errors.New("知识删除失败")
		return
	}
	return
}

//KnowledgeBuy
func (s *sPublish) KnowledgeBuy(ctx context.Context, req *desk.KnowledgeBuyReq) (res *desk.KnowledgeBuyRes, err error) {
	res = new(desk.KnowledgeBuyRes)
	count := 0
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.HisKnowledge.Ctx(ctx)
		m = m.Where("user_id = ?", req.UserId)
		m = m.Where("knowledge_id = ?", req.KnowledgeId)
		count, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取his知识失败")
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	if count > 0 {
		res.IsBuy = true
	} else {
		res.IsBuy = false
	}
	return
}
