package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"tender/api/v1/common"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/do"
	"tender/internal/packed/websocket"
)

var Msg = msgController{}

type msgController struct {
	BaseController
}

func (c *msgController) SendWs(ctx context.Context, req *common.SendWsReq) (*common.SendWsRes, error) {

	msgId := uuid.New().String()

	_, err := dao.SysWsMsg.Ctx(ctx).Insert(do.SysWsMsg{
		MessageId: msgId,
		UserId:    req.UserId,
		Content:   req.Content,
		IsRead:    0,
		IsDel:     0,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("插入消息失败: %w", err)
	}

	websocket.SendToUser(req.UserId, &websocket.WResponse{
		Event: "context",
		Data: map[string]string{
			"messageId": msgId,
			"content":   req.Content,
		},
	})
	return &common.SendWsRes{}, nil
}
