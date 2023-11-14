// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinBetslipsDetail = "win_betslips_details"

// WinBetslipsDetail mapped from table <win_betslips_details>
type WinBetslipsDetail struct {
	ID         int64  `gorm:"column:id;type:bigint;primaryKey;comment:主键-同注单表一致" json:"id,string"`                    // 主键-同注单表一致
	XbUID      int64  `gorm:"column:xb_uid;type:int unsigned;not null;comment:对应user表id" json:"xbUid"`                // 对应user表id
	XbUsername string `gorm:"column:xb_username;type:varchar(32);not null;comment:对应user表username" json:"xbUsername"` // 对应user表username
	BetJSON    string `gorm:"column:bet_json;type:text;comment:投注原始json" json:"betJson"`                              // 投注原始json
	RewardJSON string `gorm:"column:reward_json;type:text;comment:开彩原始json" json:"rewardJson"`                        // 开彩原始json
	RefundJSON string `gorm:"column:refund_json;type:text;comment:退款原始json" json:"refundJson"`                        // 退款原始json
}

// TableName WinBetslipsDetail's table name
func (*WinBetslipsDetail) TableName() string {
	return TableNameWinBetslipsDetail
}
