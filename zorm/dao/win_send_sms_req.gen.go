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

func newWinSendSmsReq(db *gorm.DB, opts ...gen.DOOption) winSendSmsReq {
	_winSendSmsReq := winSendSmsReq{}

	_winSendSmsReq.winSendSmsReqDo.UseDB(db, opts...)
	_winSendSmsReq.winSendSmsReqDo.UseModel(&model.WinSendSmsReq{})

	tableName := _winSendSmsReq.winSendSmsReqDo.TableName()
	_winSendSmsReq.ALL = field.NewAsterisk(tableName)
	_winSendSmsReq.ID = field.NewInt64(tableName, "id")
	_winSendSmsReq.AreaCode = field.NewString(tableName, "area_code")
	_winSendSmsReq.Mobile = field.NewString(tableName, "mobile")
	_winSendSmsReq.Type = field.NewString(tableName, "type")
	_winSendSmsReq.CreateTime = field.NewTime(tableName, "create_time")

	_winSendSmsReq.fillFieldMap()

	return _winSendSmsReq
}

type winSendSmsReq struct {
	winSendSmsReqDo

	ALL        field.Asterisk
	ID         field.Int64
	AreaCode   field.String // 区号
	Mobile     field.String // 手机号
	Type       field.String // 类型
	CreateTime field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (w winSendSmsReq) Table(newTableName string) *winSendSmsReq {
	w.winSendSmsReqDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winSendSmsReq) As(alias string) *winSendSmsReq {
	w.winSendSmsReqDo.DO = *(w.winSendSmsReqDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winSendSmsReq) updateTableName(table string) *winSendSmsReq {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.AreaCode = field.NewString(table, "area_code")
	w.Mobile = field.NewString(table, "mobile")
	w.Type = field.NewString(table, "type")
	w.CreateTime = field.NewTime(table, "create_time")

	w.fillFieldMap()

	return w
}

func (w *winSendSmsReq) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winSendSmsReq) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 5)
	w.fieldMap["id"] = w.ID
	w.fieldMap["area_code"] = w.AreaCode
	w.fieldMap["mobile"] = w.Mobile
	w.fieldMap["type"] = w.Type
	w.fieldMap["create_time"] = w.CreateTime
}

func (w winSendSmsReq) clone(db *gorm.DB) winSendSmsReq {
	w.winSendSmsReqDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winSendSmsReq) replaceDB(db *gorm.DB) winSendSmsReq {
	w.winSendSmsReqDo.ReplaceDB(db)
	return w
}

type winSendSmsReqDo struct{ gen.DO }

type IWinSendSmsReqDo interface {
	gen.SubQuery
	Debug() IWinSendSmsReqDo
	WithContext(ctx context.Context) IWinSendSmsReqDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinSendSmsReqDo
	WriteDB() IWinSendSmsReqDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinSendSmsReqDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinSendSmsReqDo
	Not(conds ...gen.Condition) IWinSendSmsReqDo
	Or(conds ...gen.Condition) IWinSendSmsReqDo
	Select(conds ...field.Expr) IWinSendSmsReqDo
	Where(conds ...gen.Condition) IWinSendSmsReqDo
	Order(conds ...field.Expr) IWinSendSmsReqDo
	Distinct(cols ...field.Expr) IWinSendSmsReqDo
	Omit(cols ...field.Expr) IWinSendSmsReqDo
	Join(table schema.Tabler, on ...field.Expr) IWinSendSmsReqDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinSendSmsReqDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinSendSmsReqDo
	Group(cols ...field.Expr) IWinSendSmsReqDo
	Having(conds ...gen.Condition) IWinSendSmsReqDo
	Limit(limit int) IWinSendSmsReqDo
	Offset(offset int) IWinSendSmsReqDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinSendSmsReqDo
	Unscoped() IWinSendSmsReqDo
	Create(values ...*model.WinSendSmsReq) error
	CreateInBatches(values []*model.WinSendSmsReq, batchSize int) error
	Save(values ...*model.WinSendSmsReq) error
	First() (*model.WinSendSmsReq, error)
	Take() (*model.WinSendSmsReq, error)
	Last() (*model.WinSendSmsReq, error)
	Find() ([]*model.WinSendSmsReq, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinSendSmsReq, err error)
	FindInBatches(result *[]*model.WinSendSmsReq, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinSendSmsReq) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinSendSmsReqDo
	Assign(attrs ...field.AssignExpr) IWinSendSmsReqDo
	Joins(fields ...field.RelationField) IWinSendSmsReqDo
	Preload(fields ...field.RelationField) IWinSendSmsReqDo
	FirstOrInit() (*model.WinSendSmsReq, error)
	FirstOrCreate() (*model.WinSendSmsReq, error)
	FindByPage(offset int, limit int) (result []*model.WinSendSmsReq, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinSendSmsReqDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winSendSmsReqDo) Debug() IWinSendSmsReqDo {
	return w.withDO(w.DO.Debug())
}

func (w winSendSmsReqDo) WithContext(ctx context.Context) IWinSendSmsReqDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winSendSmsReqDo) ReadDB() IWinSendSmsReqDo {
	return w.Clauses(dbresolver.Read)
}

func (w winSendSmsReqDo) WriteDB() IWinSendSmsReqDo {
	return w.Clauses(dbresolver.Write)
}

func (w winSendSmsReqDo) Session(config *gorm.Session) IWinSendSmsReqDo {
	return w.withDO(w.DO.Session(config))
}

func (w winSendSmsReqDo) Clauses(conds ...clause.Expression) IWinSendSmsReqDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winSendSmsReqDo) Returning(value interface{}, columns ...string) IWinSendSmsReqDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winSendSmsReqDo) Not(conds ...gen.Condition) IWinSendSmsReqDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winSendSmsReqDo) Or(conds ...gen.Condition) IWinSendSmsReqDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winSendSmsReqDo) Select(conds ...field.Expr) IWinSendSmsReqDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winSendSmsReqDo) Where(conds ...gen.Condition) IWinSendSmsReqDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winSendSmsReqDo) Order(conds ...field.Expr) IWinSendSmsReqDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winSendSmsReqDo) Distinct(cols ...field.Expr) IWinSendSmsReqDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winSendSmsReqDo) Omit(cols ...field.Expr) IWinSendSmsReqDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winSendSmsReqDo) Join(table schema.Tabler, on ...field.Expr) IWinSendSmsReqDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winSendSmsReqDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinSendSmsReqDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winSendSmsReqDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinSendSmsReqDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winSendSmsReqDo) Group(cols ...field.Expr) IWinSendSmsReqDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winSendSmsReqDo) Having(conds ...gen.Condition) IWinSendSmsReqDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winSendSmsReqDo) Limit(limit int) IWinSendSmsReqDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winSendSmsReqDo) Offset(offset int) IWinSendSmsReqDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winSendSmsReqDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinSendSmsReqDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winSendSmsReqDo) Unscoped() IWinSendSmsReqDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winSendSmsReqDo) Create(values ...*model.WinSendSmsReq) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winSendSmsReqDo) CreateInBatches(values []*model.WinSendSmsReq, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winSendSmsReqDo) Save(values ...*model.WinSendSmsReq) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winSendSmsReqDo) First() (*model.WinSendSmsReq, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSendSmsReq), nil
	}
}

func (w winSendSmsReqDo) Take() (*model.WinSendSmsReq, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSendSmsReq), nil
	}
}

func (w winSendSmsReqDo) Last() (*model.WinSendSmsReq, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSendSmsReq), nil
	}
}

func (w winSendSmsReqDo) Find() ([]*model.WinSendSmsReq, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinSendSmsReq), err
}

func (w winSendSmsReqDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinSendSmsReq, err error) {
	buf := make([]*model.WinSendSmsReq, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winSendSmsReqDo) FindInBatches(result *[]*model.WinSendSmsReq, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winSendSmsReqDo) Attrs(attrs ...field.AssignExpr) IWinSendSmsReqDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winSendSmsReqDo) Assign(attrs ...field.AssignExpr) IWinSendSmsReqDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winSendSmsReqDo) Joins(fields ...field.RelationField) IWinSendSmsReqDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winSendSmsReqDo) Preload(fields ...field.RelationField) IWinSendSmsReqDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winSendSmsReqDo) FirstOrInit() (*model.WinSendSmsReq, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSendSmsReq), nil
	}
}

func (w winSendSmsReqDo) FirstOrCreate() (*model.WinSendSmsReq, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSendSmsReq), nil
	}
}

func (w winSendSmsReqDo) FindByPage(offset int, limit int) (result []*model.WinSendSmsReq, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w winSendSmsReqDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winSendSmsReqDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winSendSmsReqDo) Delete(models ...*model.WinSendSmsReq) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winSendSmsReqDo) withDO(do gen.Dao) *winSendSmsReqDo {
	w.DO = *do.(*gen.DO)
	return w
}