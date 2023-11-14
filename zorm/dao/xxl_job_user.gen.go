// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gitlab.skig.tech/zero-core/common/zorm/model"
)

func newXxlJobUser(db *gorm.DB, opts ...gen.DOOption) xxlJobUser {
	_xxlJobUser := xxlJobUser{}

	_xxlJobUser.xxlJobUserDo.UseDB(db, opts...)
	_xxlJobUser.xxlJobUserDo.UseModel(&model.XxlJobUser{})

	tableName := _xxlJobUser.xxlJobUserDo.TableName()
	_xxlJobUser.ALL = field.NewAsterisk(tableName)
	_xxlJobUser.ID = field.NewInt64(tableName, "id")
	_xxlJobUser.Username = field.NewString(tableName, "username")
	_xxlJobUser.Password = field.NewString(tableName, "password")
	_xxlJobUser.Role = field.NewInt64(tableName, "role")
	_xxlJobUser.Permission = field.NewString(tableName, "permission")

	_xxlJobUser.fillFieldMap()

	return _xxlJobUser
}

type xxlJobUser struct {
	xxlJobUserDo

	ALL        field.Asterisk
	ID         field.Int64
	Username   field.String // 账号
	Password   field.String // 密码
	Role       field.Int64  // 角色：0-普通用户、1-管理员
	Permission field.String // 权限：执行器ID列表，多个逗号分割

	fieldMap map[string]field.Expr
}

func (x xxlJobUser) Table(newTableName string) *xxlJobUser {
	x.xxlJobUserDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xxlJobUser) As(alias string) *xxlJobUser {
	x.xxlJobUserDo.DO = *(x.xxlJobUserDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xxlJobUser) updateTableName(table string) *xxlJobUser {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt64(table, "id")
	x.Username = field.NewString(table, "username")
	x.Password = field.NewString(table, "password")
	x.Role = field.NewInt64(table, "role")
	x.Permission = field.NewString(table, "permission")

	x.fillFieldMap()

	return x
}

func (x *xxlJobUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xxlJobUser) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 5)
	x.fieldMap["id"] = x.ID
	x.fieldMap["username"] = x.Username
	x.fieldMap["password"] = x.Password
	x.fieldMap["role"] = x.Role
	x.fieldMap["permission"] = x.Permission
}

func (x xxlJobUser) clone(db *gorm.DB) xxlJobUser {
	x.xxlJobUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xxlJobUser) replaceDB(db *gorm.DB) xxlJobUser {
	x.xxlJobUserDo.ReplaceDB(db)
	return x
}

type xxlJobUserDo struct{ gen.DO }

type IXxlJobUserDo interface {
	gen.SubQuery
	Debug() IXxlJobUserDo
	WithContext(ctx context.Context) IXxlJobUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXxlJobUserDo
	WriteDB() IXxlJobUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXxlJobUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXxlJobUserDo
	Not(conds ...gen.Condition) IXxlJobUserDo
	Or(conds ...gen.Condition) IXxlJobUserDo
	Select(conds ...field.Expr) IXxlJobUserDo
	Where(conds ...gen.Condition) IXxlJobUserDo
	Order(conds ...field.Expr) IXxlJobUserDo
	Distinct(cols ...field.Expr) IXxlJobUserDo
	Omit(cols ...field.Expr) IXxlJobUserDo
	Join(table schema.Tabler, on ...field.Expr) IXxlJobUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobUserDo
	Group(cols ...field.Expr) IXxlJobUserDo
	Having(conds ...gen.Condition) IXxlJobUserDo
	Limit(limit int) IXxlJobUserDo
	Offset(offset int) IXxlJobUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobUserDo
	Unscoped() IXxlJobUserDo
	Create(values ...*model.XxlJobUser) error
	CreateInBatches(values []*model.XxlJobUser, batchSize int) error
	Save(values ...*model.XxlJobUser) error
	First() (*model.XxlJobUser, error)
	Take() (*model.XxlJobUser, error)
	Last() (*model.XxlJobUser, error)
	Find() ([]*model.XxlJobUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobUser, err error)
	FindInBatches(result *[]*model.XxlJobUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XxlJobUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXxlJobUserDo
	Assign(attrs ...field.AssignExpr) IXxlJobUserDo
	Joins(fields ...field.RelationField) IXxlJobUserDo
	Preload(fields ...field.RelationField) IXxlJobUserDo
	FirstOrInit() (*model.XxlJobUser, error)
	FirstOrCreate() (*model.XxlJobUser, error)
	FindByPage(offset int, limit int) (result []*model.XxlJobUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXxlJobUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xxlJobUserDo) Debug() IXxlJobUserDo {
	return x.withDO(x.DO.Debug())
}

func (x xxlJobUserDo) WithContext(ctx context.Context) IXxlJobUserDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xxlJobUserDo) ReadDB() IXxlJobUserDo {
	return x.Clauses(dbresolver.Read)
}

func (x xxlJobUserDo) WriteDB() IXxlJobUserDo {
	return x.Clauses(dbresolver.Write)
}

func (x xxlJobUserDo) Session(config *gorm.Session) IXxlJobUserDo {
	return x.withDO(x.DO.Session(config))
}

func (x xxlJobUserDo) Clauses(conds ...clause.Expression) IXxlJobUserDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xxlJobUserDo) Returning(value interface{}, columns ...string) IXxlJobUserDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xxlJobUserDo) Not(conds ...gen.Condition) IXxlJobUserDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xxlJobUserDo) Or(conds ...gen.Condition) IXxlJobUserDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xxlJobUserDo) Select(conds ...field.Expr) IXxlJobUserDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xxlJobUserDo) Where(conds ...gen.Condition) IXxlJobUserDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xxlJobUserDo) Order(conds ...field.Expr) IXxlJobUserDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xxlJobUserDo) Distinct(cols ...field.Expr) IXxlJobUserDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xxlJobUserDo) Omit(cols ...field.Expr) IXxlJobUserDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xxlJobUserDo) Join(table schema.Tabler, on ...field.Expr) IXxlJobUserDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xxlJobUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobUserDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xxlJobUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobUserDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xxlJobUserDo) Group(cols ...field.Expr) IXxlJobUserDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xxlJobUserDo) Having(conds ...gen.Condition) IXxlJobUserDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xxlJobUserDo) Limit(limit int) IXxlJobUserDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xxlJobUserDo) Offset(offset int) IXxlJobUserDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xxlJobUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobUserDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xxlJobUserDo) Unscoped() IXxlJobUserDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xxlJobUserDo) Create(values ...*model.XxlJobUser) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xxlJobUserDo) CreateInBatches(values []*model.XxlJobUser, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xxlJobUserDo) Save(values ...*model.XxlJobUser) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xxlJobUserDo) First() (*model.XxlJobUser, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobUser), nil
	}
}

func (x xxlJobUserDo) Take() (*model.XxlJobUser, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobUser), nil
	}
}

func (x xxlJobUserDo) Last() (*model.XxlJobUser, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobUser), nil
	}
}

func (x xxlJobUserDo) Find() ([]*model.XxlJobUser, error) {
	result, err := x.DO.Find()
	return result.([]*model.XxlJobUser), err
}

func (x xxlJobUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobUser, err error) {
	buf := make([]*model.XxlJobUser, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xxlJobUserDo) FindInBatches(result *[]*model.XxlJobUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xxlJobUserDo) Attrs(attrs ...field.AssignExpr) IXxlJobUserDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xxlJobUserDo) Assign(attrs ...field.AssignExpr) IXxlJobUserDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xxlJobUserDo) Joins(fields ...field.RelationField) IXxlJobUserDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xxlJobUserDo) Preload(fields ...field.RelationField) IXxlJobUserDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xxlJobUserDo) FirstOrInit() (*model.XxlJobUser, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobUser), nil
	}
}

func (x xxlJobUserDo) FirstOrCreate() (*model.XxlJobUser, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobUser), nil
	}
}

func (x xxlJobUserDo) FindByPage(offset int, limit int) (result []*model.XxlJobUser, count int64, err error) {
	result, err = x.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = x.Offset(-1).Limit(-1).Count()
	return
}

func (x xxlJobUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xxlJobUserDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xxlJobUserDo) Delete(models ...*model.XxlJobUser) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xxlJobUserDo) withDO(do gen.Dao) *xxlJobUserDo {
	x.DO = *do.(*gen.DO)
	return x
}
