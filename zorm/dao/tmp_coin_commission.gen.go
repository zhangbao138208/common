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

func newTmpCoinCommission(db *gorm.DB, opts ...gen.DOOption) tmpCoinCommission {
	_tmpCoinCommission := tmpCoinCommission{}

	_tmpCoinCommission.tmpCoinCommissionDo.UseDB(db, opts...)
	_tmpCoinCommission.tmpCoinCommissionDo.UseModel(&model.TmpCoinCommission{})

	tableName := _tmpCoinCommission.tmpCoinCommissionDo.TableName()
	_tmpCoinCommission.ALL = field.NewAsterisk(tableName)
	_tmpCoinCommission.ID = field.NewInt64(tableName, "id")
	_tmpCoinCommission.UID = field.NewInt64(tableName, "uid")
	_tmpCoinCommission.SubUID = field.NewInt64(tableName, "sub_uid")
	_tmpCoinCommission.Riqi = field.NewInt64(tableName, "riqi")
	_tmpCoinCommission.Coun = field.NewInt64(tableName, "coun")
	_tmpCoinCommission.Coin = field.NewField(tableName, "coin")

	_tmpCoinCommission.fillFieldMap()

	return _tmpCoinCommission
}

type tmpCoinCommission struct {
	tmpCoinCommissionDo

	ALL    field.Asterisk
	ID     field.Int64
	UID    field.Int64 // 代理UID
	SubUID field.Int64 // 下级UID
	Riqi   field.Int64 // 佣金时间
	Coun   field.Int64
	Coin   field.Field

	fieldMap map[string]field.Expr
}

func (t tmpCoinCommission) Table(newTableName string) *tmpCoinCommission {
	t.tmpCoinCommissionDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tmpCoinCommission) As(alias string) *tmpCoinCommission {
	t.tmpCoinCommissionDo.DO = *(t.tmpCoinCommissionDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tmpCoinCommission) updateTableName(table string) *tmpCoinCommission {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.UID = field.NewInt64(table, "uid")
	t.SubUID = field.NewInt64(table, "sub_uid")
	t.Riqi = field.NewInt64(table, "riqi")
	t.Coun = field.NewInt64(table, "coun")
	t.Coin = field.NewField(table, "coin")

	t.fillFieldMap()

	return t
}

func (t *tmpCoinCommission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tmpCoinCommission) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["uid"] = t.UID
	t.fieldMap["sub_uid"] = t.SubUID
	t.fieldMap["riqi"] = t.Riqi
	t.fieldMap["coun"] = t.Coun
	t.fieldMap["coin"] = t.Coin
}

func (t tmpCoinCommission) clone(db *gorm.DB) tmpCoinCommission {
	t.tmpCoinCommissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tmpCoinCommission) replaceDB(db *gorm.DB) tmpCoinCommission {
	t.tmpCoinCommissionDo.ReplaceDB(db)
	return t
}

type tmpCoinCommissionDo struct{ gen.DO }

type ITmpCoinCommissionDo interface {
	gen.SubQuery
	Debug() ITmpCoinCommissionDo
	WithContext(ctx context.Context) ITmpCoinCommissionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITmpCoinCommissionDo
	WriteDB() ITmpCoinCommissionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITmpCoinCommissionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITmpCoinCommissionDo
	Not(conds ...gen.Condition) ITmpCoinCommissionDo
	Or(conds ...gen.Condition) ITmpCoinCommissionDo
	Select(conds ...field.Expr) ITmpCoinCommissionDo
	Where(conds ...gen.Condition) ITmpCoinCommissionDo
	Order(conds ...field.Expr) ITmpCoinCommissionDo
	Distinct(cols ...field.Expr) ITmpCoinCommissionDo
	Omit(cols ...field.Expr) ITmpCoinCommissionDo
	Join(table schema.Tabler, on ...field.Expr) ITmpCoinCommissionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITmpCoinCommissionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITmpCoinCommissionDo
	Group(cols ...field.Expr) ITmpCoinCommissionDo
	Having(conds ...gen.Condition) ITmpCoinCommissionDo
	Limit(limit int) ITmpCoinCommissionDo
	Offset(offset int) ITmpCoinCommissionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITmpCoinCommissionDo
	Unscoped() ITmpCoinCommissionDo
	Create(values ...*model.TmpCoinCommission) error
	CreateInBatches(values []*model.TmpCoinCommission, batchSize int) error
	Save(values ...*model.TmpCoinCommission) error
	First() (*model.TmpCoinCommission, error)
	Take() (*model.TmpCoinCommission, error)
	Last() (*model.TmpCoinCommission, error)
	Find() ([]*model.TmpCoinCommission, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TmpCoinCommission, err error)
	FindInBatches(result *[]*model.TmpCoinCommission, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TmpCoinCommission) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITmpCoinCommissionDo
	Assign(attrs ...field.AssignExpr) ITmpCoinCommissionDo
	Joins(fields ...field.RelationField) ITmpCoinCommissionDo
	Preload(fields ...field.RelationField) ITmpCoinCommissionDo
	FirstOrInit() (*model.TmpCoinCommission, error)
	FirstOrCreate() (*model.TmpCoinCommission, error)
	FindByPage(offset int, limit int) (result []*model.TmpCoinCommission, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITmpCoinCommissionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tmpCoinCommissionDo) Debug() ITmpCoinCommissionDo {
	return t.withDO(t.DO.Debug())
}

func (t tmpCoinCommissionDo) WithContext(ctx context.Context) ITmpCoinCommissionDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tmpCoinCommissionDo) ReadDB() ITmpCoinCommissionDo {
	return t.Clauses(dbresolver.Read)
}

func (t tmpCoinCommissionDo) WriteDB() ITmpCoinCommissionDo {
	return t.Clauses(dbresolver.Write)
}

func (t tmpCoinCommissionDo) Session(config *gorm.Session) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Session(config))
}

func (t tmpCoinCommissionDo) Clauses(conds ...clause.Expression) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tmpCoinCommissionDo) Returning(value interface{}, columns ...string) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tmpCoinCommissionDo) Not(conds ...gen.Condition) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tmpCoinCommissionDo) Or(conds ...gen.Condition) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tmpCoinCommissionDo) Select(conds ...field.Expr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tmpCoinCommissionDo) Where(conds ...gen.Condition) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tmpCoinCommissionDo) Order(conds ...field.Expr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tmpCoinCommissionDo) Distinct(cols ...field.Expr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tmpCoinCommissionDo) Omit(cols ...field.Expr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tmpCoinCommissionDo) Join(table schema.Tabler, on ...field.Expr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tmpCoinCommissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tmpCoinCommissionDo) RightJoin(table schema.Tabler, on ...field.Expr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tmpCoinCommissionDo) Group(cols ...field.Expr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tmpCoinCommissionDo) Having(conds ...gen.Condition) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tmpCoinCommissionDo) Limit(limit int) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tmpCoinCommissionDo) Offset(offset int) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tmpCoinCommissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tmpCoinCommissionDo) Unscoped() ITmpCoinCommissionDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tmpCoinCommissionDo) Create(values ...*model.TmpCoinCommission) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tmpCoinCommissionDo) CreateInBatches(values []*model.TmpCoinCommission, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tmpCoinCommissionDo) Save(values ...*model.TmpCoinCommission) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tmpCoinCommissionDo) First() (*model.TmpCoinCommission, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmpCoinCommission), nil
	}
}

func (t tmpCoinCommissionDo) Take() (*model.TmpCoinCommission, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmpCoinCommission), nil
	}
}

func (t tmpCoinCommissionDo) Last() (*model.TmpCoinCommission, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmpCoinCommission), nil
	}
}

func (t tmpCoinCommissionDo) Find() ([]*model.TmpCoinCommission, error) {
	result, err := t.DO.Find()
	return result.([]*model.TmpCoinCommission), err
}

func (t tmpCoinCommissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TmpCoinCommission, err error) {
	buf := make([]*model.TmpCoinCommission, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tmpCoinCommissionDo) FindInBatches(result *[]*model.TmpCoinCommission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tmpCoinCommissionDo) Attrs(attrs ...field.AssignExpr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tmpCoinCommissionDo) Assign(attrs ...field.AssignExpr) ITmpCoinCommissionDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tmpCoinCommissionDo) Joins(fields ...field.RelationField) ITmpCoinCommissionDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tmpCoinCommissionDo) Preload(fields ...field.RelationField) ITmpCoinCommissionDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tmpCoinCommissionDo) FirstOrInit() (*model.TmpCoinCommission, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmpCoinCommission), nil
	}
}

func (t tmpCoinCommissionDo) FirstOrCreate() (*model.TmpCoinCommission, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmpCoinCommission), nil
	}
}

func (t tmpCoinCommissionDo) FindByPage(offset int, limit int) (result []*model.TmpCoinCommission, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tmpCoinCommissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tmpCoinCommissionDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tmpCoinCommissionDo) Delete(models ...*model.TmpCoinCommission) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tmpCoinCommissionDo) withDO(do gen.Dao) *tmpCoinCommissionDo {
	t.DO = *do.(*gen.DO)
	return t
}
