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

func newWinUserLevel(db *gorm.DB, opts ...gen.DOOption) winUserLevel {
	_winUserLevel := winUserLevel{}

	_winUserLevel.winUserLevelDo.UseDB(db, opts...)
	_winUserLevel.winUserLevelDo.UseModel(&model.WinUserLevel{})

	tableName := _winUserLevel.winUserLevelDo.TableName()
	_winUserLevel.ALL = field.NewAsterisk(tableName)
	_winUserLevel.ID = field.NewInt64(tableName, "id")
	_winUserLevel.Code = field.NewString(tableName, "code")
	_winUserLevel.Name = field.NewString(tableName, "name")
	_winUserLevel.Icon = field.NewString(tableName, "icon")
	_winUserLevel.DepositChannel = field.NewString(tableName, "deposit_channel")
	_winUserLevel.ScoreRelegation = field.NewInt64(tableName, "score_relegation")
	_winUserLevel.RelegationDay = field.NewInt64(tableName, "relegation_day")
	_winUserLevel.BetSum = field.NewInt64(tableName, "bet_sum")
	_winUserLevel.DepositSum = field.NewInt64(tableName, "deposit_sum")
	_winUserLevel.WithdrawalCount = field.NewInt64(tableName, "withdrawal_count")
	_winUserLevel.WithdrawalCoin = field.NewInt64(tableName, "withdrawal_coin")
	_winUserLevel.RebateMax = field.NewInt64(tableName, "rebate_max")
	_winUserLevel.UpgradeReward = field.NewInt64(tableName, "upgrade_reward")
	_winUserLevel.BirthdayReward = field.NewInt64(tableName, "birthday_reward")
	_winUserLevel.WeekReward = field.NewInt64(tableName, "week_reward")
	_winUserLevel.MonthReward = field.NewInt64(tableName, "month_reward")
	_winUserLevel.FlowClaim = field.NewInt64(tableName, "flow_claim")
	_winUserLevel.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserLevel.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winUserLevel.UpdatedUser = field.NewString(tableName, "updated_user")
	_winUserLevel.ScoreUpgradeMin = field.NewInt64(tableName, "score_upgrade_min")
	_winUserLevel.ScoreUpgradeMax = field.NewInt64(tableName, "score_upgrade_max")
	_winUserLevel.ScoreUpgradeRate = field.NewInt64(tableName, "score_upgrade_rate")

	_winUserLevel.fillFieldMap()

	return _winUserLevel
}

type winUserLevel struct {
	winUserLevelDo

	ALL              field.Asterisk
	ID               field.Int64
	Code             field.String // 会员等级
	Name             field.String // 等级名称
	Icon             field.String // ICON
	DepositChannel   field.String // 存款通道
	ScoreRelegation  field.Int64  // 保级有效投注
	RelegationDay    field.Int64  // 保级有效天数
	BetSum           field.Int64  // 累计投注
	DepositSum       field.Int64  // 累计存款
	WithdrawalCount  field.Int64  // 每日提款次数
	WithdrawalCoin   field.Int64  // 每日提款额度
	RebateMax        field.Int64  // 每日返水上限
	UpgradeReward    field.Int64  // 升级礼金
	BirthdayReward   field.Int64  // 生日礼金
	WeekReward       field.Int64  // 周礼金
	MonthReward      field.Int64  // 月礼金
	FlowClaim        field.Int64  // 打码倍数
	CreatedAt        field.Int64
	UpdatedAt        field.Int64
	UpdatedUser      field.String // 最后修改人
	ScoreUpgradeMin  field.Int64
	ScoreUpgradeMax  field.Int64
	ScoreUpgradeRate field.Int64

	fieldMap map[string]field.Expr
}

func (w winUserLevel) Table(newTableName string) *winUserLevel {
	w.winUserLevelDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserLevel) As(alias string) *winUserLevel {
	w.winUserLevelDo.DO = *(w.winUserLevelDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserLevel) updateTableName(table string) *winUserLevel {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Code = field.NewString(table, "code")
	w.Name = field.NewString(table, "name")
	w.Icon = field.NewString(table, "icon")
	w.DepositChannel = field.NewString(table, "deposit_channel")
	w.ScoreRelegation = field.NewInt64(table, "score_relegation")
	w.RelegationDay = field.NewInt64(table, "relegation_day")
	w.BetSum = field.NewInt64(table, "bet_sum")
	w.DepositSum = field.NewInt64(table, "deposit_sum")
	w.WithdrawalCount = field.NewInt64(table, "withdrawal_count")
	w.WithdrawalCoin = field.NewInt64(table, "withdrawal_coin")
	w.RebateMax = field.NewInt64(table, "rebate_max")
	w.UpgradeReward = field.NewInt64(table, "upgrade_reward")
	w.BirthdayReward = field.NewInt64(table, "birthday_reward")
	w.WeekReward = field.NewInt64(table, "week_reward")
	w.MonthReward = field.NewInt64(table, "month_reward")
	w.FlowClaim = field.NewInt64(table, "flow_claim")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.UpdatedUser = field.NewString(table, "updated_user")
	w.ScoreUpgradeMin = field.NewInt64(table, "score_upgrade_min")
	w.ScoreUpgradeMax = field.NewInt64(table, "score_upgrade_max")
	w.ScoreUpgradeRate = field.NewInt64(table, "score_upgrade_rate")

	w.fillFieldMap()

	return w
}

func (w *winUserLevel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserLevel) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 23)
	w.fieldMap["id"] = w.ID
	w.fieldMap["code"] = w.Code
	w.fieldMap["name"] = w.Name
	w.fieldMap["icon"] = w.Icon
	w.fieldMap["deposit_channel"] = w.DepositChannel
	w.fieldMap["score_relegation"] = w.ScoreRelegation
	w.fieldMap["relegation_day"] = w.RelegationDay
	w.fieldMap["bet_sum"] = w.BetSum
	w.fieldMap["deposit_sum"] = w.DepositSum
	w.fieldMap["withdrawal_count"] = w.WithdrawalCount
	w.fieldMap["withdrawal_coin"] = w.WithdrawalCoin
	w.fieldMap["rebate_max"] = w.RebateMax
	w.fieldMap["upgrade_reward"] = w.UpgradeReward
	w.fieldMap["birthday_reward"] = w.BirthdayReward
	w.fieldMap["week_reward"] = w.WeekReward
	w.fieldMap["month_reward"] = w.MonthReward
	w.fieldMap["flow_claim"] = w.FlowClaim
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["updated_user"] = w.UpdatedUser
	w.fieldMap["score_upgrade_min"] = w.ScoreUpgradeMin
	w.fieldMap["score_upgrade_max"] = w.ScoreUpgradeMax
	w.fieldMap["score_upgrade_rate"] = w.ScoreUpgradeRate
}

func (w winUserLevel) clone(db *gorm.DB) winUserLevel {
	w.winUserLevelDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserLevel) replaceDB(db *gorm.DB) winUserLevel {
	w.winUserLevelDo.ReplaceDB(db)
	return w
}

type winUserLevelDo struct{ gen.DO }

type IWinUserLevelDo interface {
	gen.SubQuery
	Debug() IWinUserLevelDo
	WithContext(ctx context.Context) IWinUserLevelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserLevelDo
	WriteDB() IWinUserLevelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserLevelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserLevelDo
	Not(conds ...gen.Condition) IWinUserLevelDo
	Or(conds ...gen.Condition) IWinUserLevelDo
	Select(conds ...field.Expr) IWinUserLevelDo
	Where(conds ...gen.Condition) IWinUserLevelDo
	Order(conds ...field.Expr) IWinUserLevelDo
	Distinct(cols ...field.Expr) IWinUserLevelDo
	Omit(cols ...field.Expr) IWinUserLevelDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserLevelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelDo
	Group(cols ...field.Expr) IWinUserLevelDo
	Having(conds ...gen.Condition) IWinUserLevelDo
	Limit(limit int) IWinUserLevelDo
	Offset(offset int) IWinUserLevelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLevelDo
	Unscoped() IWinUserLevelDo
	Create(values ...*model.WinUserLevel) error
	CreateInBatches(values []*model.WinUserLevel, batchSize int) error
	Save(values ...*model.WinUserLevel) error
	First() (*model.WinUserLevel, error)
	Take() (*model.WinUserLevel, error)
	Last() (*model.WinUserLevel, error)
	Find() ([]*model.WinUserLevel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLevel, err error)
	FindInBatches(result *[]*model.WinUserLevel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserLevel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserLevelDo
	Assign(attrs ...field.AssignExpr) IWinUserLevelDo
	Joins(fields ...field.RelationField) IWinUserLevelDo
	Preload(fields ...field.RelationField) IWinUserLevelDo
	FirstOrInit() (*model.WinUserLevel, error)
	FirstOrCreate() (*model.WinUserLevel, error)
	FindByPage(offset int, limit int) (result []*model.WinUserLevel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserLevelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserLevelDo) Debug() IWinUserLevelDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserLevelDo) WithContext(ctx context.Context) IWinUserLevelDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserLevelDo) ReadDB() IWinUserLevelDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserLevelDo) WriteDB() IWinUserLevelDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserLevelDo) Session(config *gorm.Session) IWinUserLevelDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserLevelDo) Clauses(conds ...clause.Expression) IWinUserLevelDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserLevelDo) Returning(value interface{}, columns ...string) IWinUserLevelDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserLevelDo) Not(conds ...gen.Condition) IWinUserLevelDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserLevelDo) Or(conds ...gen.Condition) IWinUserLevelDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserLevelDo) Select(conds ...field.Expr) IWinUserLevelDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserLevelDo) Where(conds ...gen.Condition) IWinUserLevelDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserLevelDo) Order(conds ...field.Expr) IWinUserLevelDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserLevelDo) Distinct(cols ...field.Expr) IWinUserLevelDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserLevelDo) Omit(cols ...field.Expr) IWinUserLevelDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserLevelDo) Join(table schema.Tabler, on ...field.Expr) IWinUserLevelDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserLevelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserLevelDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserLevelDo) Group(cols ...field.Expr) IWinUserLevelDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserLevelDo) Having(conds ...gen.Condition) IWinUserLevelDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserLevelDo) Limit(limit int) IWinUserLevelDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserLevelDo) Offset(offset int) IWinUserLevelDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserLevelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLevelDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserLevelDo) Unscoped() IWinUserLevelDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserLevelDo) Create(values ...*model.WinUserLevel) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserLevelDo) CreateInBatches(values []*model.WinUserLevel, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserLevelDo) Save(values ...*model.WinUserLevel) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserLevelDo) First() (*model.WinUserLevel, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevel), nil
	}
}

func (w winUserLevelDo) Take() (*model.WinUserLevel, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevel), nil
	}
}

func (w winUserLevelDo) Last() (*model.WinUserLevel, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevel), nil
	}
}

func (w winUserLevelDo) Find() ([]*model.WinUserLevel, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserLevel), err
}

func (w winUserLevelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLevel, err error) {
	buf := make([]*model.WinUserLevel, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserLevelDo) FindInBatches(result *[]*model.WinUserLevel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserLevelDo) Attrs(attrs ...field.AssignExpr) IWinUserLevelDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserLevelDo) Assign(attrs ...field.AssignExpr) IWinUserLevelDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserLevelDo) Joins(fields ...field.RelationField) IWinUserLevelDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserLevelDo) Preload(fields ...field.RelationField) IWinUserLevelDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserLevelDo) FirstOrInit() (*model.WinUserLevel, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevel), nil
	}
}

func (w winUserLevelDo) FirstOrCreate() (*model.WinUserLevel, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevel), nil
	}
}

func (w winUserLevelDo) FindByPage(offset int, limit int) (result []*model.WinUserLevel, count int64, err error) {
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

func (w winUserLevelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserLevelDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserLevelDo) Delete(models ...*model.WinUserLevel) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserLevelDo) withDO(do gen.Dao) *winUserLevelDo {
	w.DO = *do.(*gen.DO)
	return w
}