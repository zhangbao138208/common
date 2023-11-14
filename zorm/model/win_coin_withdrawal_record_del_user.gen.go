// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinCoinWithdrawalRecordDelUser = "win_coin_withdrawal_record_del_user"

// WinCoinWithdrawalRecordDelUser mapped from table <win_coin_withdrawal_record_del_user>
type WinCoinWithdrawalRecordDelUser struct {
	ID                      int64           `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	OrderID                 string          `gorm:"column:order_id;type:varchar(64);not null;comment:订单号(三方平台用)" json:"orderId"`                                                          // 订单号(三方平台用)
	PlatOrderID             string          `gorm:"column:plat_order_id;type:varchar(64);not null;comment:三方平台订单号" json:"platOrderId"`                                                    // 三方平台订单号
	UID                     int64           `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                                                                                  // UID
	Username                string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                                                                // 用户名
	Code                    string          `gorm:"column:code;type:varchar(50);comment:支付通道编码" json:"code"`                                                                              // 支付通道编码
	PlatName                string          `gorm:"column:plat_name;type:varchar(100);comment:平台名称" json:"platName"`                                                                      // 平台名称
	PlatNickName            string          `gorm:"column:plat_nick_name;type:varchar(100);comment:平台自定义名称" json:"platNickName"`                                                          // 平台自定义名称
	WithdrawalAddress       string          `gorm:"column:withdrawal_address;type:text;not null;comment:加密地址" json:"withdrawalAddress"`                                                   // 加密地址
	WithdrawalAmount        decimal.Decimal `gorm:"column:withdrawal_amount;type:decimal(15,4);not null;default:0.0000;comment:提款金额" json:"withdrawalAmount"`                             // 提款金额
	ExchangeRate            decimal.Decimal `gorm:"column:exchange_rate;type:decimal(15,4);not null;default:0.0000;comment:汇率" json:"exchangeRate"`                                       // 汇率
	RealAmount              decimal.Decimal `gorm:"column:real_amount;type:decimal(15,4);not null;default:0.0000;comment:到账金额" json:"realAmount"`                                         // 到账金额
	CoinBefore              decimal.Decimal `gorm:"column:coin_before;type:decimal(15,4);not null;default:0.0000;comment:提款前用户金额" json:"coinBefore"`                                      // 提款前用户金额
	MainNetFees             decimal.Decimal `gorm:"column:main_net_fees;type:decimal(15,4);not null;default:0.0000;comment:主网费" json:"mainNetFees"`                                       // 主网费
	Currency                string          `gorm:"column:currency;type:varchar(50);not null;comment:币种" json:"currency"`                                                                 // 币种
	CategoryCurrency        int64           `gorm:"column:category_currency;type:tinyint(1);not null;default:1;comment:货币类型:0-数字货币 1-法币" json:"categoryCurrency"`                         // 货币类型:0-数字货币 1-法币
	CategoryTransfer        int64           `gorm:"column:category_transfer;type:tinyint;not null;default:3;comment:转账类型:1-TRC-20 2-ERC-20 3-BANK 4-PIX 5-GCASH" json:"categoryTransfer"` // 转账类型:1-TRC-20 2-ERC-20 3-BANK 4-PIX 5-GCASH
	Status                  int64           `gorm:"column:status;type:tinyint;not null;comment:状态: 0-申请中，1-提款成功，2-提款失败，3-稽核成功 4-代付种" json:"status"`                                       // 状态: 0-申请中，1-提款成功，2-提款失败，3-稽核成功 4-代付种
	AuditType               int64           `gorm:"column:audit_type;type:tinyint;comment:审核类型 0人工，1自动" json:"auditType"`                                                                 // 审核类型 0人工，1自动
	AdminUsername           string          `gorm:"column:admin_username;type:varchar(32);not null;comment:操作人" json:"adminUsername"`                                                     // 操作人
	Mark                    string          `gorm:"column:mark;type:varchar(128);comment:审核备注" json:"mark"`                                                                               // 审核备注
	WithdrawalAdminUsername string          `gorm:"column:withdrawal_admin_username;type:varchar(32);comment:提现操作人" json:"withdrawalAdminUsername"`                                       // 提现操作人
	CustomizeNoticeID       int64           `gorm:"column:customize_notice_Id;type:int;comment:自定义通知客户信息Id" json:"customizeNoticeId"`                                                     // 自定义通知客户信息Id
	CreatedAt               int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt               int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	FinanceOperatorAt       int64           `gorm:"column:finance_operator_at;type:int;comment:财务操作时间" json:"financeOperatorAt"` // 财务操作时间
}

// TableName WinCoinWithdrawalRecordDelUser's table name
func (*WinCoinWithdrawalRecordDelUser) TableName() string {
	return TableNameWinCoinWithdrawalRecordDelUser
}