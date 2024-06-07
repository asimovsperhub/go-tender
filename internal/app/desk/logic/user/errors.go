package user

import "github.com/gogf/gf/v2/errors/gerror"

var ErrUserNameOrMobileExists = gerror.New("手机号已存在")
