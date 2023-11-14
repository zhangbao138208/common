// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinNotice = "win_notice"

// WinNotice mapped from table <win_notice>
type WinNotice struct {
	ID           int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Title        string `gorm:"column:title;type:varchar(32);not null;comment:标题" json:"title"`                       // 标题
	Content      string `gorm:"column:content;type:text;comment:内容" json:"content"`                                   // 内容
	Uids         string `gorm:"column:uids;type:text;comment:系统公告制定用户" json:"uids"`                                   // 系统公告制定用户
	Category     int64  `gorm:"column:category;type:tinyint;default:4;comment:类型:1-系统公告2-站内信 3-系统消息" json:"category"` // 类型:1-系统公告2-站内信 3-系统消息
	Status       int64  `gorm:"column:status;type:tinyint;not null;default:1;comment:状态:1-启用 0-停用" json:"status"`     // 状态:1-启用 0-停用
	Sort         int64  `gorm:"column:sort;type:smallint;not null;default:1;comment:排序:从大到小" json:"sort"`             // 排序:从大到小
	CreatedAt    int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt    int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	OperatorName string `gorm:"column:operator_name;type:varchar(32);comment:操作人姓名" json:"operatorName"` // 操作人姓名
}

// TableName WinNotice's table name
func (*WinNotice) TableName() string {
	return TableNameWinNotice
}