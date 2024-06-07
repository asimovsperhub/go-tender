package sys_finance

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	"tender/api/v1/system"
	"tender/internal/app/system/consts"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/library/libUtils"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
	"tender/library/liberr"
)

func init() {
	service.RegisterFinance(New())
}

func New() *sSysFinance {
	return &sSysFinance{}
}

type sSysFinance struct {
}

func (s *sSysFinance) QueryPaySettings(ctx context.Context, req *system.PaySettingsReq) (res *system.PaySettingsRes, err error) {
	res = new(system.PaySettingsRes)

	var settings *entity.PaySettings
	err = dao.PaySettings.Ctx(ctx).Scan(&settings)
	if err != nil {
		return nil, err
	}

	res.Settings = settings
	return
}

func (s *sSysFinance) UpdatePaySettings(ctx context.Context, req *system.UpdatePaySettingsReq) (res *system.UpdatePaySettingsRes, err error) {

	var settings *entity.PaySettings
	err = dao.PaySettings.Ctx(ctx).Scan(&settings)
	if err != nil {
		return nil, err
	}
	data := do.PaySettings{}
	if req.WeixinMchid != nil {
		data.WeixinMchid = *req.WeixinMchid
	}
	if req.WeixinAppid != nil {
		data.WeixinAppid = *req.WeixinAppid
	}
	if req.WeixinApikey != nil {
		data.WeixinApikey = *req.WeixinApikey
	}
	if req.WeixinSerialno != nil {
		data.WeixinSerialno = *req.WeixinSerialno
	}
	if req.WeixinPrivatekey != nil {
		data.WeixinPrivatekey = *req.WeixinPrivatekey
	}
	if req.AlipayAppid != nil {
		data.AlipayAppid = *req.AlipayAppid
	}
	if req.AlipayPrivatekey != nil {
		data.AlipayPrivatekey = *req.AlipayPrivatekey
	}
	if req.AlipayAppCertPublicKey != nil {
		data.AlipayAppCertPublicKey = *req.AlipayAppCertPublicKey
	}
	if req.AlipayRootCert != nil {
		data.AlipayRootCert = *req.AlipayRootCert
	}
	if req.AlipayPublicCert != nil {
		data.AlipayPublicCert = *req.AlipayPublicCert
	}
	if settings.Id > 0 {
		dao.PaySettings.Ctx(ctx).Where("id = ?", settings.Id).Update(data)
	} else {
		dao.PaySettings.Ctx(ctx).Insert(do.PaySettings{
			WeixinMchid:            *req.WeixinMchid,
			WeixinAppid:            *req.WeixinAppid,
			WeixinApikey:           *req.WeixinApikey,
			WeixinSerialno:         *req.WeixinSerialno,
			WeixinPrivatekey:       *req.WeixinPrivatekey,
			AlipayAppid:            *req.AlipayAppid,
			AlipayPrivatekey:       *req.AlipayPrivatekey,
			AlipayAppCertPublicKey: *req.AlipayAppCertPublicKey,
			AlipayRootCert:         *req.AlipayRootCert,
			AlipayPublicCert:       *req.AlipayPublicCert,
		})
	}

	return
}

func (s *sSysFinance) PurchaseList(ctx context.Context, req *system.FinanceReq) (res *system.FinanceRes, err error) {
	res = new(system.FinanceRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.PurchaseLog.Ctx(ctx)
	if req.Mobile != "" {
		m = m.Where("mobile like ?", "%"+req.Mobile+"%")
	}
	if len(req.DateRange) > 0 {
		m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0]+" 00:00:00", req.DateRange[1]+" 23:59:59")
	}
	order := "created_at DESC"
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取最新列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.PurchaseLogList)
		log.Println("获取最新列表-------->", res.PurchaseLogList)
		liberr.ErrIsNil(ctx, err, "获取最新列表数据失败")
	})
	return
}

func (s *sSysFinance) PurchaseInfo(ctx context.Context, req *system.PurchaseInfoReq) (res *system.PurchaseInfoRes, err error) {
	res = new(system.PurchaseInfoRes)

	m := dao.PurchaseLog.Ctx(ctx)

	todayAmount, _ := m.Where("created_at >=? AND created_at <=?", libUtils.GetDayFirstTime(), libUtils.GetDayLastTime()).Sum("amount")
	res.TodayAmount = float64(todayAmount) / 100

	weekStartTime, weekEndTime := libUtils.WeekIntervalTime(0)
	thisWeekAmount, _ := m.Where("created_at >=? AND created_at <=?", libUtils.ConvertTsToDate(weekStartTime), libUtils.ConvertTsToDate(weekEndTime)).Sum("amount")
	res.ThisWeekAmount = float64(thisWeekAmount) / 100

	monthStartTime, monthEndTime := libUtils.MonthIntervalTime(0)

	thisMonthAmount, _ := m.Where("created_at >=? AND created_at <=?", libUtils.ConvertTsToDate(monthStartTime), libUtils.ConvertTsToDate(monthEndTime)).Sum("amount")
	res.ThisMonthAmount = float64(thisMonthAmount) / 100
	return
}

func (s *sSysFinance) TrendingInfo(ctx context.Context, req *system.TrendingInfoReq) (res *system.TrendingInfoRes, err error) {
	res = new(system.TrendingInfoRes)

	searchType := req.Type

	if searchType == "day" {
		m := dao.PurchaseLog.Ctx(ctx)
		if req.Start != "" {
			m = m.Where("DATE_FORMAT(created_at, '%Y-%m-%d') >=?", req.Start)
		}
		if req.End != "" {
			m = m.Where("DATE_FORMAT(created_at, '%Y-%m-%d') <=?", req.End)
		}
		result, err := m.Fields("DATE_FORMAT(created_at, '%Y-%m-%d') as date, sum(amount) as total_amount").Group("date").Order("date").All()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println(result)
		var resultList []*system.TrendingItem
		for _, item := range result {
			resultItem := system.TrendingItem{}
			row := item.Map()
			resultItem.Name = row["date"].(string)
			resultItem.Value = row["total_amount"].(float64)
			resultList = append(resultList, &resultItem)
		}

		res.TrendingItemList = resultList
	} else if searchType == "month" {
		m := dao.PurchaseLog.Ctx(ctx)
		if req.Start != "" {
			m = m.Where("DATE_FORMAT(created_at, '%Y-%m') >=?", req.Start)
		}
		if req.End != "" {
			m = m.Where("DATE_FORMAT(created_at, '%Y-%m') <=?", req.End)
		}
		result, err := m.Fields("DATE_FORMAT(created_at, '%Y-%m') as date, sum(amount) as total_amount").Group("date").Order("date").All()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println(result)
		var resultList []*system.TrendingItem
		for _, item := range result {
			resultItem := system.TrendingItem{}
			row := item.Map()
			resultItem.Name = row["date"].(string)
			resultItem.Value = row["total_amount"].(float64)
			resultList = append(resultList, &resultItem)
		}

		res.TrendingItemList = resultList
	}

	return
}
