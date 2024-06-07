package service

import (
	"context"
	"tender/api/v1/system"
)

type ISysMsg interface {
	List(ctx context.Context, req *system.WsMsgListReq) (res *system.WsMsgListRes, err error)
	MarkAsRead(ctx context.Context, req *system.MarkAsReadReq) (res *system.MarkAsReadRes, err error)
	MarkAsAllRead(ctx context.Context, req *system.MarkAsAllReadReq) (res *system.MarkAsAllReadRes, err error)
	DeleteMsg(ctx context.Context, req *system.DeleteMsgReq) (res *system.DeleteMsgRes, err error)
	ClearMsg(ctx context.Context, req *system.ClearMsgReq) (res *system.ClearMsgRes, err error)
	UnreadCount(ctx context.Context, req *system.UnreadCountReq) (res *system.UnreadCountRes, err error)
}

var localSysMsg ISysMsg

func SysMsg() ISysMsg {
	if localSysMsg == nil {
		panic("implement not found for interface SysMsg, forgot register?")
	}
	return localSysMsg
}

func RegisterSysMsg(i ISysMsg) {
	localSysMsg = i
}
