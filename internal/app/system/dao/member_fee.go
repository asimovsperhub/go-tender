// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"tender/internal/app/system/dao/internal"
)

// internalMemberFeeDao is internal type for wrapping internal DAO implements.
type internalMemberFeeDao = *internal.MemberFeeDao

// memberFeeDao is the data access object for table member_fee.
// You can define custom methods on it to extend its functionality as you wish.
type memberFeeDao struct {
	internalMemberFeeDao
}

var (
	// MemberFee is globally public accessible object for table member_fee operations.
	MemberFee = memberFeeDao{
		internal.NewMemberFeeDao(),
	}
)

// Fill with you ideas below.
