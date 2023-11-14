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

func newWinBetslipsDetail(db *gorm.DB, opts ...gen.DOOption) winBetslipsDetail {
	_winBetslipsDetail := winBetslipsDetail{}

	_winBetslipsDetail.winBetslipsDetailDo.UseDB(db, opts...)
	_winBetslipsDetail.winBetslipsDetailDo.UseModel(&model.WinBetslipsDetail{})

	tableName := _winBetslipsDetail.winBetslipsDetailDo.TableName()
	_winBetslipsDetail.ALL = field.NewAsterisk(tableName)
	_winBetslipsDetail.ID = field.NewInt64(tableName, "id")
	_winBetslipsDetail.XbUID = field.NewInt64(tableName, "xb_uid")
	_winBetslipsDetail.XbUsername = field.NewString(tableName, "xb_username")
	_winBetslipsDetail.BetJSON = field.NewString(tableName, "bet_json")
	_winBetslipsDetail.RewardJSON = field.NewString(tableName, "reward_json")
	_winBetslipsDetail.RefundJSON = field.NewString(tableName, "refund_json")

	_winBetslipsDetail.fillFieldMap()

	return _winBetslipsDetail
}

type winBetslipsDetail struct {
	winBetslipsDetailDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键-同注单表一致
	XbUID      field.Int64  // 对应user表id
	XbUsername field.String // 对应user表username
	BetJSON    field.String // 投注原始json
	RewardJSON field.String // 开彩原始json
	RefundJSON field.String // 退款原始json

	fieldMap map[string]field.Expr
}

func (w winBetslipsDetail) Table(newTableName string) *winBetslipsDetail {
	w.winBetslipsDetailDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBetslipsDetail) As(alias string) *winBetslipsDetail {
	w.winBetslipsDetailDo.DO = *(w.winBetslipsDetailDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBetslipsDetail) updateTableName(table string) *winBetslipsDetail {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.XbUID = field.NewInt64(table, "xb_uid")
	w.XbUsername = field.NewString(table, "xb_username")
	w.BetJSON = field.NewString(table, "bet_json")
	w.RewardJSON = field.NewString(table, "reward_json")
	w.RefundJSON = field.NewString(table, "refund_json")

	w.fillFieldMap()

	return w
}

func (w *winBetslipsDetail) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBetslipsDetail) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 6)
	w.fieldMap["id"] = w.ID
	w.fieldMap["xb_uid"] = w.XbUID
	w.fieldMap["xb_username"] = w.XbUsername
	w.fieldMap["bet_json"] = w.BetJSON
	w.fieldMap["reward_json"] = w.RewardJSON
	w.fieldMap["refund_json"] = w.RefundJSON
}

func (w winBetslipsDetail) clone(db *gorm.DB) winBetslipsDetail {
	w.winBetslipsDetailDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBetslipsDetail) replaceDB(db *gorm.DB) winBetslipsDetail {
	w.winBetslipsDetailDo.ReplaceDB(db)
	return w
}

type winBetslipsDetailDo struct{ gen.DO }

type IWinBetslipsDetailDo interface {
	gen.SubQuery
	Debug() IWinBetslipsDetailDo
	WithContext(ctx context.Context) IWinBetslipsDetailDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinBetslipsDetailDo
	WriteDB() IWinBetslipsDetailDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinBetslipsDetailDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinBetslipsDetailDo
	Not(conds ...gen.Condition) IWinBetslipsDetailDo
	Or(conds ...gen.Condition) IWinBetslipsDetailDo
	Select(conds ...field.Expr) IWinBetslipsDetailDo
	Where(conds ...gen.Condition) IWinBetslipsDetailDo
	Order(conds ...field.Expr) IWinBetslipsDetailDo
	Distinct(cols ...field.Expr) IWinBetslipsDetailDo
	Omit(cols ...field.Expr) IWinBetslipsDetailDo
	Join(table schema.Tabler, on ...field.Expr) IWinBetslipsDetailDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetailDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetailDo
	Group(cols ...field.Expr) IWinBetslipsDetailDo
	Having(conds ...gen.Condition) IWinBetslipsDetailDo
	Limit(limit int) IWinBetslipsDetailDo
	Offset(offset int) IWinBetslipsDetailDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipsDetailDo
	Unscoped() IWinBetslipsDetailDo
	Create(values ...*model.WinBetslipsDetail) error
	CreateInBatches(values []*model.WinBetslipsDetail, batchSize int) error
	Save(values ...*model.WinBetslipsDetail) error
	First() (*model.WinBetslipsDetail, error)
	Take() (*model.WinBetslipsDetail, error)
	Last() (*model.WinBetslipsDetail, error)
	Find() ([]*model.WinBetslipsDetail, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslipsDetail, err error)
	FindInBatches(result *[]*model.WinBetslipsDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinBetslipsDetail) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinBetslipsDetailDo
	Assign(attrs ...field.AssignExpr) IWinBetslipsDetailDo
	Joins(fields ...field.RelationField) IWinBetslipsDetailDo
	Preload(fields ...field.RelationField) IWinBetslipsDetailDo
	FirstOrInit() (*model.WinBetslipsDetail, error)
	FirstOrCreate() (*model.WinBetslipsDetail, error)
	FindByPage(offset int, limit int) (result []*model.WinBetslipsDetail, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinBetslipsDetailDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winBetslipsDetailDo) Debug() IWinBetslipsDetailDo {
	return w.withDO(w.DO.Debug())
}

func (w winBetslipsDetailDo) WithContext(ctx context.Context) IWinBetslipsDetailDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBetslipsDetailDo) ReadDB() IWinBetslipsDetailDo {
	return w.Clauses(dbresolver.Read)
}

func (w winBetslipsDetailDo) WriteDB() IWinBetslipsDetailDo {
	return w.Clauses(dbresolver.Write)
}

func (w winBetslipsDetailDo) Session(config *gorm.Session) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Session(config))
}

func (w winBetslipsDetailDo) Clauses(conds ...clause.Expression) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBetslipsDetailDo) Returning(value interface{}, columns ...string) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBetslipsDetailDo) Not(conds ...gen.Condition) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBetslipsDetailDo) Or(conds ...gen.Condition) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBetslipsDetailDo) Select(conds ...field.Expr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBetslipsDetailDo) Where(conds ...gen.Condition) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBetslipsDetailDo) Order(conds ...field.Expr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBetslipsDetailDo) Distinct(cols ...field.Expr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBetslipsDetailDo) Omit(cols ...field.Expr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBetslipsDetailDo) Join(table schema.Tabler, on ...field.Expr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBetslipsDetailDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBetslipsDetailDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBetslipsDetailDo) Group(cols ...field.Expr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBetslipsDetailDo) Having(conds ...gen.Condition) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBetslipsDetailDo) Limit(limit int) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBetslipsDetailDo) Offset(offset int) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBetslipsDetailDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBetslipsDetailDo) Unscoped() IWinBetslipsDetailDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winBetslipsDetailDo) Create(values ...*model.WinBetslipsDetail) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBetslipsDetailDo) CreateInBatches(values []*model.WinBetslipsDetail, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBetslipsDetailDo) Save(values ...*model.WinBetslipsDetail) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBetslipsDetailDo) First() (*model.WinBetslipsDetail, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetail), nil
	}
}

func (w winBetslipsDetailDo) Take() (*model.WinBetslipsDetail, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetail), nil
	}
}

func (w winBetslipsDetailDo) Last() (*model.WinBetslipsDetail, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetail), nil
	}
}

func (w winBetslipsDetailDo) Find() ([]*model.WinBetslipsDetail, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBetslipsDetail), err
}

func (w winBetslipsDetailDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslipsDetail, err error) {
	buf := make([]*model.WinBetslipsDetail, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBetslipsDetailDo) FindInBatches(result *[]*model.WinBetslipsDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBetslipsDetailDo) Attrs(attrs ...field.AssignExpr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBetslipsDetailDo) Assign(attrs ...field.AssignExpr) IWinBetslipsDetailDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBetslipsDetailDo) Joins(fields ...field.RelationField) IWinBetslipsDetailDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBetslipsDetailDo) Preload(fields ...field.RelationField) IWinBetslipsDetailDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBetslipsDetailDo) FirstOrInit() (*model.WinBetslipsDetail, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetail), nil
	}
}

func (w winBetslipsDetailDo) FirstOrCreate() (*model.WinBetslipsDetail, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetail), nil
	}
}

func (w winBetslipsDetailDo) FindByPage(offset int, limit int) (result []*model.WinBetslipsDetail, count int64, err error) {
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

func (w winBetslipsDetailDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBetslipsDetailDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBetslipsDetailDo) Delete(models ...*model.WinBetslipsDetail) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBetslipsDetailDo) withDO(do gen.Dao) *winBetslipsDetailDo {
	w.DO = *do.(*gen.DO)
	return w
}