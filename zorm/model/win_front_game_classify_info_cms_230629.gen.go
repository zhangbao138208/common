// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinFrontGameClassifyInfoCms230629 = "win_front_game_classify_info_cms_230629"

// WinFrontGameClassifyInfoCms230629 mapped from table <win_front_game_classify_info_cms_230629>
type WinFrontGameClassifyInfoCms230629 struct {
	ID          int64  `gorm:"column:id;type:int;not null" json:"id,string"`
	CID         int64  `gorm:"column:c_id;type:int;not null;comment:分类id" json:"cId"`                            // 分类id
	SlotID      string `gorm:"column:slot_id;type:varchar(255);not null;default:0;comment:slotId" json:"slotId"` // slotId
	GameID      int64  `gorm:"column:game_id;type:int;not null;comment:游戏ID(关联game_list)" json:"gameId"`         // 游戏ID(关联game_list)
	GameGroupID int64  `gorm:"column:game_group_id;type:int;not null;comment:游戏大类类型" json:"gameGroupId"`         // 游戏大类类型
	PlatID      int64  `gorm:"column:plat_id;type:int;not null;comment:游戏平台id" json:"platId"`                    // 游戏平台id
	Sort        int64  `gorm:"column:sort;type:int;default:99;comment:排序" json:"sort"`                           // 排序
	CreatedAt   int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                  // 创建时间
	CreateUser  string `gorm:"column:create_user;type:varchar(255);not null;comment:创建人" json:"createUser"`      // 创建人
	UpdatedAt   int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`                                  // 修改人
	UpdateUser  string `gorm:"column:update_user;type:varchar(255);comment:修改人" json:"updateUser"`               // 修改人
}

// TableName WinFrontGameClassifyInfoCms230629's table name
func (*WinFrontGameClassifyInfoCms230629) TableName() string {
	return TableNameWinFrontGameClassifyInfoCms230629
}
