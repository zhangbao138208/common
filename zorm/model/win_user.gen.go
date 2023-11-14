// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinUser = "win_user"

// WinUser mapped from table <win_user>
type WinUser struct {
	ID                int64           `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Username          string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                                   // 用户名
	MerchantID        int64           `gorm:"column:merchant_id;type:int;not null;comment:商户id" json:"merchantId"`                                     // 商户id
	Avatar            string          `gorm:"column:avatar;type:varchar(64);comment:头像" json:"avatar"`                                                 // 头像
	Fcoin             decimal.Decimal `gorm:"column:fcoin;type:decimal(15,4);not null;default:0.0000;comment:冻结金额" json:"fcoin"`                       // 冻结金额
	CoinCommission    decimal.Decimal `gorm:"column:coin_commission;type:decimal(15,4);not null;default:0.0000;comment:佣金可提现金额" json:"coinCommission"` // 佣金可提现金额
	LevelID           int64           `gorm:"column:level_id;type:tinyint;not null;default:1;comment:会员等级" json:"levelId"`                             // 会员等级
	Role              int64           `gorm:"column:role;type:tinyint;not null;comment:角色:0-会员 1-代理 2-总代理 4-测试" json:"role"`                           // 角色:0-会员 1-代理 2-总代理 4-测试
	IsPromoter        int64           `gorm:"column:is_promoter;type:tinyint(1);not null;comment:是否推广:0-不是 1-是" json:"isPromoter"`                     // 是否推广:0-不是 1-是
	Flag              int64           `gorm:"column:flag;type:int unsigned;not null;comment:会员旗" json:"flag"`                                          // 会员旗
	RealName          string          `gorm:"column:real_name;type:varchar(32);not null;comment:真实姓名" json:"realName"`                                 // 真实姓名
	Signature         string          `gorm:"column:signature;type:varchar(128);not null;comment:个性签名" json:"signature"`                               // 个性签名
	Birthday          string          `gorm:"column:birthday;type:varchar(16);comment:生日" json:"birthday"`                                             // 生日
	AreaCode          string          `gorm:"column:area_code;type:varchar(8);not null;comment:区号" json:"areaCode"`                                    // 区号
	Mobile            string          `gorm:"column:mobile;type:varchar(16);not null;comment:手机号码" json:"mobile"`                                      // 手机号码
	Email             string          `gorm:"column:email;type:varchar(32);comment:邮箱" json:"email"`                                                   // 邮箱
	Sex               int64           `gorm:"column:sex;type:tinyint(1);not null;default:1;comment:性别:1-男 0-女 2-未知" json:"sex"`                        // 性别:1-男 0-女 2-未知
	BindBank          int64           `gorm:"column:bind_bank;type:tinyint(1);not null;comment:是否绑定银行卡:1-已绑定 0-未绑定" json:"bindBank"`                   // 是否绑定银行卡:1-已绑定 0-未绑定
	Address           string          `gorm:"column:address;type:varchar(256);not null;comment:家庭地址" json:"address"`                                   // 家庭地址
	Score             int64           `gorm:"column:score;type:int unsigned;not null;comment:积分" json:"score"`                                         // 积分
	PromoCode         string          `gorm:"column:promo_code;type:varchar(8);not null;comment:推广码" json:"promoCode"`                                 // 推广码
	SupUid1           int64           `gorm:"column:sup_uid_1;type:int unsigned;not null;comment:上1级代理" json:"supUid1"`                                // 上1级代理
	SupUsername1      string          `gorm:"column:sup_username_1;type:varchar(32);not null;comment:上1级代理" json:"supUsername1"`                       // 上1级代理
	SupUid2           int64           `gorm:"column:sup_uid_2;type:int unsigned;not null;comment:上2级代理" json:"supUid2"`                                // 上2级代理
	SupUid3           int64           `gorm:"column:sup_uid_3;type:int unsigned;not null;comment:上3级代理" json:"supUid3"`                                // 上3级代理
	SupUid4           int64           `gorm:"column:sup_uid_4;type:int unsigned;not null;comment:上4级代理" json:"supUid4"`                                // 上4级代理
	SupUid5           int64           `gorm:"column:sup_uid_5;type:int unsigned;not null;comment:上5级代理" json:"supUid5"`                                // 上5级代理
	SupUid6           int64           `gorm:"column:sup_uid_6;type:int unsigned;not null;comment:上6级代理" json:"supUid6"`                                // 上6级代理
	SupUIDTop         int64           `gorm:"column:sup_uid_top;type:int;not null;comment:顶级推广用户名" json:"supUidTop"`                                   // 顶级推广用户名
	SupUsernameTop    string          `gorm:"column:sup_username_top;type:varchar(32);not null;comment:顶级推广用户名" json:"supUsernameTop"`                 // 顶级推广用户名
	SupLevelTop       int64           `gorm:"column:sup_level_top;type:int;not null;default:-1;comment:顶级推广层级" json:"supLevelTop"`                     // 顶级推广层级
	PasswordHash      string          `gorm:"column:password_hash;type:varchar(255);not null;comment:登录密码" json:"passwordHash"`                        // 登录密码
	PasswordCoin      string          `gorm:"column:password_coin;type:varchar(255);not null;comment:取款密码" json:"passwordCoin"`                        // 取款密码
	IP                string          `gorm:"column:ip;type:varchar(128);not null;comment:IP地址" json:"ip"`                                             // IP地址
	IPRegion          string          `gorm:"column:ip_region;type:varchar(255);comment:IP归属地" json:"ipRegion"`                                        // IP归属地
	ThirdLoginType    string          `gorm:"column:third_login_type;type:varchar(128);not null;comment:三方登陆类型" json:"thirdLoginType"`                 // 三方登陆类型
	FreezeCause       string          `gorm:"column:freeze_cause;type:varchar(1000)" json:"freezeCause"`
	FreezeAt          int64           `gorm:"column:freeze_at;type:int;comment:冻结时间" json:"freezeAt"`                                     // 冻结时间
	OperatorName      string          `gorm:"column:operator_name;type:varchar(32);comment:操作人姓名" json:"operatorName"`                    // 操作人姓名
	CreatedName       string          `gorm:"column:created_name;type:varchar(32);comment:创建人" json:"createdName"`                        // 创建人
	Status            int64           `gorm:"column:status;type:tinyint(1);not null;default:10;comment:状态:10-正常 9-冻结 8-删除" json:"status"` // 状态:10-正常 9-冻结 8-删除
	LastLoginIP       string          `gorm:"column:last_login_ip;type:varchar(128);comment:最后登陆ip" json:"lastLoginIp"`                   // 最后登陆ip
	LastLoginIPRegion string          `gorm:"column:last_login_ip_region;type:varchar(255);comment:最后登录IP归属地" json:"lastLoginIpRegion"`   // 最后登录IP归属地
	LastLoginTime     int64           `gorm:"column:last_login_time;type:int;comment:最后登陆时间" json:"lastLoginTime"`                        // 最后登陆时间
	LastLoginDeviceID string          `gorm:"column:last_login_device_id;type:varchar(64);comment:最后登录设备id" json:"lastLoginDeviceId"`     // 最后登录设备id
	CreatedAt         int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt         int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinUser's table name
func (*WinUser) TableName() string {
	return TableNameWinUser
}