package model

import "github.com/gogf/gf/v2/os/gtime"

type Bbs struct {
	Id             uint64      `json:"id"             description:""`
	Title          string      `json:"title"          description:"标题"`
	Abstract       string      `json:"abstract"       description:"摘要"`
	ReviewMessage  string      `json:"reviewMessage"  description:"审核留言"`
	ReviewStatus   uint        `json:"reviewStatus"   description:"审核状态;0:待审核,1:通过，2:未通过"`
	Views          int         `json:"views"          description:"浏览量"`
	ReplyCount     int         `json:"replyCount"     description:"回复量"`
	LikeCount      int         `json:"likeCount"      description:"点赞量"`
	CollectCount   int         `json:"collectCount"      description:"收藏量"`
	ReleaseCount   int         `json:"releaseCount"      description:"发布帖子量量"`
	UserId         int64       `json:"userId"         description:"用户id"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"发布时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
	Classification string      `json:"classification" description:"所属分类"`
	Status         uint        `json:"status"         description:"审核状态;0:软删除,1:正常，2:永久删除"`
	Rank           string      `json:"rank" description:"排序"`
	Content        string      `json:"content" description:"内容"`
}
