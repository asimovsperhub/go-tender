package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var (
	Index = IndexController{}
)

type IndexController struct {
	BaseController
}

// 招标搜索
func (c *IndexController) BidSearchList(ctx context.Context, req *desk.BidSearchReq) (res *desk.BidSearchRes, err error) {
	res, err = service.IndexManger().BidSearchList(ctx, req)
	return
}

// 企业搜索
func (c *IndexController) EnterpriseSearchList(ctx context.Context, req *desk.EnterpriseSearchReq) (res *desk.EnterpriseSearchRes, err error) {
	res, err = service.IndexManger().EnterpriseSearchList(ctx, req)
	return
}

//EnterpriseDetailsReq
func (c *IndexController) EnterpriseDetails(ctx context.Context, req *desk.EnterpriseCommerceReq) (res *desk.EnterpriseCommerceRes, err error) {
	res, err = service.IndexManger().EnterpriseCommerce(ctx, req)
	return
}

//EnterprisePunishmentReq
func (c *IndexController) EnterprisePunishment(ctx context.Context, req *desk.EnterprisePunishmentReq) (res *desk.EnterprisePunishmentRes, err error) {
	res, err = service.IndexManger().EnterprisePunishment(ctx, req)
	return
}

//EnterpriseQualificationReq
func (c *IndexController) EnterpriseQualification(ctx context.Context, req *desk.EnterpriseQualificationReq) (res *desk.EnterpriseQualificationRes, err error) {
	res, err = service.IndexManger().EnterpriseQualification(ctx, req)
	return
}

//EnterpriseLawSuitReq
func (c *IndexController) EnterpriseLawSuit(ctx context.Context, req *desk.EnterpriseLawSuitReq) (res *desk.EnterpriseLawSuitRes, err error) {
	res, err = service.IndexManger().EnterpriseLawSuit(ctx, req)
	return
}

//EnterpriseBiddingReq
func (c *IndexController) EnterpriseBidding(ctx context.Context, req *desk.EnterpriseBiddingReq) (res *desk.EnterpriseBiddingRes, err error) {
	res, err = service.IndexManger().EnterpriseBidding(ctx, req)
	return
}

// 最新公告
func (c *IndexController) AnnouncementList(ctx context.Context, req *desk.AnnouncementListReq) (res *desk.AnnouncementListRes, err error) {
	res, err = service.IndexManger().AnnouncementList(ctx, req)
	return
}

//AnnouncementReq
func (c *IndexController) Announcement(ctx context.Context, req *desk.AnnouncementReq) (res *desk.AnnouncementRes, err error) {
	res, err = service.IndexManger().Announcement(ctx, req)
	return
}

//ConsulListReq
//咨询
func (c *IndexController) ConsulList(ctx context.Context, req *desk.ConsulListReq) (res *desk.ConsulListRes, err error) {
	res, err = service.IndexManger().ConsulList(ctx, req)
	return
}

//SettingGetReq
func (c *IndexController) SettingGet(ctx context.Context, req *desk.SettingGetReq) (res *desk.SettingGetRes, err error) {
	res, err = service.IndexManger().SettingGet(ctx, req)
	return
}

// 咨询 非自营 ConsulReq

func (c *IndexController) Consul(ctx context.Context, req *desk.ConsulReq) (res *desk.ConsulRes, err error) {
	res, err = service.IndexManger().Consul(ctx, req)
	return
}

//政策
func (c *IndexController) ConsultationList(ctx context.Context, req *desk.ConsultationListReq) (res *desk.ConsultationListRes, err error) {
	res, err = service.IndexManger().ConsultationList(ctx, req)
	return
}

// 单个政策 ConsultationReq
func (c *IndexController) Consultation(ctx context.Context, req *desk.ConsultationReq) (res *desk.ConsultationRes, err error) {
	res, err = service.IndexManger().Consultation(ctx, req)
	return
}

// 在线知库
//KnowledgeListReq
func (c *IndexController) KnowledgeList(ctx context.Context, req *desk.KnowledgeListReq) (res *desk.KnowledgeListRes, err error) {
	res, err = service.IndexManger().KnowledgeList(ctx, req)
	return
}

//KnowledgeBrowseReq
func (c *IndexController) KnowledgeBrowse(ctx context.Context, req *desk.KnowledgeBrowseReq) (res *desk.KnowledgeBrowseRes, err error) {
	res, err = service.IndexManger().KnowledgeBrowse(ctx, req)
	return
}

//文章收藏量 CollectCountGetReq
func (c *IndexController) CollectCountGet(ctx context.Context, req *desk.CollectCountGetReq) (res *desk.CollectCountGetRes, err error) {
	res, err = service.Collect().CollectCountGet(ctx, req)
	return
}

//ConsultationBrowseReq
func (c *IndexController) ConsultationBrowse(ctx context.Context, req *desk.StatisticsBrowseReq) (res *desk.StatisticsBrowseRes, err error) {
	res, err = service.IndexManger().StatisticsBrowse(ctx, req)
	return
}

//StatisticsGetReq
func (c *IndexController) StatisticsGet(ctx context.Context, req *desk.StatisticsGetReq) (res *desk.StatisticsGetRes, err error) {
	res, err = service.IndexManger().StatisticsGet(ctx, req)
	return
}

//MemberInFindReq
func (c *IndexController) MemberInFind(ctx context.Context, req *desk.MemberInFindReq) (res *desk.MemberInFindRes, err error) {
	res, err = service.IndexManger().MemberInFind(ctx, req)
	return
}

//SearchReq
func (c *IndexController) Search(ctx context.Context, req *desk.SearchReq) (res *desk.SearchRes, err error) {
	res, err = service.IndexManger().Search(ctx, req)
	return
}
