// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinPromotions = "win_promotions"

// WinPromotions mapped from table <win_promotions>
type WinPromotions struct {
	ID             int64           `gorm:"column:id;type:int(11);primaryKey" json:"id,string"`
	Amount         decimal.Decimal `gorm:"column:amount;type:decimal(15,4);not null;comment:总预算" json:"amount"`                                   // 总预算
	Balance        decimal.Decimal `gorm:"column:balance;type:decimal(15,4);not null;comment:总预算-剩余金额" json:"balance"`                            // 总预算-剩余金额
	DescriptZh     string          `gorm:"column:descript_zh;type:longtext;comment:详细描述-中文" json:"descriptZh"`                                    // 详细描述-中文
	Img            string          `gorm:"column:img;type:varchar(1024);not null;comment:图片" json:"img"`                                          // 图片
	Category       string          `gorm:"column:category;type:varchar(16);not null;default:0;comment:类型:1-充值优惠 2-豪礼赠送 3-新活动" json:"category"`    // 类型:1-充值优惠 2-豪礼赠送 3-新活动
	GameType       int64           `gorm:"column:game_type;type:tinyint(4);not null;comment:活动游戏类型，见字典dic_promotion_game_type" json:"gameType"`   // 活动游戏类型，见字典dic_promotion_game_type
	Info           string          `gorm:"column:info;type:longtext;comment:补充信息" json:"info"`                                                    // 补充信息
	Descript       string          `gorm:"column:descript;type:longtext;comment:详情描述" json:"descript"`                                            // 详情描述
	StartedAt      int64           `gorm:"column:started_at;type:int(11);comment:开始时间" json:"startedAt"`                                          // 开始时间
	Ladder         string          `gorm:"column:ladder;type:mediumtext;comment:新活动阶梯" json:"ladder"`                                             // 新活动阶梯
	PayoutCategory int64           `gorm:"column:payout_category;type:int(11);not null;comment:派彩类型: 0-自动派彩 1-人工派彩 2-手动派彩" json:"payoutCategory"` // 派彩类型: 0-自动派彩 1-人工派彩 2-手动派彩
	EndedAt        int64           `gorm:"column:ended_at;type:int(11);comment:结算时间" json:"endedAt"`                                              // 结算时间
	Sort           int64           `gorm:"column:sort;type:int(11);not null;default:100;comment:排序(从高到底、ID降序)" json:"sort"`                       // 排序(从高到底、ID降序)
	Status         int64           `gorm:"column:status;type:tinyint(4);not null;default:1;comment:状态:1-启用 0-停用" json:"status"`                   // 状态:1-启用 0-停用
	CreatedAt      int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt      int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	OperatorName   string          `gorm:"column:operator_name;type:varchar(32);comment:操作人姓名" json:"operatorName"`  // 操作人姓名
	Lang           string          `gorm:"column:lang;type:varchar(100);not null;default:en;comment:语言" json:"lang"` // 语言
}

// TableName WinPromotions's table name
func (*WinPromotions) TableName() string {
	return TableNameWinPromotions
}
