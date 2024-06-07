package index

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"log"
	"strconv"
	"strings"
	"tender/api/v1/desk"
	CommonBidService "tender/internal/app/common/service"
	"tender/internal/app/desk/consts"
	deskdao "tender/internal/app/desk/dao"
	"tender/internal/app/desk/model"
	"tender/internal/app/desk/model/do"
	deskentity "tender/internal/app/desk/model/entity"
	"tender/internal/app/desk/service"
	"tender/internal/app/system/dao"
	"tender/library/liberr"
	"tender/library/libmessage"
	"time"
)

func init() {
	service.RegisterIndexManger(New())
}

func New() *sIndex {
	return &sIndex{}
}

type sIndex struct {
}

// 招标数据
func (s *sIndex) BidSearchList(ctx context.Context, req *desk.BidSearchReq) (res *desk.BidSearchRes, err error) {
	res = new(desk.BidSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	// m := g.DB("data").Schema("crawldata").Model("bidding_information")
	m := g.DB("data").Schema("crawldata").Model("bid")
	//m := dao.BiddingInformation.Ctx(ctx)
	// order := "id"
	if req.KeyWords != "" {
		m = m.Where(fmt.Sprintf("MATCH(title,announcement_content) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		err = m.Page(req.PageNum, req.PageSize).Scan(&res.BiddingInformation)
		log.Println("获取数据列表-------->", res.BiddingInformation)
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
	})
	return
}

// 企业数据
func (s *sIndex) EnterpriseSearchList(ctx context.Context, req *desk.EnterpriseSearchReq) (res *desk.EnterpriseSearchRes, err error) {
	res = new(desk.EnterpriseSearchRes)
	tyc_m := deskdao.Tyc.Ctx(ctx)
	tyc := (*deskentity.Tyc)(nil)
	if req.KeyWords != "" {
		err = tyc_m.Where("name = ? and  number = ? and size = ? and type = ?", req.KeyWords, req.PageNum, req.PageSize, "keywords").Scan(&tyc)
		liberr.ErrIsNil(ctx, err, "获取企业信息数据失败")
		if tyc != nil {
			json.Unmarshal([]byte(tyc.Body), &res.EnterpriseInformation)
		} else {
			EnterpriseInformation := libmessage.OpenTanYanChaSearch(req.KeyWords, strconv.Itoa(req.PageSize), strconv.Itoa(req.PageNum))
			log.Println("企业信息:", EnterpriseInformation)
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				err = g.Try(ctx, func(ctx context.Context) {
					_, e := deskdao.Tyc.Ctx(ctx).TX(tx).Insert(do.Tyc{
						Name:      req.KeyWords,
						Number:    req.PageNum,
						Size:      req.PageSize,
						Type:      "keywords",
						Body:      EnterpriseInformation,
						CreatedAt: gtime.Now(),
					})
					liberr.ErrIsNil(ctx, e, "插入企业信息失败")
					return
				})
				return err
			})
			res.EnterpriseInformation = EnterpriseInformation
		}

	}
	return
}

//EnterpriseCommerce
func (s *sIndex) EnterpriseCommerce(ctx context.Context, req *desk.EnterpriseCommerceReq) (res *desk.EnterpriseCommerceRes, err error) {
	res = new(desk.EnterpriseCommerceRes)
	tyc_m := deskdao.Tyc.Ctx(ctx)
	tyc := (*deskentity.Tyc)(nil)
	if req.Name != nil {
		err = tyc_m.Where("name = ? and type = ?", *req.Name, "commerce").Scan(&tyc)
		liberr.ErrIsNil(ctx, err, "获取工商和主要人员数据失败")
		if tyc != nil {
			json.Unmarshal([]byte(tyc.Body), &res.Commerce)
		} else {
			Commerce, code := libmessage.IndustryCommerceAndPersonnel(*req.Name)
			log.Println("工商信息:", Commerce)
			if code == 0 {
				err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					err = g.Try(ctx, func(ctx context.Context) {
						_, e := deskdao.Tyc.Ctx(ctx).TX(tx).Insert(do.Tyc{
							Name:      *req.Name,
							Type:      "commerce",
							Body:      Commerce,
							CreatedAt: gtime.Now(),
						})
						liberr.ErrIsNil(ctx, e, "插入企业信息失败")
						return
					})
					return err
				})
			}
			res.Commerce = Commerce
		}

	}
	return
}

//EnterprisePunishment
func (s *sIndex) EnterprisePunishment(ctx context.Context, req *desk.EnterprisePunishmentReq) (res *desk.EnterprisePunishmentRes, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	res = new(desk.EnterprisePunishmentRes)
	tyc_m := deskdao.Tyc.Ctx(ctx)
	tyc := (*deskentity.Tyc)(nil)
	if req.Name != nil {
		err = tyc_m.Where("name = ? and  number = ? and size = ? and type = ?", *req.Name, req.PageNum, req.PageSize, "punishment").Scan(&tyc)
		liberr.ErrIsNil(ctx, err, "获取经营风险数据失败")
		if tyc != nil {
			json.Unmarshal([]byte(tyc.Body), &res.Punishment)
		} else {
			Punishment, code := libmessage.PunishmentInfo(*req.Name, strconv.Itoa(req.PageSize), strconv.Itoa(req.PageNum))
			if code == 0 {
				log.Println("经营风险:", Punishment)
				err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					err = g.Try(ctx, func(ctx context.Context) {
						_, e := deskdao.Tyc.Ctx(ctx).TX(tx).Insert(do.Tyc{
							Name:      *req.Name,
							Number:    req.PageNum,
							Size:      req.PageSize,
							Type:      "punishment",
							Body:      Punishment,
							CreatedAt: gtime.Now(),
						})
						liberr.ErrIsNil(ctx, e, "插入企业信息失败")
						return
					})
					return err
				})
			}
			res.Punishment = Punishment
		}
	}
	return
}

//EnterpriseQualification
func (s *sIndex) EnterpriseQualification(ctx context.Context, req *desk.EnterpriseQualificationReq) (res *desk.EnterpriseQualificationRes, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	res = new(desk.EnterpriseQualificationRes)
	tyc_m := deskdao.Tyc.Ctx(ctx)
	tyc := (*deskentity.Tyc)(nil)
	if req.Name != nil {
		err = tyc_m.Where("name = ? and  number = ? and size = ? and type = ?", *req.Name, req.PageNum, req.PageSize, "qualification").Scan(&tyc)
		liberr.ErrIsNil(ctx, err, "获取资质数据失败")
		if tyc != nil {
			json.Unmarshal([]byte(tyc.Body), &res.Qualification)
		} else {
			Qualification, code := libmessage.Qualification(*req.Name, strconv.Itoa(req.PageSize), strconv.Itoa(req.PageNum))
			log.Println("建筑资质:", Qualification)
			if code == 0 {
				err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					err = g.Try(ctx, func(ctx context.Context) {
						_, e := deskdao.Tyc.Ctx(ctx).TX(tx).Insert(do.Tyc{
							Name:      *req.Name,
							Number:    req.PageNum,
							Size:      req.PageSize,
							Type:      "qualification",
							Body:      Qualification,
							CreatedAt: gtime.Now(),
						})
						liberr.ErrIsNil(ctx, e, "插入企业信息失败")
						return
					})
					return err
				})
			}
			res.Qualification = Qualification
		}
	}
	return
}

//EnterpriseLawSuit
func (s *sIndex) EnterpriseLawSuit(ctx context.Context, req *desk.EnterpriseLawSuitReq) (res *desk.EnterpriseLawSuitRes, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	res = new(desk.EnterpriseLawSuitRes)
	tyc_m := deskdao.Tyc.Ctx(ctx)
	tyc := (*deskentity.Tyc)(nil)
	if req.Name != nil {
		err = tyc_m.Where("name = ? and  number = ? and size = ? and type = ?", *req.Name, req.PageNum, req.PageSize, "lawsuit").Scan(&tyc)
		liberr.ErrIsNil(ctx, err, "获取司法风险数据失败")
		if tyc != nil {
			json.Unmarshal([]byte(tyc.Body), &res.LawSuit)
		} else {
			LawSuit, code := libmessage.LawSuit(*req.Name, strconv.Itoa(req.PageSize), strconv.Itoa(req.PageNum))
			log.Println("司法风险:", LawSuit)
			if code == 0 {
				err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					err = g.Try(ctx, func(ctx context.Context) {
						_, e := deskdao.Tyc.Ctx(ctx).TX(tx).Insert(do.Tyc{
							Name:      *req.Name,
							Number:    req.PageNum,
							Size:      req.PageSize,
							Type:      "lawsuit",
							Body:      LawSuit,
							CreatedAt: gtime.Now(),
						})
						liberr.ErrIsNil(ctx, e, "插入企业信息失败")
						return
					})
					return err
				})
			}
			res.LawSuit = LawSuit
		}

	}
	return
}

//EnterpriseBidding
func (s *sIndex) EnterpriseBidding(ctx context.Context, req *desk.EnterpriseBiddingReq) (res *desk.EnterpriseBiddingRes, err error) {
	res = new(desk.EnterpriseBiddingRes)
	bidding := [](*model.BiddingEnterprise)(nil)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	if req.Name != nil {
		bidding_m := g.DB("data").Schema("crawldata").Model("bid")
		bidding_m = bidding_m.Where("bulletin_type = ? ", "中标公告")
		bidding_m = bidding_m.Where(fmt.Sprintf("MATCH(win_name) AGAINST('%s*' IN BOOLEAN MODE)", *req.Name))
		res.Total, err = bidding_m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		err = bidding_m.Page(req.PageNum, req.PageSize).Order("release_time desc").Scan(&bidding)
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
		res.Bidding = bidding
	}
	return
}

// 最新公告
//AnnouncementList
func (s *sIndex) AnnouncementList(ctx context.Context, req *desk.AnnouncementListReq) (res *desk.AnnouncementListRes, err error) {
	res = new(desk.AnnouncementListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	// m := g.DB("data").Schema("crawldata").Model("bidding_information")
	m := g.DB("data").Schema("crawldata").Model("bid")
	// order := "release_time desc"
	//if req.Type != "" {
	//	//bulletin_type  notice_nature
	//	if req.Type == "在建公告" {
	//		// 对于长度不同的字符串进行匹配时。较短字符串比较完后还没有大小之分。则较长的字符串较大
	//		m = m.Where("notice_nature = '正常公告'")
	//		m = m.Where("bidopening_time > ?", time.Now().Format("2006-01-02"))
	//	} else if req.Type == "招标公告" {
	//		m = m.Where("notice_nature = '正常公告'")
	//		m = m.Where("bidopening_time < ?", time.Now().Format("2006-01-02"))
	//	} else if req.Type == "中标公告" {
	//		m = m.Where("notice_nature = '中标公告'")
	//	} else if req.Type == "其他公告" {
	//		m = m.Where("notice_nature = '更正公告'")
	//	} else {
	//		m = m.Where("bulletin_type = ?", req.Type)
	//	}
	//}

	if req.Type != "" {
		if req.Type == "在建公告" {
			m = m.Where("bulletin_type = ?", "招标公告")
			m = m.Where("bidopening_time > ?", time.Now().Format("2006-01-02"))
			//} else if req.Type == "招标公告" {
			//	m = m.Where("bulletin_type = ?", "招标公告")
			//	m = m.Where("bidopening_time < ?", time.Now().Format("2006-01-02"))
		} else {
			m = m.Where("bulletin_type = ?", req.Type)
		}
	}
	if req.Industry != "" {
		m = m.Where("industry_classification = ?", req.Industry)
	}
	if req.KeyWords != "" {
		// MATCH(column_name1,column_name2) AGAINST('keyword' IN BOOLEAN MODE)
		// m = m.Where("title = ?", req.KeyWords)
		m = m.Where(fmt.Sprintf("MATCH(title,announcement_content) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
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
	if req.City != "" {
		if strings.HasSuffix(req.City, "省") {
			m = m.Where("province = ?", strings.Replace(req.City, "省", "", -1))
		} else {
			m = m.Where("city = ?", req.City)
		}

	}
	m = m.Where("display = 1")
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		if req.Type == "中标公告" {
			err = m.Fields("id,title,release_day,release_time,city,bulletin_type,bidopening_time,bid_amount,industry_classification,abstract,rank,crawler").Page(req.PageNum, req.PageSize).Order("release_day desc,area desc,bid_amount desc,rank desc").Scan(&res.BiddingInformation)
			liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
		} else {
			err = m.Fields("id,title,release_day,release_time,city,bulletin_type,bidopening_time,amount,industry_classification,abstract,rank,crawler").Page(req.PageNum, req.PageSize).Order("release_day desc,area desc,amount desc,rank desc").Scan(&res.BiddingInformation)
			liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
		}
	})
	if req.Type == "在建公告" {
		for i := 0; i < len(res.BiddingInformation); i++ {
			res.BiddingInformation[i].BulletinType = "在建公告"
		}
	}
	return
}

//Announcement 获取单个数据
func (s *sIndex) Announcement(ctx context.Context, req *desk.AnnouncementReq) (res *desk.AnnouncementRes, err error) {
	res = new(desk.AnnouncementRes)
	m := g.DB("data").Schema("crawldata").Model("bid")
	if req.Id != "" {
		m = m.Where("id = ?", req.Id)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.BiddingInformation)
		liberr.ErrIsNil(ctx, err, "获取数据数据失败")
	})
	return
}

//ConsulList
// 咨询
func (s *sIndex) ConsulList(ctx context.Context, req *desk.ConsulListReq) (res *desk.ConsulListRes, err error) {
	res = new(desk.ConsulListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	// m := g.DB("data").Schema("crawldata").Model("bidding_information")
	m := g.DB("data").Schema("crawldata").Model("bid")
	//m := dao.BiddingInformation.Ctx(ctx)
	// order := "release_time asc"
	// 服务类型
	if req.Type != "" {
		//bulletin_type  notice_nature
		if req.Type == "投标咨询" {
			req.Type = "招标公告"
		}
		if req.Type == "采购咨询" {
			req.Type = "采购公告"
		}
		//if req.Type == "采购咨询" {
		//	req.Type = "采购公告"
		//}
		m = m.Where("bulletin_type = ?", req.Type)
	}
	// 行业类型
	if req.Industry != "" {
		if req.Industry == "工程建筑" {
			req.Industry = "工程建设"
		}
		m = m.Where("industry_classification = ?", req.Industry)
	}
	if req.KeyWords != "" {
		// m = m.Where("title = ?", req.KeyWords)
		m = m.Where(fmt.Sprintf("MATCH(title,announcement_content) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
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
	if req.City != "" {
		m = m.Where("city like ?", req.City+"%")
	}
	// 对于长度不同的字符串进行匹配时。较短字符串比较完后还没有大小之分。则较长的字符串较大
	// 咨询服务 是开标时间大于当前时间的招标数据
	m = m.Where("bidopening_time > ?", time.Now().Format("2006-01-02"))
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		//m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.BiddingInformation)
		err = m.Page(req.PageNum, req.PageSize).Order("release_time desc").Scan(&res.BiddingInformation)
		log.Println("获取数据列表-------->", res.BiddingInformation)
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
	})
	return
}

//SettingGet
func (s *sIndex) SettingGet(ctx context.Context, req *desk.SettingGetReq) (res *desk.SettingGetRes, err error) {
	res = new(desk.SettingGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDataset.Ctx(ctx).Scan(&res.Result)
		liberr.ErrIsNil(ctx, err, "获取数据中心设置失败")
	})
	return
}

// 咨询服务非自营 Consul
func (s *sIndex) Consul(ctx context.Context, req *desk.ConsulReq) (res *desk.ConsulRes, err error) {
	res = new(desk.ConsulRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysEnterprise.Ctx(ctx)
	if req.Industry != "" {
		m = m.Where("industry = ?", req.Industry)
	}
	if req.City != "" {
		m = m.Where("location = ?", req.City)
	}
	// order := "id"
	if req.KeyWords != "" {
		m = m.Where("name like ?", "%"+req.KeyWords+"%")
	}
	if req.Start != "" {
		start, _ := strconv.Atoi(req.Start)
		m = m.Where("establishment_at < ?", gtime.Now().AddDate(-start, 0, 0))
	}
	if req.End != "" {
		end, _ := strconv.Atoi(req.End)
		m = m.Where("establishment_at > ?", gtime.Now().AddDate(-end, 0, 0))
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取列表失败")
		err = m.Page(req.PageNum, req.PageSize).Scan(&res.EnterpriseInformation)
		log.Println("获取列表-------->", res.EnterpriseInformation)
		liberr.ErrIsNil(ctx, err, "获取列表数据失败")
	})
	return
}

//ConsultationList
// 政策资讯
func (s *sIndex) ConsultationList(ctx context.Context, req *desk.ConsultationListReq) (res *desk.ConsultationListRes, err error) {
	res = new(desk.ConsultationListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := g.DB("data").Schema("crawldata").Model("consultation")
	// order := "publish asc"
	//if req.Classification == "法律法规" {
	//	if req.Type == "常用" {
	//		m = dao.Law.Ctx(ctx)
	//	} else {
	//		m = g.DB("data").Schema("crawldata").Model("law")
	//	}
	//} else {
	//	if req.Classification != "" {
	//		m = m.Where("type = ?", req.Classification)
	//	}
	//}
	if req.Type == "常用" {
		m = deskdao.Consultation.Ctx(ctx)
	}
	if req.Classification != "" {
		m = m.Where("type = ?", req.Classification)
	}
	if req.KeyWords != "" {
		m = m.Where("title like ?", "%"+req.KeyWords+"%")
	}
	if len(req.DateRange) > 0 {
		m = m.Where("publish >=? AND publish <=?", req.DateRange[0]+" 00:00:00", req.DateRange[1]+" 23:59:59")
	} else {
		if req.Start != "" {
			m = m.Where("publish >=?", req.Start)
		}
		if req.End != "" {
			m = m.Where("publish <=?", req.End)
		}
	}
	m = m.Where("display = ?", 1)
	order := "publish "
	if req.OrderBy != "" {
		order = order + req.OrderBy
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.ConsultationInformation)
		log.Println("获取数据列表-------->", res.ConsultationInformation)
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
	})
	return
}

// 单个政策资讯
func (s *sIndex) Consultation(ctx context.Context, req *desk.ConsultationReq) (res *desk.ConsultationRes, err error) {
	res = new(desk.ConsultationRes)
	m := g.DB("data").Schema("crawldata").Model("consultation")
	// order := "publish asc"
	// order := "publish asc"
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
	if req.Id != "" {
		//bulletin_type  notice_nature
		m = m.Where("id = ?", req.Id)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&res.ConsultationInformation)
		log.Println("获取数据-------->", res.ConsultationInformation)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	if err != nil {
		return
	}
	//if res.ConsultationInformation.Pdf != "" {
	//	res.ConsultationInformation.Pdf = "https://biaoziku.com/law/" + res.ConsultationInformation.Pdf
	//} else {
	//	if res.ConsultationInformation.Word != "" {
	//		list := strings.Split(res.ConsultationInformation.Word, ".")
	//		res.ConsultationInformation.Pdf = "https://biaoziku.com/law/" + list[0] + ".pdf"
	//	}
	//}
	//if res.ConsultationInformation.Word != "" {
	//	res.ConsultationInformation.Word = "https://biaoziku.com/law/" + res.ConsultationInformation.Word
	//}
	return
}

//KnowledgeList
func (s *sIndex) KnowledgeList(ctx context.Context, req *desk.KnowledgeListReq) (res *desk.KnowledgeListRes, err error) {
	MemberKnowledgeList := ([]*deskentity.MemberKnowledge)(nil)
	res = new(desk.KnowledgeListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := deskdao.MemberKnowledge.Ctx(ctx)
	if req.Type != "" {
		m = m.Where("primary_classification = ?", req.Type)
	}
	if req.Classification != "" {
		m = m.Where("secondary_classification = ?", req.Classification)
	}
	if req.KeyWords != "" {
		// m = m.Where("title like ?", "%"+req.KeyWords+"%")
		m = m.Where(fmt.Sprintf("MATCH(title,abstract) AGAINST('%s*' IN BOOLEAN MODE)", req.KeyWords))
	}
	if req.ArticleType != "" {

	}
	m = m.Where("review_status = ?", 1)
	m = m.Where("display = ?", 1)
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order("created_at desc").Scan(&MemberKnowledgeList)
		log.Println("获取数据列表-------->", &MemberKnowledgeList)
		liberr.ErrIsNil(ctx, err, "获取数据列表数据失败")
	})
	if MemberKnowledgeList != nil {
		for i := 0; i < len(MemberKnowledgeList); i++ {
			var count int
			m_collect := deskdao.MemberCollect.Ctx(ctx)
			if req.ArticleType != "" {
				m_collect = m_collect.Where("type = ?", req.ArticleType)
			}
			m_collect = m_collect.Where("article_id = ?", MemberKnowledgeList[i].Id)
			err = g.Try(ctx, func(ctx context.Context) {
				count, err = m_collect.Count()
				liberr.ErrIsNil(ctx, err, "获取数据失败")
			})
			res.MemberKnowledge = append(res.MemberKnowledge, desk.Knowledge{
				Knowledge: MemberKnowledgeList[i],
				Count:     count,
			})
		}
	}
	return
}
func (s *sIndex) KnowledgeBrowse(ctx context.Context, req *desk.KnowledgeBrowseReq) (res *desk.KnowledgeBrowseRes, err error) {
	knowledge := (*deskentity.MemberKnowledge)(nil)
	m := deskdao.MemberKnowledge.Ctx(ctx)
	m = m.Where("id = ?", req.KnowledgeId)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&knowledge)
		liberr.ErrIsNil(ctx, err, "获取知识数据失败")
	})
	if err != nil {
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = m.Update(do.MemberKnowledge{
			Views: knowledge.Views + 1,
		})
		liberr.ErrIsNil(ctx, err, "更新知识数据失败")
	})
	return
}

//ConsultationBrowse
func (s *sIndex) StatisticsBrowse(ctx context.Context, req *desk.StatisticsBrowseReq) (res *desk.StatisticsBrowseRes, err error) {
	statistics := (*deskentity.Statistics)(nil)
	m := deskdao.Statistics.Ctx(ctx)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Limit(1).Scan(&statistics)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	if err != nil {
		return
	}
	data := do.Statistics{}
	if statistics != nil {
		if req.Type == 0 {
			data.Consultation = statistics.Consultation + 1
		}
		if req.Type == 1 {
			data.Knowledge = statistics.Knowledge + 1
		}
		err = g.Try(ctx, func(ctx context.Context) {
			log.Println(err)
			_, err = m.Where("id = ?", statistics.Id).Update(data)
			liberr.ErrIsNil(ctx, err, "更新数据失败")
		})
	} else {
		if req.Type == 0 {
			data.Consultation = 1
			data.Knowledge = 0
		}
		if req.Type == 1 {
			data.Knowledge = 1
			data.Consultation = 0
		}
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = m.Insert(data)
			liberr.ErrIsNil(ctx, err, "插入数据失败")
		})
	}
	return
}

//StatisticsGet
func (s *sIndex) StatisticsGet(ctx context.Context, req *desk.StatisticsGetReq) (res *desk.StatisticsGetRes, err error) {
	res = new(desk.StatisticsGetRes)
	statistics := (*deskentity.Statistics)(nil)
	// 招标
	// bidding_information
	tender_m := g.DB("data").Schema("crawldata").Model("bid")
	// 企业信息
	enterprise_m := dao.SysEnterprise.Ctx(ctx)
	// 	今日
	statistics_m := deskdao.Statistics.Ctx(ctx)
	information_m := g.DB("data").Schema("crawldata").Model("consultation")
	// 招标总数
	tender, tender_count, information := 0, 0, 0
	enterprise_count := 0
	// 历史招标
	err = g.Try(ctx, func(ctx context.Context) {
		tender_count, _ = tender_m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	//政策咨询数量
	err = g.Try(ctx, func(ctx context.Context) {
		information, err = information_m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据列表失败")
	})
	// 今日招标
	err = g.Try(ctx, func(ctx context.Context) {
		tender_m = tender_m.Where("release_time > ?", time.Now().Format("2006-01-02"))
		tender, _ = tender_m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	// 今日企业
	err = g.Try(ctx, func(ctx context.Context) {
		enterprise_m = enterprise_m.Where("created_at > ?", time.Now().Format("2006-01-02"))
		enterprise_count, _ = enterprise_m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	err = g.Try(ctx, func(ctx context.Context) {
		err = statistics_m.Scan(&statistics)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	consultation_count := 0
	knowledge_count := 0
	if statistics != nil {
		consultation_count = statistics.Consultation
		knowledge_count = statistics.Knowledge
	}
	//爬取网站
	m := g.DB("data").Schema("crawldata").Model("bid")
	m.Fields("site").Where("site <> ''").Distinct()
	site_count := 0
	err = g.Try(ctx, func(ctx context.Context) {
		site_count, _ = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	res.HisConsultationCount = consultation_count
	res.HisDownload = knowledge_count
	res.TodayEnterprise = enterprise_count
	res.TodayTenderCount = tender
	res.HisTenderCount = tender_count
	res.Website = site_count + 3
	res.InformationCount = information
	return
}

//MemberInFind
func (s *sIndex) MemberInFind(ctx context.Context, req *desk.MemberInFindReq) (res *desk.MemberInFindRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.MemberIntegral.Ctx(ctx).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取积分信息失败")
	})
	return
}

//Search
func (s *sIndex) Search(ctx context.Context, req *desk.SearchReq) (res *desk.SearchRes, err error) {
	res = new(desk.SearchRes)
	type q struct {
		Query map[string]map[string]string
		From  int
		Size  int
	}
	// q = {"query": {"query_string": {"query": "深圳网"}}, "from": 0, "size": 20}
	var result interface{}
	result, err = CommonBidService.BD.Query("crawler", &q{Query: map[string]map[string]string{"query_string": {"query": fmt.Sprintf("%s", req.KeyWords)}},
		From: req.PageNum, Size: req.PageSize})
	res.Result = result
	return
}
