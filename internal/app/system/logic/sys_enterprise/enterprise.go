package sys_enterprise

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"log"
	"tender/api/v1/system"
	"tender/internal/app/system/consts"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
	"tender/internal/packed/websocket"
	"tender/library/liberr"
)

func init() {
	service.RegisterEnterpriseManger(New())
}

func New() *sSysEnterpriseManger {
	return &sSysEnterpriseManger{}
}

type sSysEnterpriseManger struct {
}

func (s *sSysEnterpriseManger) List(ctx context.Context, req *system.EnterpriseSearchReq) (res *system.EnterpriseSearchRes, err error) {
	res = new(system.EnterpriseSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysEnterprise.Ctx(ctx)
	order := "created_at desc"
	if req.Name != "" {
		m = m.Where("name like ?", "%"+req.Name+"%")
	}
	if req.City != "" {
		m = m.Where("location like ?", "%"+req.City+"%")
	}
	if req.Industry != "" {
		m = m.Where("industry like ?", "%"+req.Industry+"%")
	}
	if req.Contact != "" {
		m = m.Where("contact like ?", "%"+req.Contact+"%")
	}
	if len(req.DateRange) > 0 {
		m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0]+" 00:00:00", req.DateRange[1]+" 23:59:59")
	} else {
		if req.Start != "" {
			m = m.Where("created_at >=?", req.Start+" 00:00:00")
		}
		if req.End != "" {
			m = m.Where("created_at <=?", req.End+" 23:59:59")
		}
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取企业列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.EnterpriseList)
		log.Println("获取会员列表-------->", res.EnterpriseList)
		liberr.ErrIsNil(ctx, err, "获取企业列表数据失败")
	})
	return
}

func (s *sSysEnterpriseManger) LicenseList(ctx context.Context, req *system.EnterpriselLcenseSearchReq) (res *system.EnterpriseSearchRes, err error) {
	res = new(system.EnterpriseSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysEnterprise.Ctx(ctx)
	order := "created_at desc"
	if req.KeyWords != "" {
		m = m.Where("name like ?", "%"+req.KeyWords+"%")
		m = m.Where("license_status <> 1")
	} else {
		m = m.Where("license_status <> 1")
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取企业列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.EnterpriseList)
		log.Println("获取营业执照审查企业列表-------->", res.EnterpriseList)
		liberr.ErrIsNil(ctx, err, "获取企业列表数据失败")
	})
	return
}

func (s *sSysEnterpriseManger) CertificateList(ctx context.Context, req *system.EnterpriseCertificateSearchReq) (res *system.EnterpriseSearchRes, err error) {
	res = new(system.EnterpriseSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysEnterprise.Ctx(ctx)
	order := "created_at desc"
	if req.KeyWords != "" {
		m = m.Where("name like ?", "%"+req.KeyWords+"%")
		m = m.Where("certificate_status <> 1")
	} else {
		m = m.Where("certificate_status <> 1")
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取企业列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.EnterpriseList)
		log.Println("获取证书审查企业-------->", res.EnterpriseList)
		liberr.ErrIsNil(ctx, err, "获取企业列表数据失败")
	})
	return
}

func (s *sSysEnterpriseManger) CertificateReview(ctx context.Context, req *system.EnterpriseCertificateReviewReq) (res *system.EnterpriseCertificateReviewRes, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysEnterprise.Ctx(ctx).TX(tx).WherePri(req.Id).Update(do.SysEnterprise{
				CertificateStatus:    req.Status,
				OpcertificateMessage: req.OpCertificateMessage,
				UpdatedAt:            gtime.Now(),
				// 操作管理员id
				OperationId: service.Context().GetUserId(ctx),
			})
			liberr.ErrIsNil(ctx, err, "修改企业证明证书状态失败")
		})
		return err
	})
	enterprise := (*entity.SysEnterprise)(nil)
	m := dao.SysEnterprise.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&enterprise)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	switch req.Status {
	case 1:
		req.OpCertificateMessage = "证明书审核通过: " + req.OpCertificateMessage
	case 2:
		req.OpCertificateMessage = "证明书审核未通过: " + req.OpCertificateMessage
	}
	if req.Status != 0 {
		msgId := uuid.New().String()
		_, err = dao.SysWsMsg.Ctx(ctx).Insert(do.SysWsMsg{
			MessageId: msgId,
			UserId:    enterprise.UserId,
			Content:   req.OpCertificateMessage,
			IsRead:    0,
			IsDel:     0,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		})
		if err != nil {
			return nil, fmt.Errorf("插入消息失败: %w", err)
		}

		websocket.SendToUser(uint64(enterprise.UserId), &websocket.WResponse{
			Event: "context",
			Data: map[string]string{
				"messageId": msgId,
				"content":   req.OpCertificateMessage,
			},
		})
	}
	return
}

func (s *sSysEnterpriseManger) LicenseReview(ctx context.Context, req *system.EnterpriselLcenseReviewReq) (res *system.EnterpriselLcenseReviewRes, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysEnterprise.Ctx(ctx).TX(tx).WherePri(req.Id).Update(do.SysEnterprise{
				LicenseStatus:    req.Status,
				OplicenseMessage: req.OpLicenseMessage,
				UpdatedAt:        gtime.Now(),
				// 操作管理员id
				OperationId: service.Context().GetUserId(ctx),
			})
			liberr.ErrIsNil(ctx, err, "修改企业营业执照状态失败")
		})
		return err
	})
	enterprise := (*entity.SysEnterprise)(nil)
	m := dao.SysEnterprise.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&enterprise)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	switch req.Status {
	case 1:
		req.OpLicenseMessage = "营业执照审核通过: " + req.OpLicenseMessage
	case 2:
		req.OpLicenseMessage = "营业执照审核未通过: " + req.OpLicenseMessage
	}
	if req.Status != 0 {
		msgId := uuid.New().String()
		_, err = dao.SysWsMsg.Ctx(ctx).Insert(do.SysWsMsg{
			MessageId: msgId,
			UserId:    enterprise.UserId,
			Content:   req.OpLicenseMessage,
			IsRead:    0,
			IsDel:     0,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		})
		if err != nil {
			return nil, fmt.Errorf("插入消息失败: %w", err)
		}

		websocket.SendToUser(uint64(enterprise.UserId), &websocket.WResponse{
			Event: "context",
			Data: map[string]string{
				"messageId": msgId,
				"content":   req.OpLicenseMessage,
			},
		})
	}
	return
}
