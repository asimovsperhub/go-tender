// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BbsContent is the golang structure for table bbs_content.
type BbsContent struct {
	Id      uint64 `json:"id"      description:""`
	BbsId   int64  `json:"bbsId"   description:"论坛id"`
	Content string `json:"content" description:"内容"`
}
