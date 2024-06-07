package sys_data

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"log"
	"strings"
	"tender/api/v1/system"
	deskdao "tender/internal/app/desk/dao"
	desk_do "tender/internal/app/desk/model/do"
	deskentity "tender/internal/app/desk/model/entity"
	"tender/internal/app/system/consts"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
	"tender/internal/packed/websocket"
	"tender/library/libUtils"
	"tender/library/liberr"
	"time"
)

func init() {
	service.RegisterDataManger(New())
}

func New() *sDataManger {
	return &sDataManger{}
}

type sDataManger struct {
}

// 招标数据
func (s *sDataManger) List(ctx context.Context, req *system.DataSearchReq) (res *system.DataSearchRes, err error) {
	res = new(system.DataSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := g.DB("data").Schema("crawldata").Model("bid")
	// order := "release_time desc"
	if req.BulletinType != "" {
		if req.BulletinType == "在建公告" {
			m = m.Where("bulletin_type = ?", "招标公告")
			m = m.Where("bidopening_time > ?", time.Now().Format("2006-01-02"))
		} else {
			m = m.Where("bulletin_type = ?", req.BulletinType)
		}
	}
	if req.IndustryType != "" {
		m = m.Where("industry_classification = ?", req.IndustryType)
	}
	if req.City != "" {
		if strings.HasSuffix(req.City, "省") {
			m = m.Where("province = ?", strings.Replace(req.City, "省", "", -1))
		} else {
			m = m.Where("city = ?", req.City)
		}
	}
	if req.KeyWords != "" {
		m = m.Where(fmt.Sprintf("MATCH(title,announcement_content) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
	}
	if req.Attachment != nil {
		if *req.Attachment == 1 {
			m = m.Where("attachment <> ''")
		} else if *req.Attachment == 0 {
			m = m.Where("attachment = ''")
		}
	}
	if req.OriginalBulletin != "" {
		m = m.Where("original_type = ?", req.OriginalBulletin)
	}
	if req.OriginalIndustry != "" {
		m = m.Where("original_classification = ?", req.OriginalIndustry)
	}
	if req.Site != "" {
		m = m.Where("site = ?", req.Site)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		if req.BulletinType == "中标公告" {
			err = m.Fields("id,title,release_day,release_time,city,bulletin_type,bidopening_time,bid_amount,industry_classification,abstract,rank,attachment,contact_person,display,crawler").Page(req.PageNum, req.PageSize).Order("release_day desc,area desc,bid_amount desc,rank desc").Scan(&res.BiddingInformation)
			liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
		} else {
			err = m.Fields("id,title,release_day,release_time,city,bulletin_type,bidopening_time,amount,industry_classification,abstract,rank,attachment,contact_person,display,crawler").Page(req.PageNum, req.PageSize).Order("release_day desc,area desc,amount desc,rank desc").Scan(&res.BiddingInformation)
			liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
		}
	})
	for i := 0; i < len(res.BiddingInformation); i++ {
		att := []string{}
		if res.BiddingInformation[i].Attachment != "" {
			for _, url := range strings.Split(res.BiddingInformation[i].Attachment, ",") {
				m0 := strings.Split(url, "/")
				m1 := m0[len(m0)-1]
				m2 := m0[len(m0)-2]
				att = append(att, fmt.Sprintf("https://biaoziku.com/api/v1/system/tender/single/download?directory=%s&name=%s", m2, m1))
			}
		}
		res.BiddingInformation[i].Attachment = strings.Join(att, ",")
	}
	return
}

//TenderGet
func (s *sDataManger) TenderGet(ctx context.Context, req *system.TenderGetReq) (res *system.TenderGetRes, err error) {
	res = new(system.TenderGetRes)
	// m := g.DB("data").Schema("crawldata").Model("bidding_information")
	m := g.DB("data").Schema("crawldata").Model("bid")
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.BiddingInformation)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

//TenderEdit
func (s *sDataManger) TenderEdit(ctx context.Context, req *system.TenderEditReq) (res *system.TenderEditRes, err error) {
	data := desk_do.Bid{}
	if req.BulletinType != "" {
		data.BulletinType = req.BulletinType
	}
	if req.IndustryType != "" {
		data.IndustryClassification = req.IndustryType
	}
	res = new(system.TenderEditRes)
	m := g.DB("data").Schema("crawldata").Model("bid")
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Update(data)
		liberr.ErrIsNil(ctx, err, "更新数据失败")
	})
	return
}

//TenderTypeGet
func (s *sDataManger) TenderTypeGet(ctx context.Context, req *system.TenderTypeGetReq) (res *system.TenderTypeGetRes, err error) {
	res = new(system.TenderTypeGetRes)
	m := g.DB("data").Schema("crawldata").Model("bid")
	m_ := g.DB("data").Schema("crawldata").Model("bid")
	err = g.Try(ctx, func(ctx context.Context) {
		original_type, err := m.Fields("original_type").Distinct().All()
		log.Println(original_type)
		if original_type != nil {
			for _, v := range original_type {
				mp := v.Map()
				if mp != nil {
					if mp["original_type"].(string) != "" {
						res.OriginalType = append(res.OriginalType, mp["original_type"].(string))
					}
				}
			}
		}
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	err = g.Try(ctx, func(ctx context.Context) {
		classification, err := m_.Fields("original_classification").Distinct().All()
		log.Println(classification)
		if classification != nil {
			for _, v := range classification {
				mp := v.Map()
				if mp != nil {
					if mp["original_classification"].(string) != "" {
						res.OriginalClassification = append(res.OriginalClassification, mp["original_classification"].(string))
					}
				}
			}
		}
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

// TenderAddSort
func (s *sDataManger) TenderAddSort(ctx context.Context, req *system.TenderAddSortReq) (res *system.TenderAddSortRes, err error) {
	res = new(system.TenderAddSortRes)
	err = g.Try(ctx, func(ctx context.Context) {
		for _, id := range req.Ids {
			rank := "a"
			if req.Height == 0 {
				rank = "w" + id
			} else if req.Height == 1 {
				rank = "m" + id
			} else if req.Height == 2 {
				rank = "c" + id
			}
			switch req.Type {
			case 0:
				m := g.DB("data").Schema("crawldata").Model("bid")
				_, err = m.Where("id = ?", id).Update(desk_do.Bid{Height: req.Height, Rank: rank})
				if err != nil {
					log.Println("更新排名出错")
				}
			case 1:
				m := g.DB("data").Schema("crawldata").Model("consultation")
				_, err = m.Where("id = ?", id).Update(desk_do.Consultation{Height: req.Height, Rank: rank})
				if err != nil {
					log.Println("更新排名出错")
				}
			case 2:
				m := deskdao.MemberKnowledge.Ctx(ctx)
				_, err = m.Where("id = ?", id).Update(desk_do.MemberKnowledge{Height: req.Height, Rank: rank})
				if err != nil {
					log.Println("更新排名出错")
				}
			}
		}
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

//TenderDelSort
func (s *sDataManger) TenderDelSort(ctx context.Context, req *system.TenderDelSortReq) (res *system.TenderDelSortRes, err error) {
	res = new(system.TenderDelSortRes)
	switch req.Type {
	case 0:
		m := g.DB("data").Schema("crawldata").Model("bid")
		_, err = m.Where("id in (?)", req.Ids).Update(desk_do.Bid{Height: "", Rank: ""})
		if err != nil {
			log.Println("更新排名出错")
		}
	case 1:
		m := g.DB("data").Schema("crawldata").Model("consultation")
		_, err = m.Where("id in (?)", req.Ids).Update(desk_do.Consultation{Height: "", Rank: ""})
		if err != nil {
			log.Println("更新排名出错")
		}
	case 2:
		m := deskdao.MemberKnowledge.Ctx(ctx)
		_, err = m.Where("id in (?)", req.Ids).Update(desk_do.MemberKnowledge{Height: "", Rank: ""})
		if err != nil {
			log.Println("更新排名出错")
		}
	}
	return
}

//TenderDrag
func (s *sDataManger) TenderDrag(ctx context.Context, req *system.TenderDragReq) (res *system.TenderDragRes, err error) {
	type Bid struct {
		Id     uint64 `json:"id"                     description:""`
		Height string `json:"height"                 description:"0高1中2低"`
		Rank   string `json:"rank"                   description:"排序"`
	}
	res = new(system.TenderDragRes)
	bid := []*Bid{}
	m := g.DB("data").Schema("crawldata").Model("bid")
	mp := g.DB("data").Schema("crawldata").Model("bid")
	switch req.Type {
	//case 0:
	//	m = g.DB("data").Schema("crawldata").Model("bid")
	//	mp = g.DB("data").Schema("crawldata").Model("bid")
	case 1:
		m = g.DB("data").Schema("crawldata").Model("consultation")
		mp = g.DB("data").Schema("crawldata").Model("consultation")

	case 2:
		m = deskdao.MemberKnowledge.Ctx(ctx)
		mp = deskdao.MemberKnowledge.Ctx(ctx)
	}
	// 下移
	if req.Direction == 0 {
		m = m.Fields("rank").Where("rank < ?", req.Rank).Order("rank desc").Limit(1, 2)
		err = m.Scan(&bid)
		liberr.ErrIsNil(ctx, err, "获取招标数据失败数据失败")
		switch len(bid) {
		case 2:
			prev := bid[0].Rank
			next := bid[1].Rank
			rank, ok := libUtils.Rank(next, prev)
			log.Println("向下------------->prev: %s,next:%s,rank:%s", prev, next, rank)
			if !ok {
				log.Println("需要重新排序 不满足生成的rank>next........")
			}
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = mp.Where("id = ?", req.Id).Update(desk_do.Bid{Rank: rank})
				if err != nil {
					log.Println("更新排名出错", err)
				}
				liberr.ErrIsNil(ctx, err, "更新数据失败")
			})
			return
		case 1:
			prev := bid[0].Rank
			rank, ok := libUtils.Rank(prev, prev+"n")
			if !ok {
				log.Println("需要重新排序 不满足生成的rank>next........")
			}
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = mp.Where("id = ?", req.Id).Update(desk_do.Bid{Rank: rank})
				if err != nil {
					log.Println("更新排名出错", err)
				}
				liberr.ErrIsNil(ctx, err, "更新数据失败")
			})
			return
		case 0:
			return
		}
	} else if req.Direction == 1 {
		m = m.Fields("rank").Where("rank > ?", req.Rank).Order("rank").Limit(1, 2)
		err = m.Scan(&bid)
		liberr.ErrIsNil(ctx, err, "获取招标数据失败数据失败")
		switch len(bid) {
		case 2:
			next := bid[0].Rank
			prev := bid[1].Rank
			rank, ok := libUtils.Rank(next, prev)
			log.Println("向上------------->prev: %s,next:%s,rank:%s", prev, next, rank)
			if !ok {
				log.Println("需要重新排序 不满足生成的rank>next........")
			}
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = mp.Where("id = ?", req.Id).Update(desk_do.Bid{Rank: rank})
				if err != nil {
					log.Println("更新排名出错", err)
				}
				liberr.ErrIsNil(ctx, err, "更新数据失败")
			})
			return
		case 1:
			next := bid[0].Rank
			rank, ok := libUtils.Rank(next, next+"n")
			if !ok {
				log.Println("需要重新排序 不满足生成的rank>next........")
			}
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = mp.Where("id = ?", req.Id).Update(desk_do.Bid{Rank: rank})
				if err != nil {
					log.Println("更新排名出错", err)
				}
				liberr.ErrIsNil(ctx, err, "更新数据失败")
			})
			return
		case 0:
			return
		}
	}
	return
}

//TenderDel
func (s *sDataManger) TenderDel(ctx context.Context, req *system.TenderDelReq) (res *system.TenderDelRes, err error) {
	//m := g.DB("data").Schema("crawldata").Model("bidding_information")
	m := g.DB("data").Schema("crawldata").Model("bid")
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Where("id = ?", req.Id).Delete()
		liberr.ErrIsNil(ctx, err, "删除数据失败")
	})
	return
}

// 所有法律数据
func (s *sDataManger) AllLawList(ctx context.Context, req *system.AllLawSearchReq) (res *system.AllLawSearchRes, err error) {
	res = new(system.AllLawSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := g.DB("data").Schema("crawldata").Model("law")
	// m := dao.Law.Ctx(ctx)
	order := "id"
	if req.KeyWords != "" {
		m = m.Where("title like ?", "%"+req.KeyWords+"%")
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取法律数据列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.Law)
		log.Println("获取数据列表-------->", res.Law)
		liberr.ErrIsNil(ctx, err, "获取法律数据列表数据失败")
	})
	return
}

// 常用法律数据
func (s *sDataManger) LawList(ctx context.Context, req *system.LawSearchReq) (res *system.LawSearchRes, err error) {
	res = new(system.LawSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	// m := g.DB("data").Schema("crawldata").Model("law")
	m := deskdao.Consultation.Ctx(ctx)
	order := "id"
	if req.KeyWords != "" {
		m = m.Where("title like ?", "%"+req.KeyWords+"%")
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取法律数据列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.Law)
		log.Println("获取数据列表-------->", res.Law)
		liberr.ErrIsNil(ctx, err, "获取法律数据列表数据失败")
	})
	//for i := 0; i < len(res.Law); i++ {
	//	if res.Law[i].Pdf != "" {
	//		res.Law[i].Pdf = "https://biaoziku.com/law/" + res.Law[i].Pdf
	//	} else {
	//		if res.Law[i].Word != "" {
	//			list := strings.Split(res.Law[i].Word, ".")
	//			res.Law[i].Pdf = "https://biaoziku.com/law/" + list[0] + ".pdf"
	//		}
	//	}
	//	if res.Law[i].Word != "" {
	//		res.Law[i].Word = "https://biaoziku.com/law/" + res.Law[i].Word
	//	}
	//}
	return
}

// 添加常用法律
func (s *sDataManger) AddLaw(ctx context.Context, req *system.LawAddReq) (res *system.LawAddRes, err error) {
	if req.Id == 0 {
		err = errors.New("请输入正确的法律id")
		return
	}
	var res_law *deskentity.Consultation
	m := g.DB("data").Schema("crawldata").Model("consultation")
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Limit(1).Scan(&res_law)
		liberr.ErrIsNil(ctx, err, "获取法律数据列表数据失败")
	})
	count := 0
	err = g.Try(ctx, func(ctx context.Context) {
		count, err = deskdao.Consultation.Ctx(ctx).Where("id = ?", res_law.Id).Count()
		liberr.ErrIsNil(ctx, err, "获取常用法律失败")
	})
	if count > 0 {
		err = errors.New("改该法律id已存在常用法律")
		return
	}
	if res_law != nil {
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err = g.Try(ctx, func(ctx context.Context) {
				//do.Law{Id: res_law.Id}
				_, err := deskdao.Consultation.Ctx(ctx).TX(tx).Insert(res_law)
				liberr.ErrIsNil(ctx, err, "添加常用法律失败")
			})
			return err
		})
	}
	return
}

//删除常用法律
func (s *sDataManger) DelLaw(ctx context.Context, req *system.LawDelReq) (res *system.LawDelRes, err error) {
	if req.Id == 0 {
		err = errors.New("请输入正确的法律id")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err := deskdao.Consultation.Ctx(ctx).TX(tx).Where("id = ?", req.Id).Delete()
			liberr.ErrIsNil(ctx, err, "删除常用法律失败")
		})
		return err
	})
	return
}

////	Id           string //
////	Consultation string // 咨询服务非营利类目
////	Purchase     string // 采购咨询分类
////	Bid          string // 投标咨询分类
////	Industry     string // 行业咨询分类
////	Market       string // 市场研究分类
////	OperationId  string // 操作管理员id
////	CreatedAt    string // 创建日期
////	UpdatedAt    string // 修改日期
func (s *sDataManger) Setting(ctx context.Context, req *system.SettingReq) (res *system.SettingRes, err error) {
	SysEnterprise := [](*entity.SysEnterprise)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysEnterprise.Ctx(ctx).Fields("industry").Distinct().Scan(&SysEnterprise)
		liberr.ErrIsNil(ctx, err, "获取企业列表数据失败")
	})
	// 咨询服务非营利类目
	if req.DelConsultation != nil {
		cons := strings.Split(*req.DelConsultation, ",")
		for i := 0; i < len(cons); i++ {
			if cons[i] != "" {
				for _, enter := range SysEnterprise {
					if cons[i] == enter.Industry {
						err = errors.New(fmt.Sprintf("已有企业属于该 %s 行业,不能删除", enter.Industry))
						return
					}
				}
			}
		}
	}
	var res_id *entity.SysDataset
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDataset.Ctx(ctx).Scan(&res_id)
		liberr.ErrIsNil(ctx, err, "获取数据中心设置失败")
	})
	opt := service.Context().GetUserId(ctx)
	if res_id != nil {
		data := do.SysDataset{}
		if req.Consultation != nil {
			data.Consultation = *req.Consultation
		}
		if req.Purchase != nil {
			data.Purchase = *req.Purchase
		}
		if req.PurchaseType != nil {
			data.PurchaseType = *req.PurchaseType
		}
		if req.Bid != nil {
			data.Bid = *req.Bid
		}
		if req.Industry != nil {
			data.Industry = *req.Industry
		}
		if req.IndustryType != nil {
			data.IndustryType = *req.IndustryType
		}
		if req.Market != nil {
			data.Market = *req.Market
		}
		if req.MarketType != nil {
			data.MarketType = *req.MarketType
		}
		if req.KeyWords != nil {
			data.Keywords = *req.KeyWords
		}
		data.OperationId = opt
		data.UpdatedAt = gtime.Now()
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysDataset.Ctx(ctx).WherePri(res_id.Id).Update(data)
			liberr.ErrIsNil(ctx, err, "修改设置失败")
		})
		return
	} else {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysDataset.Ctx(ctx).Insert(do.SysDataset{
				Consultation: *req.Consultation,
				Purchase:     *req.Purchase,
				PurchaseType: *req.PurchaseType,
				Bid:          *req.Bid,
				Industry:     *req.Industry,
				IndustryType: *req.IndustryType,
				Market:       *req.Market,
				MarketType:   *req.MarketType,
				Keywords:     *req.KeyWords,
				OperationId:  opt,
				CreatedAt:    gtime.Now(),
			})
			liberr.ErrIsNil(ctx, err, "添加设置失败")
		})
	}
	return
}

// GetSetting

func (s *sDataManger) GetSetting(ctx context.Context, req *system.SettingGetReq) (res *system.SettingGetRes, err error) {
	var res_id *entity.SysDataset
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDataset.Ctx(ctx).Scan(&res_id)
		liberr.ErrIsNil(ctx, err, "获取数据中心设置失败")
	})
	res = &system.SettingGetRes{
		Result: res_id,
	}
	return
}

//EnterpriseList
func (s *sDataManger) EnterpriseList(ctx context.Context, req *system.EnterpriseListReq) (res *system.EnterpriseListRes, err error) {
	res = new(system.EnterpriseListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysEnterprise.Ctx(ctx)
	if req.KeyWords != "" {
		m = m.Where("name like ?", "%"+req.KeyWords+"%")
	}
	m = m.Where("license_status = 1 and certificate_status = 1")
	total := 0
	err = g.Try(ctx, func(ctx context.Context) {
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取列表失败")
		err = m.Page(req.PageNum, req.PageSize).Scan(&res.EnterpriseInformation)
		log.Println("获取列表-------->", res.EnterpriseInformation)
		liberr.ErrIsNil(ctx, err, "获取列表数据失败")
	})
	res.Total = total
	//enterprise_result, err := m.Fields("COUNT(Id) total,location").Group("location").All()
	//if err != nil {
	//	liberr.ErrIsNil(ctx, err, "获取企业数据失败")
	//	return
	//}
	//data := map[string]int64{}
	//if enterprise_result != nil {
	//	for i := 0; i < len(enterprise_result); i++ {
	//		tm := enterprise_result[i].Map()
	//		if tm != nil {
	//			if data != nil {
	//				v, ok := data[tm["location"].(string)]
	//				if ok {
	//					data[tm["location"].(string)] = v + tm["total"].(int64)
	//				} else {
	//					data[tm["location"].(string)] = tm["total"].(int64)
	//				}
	//				//if data[tm["location"].(string)]
	//			}
	//			log.Println(tm["location"].(string), tm["total"].(int64))
	//		}
	//	}
	//}
	//_, ok := data["深圳市"]
	//if ok {
	//	v, ok1 := data["深圳"]
	//	if ok1 {
	//		data["深圳"] = data["深圳市"] + v
	//	}
	//	delete(data, "深圳市")
	//}
	//res.Statistics.City = data
	//res.Statistics.Count = total
	return
}

//EnterpriseGet
func (s *sDataManger) EnterpriseGet(ctx context.Context, req *system.EnterpriseGetReq) (res *system.EnterpriseGetRes, err error) {
	res = new(system.EnterpriseGetRes)
	m := dao.SysEnterprise.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.EnterpriseInformation)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

//EnterpriseDel
func (s *sDataManger) EnterpriseDel(ctx context.Context, req *system.EnterpriseDelReq) (res *system.EnterpriseDelRes, err error) {
	m := dao.SysEnterprise.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Delete()
		liberr.ErrIsNil(ctx, err, "删除数据失败")
	})
	return
}

// 在线知库  Knowledge
func (s *sDataManger) KnowledgeList(ctx context.Context, req *system.KnowledgeListReq) (res *system.KnowledgeListRes, err error) {
	res = new(system.KnowledgeListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := deskdao.MemberKnowledge.Ctx(ctx)

	if req.Primary != "" {
		m = m.Where("primary_classification = ?", req.Primary)
	}
	if req.Secondary != "" {
		m = m.Where("secondary_classification = ?", req.Secondary)
	}
	if req.KeyWords != "" {
		m = m.Where("title like ?", "%"+req.KeyWords+"%")
	}
	if req.Attachment != nil {
		if *req.Attachment == 1 {
			m = m.Where("attachment_url <> ''")
		} else if *req.Attachment == 0 {
			m = m.Where("attachment_url = ''")
		}
	}
	if req.OriginalIndustry != "" {
		m = m.Where("original_type = ?", req.OriginalIndustry)
	}
	if req.Site != "" {
		m = m.Where("site = ?", req.Site)
	}
	m = m.Where("review_status = 1")
	total := 0
	err = g.Try(ctx, func(ctx context.Context) {
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order("rank desc,created_at desc").Scan(&res.MemberKnowledge)
		log.Println("获取列表-------->", res.MemberKnowledge)
		liberr.ErrIsNil(ctx, err, "获取列表数据失败")
	})
	res.ListRes.Total = total
	//res.Total = total
	//res.Statistics.Count = total
	return
}

//KnowledgeGet
func (s *sDataManger) KnowledgeGet(ctx context.Context, req *system.KnowledgeGetReq) (res *system.KnowledgeGetRes, err error) {
	res = new(system.KnowledgeGetRes)
	m := deskdao.MemberKnowledge.Ctx(ctx)

	if req.Id != "" {
		m = m.Where("id = ?", req.Id)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.MemberKnowledge)
		log.Println("获取-------->", res.MemberKnowledge)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

//InformationList
func (s *sDataManger) InformationList(ctx context.Context, req *system.InformationListReq) (res *system.InformationListRes, err error) {
	res = new(system.InformationListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := g.DB("data").Schema("crawldata").Model("consultation")
	//if req.Classification == "法律法规" {
	//	m = g.DB("data").Schema("crawldata").Model("law")
	//} else {
	//	if req.Classification != "" {
	//		m = m.Where("type = ?", req.Classification)
	//	}
	//}
	if req.Classification != "" {
		m = m.Where("type = ?", req.Classification)
	}
	if req.KeyWords != "" {
		m = m.Where("title like ?", "%"+req.KeyWords+"%")
	}
	if req.Attachment != nil {
		if *req.Attachment == 1 {
			m = m.Where("attachment <> ''")
		} else if *req.Attachment == 0 {
			m = m.Where("attachment = ''")
		}
	}
	if req.Site != "" {
		m = m.Where("site = ?", req.Site)
	}
	total := 0
	err = g.Try(ctx, func(ctx context.Context) {
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order("rank desc,publish desc").Scan(&res.ConsultationInformation)
		log.Println("获取数据列表-------->", res.ConsultationInformation)
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
	})
	res.Total = total
	for i := 0; i < len(res.ConsultationInformation); i++ {
		if res.ConsultationInformation[i].Pdf != "" {
			res.ConsultationInformation[i].Pdf = "https://biaoziku.com/law/" + res.ConsultationInformation[i].Pdf
		} else {
			if res.ConsultationInformation[i].Word != "" {
				list := strings.Split(res.ConsultationInformation[i].Word, ".")
				res.ConsultationInformation[i].Pdf = "https://biaoziku.com/law/" + list[0] + ".pdf"
			}
		}
		if res.ConsultationInformation[i].Word != "" {
			res.ConsultationInformation[i].Word = "https://biaoziku.com/law/" + res.ConsultationInformation[i].Word
		}
	}
	return
}

//InformationGet
func (s *sDataManger) InformationGet(ctx context.Context, req *system.InformationGetReq) (res *system.InformationGetRes, err error) {
	res = new(system.InformationGetRes)
	//m := g.DB("data").Schema("crawldata").Model("consultation")
	//// order := "publish asc"
	//if req.Id != "" {
	//	//bulletin_type  notice_nature
	//	m = m.Where("id = ?", req.Id)
	//}
	//err = g.Try(ctx, func(ctx context.Context) {
	//	err = m.Scan(&res.Information)
	//	log.Println("获取数据-------->", res.Information)
	//	liberr.ErrIsNil(ctx, err, "获取数据失败")
	//})
	m := g.DB("data").Schema("crawldata").Model("consultation")
	//if req.Classification == "法律法规" {
	//	m = g.DB("data").Schema("crawldata").Model("law")
	//}
	if req.Id != "" {
		m = m.Where("id = ?", req.Id)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.Information)
		liberr.ErrIsNil(ctx, err, "获取法律数据数据失败")
	})
	//if res.Information.Pdf != "" {
	//	res.Information.Pdf = "https://biaoziku.com/law/" + res.Information.Pdf
	//} else {
	//	if res.Information.Word != "" {
	//		list := strings.Split(res.Information.Word, ".")
	//		res.Information.Pdf = "https://biaoziku.com/law/" + list[0] + ".pdf"
	//	}
	//}
	//if res.Information.Word != "" {
	//	res.Information.Word = "https://biaoziku.com/law/" + res.Information.Word
	//}
	return
}

// InformationEdit

func (s *sDataManger) InformationEdit(ctx context.Context, req *system.InformationEditReq) (res *system.InformationEditRes, err error) {
	res = new(system.InformationEditRes)
	m := g.DB("data").Schema("crawldata").Model("consultation")
	if req.Id != "" {
		m = m.Where("id = ?", req.Id)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Update(desk_do.Consultation{Type: req.Type})
		liberr.ErrIsNil(ctx, err, "修改数据失败")
	})
	return
}

//InformationDel
func (s *sDataManger) InformationDel(ctx context.Context, req *system.InformationDelReq) (res *system.InformationDelRes, err error) {
	//m := g.DB("data").Schema("crawldata").Model("consultation")
	//m = m.Where("id = ?", req.Id)
	//err = g.Try(ctx, func(ctx context.Context) {
	//	_, err = m.Delete()
	//	liberr.ErrIsNil(ctx, err, "删除数据失败")
	//})
	m := g.DB("data").Schema("crawldata").Model("consultation")
	//if req.Classification == "法律法规" {
	//	m = g.DB("data").Schema("crawldata").Model("law")
	//}
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Delete()
		liberr.ErrIsNil(ctx, err, "删除数据失败")
	})
	return
}

//Statistics

func (s *sDataManger) Statistics(ctx context.Context, req *system.StatisticsReq) (res *system.StatisticsRes, err error) {
	// 招标
	tender_m := g.DB("data").Schema("crawldata").Model("bid")
	// 政策
	consultation_m := g.DB("data").Schema("crawldata").Model("consultation")
	// 知库
	knowledge_m := deskdao.MemberKnowledge.Ctx(ctx)
	// 企业信息
	enterprise_m := dao.SysEnterprise.Ctx(ctx)
	res = new(system.StatisticsRes)
	data := map[string]int64{}
	count := 0
	switch req.Type {
	case 0:
		if req.BulletinType != "" {
			if req.BulletinType == "在建公告" {
				tender_m = tender_m.Where("bulletin_type = ?", "招标公告")
				tender_m = tender_m.Where("bidopening_time > ?", time.Now().Format("2006-01-02"))
			} else {
				tender_m = tender_m.Where("bulletin_type = ?", req.BulletinType)
			}
		}
		if req.IndustryType != "" {
			tender_m = tender_m.Where("industry_classification = ?", req.IndustryType)
		}
		if req.City != "" {
			if strings.HasSuffix(req.City, "省") {
				tender_m = tender_m.Where("province = ?", strings.Replace(req.City, "省", "", -1))
			} else {
				tender_m = tender_m.Where("city = ?", req.City)
			}
		}
		if req.KeyWords != "" {
			tender_m = tender_m.Where(fmt.Sprintf("MATCH(title,announcement_content) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
		}
		if req.Attachment != nil {
			if *req.Attachment == 1 {
				tender_m = tender_m.Where("attachment <> ''")
			} else if *req.Attachment == 0 {
				tender_m = tender_m.Where("attachment = ''")
			}
		}
		if req.OriginalBulletin != "" {
			tender_m = tender_m.Where("original_type = ?", req.OriginalBulletin)
		}
		if req.OriginalIndustry != "" {
			tender_m = tender_m.Where("original_classification = ?", req.OriginalIndustry)
		}
		if req.Site != "" {
			tender_m = tender_m.Where("site = ?", req.Site)
		}
		tender_result, err := tender_m.Fields("COUNT(Id) total,city").Group("city").All()
		liberr.ErrIsNil(ctx, err, "获取招标数据失败")
		if tender_result != nil {
			log.Println("tender_result------------>", tender_result)
			for i := 0; i < len(tender_result); i++ {
				tm := tender_result[i].Map()
				if tm != nil {
					//fmt.Sprintf("%v:%v", arguments["<host>"], arguments["<port>"])
					data[tm["city"].(string)] = tm["total"].(int64)
				}
			}
		}
		_, ok := data["深圳市"]
		if ok {
			v, ok1 := data["深圳"]
			if ok1 {
				data["深圳"] = data["深圳市"] + v
			}
			delete(data, "深圳市")
		}
		for _, v := range data {
			count += int(v)
		}
		res.City = data
		tender_data, consultation_data, knowledge_data, enterprise_data := count, 0, 0, 0
		err = g.Try(ctx, func(ctx context.Context) {
			consultation_data, err = consultation_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			knowledge_data, err = knowledge_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			enterprise_data, err = enterprise_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		res.Plate = map[int]int{0: tender_data, 1: consultation_data, 2: knowledge_data, 3: enterprise_data}
	case 1:
		if req.KeyWords != "" {
			consultation_m = consultation_m.Where("title like ?", "%"+req.KeyWords+"%")
		}
		if req.Classification == "法律法规" {
			consultation_m = g.DB("data").Schema("crawldata").Model("law")
		} else {
			if req.Classification != "" {
				consultation_m = consultation_m.Where("type = ?", req.Classification)
			}
		}
		if req.Attachment != nil {
			if *req.Attachment == 1 {
				consultation_m = consultation_m.Where("attachment <> ''")
			} else if *req.Attachment == 0 {
				consultation_m = consultation_m.Where("attachment = ''")
			}
		}
		if req.Site != "" {
			consultation_m = consultation_m.Where("site = ?", req.Site)
		}
		tender_data, consultation_data, knowledge_data, enterprise_data := 0, 0, 0, 0
		err = g.Try(ctx, func(ctx context.Context) {
			tender_data, err = tender_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			consultation_data, err = consultation_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			knowledge_data, err = knowledge_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			enterprise_data, err = enterprise_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		res.Plate = map[int]int{0: tender_data, 1: consultation_data, 2: knowledge_data, 3: enterprise_data}
	case 2:
		if req.Primary != "" {
			knowledge_m = knowledge_m.Where("primary_classification = ?", req.Primary)
		}
		if req.Secondary != "" {
			knowledge_m = knowledge_m.Where("secondary_classification = ?", req.Secondary)
		}
		if req.KeyWords != "" {
			knowledge_m = knowledge_m.Where("title like ?", "%"+req.KeyWords+"%")
		}
		if req.Attachment != nil {
			if *req.Attachment == 1 {
				knowledge_m = knowledge_m.Where("attachment_url <> ''")
			} else if *req.Attachment == 0 {
				knowledge_m = knowledge_m.Where("attachment_url = ''")
			}
		}
		if req.OriginalIndustry != "" {
			knowledge_m = knowledge_m.Where("original_type = ?", req.OriginalIndustry)
		}
		if req.Site != "" {
			knowledge_m = knowledge_m.Where("site = ?", req.Site)
		}
		tender_data, consultation_data, knowledge_data, enterprise_data := 0, 0, 0, 0
		err = g.Try(ctx, func(ctx context.Context) {
			tender_data, err = tender_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			consultation_data, err = consultation_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			knowledge_data, err = knowledge_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			enterprise_data, err = enterprise_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		res.Plate = map[int]int{0: tender_data, 1: consultation_data, 2: knowledge_data, 3: enterprise_data}
	case 3:
		if req.KeyWords != "" {
			enterprise_m = enterprise_m.Where("name like ?", "%"+req.KeyWords+"%")
		}
		//m = m.Where("review_status = 1")
		enterprise_result, err := enterprise_m.Fields("COUNT(Id) total,location").Group("location").All()
		liberr.ErrIsNil(ctx, err, "获取企业数据失败")
		// 企业信息
		if enterprise_result != nil {
			for i := 0; i < len(enterprise_result); i++ {
				tm := enterprise_result[i].Map()
				if tm != nil {
					v, ok := data[tm["location"].(string)]
					if ok {
						data[tm["location"].(string)] = v + tm["total"].(int64)
					} else {
						data[tm["location"].(string)] = tm["total"].(int64)
					}
					//if data[tm["location"].(string)]
					log.Println(tm["location"].(string), tm["total"].(int64))
				}
			}
		}
		_, ok := data["深圳市"]
		if ok {
			v, ok1 := data["深圳"]
			if ok1 {
				data["深圳"] = data["深圳市"] + v
			}
			delete(data, "深圳市")
		}
		for _, v := range data {
			count += int(v)
		}
		res.City = data
		tender_data, consultation_data, knowledge_data, enterprise_data := 0, 0, 0, count
		err = g.Try(ctx, func(ctx context.Context) {
			tender_data, err = tender_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			consultation_data, err = consultation_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		err = g.Try(ctx, func(ctx context.Context) {
			knowledge_data, err = knowledge_m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		})
		res.Plate = map[int]int{0: tender_data, 1: consultation_data, 2: knowledge_data, 3: enterprise_data}
	}
	err = nil
	return
}

//InformationDisplay
func (s *sDataManger) InformationDisplay(ctx context.Context, req *system.InformationDisplayReq) (res *system.InformationDisplayRes, err error) {
	res = new(system.InformationDisplayRes)
	m := g.DB("data").Schema("crawldata").Model("consultation")
	display := 0
	if req.Switch == 1 {
		display = 1
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Where("id in (?)", req.Ids).Update(desk_do.Consultation{Display: display})
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
	})
	return
}

//KnowledgeDisplay
func (s *sDataManger) KnowledgeDisplay(ctx context.Context, req *system.KnowledgeDisplayReq) (res *system.KnowledgeDisplayRes, err error) {
	res = new(system.KnowledgeDisplayRes)
	display := 0
	if req.Switch == 1 {
		display = 1
	}
	switch req.Type {
	case 0:
		m := g.DB("data").Schema("crawldata").Model("bid")
		_, err = m.Where("id in (?)", req.Ids).Update(desk_do.Bid{Display: display})
		if err != nil {
			log.Println("更新是否隐藏出错")
		}
	case 1:
		m := g.DB("data").Schema("crawldata").Model("consultation")
		_, err = m.Where("id in (?)", req.Ids).Update(desk_do.Consultation{Display: display})
		if err != nil {
			log.Println("更新是否隐藏出错")
		}
	case 2:
		m := deskdao.MemberKnowledge.Ctx(ctx)
		_, err = m.Where("id in (?)", req.Ids).Update(desk_do.MemberKnowledge{Display: display})
		if err != nil {
			log.Println("更新是否隐藏出错")
		}
	}
	return
}

//KnowledgeTypeGet
func (s *sDataManger) KnowledgeTypeGet(ctx context.Context, req *system.KnowledgeTypeGetReq) (res *system.KnowledgeTypeGetRes, err error) {
	res = new(system.KnowledgeTypeGetRes)
	m := deskdao.MemberKnowledge.Ctx(ctx)
	err = g.Try(ctx, func(ctx context.Context) {
		original_type, err := m.Fields("original_type").Distinct().All()
		log.Println(original_type)
		if original_type != nil {
			for _, v := range original_type {
				mp := v.Map()
				if mp != nil {
					if mp["original_type"].(string) != "" {
						res.OriginalType = append(res.OriginalType, mp["original_type"].(string))
					}
				}
			}
		}
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

//Release
func (s *sDataManger) Release(ctx context.Context, req *system.ReleaseReq) (res *system.ReleaseRes, err error) {
	industry_classification := "工程招标"
	if req.IndustryType != "" {
		industry_classification = req.IndustryType
	}
	res = new(system.ReleaseRes)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			m := g.DB("data").Schema("crawldata").Model("bid")
			_, e := m.Ctx(ctx).Insert(desk_do.Bid{
				Title:                  req.Title,
				BulletinType:           "调研公告",
				AnnouncementContent:    req.Content,
				Attachment:             req.Attachment,
				ContactPerson:          service.Context().GetLoginUser(ctx).UserName,
				ReleaseTime:            req.ReleaseTime,
				CreatedAt:              gtime.Now(),
				IndustryClassification: industry_classification,
				Display:                1,
				Crawler:                0,
			})
			liberr.ErrIsNil(ctx, e, "发布调研公告失败")
			if e != nil {
				return
			}
		})
		return err
	})
	return
}

//ReleaseEdit

func (s *sDataManger) ReleaseEdit(ctx context.Context, req *system.ReleaseEditReq) (res *system.ReleaseEditRes, err error) {
	res = new(system.ReleaseEditRes)
	data := desk_do.Bid{}
	if req.Title != nil {
		data.Title = *req.Title
	}
	if req.Attachment != nil {
		data.Attachment = *req.Attachment
	}
	if req.Content != nil {
		data.AnnouncementContent = *req.Content
	}
	if req.ReleaseTime != nil {
		data.ReleaseTime = *req.ReleaseTime
	}
	if req.IndustryType != nil {
		data.IndustryClassification = *req.IndustryType
	}
	data.ContactPerson = service.Context().GetLoginUser(ctx).UserName

	m := g.DB("data").Schema("crawldata").Model("bid")
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := m.Ctx(ctx).Where("id = ?", req.Id).Update(data)
			liberr.ErrIsNil(ctx, e, "更新调研公告失败")
			if e != nil {
				return
			}
		})
		return err
	})
	return
}

//ReleaseList
func (s *sDataManger) ReleaseList(ctx context.Context, req *system.ReleaseListReq) (res *system.ReleaseListRes, err error) {
	bid := [](*deskentity.Bid)(nil)
	res = new(system.ReleaseListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := g.DB("data").Schema("crawldata").Model("bid")
	// order := "release_time desc"
	m = m.Where("bulletin_type = '调研公告' ")
	if req.Title != "" {
		m = m.Where(fmt.Sprintf("MATCH(title,announcement_content) AGAINST('%s*' IN BOOLEAN MODE)", req.Title))
	}
	if req.ContactPerson != "" {
		m = m.Where("contact_person = ?", "%"+req.ContactPerson+"%")
	}
	if len(req.DateRange) > 0 {
		m = m.Where("release_time >=? AND release_time <=?", req.DateRange[0]+" 00:00:00", req.DateRange[1]+" 23:59:59")
	} else {
		if req.Start != "" {
			m = m.Where("release_time >=?", req.Start)
		}
		if req.End != "" {
			m = m.Where("release_time <=?", req.End)
		}
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		err = m.Fields("id,title,release_time,city,bulletin_type,bidopening_time,amount,industry_classification,abstract,rank,attachment,contact_person,display").Page(req.PageNum, req.PageSize).Order("release_day desc,area desc,amount desc,rank desc").Scan(&bid)
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
	})
	for i := 0; i < len(bid); i++ {
		id := bid[i].Id
		err = g.Try(ctx, func(ctx context.Context) {
			total, err := deskdao.Feedback.Ctx(ctx).Where("bulletin_id = ?", id).Count()
			liberr.ErrIsNil(ctx, err, "获取数据列表失败")
			res.LResearchList = append(res.LResearchList, system.LReleaseList{BiddingInformation: bid[i], Count: total})
		})
	}
	return
}

//ReleaseGet
func (s *sDataManger) ReleaseGet(ctx context.Context, req *system.ReleaseGetReq) (res *system.ReleaseGetRes, err error) {
	res = new(system.ReleaseGetRes)
	// m := g.DB("data").Schema("crawldata").Model("bidding_information")
	m := g.DB("data").Schema("crawldata").Model("bid")
	m = m.Where("id = ?", req.Id)
	m = m.Where("bulletin_type = '调研公告' ")
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.BiddingInformation)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

//ReleaseDel

func (s *sDataManger) ReleaseDel(ctx context.Context, req *system.ReleaseDelReq) (res *system.ReleaseDelRes, err error) {
	res = new(system.ReleaseDelRes)
	m := g.DB("data").Schema("crawldata").Model("bid")
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := m.Ctx(ctx).Where("id = ?", req.Id).Delete()
			liberr.ErrIsNil(ctx, e, "删除调研公告失败")
			if e != nil {
				return
			}
		})
		return err
	})
	return
}

//FeedbackList

func (s *sDataManger) FeedbackList(ctx context.Context, req *system.FeedbackListReq) (res *system.FeedbackListRes, err error) {
	res = new(system.FeedbackListRes)
	m := deskdao.Feedback.Ctx(ctx)
	m = m.Where("bulletin_id = ?", req.BulletinId)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			res.Total, err = m.Count()
			liberr.ErrIsNil(ctx, err, "获取数据列表失败")
			e := m.Page(req.PageNum, req.PageSize).Scan(&res.FeedbackList)
			liberr.ErrIsNil(ctx, e, "获取反馈数据失败")
		})
		return err
	})
	return
}

//FeedbackGet

func (s *sDataManger) FeedbackGet(ctx context.Context, req *system.FeedbackGetReq) (res *system.FeedbackGetRes, err error) {
	res = new(system.FeedbackGetRes)
	m := deskdao.Feedback.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			e := m.Scan(&res.Feedback)
			liberr.ErrIsNil(ctx, e, "获取反馈数据失败")
		})
		return err
	})
	return
}

//FeedbackDel

func (s *sDataManger) FeedbackDel(ctx context.Context, req *system.FeedbackDelReq) (res *system.FeedbackDelRes, err error) {
	res = new(system.FeedbackDelRes)
	m := deskdao.Feedback.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := m.Delete()
			liberr.ErrIsNil(ctx, e, "删除反馈数据失败")
		})
		return err
	})
	return
}

//FeedbackReview

func (s *sDataManger) FeedbackReview(ctx context.Context, req *system.FeedbackReviewReq) (res *system.FeedbackReviewRes, err error) {
	res = new(system.FeedbackReviewRes)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := deskdao.Feedback.Ctx(ctx).Where("id = ?", req.Id).Update(desk_do.Feedback{
				Status: req.Status,
				Reject: req.ReviewMessage,
			})
			liberr.ErrIsNil(ctx, e, "更新反馈数据失败")
		})
		return err
	})
	fb := (*deskentity.Feedback)(nil)
	m := deskdao.Feedback.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&fb)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	switch req.Status {
	case 1:
		req.ReviewMessage = "反馈审核通过: " + req.ReviewMessage
	case 2:
		req.ReviewMessage = "反馈审核驳回: " + req.ReviewMessage
	}
	if req.Status != 0 {
		msgId := uuid.New().String()
		_, err = dao.SysWsMsg.Ctx(ctx).Insert(do.SysWsMsg{
			MessageId: msgId,
			UserId:    fb.UserId,
			Content:   req.ReviewMessage,
			IsRead:    0,
			IsDel:     0,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		})
		if err != nil {
			return nil, fmt.Errorf("插入消息失败: %w", err)
		}

		websocket.SendToUser(uint64(fb.UserId), &websocket.WResponse{
			Event: "context",
			Data: map[string]string{
				"messageId": msgId,
				"content":   req.ReviewMessage,
			},
		})
	}
	return
}
