// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BidDao is the data access object for table bid.
type BidDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns BidColumns // columns contains all the column names of Table for convenient usage.
}

// BidColumns defines and stores column names for table bid.
type BidColumns struct {
	Id                     string //
	BulletinType           string // 公告类型
	OriginalType           string // 原公告类型
	Province               string // 所属省份
	City                   string // 所属城市
	IndustryClassification string // 行业分类
	OriginalClassification string // 原行业分类
	ReleaseTime            string // 发布时间
	BidopeningTime         string // 开标时间
	Title                  string // 标题
	Abstract               string // 摘要
	AnnouncementContent    string // 公告内容
	OriginalContent        string // 未清洗的公告内容
	Link                   string // 源链接
	OriginalLink           string // 源源链接
	Attachment             string // 招标附件
	Amount                 string // 金额
	ContactPerson          string // 联系人
	ContactInformation     string // 联系方式
	BidAmount              string // 中标金额
	TenderName             string // 招标企业
	WinName                string // 中标企业
	Site                   string // 站点
	CreatedAt              string // 爬取日期
	Height                 string // 1高2中3低
	Rank                   string // 排序
	Area                   string // 地区排序
	Display                string // 是否展示,0不展示1展示
	Crawler                string // 是否爬取,0 非爬取 1 爬取
	ReleaseDay             string // 发布时间day排序用的
}

//  bidColumns holds the columns for table bid.
var bidColumns = BidColumns{
	Id:                     "id",
	BulletinType:           "bulletin_type",
	OriginalType:           "original_type",
	Province:               "province",
	City:                   "city",
	IndustryClassification: "industry_classification",
	OriginalClassification: "original_classification",
	ReleaseTime:            "release_time",
	BidopeningTime:         "bidopening_time",
	Title:                  "title",
	Abstract:               "abstract",
	AnnouncementContent:    "announcement_content",
	OriginalContent:        "original_content",
	Link:                   "link",
	OriginalLink:           "original_link",
	Attachment:             "attachment",
	Amount:                 "amount",
	ContactPerson:          "contact_person",
	ContactInformation:     "contact_information",
	BidAmount:              "bid_amount",
	TenderName:             "tender_name",
	WinName:                "win_name",
	Site:                   "site",
	CreatedAt:              "created_at",
	Height:                 "height",
	Rank:                   "rank",
	Area:                   "area",
	Display:                "display",
	Crawler:                "crawler",
	ReleaseDay:             "release_day",
}

// NewBidDao creates and returns a new DAO object for table data access.
func NewBidDao() *BidDao {
	return &BidDao{
		group:   "default",
		table:   "bid",
		columns: bidColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BidDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BidDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BidDao) Columns() BidColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BidDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BidDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BidDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
