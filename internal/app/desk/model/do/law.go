// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Law is the golang structure of table law for DAO operations like Where/Data.
type Law struct {
	g.Meta     `orm:"table:law, do:true"`
	Id         interface{} // UID
	Fileid     interface{} // 文件id
	Publish    interface{} // 公布日期
	Expiry     interface{} // 施行日期
	Office     interface{} // 制定机关
	Title      interface{} // 标题
	Type       interface{} // 类型
	Word       interface{} // word
	Pdf        interface{} // pdf
	Url        interface{} // 外部链接
	Status     interface{} // 0 无效,1 有效 时效性
	Display    interface{} // 是否展示,0不展示1展示
	Content    interface{} // 内容
	Level      interface{} // 法律效力位阶
	Attachment interface{} // 附件
	Height     interface{} // 0高1中2低
	Rank       interface{} // 排序
	Crawler    interface{} // 是否爬取,0 非爬取 1 爬取
}
