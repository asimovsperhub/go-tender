// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BbsDao is the data access object for table bbs.
type BbsDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns BbsColumns // columns contains all the column names of Table for convenient usage.
}

// BbsColumns defines and stores column names for table bbs.
type BbsColumns struct {
	Id             string //
	Title          string // 标题
	Abstract       string // 摘要
	ReviewMessage  string // 审核留言
	ReviewStatus   string // 审核状态;0:待审核,1:通过，2:未通过
	Views          string // 浏览量
	ReplyCount     string // 回复量
	LikeCount      string // 点赞量
	UserId         string // 用户id
	CreatedAt      string // 发布时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
	Classification string // 所属分类
	Status         string // 审核状态;0:软删除,1:正常，2:永久删除
	Rank           string // 排名
	Content        string // 内容
}

//  bbsColumns holds the columns for table bbs.
var bbsColumns = BbsColumns{
	Id:             "id",
	Title:          "title",
	Abstract:       "abstract",
	ReviewMessage:  "review_message",
	ReviewStatus:   "review_status",
	Views:          "views",
	ReplyCount:     "reply_count",
	LikeCount:      "like_count",
	UserId:         "user_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	Classification: "classification",
	Status:         "status",
	Rank:           "rank",
	Content:        "content",
}

// NewBbsDao creates and returns a new DAO object for table data access.
func NewBbsDao() *BbsDao {
	return &BbsDao{
		group:   "default",
		table:   "bbs",
		columns: bbsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BbsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BbsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BbsDao) Columns() BbsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BbsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BbsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BbsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
