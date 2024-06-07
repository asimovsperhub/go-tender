// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"tender/internal/app/system/dao/internal"
)

// internalSysEnterpriseDao is internal type for wrapping internal DAO implements.
type internalSysEnterpriseDao = *internal.SysEnterpriseDao

// sysEnterpriseDao is the data access object for table sys_enterprise.
// You can define custom methods on it to extend its functionality as you wish.
type sysEnterpriseDao struct {
	internalSysEnterpriseDao
}

var (
	// SysEnterprise is globally public accessible object for table sys_enterprise operations.
	SysEnterprise = sysEnterpriseDao{
		internal.NewSysEnterpriseDao(),
	}
)

// Fill with you ideas below.
