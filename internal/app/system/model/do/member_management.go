// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MemberManagement is the golang structure of table member_management for DAO operations like Where/Data.
type MemberManagement struct {
	g.Meta                  `orm:"table:member_management, do:true"`
	Id                      interface{} //
	MonthlycardOriginal     interface{} // 月卡原价
	MonthlycardCurrent      interface{} // 月卡现价
	QuartercardOriginal     interface{} // 季卡原价
	QuartercardCurrent      interface{} // 季卡现价
	AnnualcardOriginal      interface{} // 年卡原价
	AnnualcardCurrent       interface{} // 年卡现价
	DownloadKnowledge       interface{} // 下载知识
	DownloadVideo           interface{} // 下载视频
	MonthlycardIntegral     interface{} // 月卡积分
	QuartercardIntegral     interface{} // 季卡积分
	AnnualcardIntegral      interface{} // 年卡积分
	KnowledgeIntegral       interface{} // 发布知识积分
	VideoIntegral           interface{} // 发布视频积分
	IssueIntegral           interface{} // 发行积分
	MonthlycardSubscription interface{} // 月卡可订阅数
	QuartercardSubscription interface{} // 季卡可订阅数
	AnnualcardSubscription  interface{} // 年卡可订阅数
	NewSubscription         interface{} // 新购买可订阅数
}
