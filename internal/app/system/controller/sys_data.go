package controller

import (
	"context"
	"tender/api/v1/system"
	"tender/internal/app/system/service"
)

var Data = sysDataController{}

type sysDataController struct {
	BaseController
}

// List 招标数据
func (c *sysDataController) List(ctx context.Context, req *system.DataSearchReq) (res *system.DataSearchRes, err error) {
	res, err = service.SysDataManger().List(ctx, req)
	return
}

// 单个招标数据
func (c *sysDataController) TenderGet(ctx context.Context, req *system.TenderGetReq) (res *system.TenderGetRes, err error) {
	res, err = service.SysDataManger().TenderGet(ctx, req)
	return
}

//TenderEditReq
func (c *sysDataController) TenderEdit(ctx context.Context, req *system.TenderEditReq) (res *system.TenderEditRes, err error) {
	res, err = service.SysDataManger().TenderEdit(ctx, req)
	return
}

//TenderTypeGetReq
func (c *sysDataController) TenderTypeGet(ctx context.Context, req *system.TenderTypeGetReq) (res *system.TenderTypeGetRes, err error) {
	res, err = service.SysDataManger().TenderTypeGet(ctx, req)
	return
}

//TenderAddSortReq

func (c *sysDataController) TenderAddSort(ctx context.Context, req *system.TenderAddSortReq) (res *system.TenderAddSortRes, err error) {
	res, err = service.SysDataManger().TenderAddSort(ctx, req)
	return
}

//TenderDelSortReq
func (c *sysDataController) TenderDelSort(ctx context.Context, req *system.TenderDelSortReq) (res *system.TenderDelSortRes, err error) {
	res, err = service.SysDataManger().TenderDelSort(ctx, req)
	return
}

//TenderDragReq
func (c *sysDataController) TenderDrag(ctx context.Context, req *system.TenderDragReq) (res *system.TenderDragRes, err error) {
	res, err = service.SysDataManger().TenderDrag(ctx, req)
	return
}

// TenderDelReq

func (c *sysDataController) TenderDel(ctx context.Context, req *system.TenderDelReq) (res *system.TenderDelRes, err error) {
	res, err = service.SysDataManger().TenderDel(ctx, req)
	return
}
func (c *sysDataController) LawList(ctx context.Context, req *system.LawSearchReq) (res *system.LawSearchRes, err error) {
	res, err = service.SysDataManger().LawList(ctx, req)
	return
}

// AllLawList
func (c *sysDataController) AllLawList(ctx context.Context, req *system.AllLawSearchReq) (res *system.AllLawSearchRes, err error) {
	res, err = service.SysDataManger().AllLawList(ctx, req)
	return
}

//AddLaw
func (c *sysDataController) AddLaw(ctx context.Context, req *system.LawAddReq) (res *system.LawAddRes, err error) {
	res, err = service.SysDataManger().AddLaw(ctx, req)
	return
}

//DelLaw
func (c *sysDataController) DelLaw(ctx context.Context, req *system.LawDelReq) (res *system.LawDelRes, err error) {
	res, err = service.SysDataManger().DelLaw(ctx, req)
	return
}
func (c *sysDataController) Setting(ctx context.Context, req *system.SettingReq) (res *system.SettingRes, err error) {
	res, err = service.SysDataManger().Setting(ctx, req)
	return
}

func (c *sysDataController) GetSetting(ctx context.Context, req *system.SettingGetReq) (res *system.SettingGetRes, err error) {
	res, err = service.SysDataManger().GetSetting(ctx, req)
	return
}

//EnterpriseListReq
func (c *sysDataController) EnterpriseList(ctx context.Context, req *system.EnterpriseListReq) (res *system.EnterpriseListRes, err error) {
	res, err = service.SysDataManger().EnterpriseList(ctx, req)
	return
}

//EnterpriseGetReq
func (c *sysDataController) EnterpriseGet(ctx context.Context, req *system.EnterpriseGetReq) (res *system.EnterpriseGetRes, err error) {
	res, err = service.SysDataManger().EnterpriseGet(ctx, req)
	return
}

//EnterpriseDelReq
func (c *sysDataController) EnterpriseDel(ctx context.Context, req *system.EnterpriseDelReq) (res *system.EnterpriseDelRes, err error) {
	res, err = service.SysDataManger().EnterpriseDel(ctx, req)
	return
}

//KnowledgeListReq
func (c *sysDataController) KnowledgeList(ctx context.Context, req *system.KnowledgeListReq) (res *system.KnowledgeListRes, err error) {
	res, err = service.SysDataManger().KnowledgeList(ctx, req)
	return
}

//KnowledgeGetReq
func (c *sysDataController) KnowledgeGet(ctx context.Context, req *system.KnowledgeGetReq) (res *system.KnowledgeGetRes, err error) {
	res, err = service.SysDataManger().KnowledgeGet(ctx, req)
	return
}

//KnowledgeDisplayReq
func (c *sysDataController) KnowledgeDisplay(ctx context.Context, req *system.KnowledgeDisplayReq) (res *system.KnowledgeDisplayRes, err error) {
	res, err = service.SysDataManger().KnowledgeDisplay(ctx, req)
	return
}

//KnowledgeTypeGetReq
func (c *sysDataController) KnowledgeTypeGet(ctx context.Context, req *system.KnowledgeTypeGetReq) (res *system.KnowledgeTypeGetRes, err error) {
	res, err = service.SysDataManger().KnowledgeTypeGet(ctx, req)
	return
}

//InformationListReq
func (c *sysDataController) InformationList(ctx context.Context, req *system.InformationListReq) (res *system.InformationListRes, err error) {
	res, err = service.SysDataManger().InformationList(ctx, req)
	return
}

//InformationGetReq
func (c *sysDataController) InformationGet(ctx context.Context, req *system.InformationGetReq) (res *system.InformationGetRes, err error) {
	res, err = service.SysDataManger().InformationGet(ctx, req)
	return
}

//InformationEditReq
func (c *sysDataController) InformationEdit(ctx context.Context, req *system.InformationEditReq) (res *system.InformationEditRes, err error) {
	res, err = service.SysDataManger().InformationEdit(ctx, req)
	return
}

//InformationDisplayReq
func (c *sysDataController) InformationDisplay(ctx context.Context, req *system.InformationDisplayReq) (res *system.InformationDisplayRes, err error) {
	res, err = service.SysDataManger().InformationDisplay(ctx, req)
	return
}

//InformationDelReq
func (c *sysDataController) InformationDel(ctx context.Context, req *system.InformationDelReq) (res *system.InformationDelRes, err error) {
	res, err = service.SysDataManger().InformationDel(ctx, req)
	return
}

//StatisticsReq
func (c *sysDataController) Statistics(ctx context.Context, req *system.StatisticsReq) (res *system.StatisticsRes, err error) {
	res, err = service.SysDataManger().Statistics(ctx, req)
	return
}

//  调研公告

//ReleaseReq

func (c *sysDataController) Release(ctx context.Context, req *system.ReleaseReq) (res *system.ReleaseRes, err error) {
	res, err = service.SysDataManger().Release(ctx, req)
	return
}

//ReleaseEditReq

func (c *sysDataController) ReleaseEdit(ctx context.Context, req *system.ReleaseEditReq) (res *system.ReleaseEditRes, err error) {
	res, err = service.SysDataManger().ReleaseEdit(ctx, req)
	return
}

//ReleaseListReq
func (c *sysDataController) ReleaseList(ctx context.Context, req *system.ReleaseListReq) (res *system.ReleaseListRes, err error) {
	res, err = service.SysDataManger().ReleaseList(ctx, req)
	return
}

//ReleaseGetReq
func (c *sysDataController) ReleaseGet(ctx context.Context, req *system.ReleaseGetReq) (res *system.ReleaseGetRes, err error) {
	res, err = service.SysDataManger().ReleaseGet(ctx, req)
	return
}

//ReleaseDelReq
func (c *sysDataController) ReleaseDel(ctx context.Context, req *system.ReleaseDelReq) (res *system.ReleaseDelRes, err error) {
	res, err = service.SysDataManger().ReleaseDel(ctx, req)
	return
}

//FeedbackListReq
func (c *sysDataController) FeedbackList(ctx context.Context, req *system.FeedbackListReq) (res *system.FeedbackListRes, err error) {
	res, err = service.SysDataManger().FeedbackList(ctx, req)
	return
}

//FeedbackGetReq
func (c *sysDataController) FeedbackGet(ctx context.Context, req *system.FeedbackGetReq) (res *system.FeedbackGetRes, err error) {
	res, err = service.SysDataManger().FeedbackGet(ctx, req)
	return
}

//FeedbackDelReq
func (c *sysDataController) FeedbackDel(ctx context.Context, req *system.FeedbackDelReq) (res *system.FeedbackDelRes, err error) {
	res, err = service.SysDataManger().FeedbackDel(ctx, req)
	return
}

//FeedbackReviewReq
func (c *sysDataController) FeedbackReview(ctx context.Context, req *system.FeedbackReviewReq) (res *system.FeedbackReviewRes, err error) {
	res, err = service.SysDataManger().FeedbackReview(ctx, req)
	return
}
