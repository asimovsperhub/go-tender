// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MemberManagement is the golang structure for table member_management.
type MemberManagement struct {
	Id                      uint64 `json:"id"                      description:""`
	MonthlycardOriginal     string `json:"monthlycardOriginal"     description:"月卡原价"`
	MonthlycardCurrent      string `json:"monthlycardCurrent"      description:"月卡现价"`
	QuartercardOriginal     string `json:"quartercardOriginal"     description:"季卡原价"`
	QuartercardCurrent      string `json:"quartercardCurrent"      description:"季卡现价"`
	AnnualcardOriginal      string `json:"annualcardOriginal"      description:"年卡原价"`
	AnnualcardCurrent       string `json:"annualcardCurrent"       description:"年卡现价"`
	DownloadKnowledge       string `json:"downloadKnowledge"       description:"下载知识"`
	DownloadVideo           string `json:"downloadVideo"           description:"下载视频"`
	MonthlycardIntegral     int    `json:"monthlycardIntegral"     description:"月卡积分"`
	QuartercardIntegral     int    `json:"quartercardIntegral"     description:"季卡积分"`
	AnnualcardIntegral      int    `json:"annualcardIntegral"      description:"年卡积分"`
	KnowledgeIntegral       int    `json:"knowledgeIntegral"       description:"发布知识积分"`
	VideoIntegral           int    `json:"videoIntegral"           description:"发布视频积分"`
	IssueIntegral           int    `json:"issueIntegral"           description:"发行积分"`
	MonthlycardSubscription int    `json:"monthlycardSubscription" description:"月卡可订阅数"`
	QuartercardSubscription int    `json:"quartercardSubscription" description:"季卡可订阅数"`
	AnnualcardSubscription  int    `json:"annualcardSubscription"  description:"年卡可订阅数"`
	NewSubscription         int    `json:"newSubscription"         description:"新购买可订阅数"`
}
