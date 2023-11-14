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

func newWinMonthlyTax230629(db *gorm.DB, opts ...gen.DOOption) winMonthlyTax230629 {
	_winMonthlyTax230629 := winMonthlyTax230629{}

	_winMonthlyTax230629.winMonthlyTax230629Do.UseDB(db, opts...)
	_winMonthlyTax230629.winMonthlyTax230629Do.UseModel(&model.WinMonthlyTax230629{})

	tableName := _winMonthlyTax230629.winMonthlyTax230629Do.TableName()
	_winMonthlyTax230629.ALL = field.NewAsterisk(tableName)
	_winMonthlyTax230629.ID = field.NewInt64(tableName, "id")
	_winMonthlyTax230629.GameID = field.NewInt64(tableName, "game_id")
	_winMonthlyTax230629.GameName = field.NewString(tableName, "game_name")
	_winMonthlyTax230629.Year = field.NewInt64(tableName, "year")
	_winMonthlyTax230629.Month = field.NewInt64(tableName, "month")
	_winMonthlyTax230629.CoinBet = field.NewField(tableName, "coin_bet")
	_winMonthlyTax230629.CoinProfit = field.NewField(tableName, "coin_profit")
	_winMonthlyTax230629.Rate = field.NewField(tableName, "rate")
	_winMonthlyTax230629.CoinTax = field.NewField(tableName, "coin_tax")
	_winMonthlyTax230629.CreatedAt = field.NewInt64(tableName, "created_at")
	_winMonthlyTax230629.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winMonthlyTax230629.fillFieldMap()

	return _winMonthlyTax230629
}

type winMonthlyTax230629 struct {
	winMonthlyTax230629Do

	ALL        field.Asterisk
	ID         field.Int64  // ID
	GameID     field.Int64  // 游戏ID
	GameName   field.String // 游戏名称
	Year       field.Int64  // 统计年
	Month      field.Int64  // 统计月
	CoinBet    field.Field  // 投注金额
	CoinProfit field.Field  // 负盈利金额
	Rate       field.Field  // 税收比例
	CoinTax    field.Field  // 税收金额
	CreatedAt  field.Int64  // 创建时间
	UpdatedAt  field.Int64  // 修改时间

	fieldMap map[string]field.Expr
}

func (w winMonthlyTax230629) Table(newTableName string) *winMonthlyTax230629 {
	w.winMonthlyTax230629Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winMonthlyTax230629) As(alias string) *winMonthlyTax230629 {
	w.winMonthlyTax230629Do.DO = *(w.winMonthlyTax230629Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winMonthlyTax230629) updateTableName(table string) *winMonthlyTax230629 {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.GameID = field.NewInt64(table, "game_id")
	w.GameName = field.NewString(table, "game_name")
	w.Year = field.NewInt64(table, "year")
	w.Month = field.NewInt64(table, "month")
	w.CoinBet = field.NewField(table, "coin_bet")
	w.CoinProfit = field.NewField(table, "coin_profit")
	w.Rate = field.NewField(table, "rate")
	w.CoinTax = field.NewField(table, "coin_tax")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winMonthlyTax230629) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winMonthlyTax230629) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["game_id"] = w.GameID
	w.fieldMap["game_name"] = w.GameName
	w.fieldMap["year"] = w.Year
	w.fieldMap["month"] = w.Month
	w.fieldMap["coin_bet"] = w.CoinBet
	w.fieldMap["coin_profit"] = w.CoinProfit
	w.fieldMap["rate"] = w.Rate
	w.fieldMap["coin_tax"] = w.CoinTax
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winMonthlyTax230629) clone(db *gorm.DB) winMonthlyTax230629 {
	w.winMonthlyTax230629Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winMonthlyTax230629) replaceDB(db *gorm.DB) winMonthlyTax230629 {
	w.winMonthlyTax230629Do.ReplaceDB(db)
	return w
}

type winMonthlyTax230629Do struct{ gen.DO }

type IWinMonthlyTax230629Do interface {
	gen.SubQuery
	Debug() IWinMonthlyTax230629Do
	WithContext(ctx context.Context) IWinMonthlyTax230629Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinMonthlyTax230629Do
	WriteDB() IWinMonthlyTax230629Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinMonthlyTax230629Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinMonthlyTax230629Do
	Not(conds ...gen.Condition) IWinMonthlyTax230629Do
	Or(conds ...gen.Condition) IWinMonthlyTax230629Do
	Select(conds ...field.Expr) IWinMonthlyTax230629Do
	Where(conds ...gen.Condition) IWinMonthlyTax230629Do
	Order(conds ...field.Expr) IWinMonthlyTax230629Do
	Distinct(cols ...field.Expr) IWinMonthlyTax230629Do
	Omit(cols ...field.Expr) IWinMonthlyTax230629Do
	Join(table schema.Tabler, on ...field.Expr) IWinMonthlyTax230629Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinMonthlyTax230629Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinMonthlyTax230629Do
	Group(cols ...field.Expr) IWinMonthlyTax230629Do
	Having(conds ...gen.Condition) IWinMonthlyTax230629Do
	Limit(limit int) IWinMonthlyTax230629Do
	Offset(offset int) IWinMonthlyTax230629Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinMonthlyTax230629Do
	Unscoped() IWinMonthlyTax230629Do
	Create(values ...*model.WinMonthlyTax230629) error
	CreateInBatches(values []*model.WinMonthlyTax230629, batchSize int) error
	Save(values ...*model.WinMonthlyTax230629) error
	First() (*model.WinMonthlyTax230629, error)
	Take() (*model.WinMonthlyTax230629, error)
	Last() (*model.WinMonthlyTax230629, error)
	Find() ([]*model.WinMonthlyTax230629, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinMonthlyTax230629, err error)
	FindInBatches(result *[]*model.WinMonthlyTax230629, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinMonthlyTax230629) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinMonthlyTax230629Do
	Assign(attrs ...field.AssignExpr) IWinMonthlyTax230629Do
	Joins(fields ...field.RelationField) IWinMonthlyTax230629Do
	Preload(fields ...field.RelationField) IWinMonthlyTax230629Do
	FirstOrInit() (*model.WinMonthlyTax230629, error)
	FirstOrCreate() (*model.WinMonthlyTax230629, error)
	FindByPage(offset int, limit int) (result []*model.WinMonthlyTax230629, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinMonthlyTax230629Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winMonthlyTax230629Do) Debug() IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Debug())
}

func (w winMonthlyTax230629Do) WithContext(ctx context.Context) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winMonthlyTax230629Do) ReadDB() IWinMonthlyTax230629Do {
	return w.Clauses(dbresolver.Read)
}

func (w winMonthlyTax230629Do) WriteDB() IWinMonthlyTax230629Do {
	return w.Clauses(dbresolver.Write)
}

func (w winMonthlyTax230629Do) Session(config *gorm.Session) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Session(config))
}

func (w winMonthlyTax230629Do) Clauses(conds ...clause.Expression) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winMonthlyTax230629Do) Returning(value interface{}, columns ...string) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winMonthlyTax230629Do) Not(conds ...gen.Condition) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winMonthlyTax230629Do) Or(conds ...gen.Condition) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winMonthlyTax230629Do) Select(conds ...field.Expr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winMonthlyTax230629Do) Where(conds ...gen.Condition) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winMonthlyTax230629Do) Order(conds ...field.Expr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winMonthlyTax230629Do) Distinct(cols ...field.Expr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winMonthlyTax230629Do) Omit(cols ...field.Expr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winMonthlyTax230629Do) Join(table schema.Tabler, on ...field.Expr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winMonthlyTax230629Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winMonthlyTax230629Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winMonthlyTax230629Do) Group(cols ...field.Expr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winMonthlyTax230629Do) Having(conds ...gen.Condition) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winMonthlyTax230629Do) Limit(limit int) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winMonthlyTax230629Do) Offset(offset int) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winMonthlyTax230629Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winMonthlyTax230629Do) Unscoped() IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winMonthlyTax230629Do) Create(values ...*model.WinMonthlyTax230629) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winMonthlyTax230629Do) CreateInBatches(values []*model.WinMonthlyTax230629, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winMonthlyTax230629Do) Save(values ...*model.WinMonthlyTax230629) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winMonthlyTax230629Do) First() (*model.WinMonthlyTax230629, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax230629), nil
	}
}

func (w winMonthlyTax230629Do) Take() (*model.WinMonthlyTax230629, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax230629), nil
	}
}

func (w winMonthlyTax230629Do) Last() (*model.WinMonthlyTax230629, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax230629), nil
	}
}

func (w winMonthlyTax230629Do) Find() ([]*model.WinMonthlyTax230629, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinMonthlyTax230629), err
}

func (w winMonthlyTax230629Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinMonthlyTax230629, err error) {
	buf := make([]*model.WinMonthlyTax230629, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winMonthlyTax230629Do) FindInBatches(result *[]*model.WinMonthlyTax230629, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winMonthlyTax230629Do) Attrs(attrs ...field.AssignExpr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winMonthlyTax230629Do) Assign(attrs ...field.AssignExpr) IWinMonthlyTax230629Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winMonthlyTax230629Do) Joins(fields ...field.RelationField) IWinMonthlyTax230629Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winMonthlyTax230629Do) Preload(fields ...field.RelationField) IWinMonthlyTax230629Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winMonthlyTax230629Do) FirstOrInit() (*model.WinMonthlyTax230629, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax230629), nil
	}
}

func (w winMonthlyTax230629Do) FirstOrCreate() (*model.WinMonthlyTax230629, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax230629), nil
	}
}

func (w winMonthlyTax230629Do) FindByPage(offset int, limit int) (result []*model.WinMonthlyTax230629, count int64, err error) {
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

func (w winMonthlyTax230629Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winMonthlyTax230629Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winMonthlyTax230629Do) Delete(models ...*model.WinMonthlyTax230629) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winMonthlyTax230629Do) withDO(do gen.Dao) *winMonthlyTax230629Do {
	w.DO = *do.(*gen.DO)
	return w
}
