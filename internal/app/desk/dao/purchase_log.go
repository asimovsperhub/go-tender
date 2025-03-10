// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"tender/internal/app/desk/dao/internal"
)

// internalPurchaseLogDao is internal type for wrapping internal DAO implements.
type internalPurchaseLogDao = *internal.PurchaseLogDao

// purchaseLogDao is the data access object for table purchase_log.
// You can define custom methods on it to extend its functionality as you wish.
type purchaseLogDao struct {
	internalPurchaseLogDao
}

var (
	// PurchaseLog is globally public accessible object for table purchase_log operations.
	PurchaseLog = purchaseLogDao{
		internal.NewPurchaseLogDao(),
	}
)

// Fill with you ideas below.
