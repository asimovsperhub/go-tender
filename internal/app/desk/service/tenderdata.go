package service

import (
	"context"
	"tender/api/v1/desk"
)

type IDataManger interface {
	List(ctx context.Context, req *desk.DataSearchReq) (res *desk.DataSearchRes, err error)
}

var localDataManger IDataManger

func DeskTenderDataManger() IDataManger {
	if localDataManger == nil {
		panic("implement not found for interface DeskTenderDataManger, forgot register?")
	}
	return localDataManger
}

func RegisterDataManger(i IDataManger) {
	localDataManger = i
}
