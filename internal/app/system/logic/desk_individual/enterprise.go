package desk_individual

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"log"
	"tender/api/v1/desk"
	DeskService "tender/internal/app/desk/service"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
	"tender/library/liberr"
)

func init() {
	service.RegisterEnterprise(New())
}

func New() *individualEnterprise {
	return &individualEnterprise{}
}

type individualEnterprise struct {
}

// 企业入驻
func (s *individualEnterprise) Add(ctx context.Context, req *desk.EnterpriseAddReq) (res *desk.EnterpriseAddRes, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err := dao.SysEnterprise.Ctx(ctx).TX(tx).Insert(do.SysEnterprise{
				Name:               req.Name,
				Nickname:           req.NickName,
				Location:           req.Location,
				Contact:            req.Contact,
				Content:            req.Content,
				Industry:           req.Industry,
				Introduction:       req.Introduction,
				EstablishmentAt:    req.EstablishmentAt,
				Icon:               req.Icon,
				License:            req.License,
				Certificate:        req.Certificate,
				CertificateMessage: req.CertificateMessage,
				LicenseMessage:     req.LicenseMessage,
				UserId:             DeskService.Context().GetUserId(ctx),
			})
			liberr.ErrIsNil(ctx, err, "添加企业入驻失败/企业name已存在")
		})
		return err
	})
	return
}

// 企业信息 Edit
func (s *individualEnterprise) Edit(ctx context.Context, req *desk.EnterpriseEditReq) (res *desk.EnterpriseEditRes, err error) {
	var enterprise *entity.SysEnterprise
	data := do.SysEnterprise{}
	user_id := DeskService.Context().GetUserId(ctx)
	enterprise, err = DeskService.DeskUser().GetEnterpriseByUserId(ctx, int(user_id))
	if err != nil {
		return
	}
	log.Println("enterprise信息----------------->", user_id, enterprise)
	if enterprise == nil {
		err = errors.New("当前用户没有企业")
		return
	} else {
		if enterprise.Name != req.OldName {
			err = errors.New("当前用户不属于该企业")
			return
		}

	}
	if req.Icon != "" {
		data.Icon = req.Icon
	}
	if req.License != "" {
		data.License = req.License
		data.LicenseStatus = 0
	}
	if req.Certificate != "" {
		data.Certificate = req.Certificate
		data.CertificateStatus = 0
	}
	if req.Name != "" {
		data.Name = req.Name
	}
	if req.NickName != "" {
		data.Nickname = req.NickName
	}
	if req.Location != "" {
		data.Location = req.Location
	}
	if req.Contact != "" {
		data.Contact = req.Contact
	}
	if req.Content != "" {
		data.Content = req.Content
	}
	if req.Industry != "" {
		data.Industry = req.Industry
	}
	if req.Introduction != "" {
		data.Introduction = req.Introduction
	}
	if req.EstablishmentAt != "" {
		data.EstablishmentAt = req.EstablishmentAt
	}
	if req.CertificateMessage != "" {
		data.CertificateMessage = req.CertificateMessage
	}
	if req.LicenseMessage != "" {
		data.LicenseMessage = req.LicenseMessage
	}
	data.UpdatedAt = gtime.Now()
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err := dao.SysEnterprise.Ctx(ctx).TX(tx).Where("name = ?", req.OldName).Update(data)
			liberr.ErrIsNil(ctx, err, "修改企业信息失败/企业名称不存在")
		})
		return err
	})
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysEnterprise.Ctx(ctx).Fields(enterprise).Where(dao.SysEnterprise.Columns().Name, req.Name).Scan(enterprise)
		if enterprise.Id == 0 {
			enterprise = nil
		}
	})
	if err != nil {
		return
	}
	res = &desk.EnterpriseEditRes{
		Result:         "修改企业信息成功",
		EnterpriseInfo: enterprise,
	}
	return
}

//Get
func (s *individualEnterprise) Get(ctx context.Context, req *desk.EnterpriseGetReq) (res *desk.EnterpriseGetRes, err error) {
	var enterprise *entity.SysEnterprise
	//user_id := DeskService.Context().GetUserId(ctx)
	//enterprise, err = DeskService.DeskUser().GetEnterpriseByUserId(ctx, int(user_id))
	//if err != nil {
	//	return
	//}
	//log.Println("enterprise信息----------------->", user_id, enterprise)
	//if enterprise == nil {
	//	err = errors.New("当前用户没有企业")
	//	return
	//}
	data := do.SysEnterprise{}
	if req.UserId != nil {
		data.UserId = req.UserId
	}
	if req.Id != nil {
		data.Id = req.Id
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysEnterprise.Ctx(ctx).Fields(enterprise).Where(data).Limit(1).Scan(&enterprise)
	})
	if err != nil {
		return
	}
	res = &desk.EnterpriseGetRes{
		EnterpriseInfo: enterprise,
	}
	return
}
