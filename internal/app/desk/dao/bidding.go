// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"tender/internal/app/desk/dao/internal"
)

// internalBiddingDao is internal type for wrapping internal DAO implements.
type internalBiddingDao = *internal.BiddingDao

// biddingDao is the data access object for table bidding.
// You can define custom methods on it to extend its functionality as you wish.
type biddingDao struct {
	internalBiddingDao
}

var (
	// Bidding is globally public accessible object for table bidding operations.
	Bidding = biddingDao{
		internal.NewBiddingDao(),
	}
)

// Fill with you ideas below.
