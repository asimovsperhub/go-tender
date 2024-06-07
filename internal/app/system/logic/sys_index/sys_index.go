package sys_index

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	"tender/api/v1/system"
	"tender/internal/app/system/consts"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/service"
	"tender/library/liberr"
)

func init() {
	service.RegisterIndexManger(New())
}

func New() *sSysIndexManger {
	return &sSysIndexManger{}
}

type sSysIndexManger struct {
}

func (s *sSysIndexManger) EnterpriseList(ctx context.Context, req *system.EnterpriseReq) (res *system.EnterpriseRes, err error) {
	res = new(system.EnterpriseRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		// 正常的是10条数据
		req.PageSize = consts.PageSize
	}
	m := dao.SysEnterprise.Ctx(ctx)
	// 以时间
	order := "created_at DESC"
	err = g.Try(ctx, func(ctx context.Context) {
		// 入驻企业总数
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取企业列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.EnterpriseList)
		log.Println("获取企业列表-------->", res.EnterpriseList)
		liberr.ErrIsNil(ctx, err, "获取企业列表数据失败")
	})
	return
}

func (s *sSysIndexManger) MemberList(ctx context.Context, req *system.MemberUserReq) (res *system.MemberUserRes, err error) {
	res = new(system.MemberUserRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.MemberUser.Ctx(ctx)
	// 0 是非会员
	m = m.Where("member_level <> 0")
	order := "created_at DESC"
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取最新会员列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.MemberUserList)
		log.Println("获取最新会员列表-------->", res.MemberUserList)
		liberr.ErrIsNil(ctx, err, "获取最新会员列表数据失败")
	})
	return
}

//UserCount

func (s *sSysIndexManger) UserCount(ctx context.Context, req *system.UserReq) (res *system.UserRes, err error) {
	res = new(system.UserRes)
	m := dao.MemberUser.Ctx(ctx)
	err = g.Try(ctx, func(ctx context.Context) {
		res.Count, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取用户数量失败")
	})
	return
}

func (s *sSysIndexManger) NumberCount(ctx context.Context, req *system.UserNumberReq) (res *system.UserNumberRes, err error) {
	res = new(system.UserNumberRes)
	m := dao.MemberUser.Ctx(ctx)
	err = g.Try(ctx, func(ctx context.Context) {
		res.UserCount, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取用户数量失败")
	})
	err = g.Try(ctx, func(ctx context.Context) {
		res.UserCount, err = m.Where("member_level > 0").Count()
		liberr.ErrIsNil(ctx, err, "获取会员数量失败")
	})
	err = g.Try(ctx, func(ctx context.Context) {
		res.EnterpriseCount, err = dao.SysEnterprise.Ctx(ctx).Count()
		liberr.ErrIsNil(ctx, err, "获取企业数量失败")
	})
	return
}

//Trending
func (s *sSysIndexManger) Trending(ctx context.Context, req *system.TrendingReq) (res *system.TrendingRes, err error) {
	res = new(system.TrendingRes)

	searchType := req.Type
	m := dao.MemberUser.Ctx(ctx)
	m = m.Where("member_level > 0")
	if searchType == "day" {
		if req.Start != "" {
			m = m.Where("DATE_FORMAT(created_at, '%Y-%m-%d') >=?", req.Start)
		}
		if req.End != "" {
			m = m.Where("DATE_FORMAT(created_at, '%Y-%m-%d') <=?", req.End)
		}
		result, err := m.Fields("DATE_FORMAT(created_at, '%Y-%m-%d') as date, count(id) as total").Group("date").Order("date").All()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println(result)
		var resultList []*system.Trending
		for _, item := range result {
			resultItem := system.Trending{}
			row := item.Map()
			resultItem.Name = row["date"].(string)
			resultItem.Value = row["total"].(int64)
			resultList = append(resultList, &resultItem)
		}

		res.TrendingItemList = resultList
	} else if searchType == "month" {
		if req.Start != "" {
			m = m.Where("DATE_FORMAT(created_at, '%Y-%m') >=?", req.Start)
		}
		if req.End != "" {
			m = m.Where("DATE_FORMAT(created_at, '%Y-%m') <=?", req.End)
		}
		result, err := m.Fields("DATE_FORMAT(created_at, '%Y-%m') as date, count(id) as total").Group("date").Order("date").All()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println(result)
		var resultList []*system.Trending
		for _, item := range result {
			resultItem := system.Trending{}
			row := item.Map()
			resultItem.Name = row["date"].(string)
			resultItem.Value = row["total"].(int64)
			resultList = append(resultList, &resultItem)
		}

		res.TrendingItemList = resultList
	}

	return
}
