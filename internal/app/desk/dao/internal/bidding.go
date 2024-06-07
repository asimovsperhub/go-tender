// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BiddingDao is the data access object for table bidding.
type BiddingDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns BiddingColumns // columns contains all the column names of Table for convenient usage.
}

// BiddingColumns defines and stores column names for table bidding.
type BiddingColumns struct {
	Id                     string //
	BulletinType           string // 公告类型
	City                   string // 所属城市
	IndustryClassification string // 行业分类
	ReleaseTime            string // 发布时间
	BidopeningTime         string // 开标时间
	Title                  string // 标题
	Abstract               string // 摘要
	Enterprise             string // 企业信息
	AnnouncementContent    string // 公告内容
	Link                   string // 源链接
	Attachment             string // 招标附件
	Amount                 string // 金额
	ContactPerson          string // 联系人
	ContactInformation     string // 联系方式
	CreatedAt              string // 爬取日期
	BidAmount              string // 中标金额
	TenderName             string // 招标企业
	WinName                string // 中标企业
}

//  biddingColumns holds the columns for table bidding.
var biddingColumns = BiddingColumns{
	Id:                     "id",
	BulletinType:           "bulletin_type",
	City:                   "city",
	IndustryClassification: "industry_classification",
	ReleaseTime:            "release_time",
	BidopeningTime:         "bidopening_time",
	Title:                  "title",
	Abstract:               "abstract",
	Enterprise:             "enterprise",
	AnnouncementContent:    "announcement_content",
	Link:                   "link",
	Attachment:             "attachment",
	Amount:                 "amount",
	ContactPerson:          "contact_person",
	ContactInformation:     "contact_information",
	CreatedAt:              "created_at",
	BidAmount:              "bid_amount",
	TenderName:             "tender_name",
	WinName:                "win_name",
}

// NewBiddingDao creates and returns a new DAO object for table data access.
func NewBiddingDao() *BiddingDao {
	return &BiddingDao{
		group:   "default",
		table:   "bidding",
		columns: biddingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BiddingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BiddingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BiddingDao) Columns() BiddingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BiddingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BiddingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BiddingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
