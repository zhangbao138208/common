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

func newWinAgentTransferLog(db *gorm.DB, opts ...gen.DOOption) winAgentTransferLog {
	_winAgentTransferLog := winAgentTransferLog{}

	_winAgentTransferLog.winAgentTransferLogDo.UseDB(db, opts...)
	_winAgentTransferLog.winAgentTransferLogDo.UseModel(&model.WinAgentTransferLog{})

	tableName := _winAgentTransferLog.winAgentTransferLogDo.TableName()
	_winAgentTransferLog.ALL = field.NewAsterisk(tableName)
	_winAgentTransferLog.ID = field.NewInt64(tableName, "id")
	_winAgentTransferLog.FromUID = field.NewInt64(tableName, "from_uid")
	_winAgentTransferLog.FromUsername = field.NewString(tableName, "from_username")
	_winAgentTransferLog.ToUID = field.NewInt64(tableName, "to_uid")
	_winAgentTransferLog.ToUsername = field.NewString(tableName, "to_username")
	_winAgentTransferLog.OperationID = field.NewInt64(tableName, "operation_id")
	_winAgentTransferLog.OperationName = field.NewString(tableName, "operation_name")
	_winAgentTransferLog.CreatedAt = field.NewInt64(tableName, "created_at")
	_winAgentTransferLog.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winAgentTransferLog.fillFieldMap()

	return _winAgentTransferLog
}

type winAgentTransferLog struct {
	winAgentTransferLogDo

	ALL           field.Asterisk
	ID            field.Int64  // ID
	FromUID       field.Int64  // 转移代理
	FromUsername  field.String // 转移代理名
	ToUID         field.Int64  // 转移到代理UID
	ToUsername    field.String // 转移到代理名
	OperationID   field.Int64  // 操作人
	OperationName field.String // 操作人姓名
	CreatedAt     field.Int64  // 创建时间
	UpdatedAt     field.Int64  // 修改时间

	fieldMap map[string]field.Expr
}

func (w winAgentTransferLog) Table(newTableName string) *winAgentTransferLog {
	w.winAgentTransferLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAgentTransferLog) As(alias string) *winAgentTransferLog {
	w.winAgentTransferLogDo.DO = *(w.winAgentTransferLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAgentTransferLog) updateTableName(table string) *winAgentTransferLog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.FromUID = field.NewInt64(table, "from_uid")
	w.FromUsername = field.NewString(table, "from_username")
	w.ToUID = field.NewInt64(table, "to_uid")
	w.ToUsername = field.NewString(table, "to_username")
	w.OperationID = field.NewInt64(table, "operation_id")
	w.OperationName = field.NewString(table, "operation_name")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winAgentTransferLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAgentTransferLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 9)
	w.fieldMap["id"] = w.ID
	w.fieldMap["from_uid"] = w.FromUID
	w.fieldMap["from_username"] = w.FromUsername
	w.fieldMap["to_uid"] = w.ToUID
	w.fieldMap["to_username"] = w.ToUsername
	w.fieldMap["operation_id"] = w.OperationID
	w.fieldMap["operation_name"] = w.OperationName
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winAgentTransferLog) clone(db *gorm.DB) winAgentTransferLog {
	w.winAgentTransferLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAgentTransferLog) replaceDB(db *gorm.DB) winAgentTransferLog {
	w.winAgentTransferLogDo.ReplaceDB(db)
	return w
}

type winAgentTransferLogDo struct{ gen.DO }

type IWinAgentTransferLogDo interface {
	gen.SubQuery
	Debug() IWinAgentTransferLogDo
	WithContext(ctx context.Context) IWinAgentTransferLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinAgentTransferLogDo
	WriteDB() IWinAgentTransferLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinAgentTransferLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinAgentTransferLogDo
	Not(conds ...gen.Condition) IWinAgentTransferLogDo
	Or(conds ...gen.Condition) IWinAgentTransferLogDo
	Select(conds ...field.Expr) IWinAgentTransferLogDo
	Where(conds ...gen.Condition) IWinAgentTransferLogDo
	Order(conds ...field.Expr) IWinAgentTransferLogDo
	Distinct(cols ...field.Expr) IWinAgentTransferLogDo
	Omit(cols ...field.Expr) IWinAgentTransferLogDo
	Join(table schema.Tabler, on ...field.Expr) IWinAgentTransferLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinAgentTransferLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinAgentTransferLogDo
	Group(cols ...field.Expr) IWinAgentTransferLogDo
	Having(conds ...gen.Condition) IWinAgentTransferLogDo
	Limit(limit int) IWinAgentTransferLogDo
	Offset(offset int) IWinAgentTransferLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAgentTransferLogDo
	Unscoped() IWinAgentTransferLogDo
	Create(values ...*model.WinAgentTransferLog) error
	CreateInBatches(values []*model.WinAgentTransferLog, batchSize int) error
	Save(values ...*model.WinAgentTransferLog) error
	First() (*model.WinAgentTransferLog, error)
	Take() (*model.WinAgentTransferLog, error)
	Last() (*model.WinAgentTransferLog, error)
	Find() ([]*model.WinAgentTransferLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAgentTransferLog, err error)
	FindInBatches(result *[]*model.WinAgentTransferLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinAgentTransferLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinAgentTransferLogDo
	Assign(attrs ...field.AssignExpr) IWinAgentTransferLogDo
	Joins(fields ...field.RelationField) IWinAgentTransferLogDo
	Preload(fields ...field.RelationField) IWinAgentTransferLogDo
	FirstOrInit() (*model.WinAgentTransferLog, error)
	FirstOrCreate() (*model.WinAgentTransferLog, error)
	FindByPage(offset int, limit int) (result []*model.WinAgentTransferLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinAgentTransferLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winAgentTransferLogDo) Debug() IWinAgentTransferLogDo {
	return w.withDO(w.DO.Debug())
}

func (w winAgentTransferLogDo) WithContext(ctx context.Context) IWinAgentTransferLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAgentTransferLogDo) ReadDB() IWinAgentTransferLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAgentTransferLogDo) WriteDB() IWinAgentTransferLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAgentTransferLogDo) Session(config *gorm.Session) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAgentTransferLogDo) Clauses(conds ...clause.Expression) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAgentTransferLogDo) Returning(value interface{}, columns ...string) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAgentTransferLogDo) Not(conds ...gen.Condition) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAgentTransferLogDo) Or(conds ...gen.Condition) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAgentTransferLogDo) Select(conds ...field.Expr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAgentTransferLogDo) Where(conds ...gen.Condition) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAgentTransferLogDo) Order(conds ...field.Expr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAgentTransferLogDo) Distinct(cols ...field.Expr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAgentTransferLogDo) Omit(cols ...field.Expr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAgentTransferLogDo) Join(table schema.Tabler, on ...field.Expr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAgentTransferLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAgentTransferLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAgentTransferLogDo) Group(cols ...field.Expr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAgentTransferLogDo) Having(conds ...gen.Condition) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAgentTransferLogDo) Limit(limit int) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAgentTransferLogDo) Offset(offset int) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAgentTransferLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAgentTransferLogDo) Unscoped() IWinAgentTransferLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAgentTransferLogDo) Create(values ...*model.WinAgentTransferLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAgentTransferLogDo) CreateInBatches(values []*model.WinAgentTransferLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAgentTransferLogDo) Save(values ...*model.WinAgentTransferLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAgentTransferLogDo) First() (*model.WinAgentTransferLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentTransferLog), nil
	}
}

func (w winAgentTransferLogDo) Take() (*model.WinAgentTransferLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentTransferLog), nil
	}
}

func (w winAgentTransferLogDo) Last() (*model.WinAgentTransferLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentTransferLog), nil
	}
}

func (w winAgentTransferLogDo) Find() ([]*model.WinAgentTransferLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAgentTransferLog), err
}

func (w winAgentTransferLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAgentTransferLog, err error) {
	buf := make([]*model.WinAgentTransferLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAgentTransferLogDo) FindInBatches(result *[]*model.WinAgentTransferLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAgentTransferLogDo) Attrs(attrs ...field.AssignExpr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAgentTransferLogDo) Assign(attrs ...field.AssignExpr) IWinAgentTransferLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAgentTransferLogDo) Joins(fields ...field.RelationField) IWinAgentTransferLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAgentTransferLogDo) Preload(fields ...field.RelationField) IWinAgentTransferLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAgentTransferLogDo) FirstOrInit() (*model.WinAgentTransferLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentTransferLog), nil
	}
}

func (w winAgentTransferLogDo) FirstOrCreate() (*model.WinAgentTransferLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentTransferLog), nil
	}
}

func (w winAgentTransferLogDo) FindByPage(offset int, limit int) (result []*model.WinAgentTransferLog, count int64, err error) {
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

func (w winAgentTransferLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAgentTransferLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAgentTransferLogDo) Delete(models ...*model.WinAgentTransferLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAgentTransferLogDo) withDO(do gen.Dao) *winAgentTransferLogDo {
	w.DO = *do.(*gen.DO)
	return w
}