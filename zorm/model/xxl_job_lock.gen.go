// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameXxlJobLock = "xxl_job_lock"

// XxlJobLock mapped from table <xxl_job_lock>
type XxlJobLock struct {
	LockName string `gorm:"column:lock_name;type:varchar(50);primaryKey;comment:锁名称" json:"lockName"` // 锁名称
}

// TableName XxlJobLock's table name
func (*XxlJobLock) TableName() string {
	return TableNameXxlJobLock
}