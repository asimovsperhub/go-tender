// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberCollectDao is the data access object for table member_collect.
type MemberCollectDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns MemberCollectColumns // columns contains all the column names of Table for convenient usage.
}

// MemberCollectColumns defines and stores column names for table member_collect.
type MemberCollectColumns struct {
	Id        string //
	Title     string // 标题
	Type      string // 类型
	Location  string // 所在地
	Industry  string // 行业
	UserId    string // 用户id
	ArticleId string // 文章id
	CreatedAt string // 收藏时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Url       string // 收藏url
}

//  memberCollectColumns holds the columns for table member_collect.
var memberCollectColumns = MemberCollectColumns{
	Id:        "id",
	Title:     "title",
	Type:      "type",
	Location:  "location",
	Industry:  "industry",
	UserId:    "user_id",
	ArticleId: "article_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Url:       "url",
}

// NewMemberCollectDao creates and returns a new DAO object for table data access.
func NewMemberCollectDao() *MemberCollectDao {
	return &MemberCollectDao{
		group:   "default",
		table:   "member_collect",
		columns: memberCollectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MemberCollectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MemberCollectDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MemberCollectDao) Columns() MemberCollectColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MemberCollectDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MemberCollectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MemberCollectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
