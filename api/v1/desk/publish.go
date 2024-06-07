package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/desk/model/entity"
	system_entity "tender/internal/app/system/model/entity"
)

//	Id                      uint64      `json:"id"                      description:""`
//	Title                   string      `json:"title"                   description:"标题"`
//	KnowledgeType           string      `json:"knowledgeType"           description:"知识类型"`
//	Authority               uint        `json:"authority"               description:"阅读下载权限，0:所有,1:会员"`
//	PrimaryClassification   string      `json:"primaryClassification"   description:"一级分类"`
//	SecondaryClassification string      `json:"secondaryClassification" description:"二级分类"`
//	IntegralSetting         int         `json:"integralSetting"         description:"积分设置"`
//	Content                 string      `json:"content"                 description:"内容"`
//	ReviewMessage           string      `json:"reviewMessage"           description:"审核留言"`
//	ReviewStatus            uint        `json:"reviewStatus"            description:"审核状态;0:待审核,1:通过，2:未通过"`
//	UserId                  int64       `json:"userId"                  description:"发布作者"`
//	CreatedAt               *gtime.Time `json:"createdAt"               description:"发布时间"`
//	UpdatedAt               *gtime.Time `json:"updatedAt"               description:"更新时间"`
//	DeletedAt               *gtime.Time `json:"deletedAt"               description:"删除时间"`
//	CoverUrl:                "cover_url",
//	DetailsUrl:              "details_url",
//	VideoUrl:                "video_url",
//	VideoIntroduction:       "video_introduction",

// 发布知识
type KnowledgeReq struct {
	g.Meta                  `path:"/publish/knowledge" tags:"前台个人中心" method:"post" summary:"发布知识"`
	Title                   string `p:"title" v:"max-length:80#字数超过80"`               //标题
	KnowledgeType           string `p:"knowledgeType"`                                //知识类型
	Authority               int    `p:"authority"`                                    //阅读下载权限
	PrimaryClassification   string `p:"primaryClassification"`                        //一级分类
	SecondaryClassification string `p:"secondaryClassification"`                      // 二级分类
	IntegralSetting         int    `p:"integralSetting"`                              // 下载积分设置
	Content                 string `p:"content"`                                      // 正文
	ReviewMessage           string `p:"reviewMessage" v:"max-length:400#字数超过400"`     // 审核留言
	Type                    int    `p:"type"`                                         // 知识类型
	CoverUrl                string `p:"coverUrl" `                                    //封面
	DetailsUrl              string `p:"detailsUrl" `                                  //详情图
	VideoUrl                string `p:"videoUrl" `                                    //视频
	VideoIntroduction       string `p:"videoIntroduction" v:"max-length:400#字数超过400"` //视频简介
	AttachmentUrl           string `p:"attachmentUrl" `                               //附件
	Abstract                string `p:"abstract" v:"max-length:400#字数超过400"`          //摘要
	common.Author
}
type KnowledgeRes struct {
}

// 我的知识列表
type KnowledgeGetListReq struct {
	g.Meta `path:"/publish/list" tags:"前台个人中心" method:"get" summary:"我的知识列表"`
	UserId int `p:"userId" v:"required#用户id不能为空"`
	common.PageReq
	common.Author
}
type KnowledgeGetListRes struct {
	MemberKnowledgeList []*entity.MemberKnowledge `json:"MemberKnowledgeList"`
	common.ListRes
}

// 单个知识
type KnowledgeGetReq struct {
	g.Meta      `path:"/publish/get" tags:"前台个人中心" method:"get" summary:"获取单个知识"`
	KnowledgeId int `p:"knowledgeId" v:"required#知识id不能为空"`
	common.Author
}

type MemberKnowledge struct {
	MemberKnowledge *entity.MemberKnowledge   `json:"memberKnowledge"`
	UserInfo        *system_entity.MemberUser `json:"memberUser"`
}

type KnowledgeGetRes struct {
	MemberKnowledge MemberKnowledge `json:"memberKnowledge"`
}

// 修改知识
type KnowledgeEditReq struct {
	g.Meta                  `path:"/publish/edit" tags:"前台个人中心" method:"post" summary:"修改知识"`
	KnowledgeId             int    `p:"knowledgeId" v:"required#知识id不能为空"`
	Title                   string `p:"title" `                                       //标题
	KnowledgeType           string `p:"knowledgeType"`                                //知识类型
	Authority               *int   `p:"authority"`                                    //阅读下载权限
	PrimaryClassification   string `p:"primaryClassification"`                        //一级分类
	SecondaryClassification string `p:"secondaryClassification"`                      // 二级分类
	IntegralSetting         *int   `p:"integralSetting"`                              // 下载积分设置
	Content                 string `p:"content"`                                      // 正文
	ReviewMessage           string `p:"reviewMessage" v:"max-length:400#字数超过400"`     // 审核留言
	Type                    *int   `p:"type"`                                         // 知识类型
	CoverUrl                string `p:"coverUrl" `                                    //封面
	DetailsUrl              string `p:"detailsUrl" `                                  //详情图
	VideoUrl                string `p:"videoUrl" `                                    //视频
	VideoIntroduction       string `p:"videoIntroduction" v:"max-length:400#字数超过400"` //视频简介
	AttachmentUrl           string `p:"attachmentUrl" `                               //附件
	Abstract                string `p:"abstract" v:"max-length:400#字数超过400"`          //摘要
	common.Author
}
type KnowledgeEditRes struct {
}

// 删除知识
type KnowledgeDelReq struct {
	g.Meta      `path:"/publish/del" tags:"前台个人中心" method:"post" summary:"知识删除"`
	KnowledgeId int `p:"knowledgeId" v:"required#知识id不能为空"`
	common.Author
}
type KnowledgeDelRes struct {
	MemberKnowledge *entity.MemberKnowledge `json:"memberKnowledge"`
}

// 获取非会员是否单次购买

type KnowledgeBuyReq struct {
	g.Meta      `path:"/publish/isbuy" tags:"前台个人中心" method:"get" summary:"知识是否购买"`
	UserId      int `p:"userId" v:"required#用户id不能为空"`
	KnowledgeId int `p:"knowledgeId" v:"required#知识id不能为空"`
	common.Author
}
type KnowledgeBuyRes struct {
	IsBuy bool `json:"isBuy"`
}
