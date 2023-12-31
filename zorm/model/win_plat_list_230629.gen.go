// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinPlatList230629 = "win_plat_list_230629"

// WinPlatList230629 mapped from table <win_plat_list_230629>
type WinPlatList230629 struct {
	ID        int64  `gorm:"column:id;type:int;not null" json:"id,string"`
	Code      string `gorm:"column:code;type:varchar(64);not null;comment:平台编码" json:"code"`                  // 平台编码
	Name      string `gorm:"column:name;type:varchar(64);not null;comment:平台名称" json:"name"`                  // 平台名称
	Config    string `gorm:"column:config;type:longtext;comment:配置信息" json:"config"`                          // 配置信息
	Rate      string `gorm:"column:rate;type:longtext;comment:费率" json:"rate"`                                // 费率
	Sort      int64  `gorm:"column:sort;type:tinyint;not null;default:100;comment:排序(从高到底、ID降序)" json:"sort"` // 排序(从高到底、ID降序)
	Status    int64  `gorm:"column:status;type:tinyint;not null;comment:状态: 1-启用 0-停用" json:"status"`         // 状态: 1-启用 0-停用
	CreatedAt int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinPlatList230629's table name
func (*WinPlatList230629) TableName() string {
	return TableNameWinPlatList230629
}
