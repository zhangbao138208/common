// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinCoinRate = "win_coin_rate"

// WinCoinRate mapped from table <win_coin_rate>
type WinCoinRate struct {
	ID               int64           `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:ID" json:"id,string"`            // ID
	OriginalCurrency string          `gorm:"column:original_currency;type:varchar(20);not null;comment:原始币种" json:"originalCurrency"` // 原始币种
	TransferCurrency string          `gorm:"column:transfer_currency;type:varchar(20);not null;comment:转换币种" json:"transferCurrency"` // 转换币种
	Rate             decimal.Decimal `gorm:"column:rate;type:decimal(11,8);not null;default:0.00000000;comment:汇率" json:"rate"`       // 汇率
	Status           int64           `gorm:"column:status;type:int;not null;default:1;comment:状态：0-关闭；1-开启" json:"status"`            // 状态：0-关闭；1-开启
	CreatedAt        int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                         // 创建时间
	UpdatedAt        int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`                                         // 修改时间
	OperatorName     string          `gorm:"column:operator_name;type:varchar(32);comment:操作人姓名" json:"operatorName"`                 // 操作人姓名
}

// TableName WinCoinRate's table name
func (*WinCoinRate) TableName() string {
	return TableNameWinCoinRate
}
