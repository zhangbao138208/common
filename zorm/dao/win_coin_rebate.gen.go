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

func newWinCoinRebate(db *gorm.DB, opts ...gen.DOOption) winCoinRebate {
	_winCoinRebate := winCoinRebate{}

	_winCoinRebate.winCoinRebateDo.UseDB(db, opts...)
	_winCoinRebate.winCoinRebateDo.UseModel(&model.WinCoinRebate{})

	tableName := _winCoinRebate.winCoinRebateDo.TableName()
	_winCoinRebate.ALL = field.NewAsterisk(tableName)
	_winCoinRebate.ID = field.NewInt64(tableName, "id")
	_winCoinRebate.UID = field.NewInt64(tableName, "uid")
	_winCoinRebate.Username = field.NewString(tableName, "username")
	_winCoinRebate.LevelID = field.NewInt64(tableName, "level_id")
	_winCoinRebate.SettleDate = field.NewString(tableName, "settle_date")
	_winCoinRebate.VaildStake = field.NewField(tableName, "vaild_stake")
	_winCoinRebate.GroupID = field.NewInt64(tableName, "group_id")
	_winCoinRebate.RebateRate = field.NewField(tableName, "rebate_rate")
	_winCoinRebate.RabateCoin = field.NewField(tableName, "rabate_coin")
	_winCoinRebate.Status = field.NewInt64(tableName, "status")
	_winCoinRebate.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinRebate.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winCoinRebate.fillFieldMap()

	return _winCoinRebate
}

type winCoinRebate struct {
	winCoinRebateDo

	ALL        field.Asterisk
	ID         field.Int64
	UID        field.Int64  // UID
	Username   field.String // 用户名
	LevelID    field.Int64  // 会员等级
	SettleDate field.String // 结算日期
	VaildStake field.Field  // 有效投注额
	GroupID    field.Int64  // 游戏类型
	RebateRate field.Field  // 返水比例
	RabateCoin field.Field  // 返水金额
	Status     field.Int64  // 状态:1-已发放 0-未发放
	CreatedAt  field.Int64
	UpdatedAt  field.Int64

	fieldMap map[string]field.Expr
}

func (w winCoinRebate) Table(newTableName string) *winCoinRebate {
	w.winCoinRebateDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinRebate) As(alias string) *winCoinRebate {
	w.winCoinRebateDo.DO = *(w.winCoinRebateDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinRebate) updateTableName(table string) *winCoinRebate {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.LevelID = field.NewInt64(table, "level_id")
	w.SettleDate = field.NewString(table, "settle_date")
	w.VaildStake = field.NewField(table, "vaild_stake")
	w.GroupID = field.NewInt64(table, "group_id")
	w.RebateRate = field.NewField(table, "rebate_rate")
	w.RabateCoin = field.NewField(table, "rabate_coin")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinRebate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinRebate) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 12)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["level_id"] = w.LevelID
	w.fieldMap["settle_date"] = w.SettleDate
	w.fieldMap["vaild_stake"] = w.VaildStake
	w.fieldMap["group_id"] = w.GroupID
	w.fieldMap["rebate_rate"] = w.RebateRate
	w.fieldMap["rabate_coin"] = w.RabateCoin
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCoinRebate) clone(db *gorm.DB) winCoinRebate {
	w.winCoinRebateDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinRebate) replaceDB(db *gorm.DB) winCoinRebate {
	w.winCoinRebateDo.ReplaceDB(db)
	return w
}

type winCoinRebateDo struct{ gen.DO }

type IWinCoinRebateDo interface {
	gen.SubQuery
	Debug() IWinCoinRebateDo
	WithContext(ctx context.Context) IWinCoinRebateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinRebateDo
	WriteDB() IWinCoinRebateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinRebateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinRebateDo
	Not(conds ...gen.Condition) IWinCoinRebateDo
	Or(conds ...gen.Condition) IWinCoinRebateDo
	Select(conds ...field.Expr) IWinCoinRebateDo
	Where(conds ...gen.Condition) IWinCoinRebateDo
	Order(conds ...field.Expr) IWinCoinRebateDo
	Distinct(cols ...field.Expr) IWinCoinRebateDo
	Omit(cols ...field.Expr) IWinCoinRebateDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinRebateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinRebateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinRebateDo
	Group(cols ...field.Expr) IWinCoinRebateDo
	Having(conds ...gen.Condition) IWinCoinRebateDo
	Limit(limit int) IWinCoinRebateDo
	Offset(offset int) IWinCoinRebateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinRebateDo
	Unscoped() IWinCoinRebateDo
	Create(values ...*model.WinCoinRebate) error
	CreateInBatches(values []*model.WinCoinRebate, batchSize int) error
	Save(values ...*model.WinCoinRebate) error
	First() (*model.WinCoinRebate, error)
	Take() (*model.WinCoinRebate, error)
	Last() (*model.WinCoinRebate, error)
	Find() ([]*model.WinCoinRebate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinRebate, err error)
	FindInBatches(result *[]*model.WinCoinRebate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinRebate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinRebateDo
	Assign(attrs ...field.AssignExpr) IWinCoinRebateDo
	Joins(fields ...field.RelationField) IWinCoinRebateDo
	Preload(fields ...field.RelationField) IWinCoinRebateDo
	FirstOrInit() (*model.WinCoinRebate, error)
	FirstOrCreate() (*model.WinCoinRebate, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinRebate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinRebateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinRebateDo) Debug() IWinCoinRebateDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinRebateDo) WithContext(ctx context.Context) IWinCoinRebateDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinRebateDo) ReadDB() IWinCoinRebateDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinRebateDo) WriteDB() IWinCoinRebateDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinRebateDo) Session(config *gorm.Session) IWinCoinRebateDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinRebateDo) Clauses(conds ...clause.Expression) IWinCoinRebateDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinRebateDo) Returning(value interface{}, columns ...string) IWinCoinRebateDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinRebateDo) Not(conds ...gen.Condition) IWinCoinRebateDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinRebateDo) Or(conds ...gen.Condition) IWinCoinRebateDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinRebateDo) Select(conds ...field.Expr) IWinCoinRebateDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinRebateDo) Where(conds ...gen.Condition) IWinCoinRebateDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinRebateDo) Order(conds ...field.Expr) IWinCoinRebateDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinRebateDo) Distinct(cols ...field.Expr) IWinCoinRebateDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinRebateDo) Omit(cols ...field.Expr) IWinCoinRebateDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinRebateDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinRebateDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinRebateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinRebateDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinRebateDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinRebateDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinRebateDo) Group(cols ...field.Expr) IWinCoinRebateDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinRebateDo) Having(conds ...gen.Condition) IWinCoinRebateDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinRebateDo) Limit(limit int) IWinCoinRebateDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinRebateDo) Offset(offset int) IWinCoinRebateDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinRebateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinRebateDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinRebateDo) Unscoped() IWinCoinRebateDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinRebateDo) Create(values ...*model.WinCoinRebate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinRebateDo) CreateInBatches(values []*model.WinCoinRebate, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinRebateDo) Save(values ...*model.WinCoinRebate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinRebateDo) First() (*model.WinCoinRebate, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRebate), nil
	}
}

func (w winCoinRebateDo) Take() (*model.WinCoinRebate, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRebate), nil
	}
}

func (w winCoinRebateDo) Last() (*model.WinCoinRebate, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRebate), nil
	}
}

func (w winCoinRebateDo) Find() ([]*model.WinCoinRebate, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinRebate), err
}

func (w winCoinRebateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinRebate, err error) {
	buf := make([]*model.WinCoinRebate, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinRebateDo) FindInBatches(result *[]*model.WinCoinRebate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinRebateDo) Attrs(attrs ...field.AssignExpr) IWinCoinRebateDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinRebateDo) Assign(attrs ...field.AssignExpr) IWinCoinRebateDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinRebateDo) Joins(fields ...field.RelationField) IWinCoinRebateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinRebateDo) Preload(fields ...field.RelationField) IWinCoinRebateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinRebateDo) FirstOrInit() (*model.WinCoinRebate, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRebate), nil
	}
}

func (w winCoinRebateDo) FirstOrCreate() (*model.WinCoinRebate, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRebate), nil
	}
}

func (w winCoinRebateDo) FindByPage(offset int, limit int) (result []*model.WinCoinRebate, count int64, err error) {
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

func (w winCoinRebateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinRebateDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinRebateDo) Delete(models ...*model.WinCoinRebate) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinRebateDo) withDO(do gen.Dao) *winCoinRebateDo {
	w.DO = *do.(*gen.DO)
	return w
}