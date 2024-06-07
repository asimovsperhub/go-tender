package service

import (
	"context"
	"tender/api/v1/desk"
)

type IIndexManger interface {
	BidSearchList(ctx context.Context, req *desk.BidSearchReq) (res *desk.BidSearchRes, err error)
	EnterpriseSearchList(ctx context.Context, req *desk.EnterpriseSearchReq) (res *desk.EnterpriseSearchRes, err error)
	EnterpriseCommerce(ctx context.Context, req *desk.EnterpriseCommerceReq) (res *desk.EnterpriseCommerceRes, err error)
	EnterprisePunishment(ctx context.Context, req *desk.EnterprisePunishmentReq) (res *desk.EnterprisePunishmentRes, err error)
	EnterpriseQualification(ctx context.Context, req *desk.EnterpriseQualificationReq) (res *desk.EnterpriseQualificationRes, err error)
	EnterpriseLawSuit(ctx context.Context, req *desk.EnterpriseLawSuitReq) (res *desk.EnterpriseLawSuitRes, err error)
	EnterpriseBidding(ctx context.Context, req *desk.EnterpriseBiddingReq) (res *desk.EnterpriseBiddingRes, err error)
	AnnouncementList(ctx context.Context, req *desk.AnnouncementListReq) (res *desk.AnnouncementListRes, err error)
	ConsultationList(ctx context.Context, req *desk.ConsultationListReq) (res *desk.ConsultationListRes, err error)
	ConsulList(ctx context.Context, req *desk.ConsulListReq) (res *desk.ConsulListRes, err error)
	SettingGet(ctx context.Context, req *desk.SettingGetReq) (res *desk.SettingGetRes, err error)
	Consul(ctx context.Context, req *desk.ConsulReq) (res *desk.ConsulRes, err error)
	KnowledgeList(ctx context.Context, req *desk.KnowledgeListReq) (res *desk.KnowledgeListRes, err error)
	KnowledgeBrowse(ctx context.Context, req *desk.KnowledgeBrowseReq) (res *desk.KnowledgeBrowseRes, err error)
	Announcement(ctx context.Context, req *desk.AnnouncementReq) (res *desk.AnnouncementRes, err error)
	Consultation(ctx context.Context, req *desk.ConsultationReq) (res *desk.ConsultationRes, err error)

	StatisticsBrowse(ctx context.Context, req *desk.StatisticsBrowseReq) (res *desk.StatisticsBrowseRes, err error)
	StatisticsGet(ctx context.Context, req *desk.StatisticsGetReq) (res *desk.StatisticsGetRes, err error)

	MemberInFind(ctx context.Context, req *desk.MemberInFindReq) (res *desk.MemberInFindRes, err error)

	Search(ctx context.Context, req *desk.SearchReq) (res *desk.SearchRes, err error)
}

var localIndexManger IIndexManger

func IndexManger() IIndexManger {
	if localIndexManger == nil {
		panic("implement not found for interface IndexManger, forgot register?")
	}
	return localIndexManger
}

func RegisterIndexManger(i IIndexManger) {
	localIndexManger = i
}
