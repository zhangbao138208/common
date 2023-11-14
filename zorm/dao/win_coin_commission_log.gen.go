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

func newWinCoinCommissionLog(db *gorm.DB, opts ...gen.DOOption) winCoinCommissionLog {
	_winCoinCommissionLog := winCoinCommissionLog{}

	_winCoinCommissionLog.winCoinCommissionLogDo.UseDB(db, opts...)
	_winCoinCommissionLog.winCoinCommissionLogDo.UseModel(&model.WinCoinCommissionLog{})

	tableName := _winCoinCommissionLog.winCoinCommissionLogDo.TableName()
	_winCoinCommissionLog.ALL = field.NewAsterisk(tableName)
	_winCoinCommissionLog.ID = field.NewInt64(tableName, "id")
	_winCoinCommissionLog.UID = field.NewInt64(tableName, "uid")
	_winCoinCommissionLog.Username = field.NewString(tableName, "username")
	_winCoinCommissionLog.Category = field.NewInt64(tableName, "category")
	_winCoinCommissionLog.ReferID = field.NewInt64(tableName, "refer_id")
	_winCoinCommissionLog.Coin = field.NewField(tableName, "coin")
	_winCoinCommissionLog.OutIn = field.NewInt64(tableName, "out_in")
	_winCoinCommissionLog.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinCommissionLog.CoinAfter = field.NewField(tableName, "coin_after")
	_winCoinCommissionLog.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinCommissionLog.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winCoinCommissionLog.fillFieldMap()

	return _winCoinCommissionLog
}

type winCoinCommissionLog struct {
	winCoinCommissionLogDo

	ALL        field.Asterisk
	ID         field.Int64
	UID        field.Int64  // UID
	Username   field.String // 用户名
	Category   field.Int64  // 类型:1-返佣 2- 佣金充值 3-佣金提款
	ReferID    field.Int64  // 关联ID
	Coin       field.Field  // 金额
	OutIn      field.Int64  // 收支类型:0-支出 1-收入
	CoinBefore field.Field  // 前金额
	CoinAfter  field.Field  // 帐变后金额
	CreatedAt  field.Int64
	UpdatedAt  field.Int64

	fieldMap map[string]field.Expr
}

func (w winCoinCommissionLog) Table(newTableName string) *winCoinCommissionLog {
	w.winCoinCommissionLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinCommissionLog) As(alias string) *winCoinCommissionLog {
	w.winCoinCommissionLogDo.DO = *(w.winCoinCommissionLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinCommissionLog) updateTableName(table string) *winCoinCommissionLog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Category = field.NewInt64(table, "category")
	w.ReferID = field.NewInt64(table, "refer_id")
	w.Coin = field.NewField(table, "coin")
	w.OutIn = field.NewInt64(table, "out_in")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.CoinAfter = field.NewField(table, "coin_after")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinCommissionLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinCommissionLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["category"] = w.Category
	w.fieldMap["refer_id"] = w.ReferID
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["out_in"] = w.OutIn
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["coin_after"] = w.CoinAfter
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCoinCommissionLog) clone(db *gorm.DB) winCoinCommissionLog {
	w.winCoinCommissionLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinCommissionLog) replaceDB(db *gorm.DB) winCoinCommissionLog {
	w.winCoinCommissionLogDo.ReplaceDB(db)
	return w
}

type winCoinCommissionLogDo struct{ gen.DO }

type IWinCoinCommissionLogDo interface {
	gen.SubQuery
	Debug() IWinCoinCommissionLogDo
	WithContext(ctx context.Context) IWinCoinCommissionLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinCommissionLogDo
	WriteDB() IWinCoinCommissionLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinCommissionLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinCommissionLogDo
	Not(conds ...gen.Condition) IWinCoinCommissionLogDo
	Or(conds ...gen.Condition) IWinCoinCommissionLogDo
	Select(conds ...field.Expr) IWinCoinCommissionLogDo
	Where(conds ...gen.Condition) IWinCoinCommissionLogDo
	Order(conds ...field.Expr) IWinCoinCommissionLogDo
	Distinct(cols ...field.Expr) IWinCoinCommissionLogDo
	Omit(cols ...field.Expr) IWinCoinCommissionLogDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinCommissionLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionLogDo
	Group(cols ...field.Expr) IWinCoinCommissionLogDo
	Having(conds ...gen.Condition) IWinCoinCommissionLogDo
	Limit(limit int) IWinCoinCommissionLogDo
	Offset(offset int) IWinCoinCommissionLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinCommissionLogDo
	Unscoped() IWinCoinCommissionLogDo
	Create(values ...*model.WinCoinCommissionLog) error
	CreateInBatches(values []*model.WinCoinCommissionLog, batchSize int) error
	Save(values ...*model.WinCoinCommissionLog) error
	First() (*model.WinCoinCommissionLog, error)
	Take() (*model.WinCoinCommissionLog, error)
	Last() (*model.WinCoinCommissionLog, error)
	Find() ([]*model.WinCoinCommissionLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinCommissionLog, err error)
	FindInBatches(result *[]*model.WinCoinCommissionLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinCommissionLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinCommissionLogDo
	Assign(attrs ...field.AssignExpr) IWinCoinCommissionLogDo
	Joins(fields ...field.RelationField) IWinCoinCommissionLogDo
	Preload(fields ...field.RelationField) IWinCoinCommissionLogDo
	FirstOrInit() (*model.WinCoinCommissionLog, error)
	FirstOrCreate() (*model.WinCoinCommissionLog, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinCommissionLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinCommissionLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinCommissionLogDo) Debug() IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinCommissionLogDo) WithContext(ctx context.Context) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinCommissionLogDo) ReadDB() IWinCoinCommissionLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinCommissionLogDo) WriteDB() IWinCoinCommissionLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinCommissionLogDo) Session(config *gorm.Session) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinCommissionLogDo) Clauses(conds ...clause.Expression) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinCommissionLogDo) Returning(value interface{}, columns ...string) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinCommissionLogDo) Not(conds ...gen.Condition) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinCommissionLogDo) Or(conds ...gen.Condition) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinCommissionLogDo) Select(conds ...field.Expr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinCommissionLogDo) Where(conds ...gen.Condition) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinCommissionLogDo) Order(conds ...field.Expr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinCommissionLogDo) Distinct(cols ...field.Expr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinCommissionLogDo) Omit(cols ...field.Expr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinCommissionLogDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinCommissionLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinCommissionLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinCommissionLogDo) Group(cols ...field.Expr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinCommissionLogDo) Having(conds ...gen.Condition) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinCommissionLogDo) Limit(limit int) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinCommissionLogDo) Offset(offset int) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinCommissionLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinCommissionLogDo) Unscoped() IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinCommissionLogDo) Create(values ...*model.WinCoinCommissionLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinCommissionLogDo) CreateInBatches(values []*model.WinCoinCommissionLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinCommissionLogDo) Save(values ...*model.WinCoinCommissionLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinCommissionLogDo) First() (*model.WinCoinCommissionLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionLog), nil
	}
}

func (w winCoinCommissionLogDo) Take() (*model.WinCoinCommissionLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionLog), nil
	}
}

func (w winCoinCommissionLogDo) Last() (*model.WinCoinCommissionLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionLog), nil
	}
}

func (w winCoinCommissionLogDo) Find() ([]*model.WinCoinCommissionLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinCommissionLog), err
}

func (w winCoinCommissionLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinCommissionLog, err error) {
	buf := make([]*model.WinCoinCommissionLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinCommissionLogDo) FindInBatches(result *[]*model.WinCoinCommissionLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinCommissionLogDo) Attrs(attrs ...field.AssignExpr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinCommissionLogDo) Assign(attrs ...field.AssignExpr) IWinCoinCommissionLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinCommissionLogDo) Joins(fields ...field.RelationField) IWinCoinCommissionLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinCommissionLogDo) Preload(fields ...field.RelationField) IWinCoinCommissionLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinCommissionLogDo) FirstOrInit() (*model.WinCoinCommissionLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionLog), nil
	}
}

func (w winCoinCommissionLogDo) FirstOrCreate() (*model.WinCoinCommissionLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionLog), nil
	}
}

func (w winCoinCommissionLogDo) FindByPage(offset int, limit int) (result []*model.WinCoinCommissionLog, count int64, err error) {
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

func (w winCoinCommissionLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinCommissionLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinCommissionLogDo) Delete(models ...*model.WinCoinCommissionLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinCommissionLogDo) withDO(do gen.Dao) *winCoinCommissionLogDo {
	w.DO = *do.(*gen.DO)
	return w
}