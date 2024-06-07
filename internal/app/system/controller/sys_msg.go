package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"tender/api/v1/system"
	commonService "tender/internal/app/common/service"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/service"
	"tender/internal/packed/websocket"

	"github.com/chanxuehong/wechat/mp/message/template"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

var Msg = msgController{}

type msgController struct {
	BaseController
}

func (c *msgController) SendTemplate(ctx context.Context, req *system.SendTemplateReq) (*system.SendTemplateRes, error) {

	var trueMap = map[string]map[string]string{}
	for k, v := range req.Params {
		trueMap[k] = map[string]string{
			"value": v,
		}
	}
	// 将 map 转换为 JSON 字节数组
	jsonData, err := json.Marshal(trueMap)
	if err != nil {
		return nil, err
	}
	// 将 JSON 字节数组转换为 json.RawMessage 类型
	rawMessage := json.RawMessage(jsonData)
	if _, err = template.Send(commonService.Mp(), template.TemplateMessage{
		ToUser:     req.ToUser,
		TemplateId: req.TemplateId,
		URL:        req.URL,
		Data:       rawMessage,
	}); err != nil {
		return nil, err
	}
	return &system.SendTemplateRes{}, nil
}

func (c *msgController) SendWs(ctx context.Context, req *system.SendWsReq) (*system.SendWsRes, error) {

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
	return &system.SendWsRes{}, nil
}

func (c *msgController) WsMsgList(ctx context.Context, req *system.WsMsgListReq) (*system.WsMsgListRes, error) {
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	return service.SysMsg().List(ctx, req)
}

func (c *msgController) MarkAsRead(ctx context.Context, req *system.MarkAsReadReq) (*system.MarkAsReadRes, error) {
	return service.SysMsg().MarkAsRead(ctx, req)
}

//MarkAsAllReadReq
func (c *msgController) MarkAsAllRead(ctx context.Context, req *system.MarkAsAllReadReq) (*system.MarkAsAllReadRes, error) {
	return service.SysMsg().MarkAsAllRead(ctx, req)
}
func (c *msgController) DeleteMsg(ctx context.Context, req *system.DeleteMsgReq) (*system.DeleteMsgRes, error) {
	return service.SysMsg().DeleteMsg(ctx, req)
}

func (c *msgController) ClearMsg(ctx context.Context, req *system.ClearMsgReq) (*system.ClearMsgRes, error) {
	return service.SysMsg().ClearMsg(ctx, req)
}

func (c *msgController) UnreadCount(ctx context.Context, req *system.UnreadCountReq) (*system.UnreadCountRes, error) {
	return service.SysMsg().UnreadCount(ctx, req)
}
