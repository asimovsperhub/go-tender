package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/desk/model"
	"tender/internal/app/desk/model/entity"
	system_entity "tender/internal/app/system/model/entity"
)

// bbs
//	Id            string //
//	Title         string // 标题
//	Abstract      string // 摘要
//	ReviewMessage string // 审核留言
//	ReviewStatus  string // 审核状态;0:待审核,1:通过，2:未通过
//	Views         string // 浏览量
//	ReplyCount    string // 回复量
//	LikeCount     string // 点赞量
//	UserId        string // 用户id
//	CreatedAt     string // 发布时间
//	UpdatedAt     string // 更新时间
//	DeletedAt     string // 删除时间

// content

//  Id:      "id",
//	BbsId:   "bbs_id",
//	Content: "content",

//reply

//  Id:        "id",
//	BbsId:     "bbs_id",
//	ReplyId:   "reply_id",
//	Content:   "content",
//	LikeCount: "like_count",
//	UserId:    "user_id",
//	CreatedAt: "created_at",
//	UpdatedAt: "updated_at",
//	DeletedAt: "deleted_at",

// 发布论坛
type BbsPublishReq struct {
	g.Meta         `path:"/bbs/publish" tags:"论坛" method:"post" summary:"发布论坛"`
	Title          string `p:"title" v:"max-length:80#字数超过80"`           // 标题
	Abstract       string `p:"abstract" v:"max-length:400#字数超过400"`      // 摘要
	Content        string `p:"content"`                                  // 正文
	ReviewMessage  string `p:"reviewMessage" v:"max-length:400#字数超过400"` // 审核留言
	Classification string `p:"classification"`                           // 所属类别
	common.Author
}
type BbsPublishRes struct {
}

// 我的论坛列表
type BbsGetListReq struct {
	g.Meta `path:"/bbs/list" tags:"论坛" method:"get" summary:"我的论坛列表"`
	UserId int `p:"userId" v:"required#用户id不能为空"`
	common.PageReq
	common.Author
}
type BbsGetListRes struct {
	BbsList []*model.Bbs `json:"bbsList"`
	common.ListRes
}

// 论坛列表
type BbsGetListAllReq struct {
	g.Meta   `path:"/bbs/all/list" tags:"论坛" method:"get" summary:"论坛列表"`
	KeyWords string `p:"keyWords"`
	Sort     *int   `p:"sort"`
	common.PageReq
	common.Author
}
type BbsList struct {
	Bbs *entity.Bbs `json:"bbs"`
	// BbsContent *entity.BbsContent        `json:"bbsContent"`
	UserInfo *system_entity.MemberUser `json:"memberUser"`
}

type BbsGetListAllRes struct {
	BbsList []BbsList `json:"bbsList"`
	common.ListRes
}

// 单个论坛
type BbsGetReq struct {
	g.Meta `path:"/bbs/get" tags:"论坛" method:"get" summary:"获取单个论坛"`
	BbsId  int `p:"bbsId" v:"required#论坛id不能为空"`
	common.Author
}
type BbsReply struct {
	Reply    *entity.BbsReply          `json:"reply"`
	UserInfo *system_entity.MemberUser `json:"memberUser"`
}
type BbsGetRes struct {
	Bbs *entity.Bbs `json:"bbs"`
	// BbsContent *entity.BbsContent        `json:"bbsContent"`
	BbsReply []BbsReply                `json:"bbsReply"`
	UserInfo *system_entity.MemberUser `json:"memberUser"`
}

// 修改论坛
type BbsEditReq struct {
	g.Meta        `path:"/bbs/edit" tags:"论坛" method:"post" summary:"修改论坛"`
	BbsId         int    `p:"bbsId" v:"required#论坛id不能为空"`
	Title         string `p:"title" v:"max-length:80#字数超过80"`           // 标题
	Abstract      string `p:"abstract" v:"max-length:400#字数超过400"`      // 摘要
	Content       string `p:"content"`                                  // 正文
	ReviewMessage string `p:"reviewMessage" v:"max-length:400#字数超过400"` // 审核留言
	common.Author
}
type BbsEditRes struct {
}

// 删除论坛
type BbsDelReq struct {
	g.Meta `path:"/bbs/del" tags:"论坛" method:"post" summary:"论坛删除"`
	BbsId  int `p:"bbsId" v:"required#论坛id不能为空"`
	common.Author
}
type BbsDelRes struct {
}

// 浏览量
type BbsBrowseReq struct {
	g.Meta `path:"/bbs/browse" tags:"论坛" method:"post" summary:"论坛浏览量增加接口"`
	BbsId  int `p:"bbsId" v:"required#论坛id不能为空"`
	common.Author
}
type BbsBrowseRes struct {
}

// 论坛点赞量
type BbsLikeReq struct {
	g.Meta `path:"/bbs/like" tags:"论坛" method:"post" summary:"论坛点赞接口"`
	BbsId  int `p:"bbsId" v:"required#论坛id不能为空"`
	Type   int `p:"type" v:"required#点赞1,取消点赞0不能为空"`
	common.Author
}
type BbsLikeRes struct {
}

// 评论
type BbsCommentReq struct {
	g.Meta  `path:"/bbs/comment" tags:"论坛" method:"post" summary:"评论"`
	BbsId   int    `p:"bbsId" v:"required#论坛id不能为空"`
	Content string `p:"content" v:"max-length:400#字数超过400"`
	common.Author
}
type BbsCommentRes struct {
}

// 评论点赞
type BbsCommentLikeReq struct {
	g.Meta    `path:"/bbs/comment/like" tags:"论坛" method:"post" summary:"评论点赞接口"`
	BbsId     int `p:"bbsId" v:"required#论坛id不能为空"`
	CommentId int `p:"commentId" v:"required#评论id不能为空"`
	Type      int `p:"type" v:"required#点赞1,取消点赞0不能为空"`
	common.Author
}
type BbsCommentLikeRes struct {
}

// 删除评论
type BbsCommentDelReq struct {
	g.Meta    `path:"/bbs/comment/del" tags:"论坛" method:"post" summary:"删除评论"`
	BbsId     int `p:"bbsId" v:"required#论坛id不能为空"`
	CommentId int `p:"commentId" v:"required#评论id不能为空"`
	common.Author
}
type BbsCommentDelRes struct {
}

// 回复
type BbsReplyReq struct {
	g.Meta  `path:"/bbs/reply" tags:"论坛" method:"post" summary:"回复"`
	BbsId   int    `p:"bbsId" v:"required#论坛id不能为空"`
	ReplyId int    `p:"replyId" v:"required#评论回复id不能为空"`
	Content string `p:"content" v:"max-length:400#字数超过400"`
	common.Author
}
type BbsReplyRes struct {
}

// 获取用户是否点赞
type BbsGetLikeReq struct {
	g.Meta `path:"/bbs/like/get" tags:"论坛" method:"get" summary:"获取用户是否点赞帖子/回复"`
	BbsId  *int `p:"bbsId" v:"required#论坛id不能为空"`
	UserId *int `p:"userId"`
	common.Author
}
type BbsGetLikeRes struct {
	ReplyLike []entity.BbsLike `json:"replyLike"`
	LikeBbs   bool             `json:"likeBbs"`
}

// 调研公告反馈

// 反馈
type FeedbackPublishReq struct {
	g.Meta             `path:"/feedback/add" tags:"公告管理" method:"post" summary:"添加反馈"`
	BulletinId         string `p:"bulletinId" v:"required#公告id不能为空"`
	Company            string `p:"company" v:"required#公司名称不能为空"`
	ContactPerson      string `p:"contactPerson" v:"required#联系人不能为空"`
	ContactInformation string `p:"contactInformation" v:"required#联系方式不能为空"`
	Remarks            string `p:"remarks"`
	Attachment         string `p:"attachment"`
	common.Author
}
type FeedbackPublishRes struct {
}
