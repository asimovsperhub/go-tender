package tenderdata

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	"tender/api/v1/desk"
	"tender/internal/app/desk/consts"
	"tender/internal/app/desk/service"
	"tender/library/liberr"
)

func init() {
	service.RegisterDataManger(New())
}

func New() *sTenderData {
	return &sTenderData{}
}

type sTenderData struct {
}

// 招标数据
func (s *sTenderData) List(ctx context.Context, req *desk.DataSearchReq) (res *desk.DataSearchRes, err error) {
	res = new(desk.DataSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	//fmt.Sprintf("MATCH(title,announcement_content) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords)
	// m := g.DB("data").Schema("crawldata").Model("bidding_information")
	m := g.DB("data").Schema("crawldata").Model("bid")
	//m := dao.BiddingInformation.Ctx(ctx)
	order := "id"
	if req.BulletinType != "" {
		//bulletin_type  notice_nature
		m = m.Where("bulletin_type = ?", req.BulletinType)
	}
	if req.KeyWords != "" {
		m = m.Where(fmt.Sprintf("MATCH(title,announcement_content) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.BiddingInformation)
		log.Println("获取数据列表-------->", res.BiddingInformation)
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
	})
	return
}
