// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinGameSlot = "win_game_slot"

// WinGameSlot mapped from table <win_game_slot>
type WinGameSlot struct {
	ID           string `gorm:"column:id;type:varchar(48);primaryKey;comment:ID(关联BrandGameId)" json:"id,string"`                                                // ID(关联BrandGameId)
	GameID       int64  `gorm:"column:game_id;type:int;primaryKey;comment:游戏ID(关联game_list)" json:"gameId"`                                                      // 游戏ID(关联game_list)
	GameGroupID  int64  `gorm:"column:game_group_id;type:tinyint;not null;comment:游戏大类类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能" json:"gameGroupId"` // 游戏大类类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能
	PlatID       int64  `gorm:"column:plat_id;type:int;primaryKey;comment:游戏平台id" json:"platId"`                                                                 // 游戏平台id
	Provider     string `gorm:"column:provider;type:varchar(64);comment:游戏提供者" json:"provider"`                                                                  // 游戏提供者
	Name         string `gorm:"column:name;type:varchar(64);not null;comment:简体名称" json:"name"`                                                                  // 简体名称
	NameZh       string `gorm:"column:name_zh;type:varchar(64);not null;comment:游戏名字(中文)" json:"nameZh"`                                                         // 游戏名字(中文)
	Img          string `gorm:"column:img;type:varchar(512);not null;comment:英文图片" json:"img"`                                                                   // 英文图片
	ImgNew       string `gorm:"column:img_new;type:varchar(512);comment:新版游戏图片" json:"imgNew"`                                                                   // 新版游戏图片
	IsNew        int64  `gorm:"column:is_new;type:tinyint(1);not null;comment:是否新游戏:1-是 0-否" json:"isNew"`                                                       // 是否新游戏:1-是 0-否
	IsCasino     int64  `gorm:"column:is_casino;type:tinyint;not null;comment:是否推荐主页 0否 1是" json:"isCasino"`                                                     // 是否推荐主页 0否 1是
	GameTypeID   string `gorm:"column:game_type_id;type:varchar(30);comment:游戏类型ID(0r code)" json:"gameTypeId"`                                                  // 游戏类型ID(0r code)
	GameTypeName string `gorm:"column:game_type_name;type:varchar(32);comment:游戏类型名称" json:"gameTypeName"`                                                       // 游戏类型名称
	FavoriteStar int64  `gorm:"column:favorite_star;type:int;not null;comment:收藏值" json:"favoriteStar"`                                                          // 收藏值
	HotStar      int64  `gorm:"column:hot_star;type:int;not null;comment:热度" json:"hotStar"`                                                                     // 热度
	Sort         int64  `gorm:"column:sort;type:int;not null;comment:排序" json:"sort"`                                                                            // 排序
	Status       int64  `gorm:"column:status;type:tinyint;not null;default:1;comment:状态:1-启用 0-停用" json:"status"`                                                // 状态:1-启用 0-停用
	Device       int64  `gorm:"column:device;type:tinyint;comment:设备:0-all 1-pc 2-h5" json:"device"`                                                             // 设备:0-all 1-pc 2-h5
	CreatedAt    int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt    int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	UpdatedUser  string `gorm:"column:updated_user;type:varchar(20);comment:最后更新人" json:"updatedUser"`   // 最后更新人
	Maintenance  string `gorm:"column:maintenance;type:varchar(500);comment:维护信息" json:"maintenance"`    // 维护信息
	OperatorName string `gorm:"column:operator_name;type:varchar(32);comment:操作人姓名" json:"operatorName"` // 操作人姓名
}

// TableName WinGameSlot's table name
func (*WinGameSlot) TableName() string {
	return TableNameWinGameSlot
}
