// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinCoinRewards = "win_coin_rewards"

// WinCoinRewards mapped from table <win_coin_rewards>
type WinCoinRewards struct {
	ID              int64           `gorm:"column:id;type:bigint(20);primaryKey" json:"id,string"`
	UID             int64           `gorm:"column:uid;type:int(11);not null;comment:UID" json:"uid"`                                      // UID
	Username        string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                        // 用户名
	Coin            decimal.Decimal `gorm:"column:coin;type:decimal(15,4);not null;default:0.0000;comment:金额" json:"coin"`                // 金额
	CoinBefore      decimal.Decimal `gorm:"column:coin_before;type:decimal(15,4);not null;default:0.0000;comment:即时金额" json:"coinBefore"` // 即时金额
	ReferID         int64           `gorm:"column:refer_id;type:int(11);comment:关联ID(活动表ID)" json:"referId"`                              // 关联ID(活动表ID)
	LadderName      string          `gorm:"column:ladder_name;type:varchar(255);comment:新活动，等级code" json:"ladderName"`                    // 新活动，等级code
	ReferCode       string          `gorm:"column:refer_code;type:varchar(64);not null;comment:关联Code(活动表Code)" json:"referCode"`         // 关联Code(活动表Code)
	FlowClaim       int64           `gorm:"column:flow_claim;type:int(11);comment:流水倍数" json:"flowClaim"`                                 // 流水倍数
	Codes           decimal.Decimal `gorm:"column:codes;type:decimal(15,4);comment:所需打码量" json:"codes"`                                   // 所需打码量
	EndedAt         int64           `gorm:"column:ended_at;type:int(11);comment:活动结束时间" json:"endedAt"`                                   // 活动结束时间
	Info            string          `gorm:"column:info;type:longtext;comment:备注" json:"info"`                                             // 备注
	Status          int64           `gorm:"column:status;type:tinyint(4);not null;comment:状态:0-申请中 1-已满足 2-已派发3-已结束" json:"status"`       // 状态:0-申请中 1-已满足 2-已派发3-已结束
	Detail          string          `gorm:"column:detail;type:varchar(255);not null;comment:详情信息" json:"detail"`                          // 详情信息
	CreatedAt       int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt       int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	TransferBonusAt int64           `gorm:"column:transfer_bonus_at;type:int(11);not null;comment:转主钱包时间" json:"transferBonusAt"` // 转主钱包时间
}

// TableName WinCoinRewards's table name
func (*WinCoinRewards) TableName() string {
	return TableNameWinCoinRewards
}
