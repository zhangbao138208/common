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

func newWinCoinRewards(db *gorm.DB, opts ...gen.DOOption) winCoinRewards {
	_winCoinRewards := winCoinRewards{}

	_winCoinRewards.winCoinRewardsDo.UseDB(db, opts...)
	_winCoinRewards.winCoinRewardsDo.UseModel(&model.WinCoinRewards{})

	tableName := _winCoinRewards.winCoinRewardsDo.TableName()
	_winCoinRewards.ALL = field.NewAsterisk(tableName)
	_winCoinRewards.ID = field.NewInt64(tableName, "id")
	_winCoinRewards.UID = field.NewInt64(tableName, "uid")
	_winCoinRewards.Username = field.NewString(tableName, "username")
	_winCoinRewards.Coin = field.NewField(tableName, "coin")
	_winCoinRewards.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinRewards.ReferID = field.NewInt64(tableName, "refer_id")
	_winCoinRewards.LadderName = field.NewString(tableName, "ladder_name")
	_winCoinRewards.ReferCode = field.NewString(tableName, "refer_code")
	_winCoinRewards.FlowClaim = field.NewInt64(tableName, "flow_claim")
	_winCoinRewards.Codes = field.NewField(tableName, "codes")
	_winCoinRewards.EndedAt = field.NewInt64(tableName, "ended_at")
	_winCoinRewards.Info = field.NewString(tableName, "info")
	_winCoinRewards.Status = field.NewInt64(tableName, "status")
	_winCoinRewards.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinRewards.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winCoinRewards.TransferBonusAt = field.NewInt64(tableName, "transfer_bonus_at")

	_winCoinRewards.fillFieldMap()

	return _winCoinRewards
}

type winCoinRewards struct {
	winCoinRewardsDo

	ALL             field.Asterisk
	ID              field.Int64
	UID             field.Int64  // UID
	Username        field.String // 用户名
	Coin            field.Field  // 金额
	CoinBefore      field.Field  // 即时金额
	ReferID         field.Int64  // 关联ID(活动表ID)
	LadderName      field.String // 新活动，等级code
	ReferCode       field.String // 关联Code(活动表Code)
	FlowClaim       field.Int64  // 流水倍数
	Codes           field.Field  // 所需打码量
	EndedAt         field.Int64  // 活动结束时间
	Info            field.String // 备注
	Status          field.Int64  // 状态:0-申请中 1-已满足 2-已派发3-已结束
	CreatedAt       field.Int64
	UpdatedAt       field.Int64
	TransferBonusAt field.Int64 // 转主钱包时间

	fieldMap map[string]field.Expr
}

func (w winCoinRewards) Table(newTableName string) *winCoinRewards {
	w.winCoinRewardsDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinRewards) As(alias string) *winCoinRewards {
	w.winCoinRewardsDo.DO = *(w.winCoinRewardsDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinRewards) updateTableName(table string) *winCoinRewards {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Coin = field.NewField(table, "coin")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.ReferID = field.NewInt64(table, "refer_id")
	w.LadderName = field.NewString(table, "ladder_name")
	w.ReferCode = field.NewString(table, "refer_code")
	w.FlowClaim = field.NewInt64(table, "flow_claim")
	w.Codes = field.NewField(table, "codes")
	w.EndedAt = field.NewInt64(table, "ended_at")
	w.Info = field.NewString(table, "info")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.TransferBonusAt = field.NewInt64(table, "transfer_bonus_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinRewards) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinRewards) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 16)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["refer_id"] = w.ReferID
	w.fieldMap["ladder_name"] = w.LadderName
	w.fieldMap["refer_code"] = w.ReferCode
	w.fieldMap["flow_claim"] = w.FlowClaim
	w.fieldMap["codes"] = w.Codes
	w.fieldMap["ended_at"] = w.EndedAt
	w.fieldMap["info"] = w.Info
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["transfer_bonus_at"] = w.TransferBonusAt
}

func (w winCoinRewards) clone(db *gorm.DB) winCoinRewards {
	w.winCoinRewardsDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinRewards) replaceDB(db *gorm.DB) winCoinRewards {
	w.winCoinRewardsDo.ReplaceDB(db)
	return w
}

type winCoinRewardsDo struct{ gen.DO }

type IWinCoinRewardsDo interface {
	gen.SubQuery
	Debug() IWinCoinRewardsDo
	WithContext(ctx context.Context) IWinCoinRewardsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinRewardsDo
	WriteDB() IWinCoinRewardsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinRewardsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinRewardsDo
	Not(conds ...gen.Condition) IWinCoinRewardsDo
	Or(conds ...gen.Condition) IWinCoinRewardsDo
	Select(conds ...field.Expr) IWinCoinRewardsDo
	Where(conds ...gen.Condition) IWinCoinRewardsDo
	Order(conds ...field.Expr) IWinCoinRewardsDo
	Distinct(cols ...field.Expr) IWinCoinRewardsDo
	Omit(cols ...field.Expr) IWinCoinRewardsDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinRewardsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinRewardsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinRewardsDo
	Group(cols ...field.Expr) IWinCoinRewardsDo
	Having(conds ...gen.Condition) IWinCoinRewardsDo
	Limit(limit int) IWinCoinRewardsDo
	Offset(offset int) IWinCoinRewardsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinRewardsDo
	Unscoped() IWinCoinRewardsDo
	Create(values ...*model.WinCoinRewards) error
	CreateInBatches(values []*model.WinCoinRewards, batchSize int) error
	Save(values ...*model.WinCoinRewards) error
	First() (*model.WinCoinRewards, error)
	Take() (*model.WinCoinRewards, error)
	Last() (*model.WinCoinRewards, error)
	Find() ([]*model.WinCoinRewards, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinRewards, err error)
	FindInBatches(result *[]*model.WinCoinRewards, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinRewards) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinRewardsDo
	Assign(attrs ...field.AssignExpr) IWinCoinRewardsDo
	Joins(fields ...field.RelationField) IWinCoinRewardsDo
	Preload(fields ...field.RelationField) IWinCoinRewardsDo
	FirstOrInit() (*model.WinCoinRewards, error)
	FirstOrCreate() (*model.WinCoinRewards, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinRewards, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinRewardsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinRewardsDo) Debug() IWinCoinRewardsDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinRewardsDo) WithContext(ctx context.Context) IWinCoinRewardsDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinRewardsDo) ReadDB() IWinCoinRewardsDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinRewardsDo) WriteDB() IWinCoinRewardsDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinRewardsDo) Session(config *gorm.Session) IWinCoinRewardsDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinRewardsDo) Clauses(conds ...clause.Expression) IWinCoinRewardsDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinRewardsDo) Returning(value interface{}, columns ...string) IWinCoinRewardsDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinRewardsDo) Not(conds ...gen.Condition) IWinCoinRewardsDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinRewardsDo) Or(conds ...gen.Condition) IWinCoinRewardsDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinRewardsDo) Select(conds ...field.Expr) IWinCoinRewardsDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinRewardsDo) Where(conds ...gen.Condition) IWinCoinRewardsDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinRewardsDo) Order(conds ...field.Expr) IWinCoinRewardsDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinRewardsDo) Distinct(cols ...field.Expr) IWinCoinRewardsDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinRewardsDo) Omit(cols ...field.Expr) IWinCoinRewardsDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinRewardsDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinRewardsDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinRewardsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinRewardsDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinRewardsDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinRewardsDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinRewardsDo) Group(cols ...field.Expr) IWinCoinRewardsDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinRewardsDo) Having(conds ...gen.Condition) IWinCoinRewardsDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinRewardsDo) Limit(limit int) IWinCoinRewardsDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinRewardsDo) Offset(offset int) IWinCoinRewardsDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinRewardsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinRewardsDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinRewardsDo) Unscoped() IWinCoinRewardsDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinRewardsDo) Create(values ...*model.WinCoinRewards) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinRewardsDo) CreateInBatches(values []*model.WinCoinRewards, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinRewardsDo) Save(values ...*model.WinCoinRewards) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinRewardsDo) First() (*model.WinCoinRewards, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRewards), nil
	}
}

func (w winCoinRewardsDo) Take() (*model.WinCoinRewards, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRewards), nil
	}
}

func (w winCoinRewardsDo) Last() (*model.WinCoinRewards, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRewards), nil
	}
}

func (w winCoinRewardsDo) Find() ([]*model.WinCoinRewards, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinRewards), err
}

func (w winCoinRewardsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinRewards, err error) {
	buf := make([]*model.WinCoinRewards, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinRewardsDo) FindInBatches(result *[]*model.WinCoinRewards, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinRewardsDo) Attrs(attrs ...field.AssignExpr) IWinCoinRewardsDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinRewardsDo) Assign(attrs ...field.AssignExpr) IWinCoinRewardsDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinRewardsDo) Joins(fields ...field.RelationField) IWinCoinRewardsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinRewardsDo) Preload(fields ...field.RelationField) IWinCoinRewardsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinRewardsDo) FirstOrInit() (*model.WinCoinRewards, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRewards), nil
	}
}

func (w winCoinRewardsDo) FirstOrCreate() (*model.WinCoinRewards, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRewards), nil
	}
}

func (w winCoinRewardsDo) FindByPage(offset int, limit int) (result []*model.WinCoinRewards, count int64, err error) {
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

func (w winCoinRewardsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinRewardsDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinRewardsDo) Delete(models ...*model.WinCoinRewards) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinRewardsDo) withDO(do gen.Dao) *winCoinRewardsDo {
	w.DO = *do.(*gen.DO)
	return w
}
