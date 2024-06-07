// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"tender/internal/app/desk/dao/internal"
)

// internalBbsDao is internal type for wrapping internal DAO implements.
type internalBbsDao = *internal.BbsDao

// bbsDao is the data access object for table bbs.
// You can define custom methods on it to extend its functionality as you wish.
type bbsDao struct {
	internalBbsDao
}

var (
	// Bbs is globally public accessible object for table bbs operations.
	Bbs = bbsDao{
		internal.NewBbsDao(),
	}
)

// Fill with you ideas below.
