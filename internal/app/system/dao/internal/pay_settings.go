// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PaySettingsDao is the data access object for table pay_settings.
type PaySettingsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns PaySettingsColumns // columns contains all the column names of Table for convenient usage.
}

// PaySettingsColumns defines and stores column names for table pay_settings.
type PaySettingsColumns struct {
	Id                     string //
	WeixinMchid            string //
	WeixinAppid            string //
	WeixinApikey           string //
	WeixinSerialno         string //
	WeixinPrivatekey       string //
	AlipayAppid            string //
	AlipayPrivatekey       string //
	AlipayAppCertPublicKey string //
	AlipayRootCert         string //
	AlipayPublicCert       string //
}

//  paySettingsColumns holds the columns for table pay_settings.
var paySettingsColumns = PaySettingsColumns{
	Id:                     "id",
	WeixinMchid:            "weixin_mchid",
	WeixinAppid:            "weixin_appid",
	WeixinApikey:           "weixin_apikey",
	WeixinSerialno:         "weixin_serialno",
	WeixinPrivatekey:       "weixin_privatekey",
	AlipayAppid:            "alipay_appid",
	AlipayPrivatekey:       "alipay_privatekey",
	AlipayAppCertPublicKey: "alipay_appCertPublicKey",
	AlipayRootCert:         "alipayRootCert",
	AlipayPublicCert:       "alipay_publicCert",
}

// NewPaySettingsDao creates and returns a new DAO object for table data access.
func NewPaySettingsDao() *PaySettingsDao {
	return &PaySettingsDao{
		group:   "default",
		table:   "pay_settings",
		columns: paySettingsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PaySettingsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PaySettingsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PaySettingsDao) Columns() PaySettingsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PaySettingsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PaySettingsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PaySettingsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
