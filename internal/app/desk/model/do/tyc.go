// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tyc is the golang structure of table tyc for DAO operations like Where/Data.
type Tyc struct {
	g.Meta    `orm:"table:tyc, do:true"`
	Id        interface{} //
	Name      interface{} // 企业名称
	Type      interface{} // 类型:commerce-工商信息,punishment-经营风险,qualification-资质,lawSuit-司法风险
	Number    interface{} // 页吗
	Size      interface{} // 页大小
	Body      interface{} // 内容
	CreatedAt *gtime.Time // 创建日期
	UpdatedAt *gtime.Time // 修改日期
	DeletedAt *gtime.Time // 删除时间
}
