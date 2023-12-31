// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserKyc = "win_user_kyc"

// WinUserKyc mapped from table <win_user_kyc>
type WinUserKyc struct {
	ID                int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	UID               int64  `gorm:"column:uid;type:int;not null;comment:uid" json:"uid"`                         // uid
	Username          string `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`       // 用户名
	MerchantID        int64  `gorm:"column:merchant_id;type:int;not null;comment:商户id" json:"merchantId"`         // 商户id
	FirstName         string `gorm:"column:first_name;type:varchar(50);comment:first_name" json:"firstName"`      // first_name
	MiddleName        string `gorm:"column:middle_name;type:varchar(50);comment:middle_name" json:"middleName"`   // middle_name
	LastName          string `gorm:"column:last_name;type:varchar(50);comment:last_name" json:"lastName"`         // last_name
	Birthday          int64  `gorm:"column:birthday;type:int;comment:birthday" json:"birthday"`                   // birthday
	ImgFront          string `gorm:"column:img_front;type:varchar(1024);comment:证件正面图片地址" json:"imgFront"`        // 证件正面图片地址
	ImgBack           string `gorm:"column:img_back;type:varchar(1024);comment:证件背面图片地址" json:"imgBack"`          // 证件背面图片地址
	IDType            int64  `gorm:"column:id_type;type:int;not null;comment:证件类型:1-身份证 2-护照 3-驾照" json:"idType"` // 证件类型:1-身份证 2-护照 3-驾照
	IDNumber          string `gorm:"column:id_number;type:varchar(32);comment:证件编号" json:"idNumber"`              // 证件编号
	Status            int64  `gorm:"column:status;type:int;not null;comment:状态:1-待审核 2-通过 3-拒绝" json:"status"`    // 状态:1-待审核 2-通过 3-拒绝
	Reason            string `gorm:"column:reason;type:varchar(256);comment:拒绝原因" json:"reason"`                  // 拒绝原因
	Remark            string `gorm:"column:remark;type:varchar(256);comment:备注" json:"remark"`                    // 备注
	AuditUsername     string `gorm:"column:audit_username;type:varchar(32);comment:审核人" json:"auditUsername"`     // 审核人
	AuditAt           int64  `gorm:"column:audit_at;type:bigint;comment:审核时间" json:"auditAt"`                     // 审核时间
	CreatedAt         int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt         int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	OperatorName      string `gorm:"column:operator_name;type:varchar(32);comment:操作人姓名" json:"operatorName"`                                                                                                                                                                                                                                                                                                                           // 操作人姓名
	Nationality       string `gorm:"column:nationality;type:varchar(10);comment:国籍" json:"nationality"`                                                                                                                                                                                                                                                                                                                                 // 国籍
	PlaceOfBirth      string `gorm:"column:place_of_birth;type:varchar(512);comment:出生地" json:"placeOfBirth"`                                                                                                                                                                                                                                                                                                                           // 出生地
	CurrentAddress    string `gorm:"column:current_address;type:varchar(512);comment:当前地址" json:"currentAddress"`                                                                                                                                                                                                                                                                                                                       // 当前地址
	PermanentAddress  string `gorm:"column:permanent_address;type:varchar(512);comment:长期地址" json:"permanentAddress"`                                                                                                                                                                                                                                                                                                                   // 长期地址
	EmploymentType    string `gorm:"column:employment_type;type:varchar(10);comment:工作性质1,Government Employee 2,OFW 3,Private Sector Employee 4,Self-Employed 5,Student 6,Unemployed" json:"employmentType"`                                                                                                                                                                                                                            // 工作性质1,Government Employee 2,OFW 3,Private Sector Employee 4,Self-Employed 5,Student 6,Unemployed
	MainSourceOfFunds string `gorm:"column:main_source_of_funds;type:varchar(10);comment:收入来源1，Allowance/Pension/Stipend/Honoraria2，Campaign Funds/Donation3，Cash on Hand4，Commission or Incentives5，E-Money6，Income from Business/Profession7，Inheritance8，Investment9，Loan Availments10，Personal Savings11，Proceeds of Personal or Real property sale12，Remittance from relatives13，Rental Income14，Salary" json:"mainSourceOfFunds"` // 收入来源1，Allowance/Pension/Stipend/Honoraria2，Campaign Funds/Donation3，Cash on Hand4，Commission or Incentives5，E-Money6，Income from Business/Profession7，Inheritance8，Investment9，Loan Availments10，Personal Savings11，Proceeds of Personal or Real property sale12，Remittance from relatives13，Rental Income14，Salary
	RelationStore     string `gorm:"column:relation_store;type:varchar(50);comment:关联商铺" json:"relationStore"`                                                                                                                                                                                                                                                                                                                          // 关联商铺
}

// TableName WinUserKyc's table name
func (*WinUserKyc) TableName() string {
	return TableNameWinUserKyc
}
