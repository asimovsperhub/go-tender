package service

import (
	"context"
	"tender/api/v1/system"
)

type IDataManger interface {
	List(ctx context.Context, req *system.DataSearchReq) (res *system.DataSearchRes, err error)
	TenderGet(ctx context.Context, req *system.TenderGetReq) (res *system.TenderGetRes, err error)
	TenderEdit(ctx context.Context, req *system.TenderEditReq) (res *system.TenderEditRes, err error)
	TenderTypeGet(ctx context.Context, req *system.TenderTypeGetReq) (res *system.TenderTypeGetRes, err error)
	TenderAddSort(ctx context.Context, req *system.TenderAddSortReq) (res *system.TenderAddSortRes, err error)
	TenderDelSort(ctx context.Context, req *system.TenderDelSortReq) (res *system.TenderDelSortRes, err error)
	TenderDrag(ctx context.Context, req *system.TenderDragReq) (res *system.TenderDragRes, err error)
	TenderDel(ctx context.Context, req *system.TenderDelReq) (res *system.TenderDelRes, err error)
	LawList(ctx context.Context, req *system.LawSearchReq) (res *system.LawSearchRes, err error)
	Setting(ctx context.Context, req *system.SettingReq) (res *system.SettingRes, err error)
	GetSetting(ctx context.Context, req *system.SettingGetReq) (res *system.SettingGetRes, err error)
	AddLaw(ctx context.Context, req *system.LawAddReq) (res *system.LawAddRes, err error)
	AllLawList(ctx context.Context, req *system.AllLawSearchReq) (res *system.AllLawSearchRes, err error)
	DelLaw(ctx context.Context, req *system.LawDelReq) (res *system.LawDelRes, err error)

	EnterpriseList(ctx context.Context, req *system.EnterpriseListReq) (res *system.EnterpriseListRes, err error)
	EnterpriseGet(ctx context.Context, req *system.EnterpriseGetReq) (res *system.EnterpriseGetRes, err error)
	EnterpriseDel(ctx context.Context, req *system.EnterpriseDelReq) (res *system.EnterpriseDelRes, err error)

	KnowledgeList(ctx context.Context, req *system.KnowledgeListReq) (res *system.KnowledgeListRes, err error)
	KnowledgeGet(ctx context.Context, req *system.KnowledgeGetReq) (res *system.KnowledgeGetRes, err error)
	KnowledgeDisplay(ctx context.Context, req *system.KnowledgeDisplayReq) (res *system.KnowledgeDisplayRes, err error)
	KnowledgeTypeGet(ctx context.Context, req *system.KnowledgeTypeGetReq) (res *system.KnowledgeTypeGetRes, err error)

	InformationList(ctx context.Context, req *system.InformationListReq) (res *system.InformationListRes, err error)
	InformationGet(ctx context.Context, req *system.InformationGetReq) (res *system.InformationGetRes, err error)
	InformationEdit(ctx context.Context, req *system.InformationEditReq) (res *system.InformationEditRes, err error)
	InformationDisplay(ctx context.Context, req *system.InformationDisplayReq) (res *system.InformationDisplayRes, err error)
	InformationDel(ctx context.Context, req *system.InformationDelReq) (res *system.InformationDelRes, err error)

	Statistics(ctx context.Context, req *system.StatisticsReq) (res *system.StatisticsRes, err error)

	Release(ctx context.Context, req *system.ReleaseReq) (res *system.ReleaseRes, err error)
	ReleaseEdit(ctx context.Context, req *system.ReleaseEditReq) (res *system.ReleaseEditRes, err error)
	ReleaseDel(ctx context.Context, req *system.ReleaseDelReq) (res *system.ReleaseDelRes, err error)
	ReleaseList(ctx context.Context, req *system.ReleaseListReq) (res *system.ReleaseListRes, err error)
	ReleaseGet(ctx context.Context, req *system.ReleaseGetReq) (res *system.ReleaseGetRes, err error)

	FeedbackList(ctx context.Context, req *system.FeedbackListReq) (res *system.FeedbackListRes, err error)
	FeedbackGet(ctx context.Context, req *system.FeedbackGetReq) (res *system.FeedbackGetRes, err error)
	FeedbackDel(ctx context.Context, req *system.FeedbackDelReq) (res *system.FeedbackDelRes, err error)
	FeedbackReview(ctx context.Context, req *system.FeedbackReviewReq) (res *system.FeedbackReviewRes, err error)
}

var localDataManger IDataManger

func SysDataManger() IDataManger {
	if localDataManger == nil {
		panic("implement not found for interface SysMemberManger, forgot register?")
	}
	return localDataManger
}

func RegisterDataManger(i IDataManger) {
	localDataManger = i
}
