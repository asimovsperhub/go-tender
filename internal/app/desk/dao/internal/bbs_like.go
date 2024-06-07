// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BbsLikeDao is the data access object for table bbs_like.
type BbsLikeDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns BbsLikeColumns // columns contains all the column names of Table for convenient usage.
}

// BbsLikeColumns defines and stores column names for table bbs_like.
type BbsLikeColumns struct {
	Id        string //
	BbsId     string // 论坛id
	ReplyId   string // 回复id
	UserId    string // 用户id
	CreatedAt string // 发布时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

//  bbsLikeColumns holds the columns for table bbs_like.
var bbsLikeColumns = BbsLikeColumns{
	Id:        "id",
	BbsId:     "bbs_id",
	ReplyId:   "reply_id",
	UserId:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewBbsLikeDao creates and returns a new DAO object for table data access.
func NewBbsLikeDao() *BbsLikeDao {
	return &BbsLikeDao{
		group:   "default",
		table:   "bbs_like",
		columns: bbsLikeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BbsLikeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BbsLikeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BbsLikeDao) Columns() BbsLikeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BbsLikeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BbsLikeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BbsLikeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
