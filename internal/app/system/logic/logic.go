package logic

import (
	_ "tender/internal/app/system/logic/context"
	_ "tender/internal/app/system/logic/middleware"

	_ "tender/internal/app/system/logic/sys_personal"

	_ "tender/internal/app/system/logic/sysAuthRule"

	_ "tender/internal/app/system/logic/sysDept"

	_ "tender/internal/app/system/logic/sysRole"

	_ "tender/internal/app/system/logic/sysUser"

	_ "tender/internal/app/system/logic/token"

	_ "tender/internal/app/system/logic/sys_membermanager"

	_ "tender/internal/app/system/logic/sys_enterprise"

	// desk_individual
	_ "tender/internal/app/system/logic/desk_individual"

	_ "tender/internal/app/system/logic/sys_data"

	_ "tender/internal/app/system/logic/sys_knowledge_manger"

	_ "tender/internal/app/system/logic/sys_index"

	_ "tender/internal/app/system/logic/sysMsg"
	_ "tender/internal/app/system/logic/sys_finance"
	_ "tender/internal/app/system/logic/sys_forum_manager"
)
