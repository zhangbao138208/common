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

func newWinCoinUserGameTransfer(db *gorm.DB, opts ...gen.DOOption) winCoinUserGameTransfer {
	_winCoinUserGameTransfer := winCoinUserGameTransfer{}

	_winCoinUserGameTransfer.winCoinUserGameTransferDo.UseDB(db, opts...)
	_winCoinUserGameTransfer.winCoinUserGameTransferDo.UseModel(&model.WinCoinUserGameTransfer{})

	tableName := _winCoinUserGameTransfer.winCoinUserGameTransferDo.TableName()
	_winCoinUserGameTransfer.ALL = field.NewAsterisk(tableName)
	_winCoinUserGameTransfer.ID = field.NewInt64(tableName, "id")
	_winCoinUserGameTransfer.Coin = field.NewField(tableName, "coin")
	_winCoinUserGameTransfer.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinUserGameTransfer.UID = field.NewInt64(tableName, "uid")
	_winCoinUserGameTransfer.Username = field.NewString(tableName, "username")
	_winCoinUserGameTransfer.Category = field.NewInt64(tableName, "category")
	_winCoinUserGameTransfer.TransactionID = field.NewString(tableName, "transaction_id")
	_winCoinUserGameTransfer.GameProviderID = field.NewInt64(tableName, "game_provider_id")
	_winCoinUserGameTransfer.Mark = field.NewString(tableName, "mark")
	_winCoinUserGameTransfer.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinUserGameTransfer.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winCoinUserGameTransfer.fillFieldMap()

	return _winCoinUserGameTransfer
}

type winCoinUserGameTransfer struct {
	winCoinUserGameTransferDo

	ALL            field.Asterisk
	ID             field.Int64
	Coin           field.Field
	CoinBefore     field.Field
	UID            field.Int64 // ID
	Username       field.String
	Category       field.Int64  // :0- 1- 2-
	TransactionID  field.String // ID
	GameProviderID field.Int64  // ID
	Mark           field.String
	CreatedAt      field.Int64
	UpdatedAt      field.Int64

	fieldMap map[string]field.Expr
}

func (w winCoinUserGameTransfer) Table(newTableName string) *winCoinUserGameTransfer {
	w.winCoinUserGameTransferDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinUserGameTransfer) As(alias string) *winCoinUserGameTransfer {
	w.winCoinUserGameTransferDo.DO = *(w.winCoinUserGameTransferDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinUserGameTransfer) updateTableName(table string) *winCoinUserGameTransfer {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Coin = field.NewField(table, "coin")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Category = field.NewInt64(table, "category")
	w.TransactionID = field.NewString(table, "transaction_id")
	w.GameProviderID = field.NewInt64(table, "game_provider_id")
	w.Mark = field.NewString(table, "mark")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinUserGameTransfer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinUserGameTransfer) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["category"] = w.Category
	w.fieldMap["transaction_id"] = w.TransactionID
	w.fieldMap["game_provider_id"] = w.GameProviderID
	w.fieldMap["mark"] = w.Mark
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCoinUserGameTransfer) clone(db *gorm.DB) winCoinUserGameTransfer {
	w.winCoinUserGameTransferDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinUserGameTransfer) replaceDB(db *gorm.DB) winCoinUserGameTransfer {
	w.winCoinUserGameTransferDo.ReplaceDB(db)
	return w
}

type winCoinUserGameTransferDo struct{ gen.DO }

type IWinCoinUserGameTransferDo interface {
	gen.SubQuery
	Debug() IWinCoinUserGameTransferDo
	WithContext(ctx context.Context) IWinCoinUserGameTransferDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinUserGameTransferDo
	WriteDB() IWinCoinUserGameTransferDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinUserGameTransferDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinUserGameTransferDo
	Not(conds ...gen.Condition) IWinCoinUserGameTransferDo
	Or(conds ...gen.Condition) IWinCoinUserGameTransferDo
	Select(conds ...field.Expr) IWinCoinUserGameTransferDo
	Where(conds ...gen.Condition) IWinCoinUserGameTransferDo
	Order(conds ...field.Expr) IWinCoinUserGameTransferDo
	Distinct(cols ...field.Expr) IWinCoinUserGameTransferDo
	Omit(cols ...field.Expr) IWinCoinUserGameTransferDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinUserGameTransferDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinUserGameTransferDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinUserGameTransferDo
	Group(cols ...field.Expr) IWinCoinUserGameTransferDo
	Having(conds ...gen.Condition) IWinCoinUserGameTransferDo
	Limit(limit int) IWinCoinUserGameTransferDo
	Offset(offset int) IWinCoinUserGameTransferDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinUserGameTransferDo
	Unscoped() IWinCoinUserGameTransferDo
	Create(values ...*model.WinCoinUserGameTransfer) error
	CreateInBatches(values []*model.WinCoinUserGameTransfer, batchSize int) error
	Save(values ...*model.WinCoinUserGameTransfer) error
	First() (*model.WinCoinUserGameTransfer, error)
	Take() (*model.WinCoinUserGameTransfer, error)
	Last() (*model.WinCoinUserGameTransfer, error)
	Find() ([]*model.WinCoinUserGameTransfer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinUserGameTransfer, err error)
	FindInBatches(result *[]*model.WinCoinUserGameTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinUserGameTransfer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinUserGameTransferDo
	Assign(attrs ...field.AssignExpr) IWinCoinUserGameTransferDo
	Joins(fields ...field.RelationField) IWinCoinUserGameTransferDo
	Preload(fields ...field.RelationField) IWinCoinUserGameTransferDo
	FirstOrInit() (*model.WinCoinUserGameTransfer, error)
	FirstOrCreate() (*model.WinCoinUserGameTransfer, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinUserGameTransfer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinUserGameTransferDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinUserGameTransferDo) Debug() IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinUserGameTransferDo) WithContext(ctx context.Context) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinUserGameTransferDo) ReadDB() IWinCoinUserGameTransferDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinUserGameTransferDo) WriteDB() IWinCoinUserGameTransferDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinUserGameTransferDo) Session(config *gorm.Session) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinUserGameTransferDo) Clauses(conds ...clause.Expression) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinUserGameTransferDo) Returning(value interface{}, columns ...string) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinUserGameTransferDo) Not(conds ...gen.Condition) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinUserGameTransferDo) Or(conds ...gen.Condition) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinUserGameTransferDo) Select(conds ...field.Expr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinUserGameTransferDo) Where(conds ...gen.Condition) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinUserGameTransferDo) Order(conds ...field.Expr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinUserGameTransferDo) Distinct(cols ...field.Expr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinUserGameTransferDo) Omit(cols ...field.Expr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinUserGameTransferDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinUserGameTransferDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinUserGameTransferDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinUserGameTransferDo) Group(cols ...field.Expr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinUserGameTransferDo) Having(conds ...gen.Condition) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinUserGameTransferDo) Limit(limit int) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinUserGameTransferDo) Offset(offset int) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinUserGameTransferDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinUserGameTransferDo) Unscoped() IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinUserGameTransferDo) Create(values ...*model.WinCoinUserGameTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinUserGameTransferDo) CreateInBatches(values []*model.WinCoinUserGameTransfer, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinUserGameTransferDo) Save(values ...*model.WinCoinUserGameTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinUserGameTransferDo) First() (*model.WinCoinUserGameTransfer, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinUserGameTransfer), nil
	}
}

func (w winCoinUserGameTransferDo) Take() (*model.WinCoinUserGameTransfer, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinUserGameTransfer), nil
	}
}

func (w winCoinUserGameTransferDo) Last() (*model.WinCoinUserGameTransfer, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinUserGameTransfer), nil
	}
}

func (w winCoinUserGameTransferDo) Find() ([]*model.WinCoinUserGameTransfer, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinUserGameTransfer), err
}

func (w winCoinUserGameTransferDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinUserGameTransfer, err error) {
	buf := make([]*model.WinCoinUserGameTransfer, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinUserGameTransferDo) FindInBatches(result *[]*model.WinCoinUserGameTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinUserGameTransferDo) Attrs(attrs ...field.AssignExpr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinUserGameTransferDo) Assign(attrs ...field.AssignExpr) IWinCoinUserGameTransferDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinUserGameTransferDo) Joins(fields ...field.RelationField) IWinCoinUserGameTransferDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinUserGameTransferDo) Preload(fields ...field.RelationField) IWinCoinUserGameTransferDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinUserGameTransferDo) FirstOrInit() (*model.WinCoinUserGameTransfer, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinUserGameTransfer), nil
	}
}

func (w winCoinUserGameTransferDo) FirstOrCreate() (*model.WinCoinUserGameTransfer, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinUserGameTransfer), nil
	}
}

func (w winCoinUserGameTransferDo) FindByPage(offset int, limit int) (result []*model.WinCoinUserGameTransfer, count int64, err error) {
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

func (w winCoinUserGameTransferDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinUserGameTransferDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinUserGameTransferDo) Delete(models ...*model.WinCoinUserGameTransfer) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinUserGameTransferDo) withDO(do gen.Dao) *winCoinUserGameTransferDo {
	w.DO = *do.(*gen.DO)
	return w
}
