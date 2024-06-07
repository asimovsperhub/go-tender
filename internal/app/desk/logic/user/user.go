package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"log"
	"tender/api/v1/desk"
	"tender/internal/app/desk/model"
	"tender/internal/app/desk/service"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/library/libUtils"
	"tender/library/liberr"
)

func init() {
	service.RegisterUser(New())
}

type sUser struct {
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) RegisterBind(ctx context.Context, pwd, mobile, nickname string, name string, openId string) (err error) {
	if openId == "" {
		if nickname == "" {
			nickname = mobile
		}
		_, err = s.UserNameOrMobileExists(ctx, mobile)
		if err != nil {
			return
		}
		UserSalt := grand.S(10)
		log.Println("add user passwd ", pwd)
		pwd = libUtils.EncryptPassword(pwd, UserSalt)
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err = g.Try(ctx, func(ctx context.Context) {
				_, e := dao.MemberUser.Ctx(ctx).TX(tx).InsertAndGetId(do.MemberUser{
					UserName:     name,
					Mobile:       mobile,
					UserNickname: nickname,
					UserPassword: pwd,
					UserSalt:     UserSalt,
					UserStatus:   1,
				})
				liberr.ErrIsNil(ctx, e, "注册用户失败")
			})
			return err
		})

	} else {
		wxUser := &entity.SysWxUser{}
		err = dao.SysWxUser.Ctx(ctx).Where(dao.SysWxUser.Columns().OpenId, openId).Scan(&wxUser)
		if err != nil {
			return
		}
		if wxUser.MainId != 0 {
			err = errors.New("该微信已绑定其他账号")
			return
		}
		var uid uint64
		uid, err = s.UserNameOrMobileExists(ctx, mobile)
		if err != nil && !errors.Is(err, ErrUserNameOrMobileExists) {
			return
		}
		if !errors.Is(err, ErrUserNameOrMobileExists) {
			UserSalt := grand.S(10)
			log.Println("add user passwd ", pwd)
			pwd = libUtils.EncryptPassword(pwd, UserSalt)
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = g.Try(ctx, func(ctx context.Context) {
					uid, e := dao.MemberUser.Ctx(ctx).TX(tx).InsertAndGetId(do.MemberUser{
						UserName:     name,
						Mobile:       mobile,
						UserNickname: nickname,
						UserPassword: pwd,
						UserSalt:     UserSalt,
					})
					_, err = dao.SysWxUser.Ctx(ctx).Where(dao.SysWxUser.Columns().OpenId, openId).TX(tx).Data(dao.SysWxUser.Columns().MainId, uid).Update()
					liberr.ErrIsNil(ctx, e, "注册用户失败")
				})
				return err
			})
		} else {
			UserSalt := grand.S(10)
			log.Println("add user passwd ", pwd)
			pwd = libUtils.EncryptPassword(pwd, UserSalt)
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = g.Try(ctx, func(ctx context.Context) {
					_, err = dao.SysWxUser.Ctx(ctx).Where(dao.SysWxUser.Columns().OpenId, openId).TX(tx).Data(dao.SysWxUser.Columns().MainId, uid).Update()
					liberr.ErrIsNil(ctx, err, "注册用户失败")
				})
				return err
			})
		}

	}
	return
}

func (s *sUser) UserNameOrMobileExists(ctx context.Context, mobile string, id ...int64) (uint64, error) {
	user := (*entity.MemberUser)(nil)
	err := g.Try(ctx, func(ctx context.Context) {
		m := dao.MemberUser.Ctx(ctx)
		if len(id) > 0 {
			m = m.Where(dao.MemberUser.Columns().Id+" != ", id)
		}
		m = m.Where(fmt.Sprintf("%s='%s'",
			dao.MemberUser.Columns().Mobile,
			mobile))
		err := m.Limit(1).Scan(&user)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
		if user == nil {
			return
		}
		//if user.UserName == userName {
		//	liberr.ErrIsNil(ctx, gerror.New("用户名已存在"))
		//}
		if user.Mobile == mobile {
			liberr.ErrIsNil(ctx, ErrUserNameOrMobileExists)
		}
	})
	if user == nil {
		return 0, err
	}
	return user.Id, err
}
func (s *sUser) GetUserByUsernamePassword(ctx context.Context, req *desk.UserLoginReq) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err = s.GetUserByUsername(ctx, req.Mobile)
		liberr.ErrIsNil(ctx, err)
		liberr.ValueIsNil(user, "账号或密码错误")
		if libUtils.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
			liberr.ErrIsNil(ctx, gerror.New("账号密码错误"))
			err = gerror.New("账号密码错误")
		}
	})
	return
}

// GetUserByUsername 通过用户名获取用户信息
func (s *sUser) GetUserByUsername(ctx context.Context, Mobile string) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user = &model.LoginUserRes{}
		err = dao.MemberUser.Ctx(ctx).Fields(user).Where(dao.MemberUser.Columns().Mobile, Mobile).Scan(user)
		user.Enterprise, err = s.GetEnterpriseByUserId(ctx, int(user.Id))
	})
	return
}

// 获取企业信息
func (s *sUser) GetEnterpriseByUserId(ctx context.Context, UserId int) (enterprise *entity.SysEnterprise, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enterprise = &entity.SysEnterprise{}
		err = dao.SysEnterprise.Ctx(ctx).Fields(enterprise).Where(dao.SysEnterprise.Columns().UserId, UserId).Scan(enterprise)
		if enterprise.Id == 0 {
			enterprise = nil
		}
	})
	if err != nil {
		return
	}

	return
}

// 修改用户信息
