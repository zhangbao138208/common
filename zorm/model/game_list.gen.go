// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGameList = "game_list"

// GameList mapped from table <game_list>
type GameList struct {
	ID                    int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Code                  string `gorm:"column:code;type:varchar(48);not null;comment:启动游戏编码" json:"code"`                                                    // 启动游戏编码
	GameProviderSubtypeID int64  `gorm:"column:game_provider_subtype_id;type:int;primaryKey;comment:关联game_provider_subtype表ID" json:"gameProviderSubtypeId"` // 关联game_provider_subtype表ID
	GamePagcorID          int64  `gorm:"column:game_pagcor_id;type:int;not null;comment:pagcor类型id" json:"gamePagcorId"`                                      // pagcor类型id
	GameTypeID            int64  `gorm:"column:game_type_id;type:int;not null;comment:游戏类型id" json:"gameTypeId"`                                              // 游戏类型id
	GameProviderID        int64  `gorm:"column:game_provider_id;type:int;primaryKey;comment:游戏供应商id" json:"gameProviderId"`                                   // 游戏供应商id
	Sort                  int64  `gorm:"column:sort;type:int;not null;comment:排序: 从低到高" json:"sort"`                                                          // 排序: 从低到高
	Status                int64  `gorm:"column:status;type:tinyint(1);not null;comment:状态: 1-启用 0-停用" json:"status"`                                          // 状态: 1-启用 0-停用
	Name                  string `gorm:"column:name;type:varchar(64);not null;comment:简体名称" json:"name"`                                                      // 简体名称
	OriginalIcon          string `gorm:"column:original_icon;type:varchar(512);not null;comment:英文图片" json:"originalIcon"`                                    // 英文图片
	LatestIcon            string `gorm:"column:latest_icon;type:varchar(512);comment:新版游戏图片" json:"latestIcon"`                                               // 新版游戏图片
	IsNew                 int64  `gorm:"column:is_new;type:tinyint(1);not null;comment:是否新游戏:1-是 0-否" json:"isNew"`                                           // 是否新游戏:1-是 0-否
	FavoriteStar          int64  `gorm:"column:favorite_star;type:int;not null;comment:收藏值" json:"favoriteStar"`                                              // 收藏值
	HotStar               int64  `gorm:"column:hot_star;type:int;not null;comment:热度" json:"hotStar"`                                                         // 热度
	Device                int64  `gorm:"column:device;type:tinyint;comment:设备:0-all 1-pc 2-h5" json:"device"`                                                 // 设备:0-all 1-pc 2-h5
	CreatedAt             int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt             int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	CreatedBy             string `gorm:"column:created_by;type:varchar(30);comment:操作人姓名" json:"createdBy"`    // 操作人姓名
	UpdatedBy             string `gorm:"column:updated_by;type:varchar(30);comment:最后更新人" json:"updatedBy"`    // 最后更新人
	Maintenance           string `gorm:"column:maintenance;type:varchar(255);comment:维护时间" json:"maintenance"` // 维护时间
}

// TableName GameList's table name
func (*GameList) TableName() string {
	return TableNameGameList
}