package sysMsg

import (
	"context"
	"fmt"
	"tender/api/v1/system"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysMsg(New())
}

func New() *sSysMsg {
	return &sSysMsg{}
}

type sSysMsg struct {
}

func (s *sSysMsg) List(ctx context.Context, req *system.WsMsgListReq) (*system.WsMsgListRes, error) {
	res := &system.WsMsgListRes{}
	daoIns := dao.SysWsMsg.Ctx(ctx)
	totalCount, err := daoIns.Where("is_del = ? AND user_id = ?", 0, req.UserId).Count()
	if err != nil {
		return nil, fmt.Errorf("获取总数失败: %w", err)
	}
	res.Total = totalCount                                         // 总条数
	res.TotalPage = (totalCount + req.PageSize - 1) / req.PageSize // 总页数

	err = daoIns.Where("is_del = ? AND user_id = ?", 0, req.UserId).Limit((req.PageNum-1)*req.PageSize, req.PageSize).OrderDesc("created_at").Scan(&res.List)
	if err != nil {
		return nil, fmt.Errorf("获取分页数据失败: %w", err)
	}
	return res, nil
}

func (s *sSysMsg) MarkAsRead(ctx context.Context, req *system.MarkAsReadReq) (*system.MarkAsReadRes, error) {
	res := &system.MarkAsReadRes{}
	daoIns := dao.SysWsMsg.Ctx(ctx)
	result, err := daoIns.Where("message_id = ? AND is_del = ?", req.MessageId, 0).Data("is_read", 1).Update()
	if err != nil {
		return nil, fmt.Errorf("更新失败: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("获取影响行数失败: %w", err)
	}
	if rows == 0 {
		return nil, fmt.Errorf("未找到该消息")
	}
	return res, err
}

//MarkAsAllRead
func (s *sSysMsg) MarkAsAllRead(ctx context.Context, req *system.MarkAsAllReadReq) (*system.MarkAsAllReadRes, error) {
	res := &system.MarkAsAllReadRes{}
	daoIns := dao.SysWsMsg.Ctx(ctx)
	result, err := daoIns.Where("user_id = ? AND is_del = ?", req.UserId, 0).Data("is_read", 1).Update()
	if err != nil {
		return nil, fmt.Errorf("更新失败: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("获取影响行数失败: %w", err)
	}
	if rows == 0 {
		return nil, fmt.Errorf("未找到该消息")
	}
	return res, err
}

func (s *sSysMsg) DeleteMsg(ctx context.Context, req *system.DeleteMsgReq) (*system.DeleteMsgRes, error) {
	res := &system.DeleteMsgRes{}
	daoIns := dao.SysWsMsg.Ctx(ctx)
	result, err := daoIns.Where("message_id", req.MessageId).Data("is_del", 1).Update()
	if err != nil {
		return nil, fmt.Errorf("删除失败: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("获取影响行数失败: %w", err)
	}
	if rows == 0 {
		return nil, fmt.Errorf("未找到该消息")
	}
	return res, err
}

func (s *sSysMsg) ClearMsg(ctx context.Context, req *system.ClearMsgReq) (res *system.ClearMsgRes, err error) {
	res = &system.ClearMsgRes{}
	daoIns := dao.SysWsMsg.Ctx(ctx)
	result, err := daoIns.Where("user_id = ?", req.UserId).Data("is_del", 1).Update()
	if err != nil {
		return nil, fmt.Errorf("清空失败: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("获取影响行数失败: %w", err)
	}
	if rows == 0 {
		return nil, fmt.Errorf("未找到该用户的消息")
	}
	return res, err
}

func (s *sSysMsg) UnreadCount(ctx context.Context, req *system.UnreadCountReq) (res *system.UnreadCountRes, err error) {
	res = &system.UnreadCountRes{}
	daoIns := dao.SysWsMsg.Ctx(ctx)
	ct, err := daoIns.Count(g.Map{
		dao.SysWsMsg.Columns().IsRead: 0,
		dao.SysWsMsg.Columns().IsDel:  0,
		dao.SysWsMsg.Columns().UserId: req.UserId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取未读消息数失败: %w", err)
	}
	res.Count = ct
	return res, err
}
