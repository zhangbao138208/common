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

func newAgentReportConfigHistory(db *gorm.DB, opts ...gen.DOOption) agentReportConfigHistory {
	_agentReportConfigHistory := agentReportConfigHistory{}

	_agentReportConfigHistory.agentReportConfigHistoryDo.UseDB(db, opts...)
	_agentReportConfigHistory.agentReportConfigHistoryDo.UseModel(&model.AgentReportConfigHistory{})

	tableName := _agentReportConfigHistory.agentReportConfigHistoryDo.TableName()
	_agentReportConfigHistory.ALL = field.NewAsterisk(tableName)
	_agentReportConfigHistory.ID = field.NewInt64(tableName, "id")
	_agentReportConfigHistory.AgentID = field.NewInt64(tableName, "agent_id")
	_agentReportConfigHistory.Agentname = field.NewString(tableName, "agentname")
	_agentReportConfigHistory.AgentLevel = field.NewInt64(tableName, "agent_level")
	_agentReportConfigHistory.RateSport = field.NewField(tableName, "rate_sport")
	_agentReportConfigHistory.RateSlot = field.NewField(tableName, "rate_slot")
	_agentReportConfigHistory.RateReal = field.NewField(tableName, "rate_real")
	_agentReportConfigHistory.RateMachine = field.NewField(tableName, "rate_machine")
	_agentReportConfigHistory.RateWithdrawal = field.NewField(tableName, "rate_withdrawal")
	_agentReportConfigHistory.ExpiredAt = field.NewInt64(tableName, "expired_at")
	_agentReportConfigHistory.Tag = field.NewInt64(tableName, "tag")
	_agentReportConfigHistory.CreatedAt = field.NewInt64(tableName, "created_at")
	_agentReportConfigHistory.CommissionRate = field.NewField(tableName, "commission_rate")

	_agentReportConfigHistory.fillFieldMap()

	return _agentReportConfigHistory
}

type agentReportConfigHistory struct {
	agentReportConfigHistoryDo

	ALL            field.Asterisk
	ID             field.Int64  // 主键ID
	AgentID        field.Int64  // 代理ID
	Agentname      field.String // 代理名称
	AgentLevel     field.Int64  // 代理层级
	RateSport      field.Field  // 厂商费率-体育
	RateSlot       field.Field  // 厂商费率-老虎机
	RateReal       field.Field  // 厂商费率-真人
	RateMachine    field.Field  // 厂商费率-实体机
	RateWithdrawal field.Field  // 提款费率
	ExpiredAt      field.Int64  // 代理关系过期天数
	Tag            field.Int64  // 标记: 0-厂商费率, 1-佣金比例
	CreatedAt      field.Int64  // 创建时间
	CommissionRate field.Field  // 佣金比例

	fieldMap map[string]field.Expr
}

func (a agentReportConfigHistory) Table(newTableName string) *agentReportConfigHistory {
	a.agentReportConfigHistoryDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a agentReportConfigHistory) As(alias string) *agentReportConfigHistory {
	a.agentReportConfigHistoryDo.DO = *(a.agentReportConfigHistoryDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *agentReportConfigHistory) updateTableName(table string) *agentReportConfigHistory {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.AgentID = field.NewInt64(table, "agent_id")
	a.Agentname = field.NewString(table, "agentname")
	a.AgentLevel = field.NewInt64(table, "agent_level")
	a.RateSport = field.NewField(table, "rate_sport")
	a.RateSlot = field.NewField(table, "rate_slot")
	a.RateReal = field.NewField(table, "rate_real")
	a.RateMachine = field.NewField(table, "rate_machine")
	a.RateWithdrawal = field.NewField(table, "rate_withdrawal")
	a.ExpiredAt = field.NewInt64(table, "expired_at")
	a.Tag = field.NewInt64(table, "tag")
	a.CreatedAt = field.NewInt64(table, "created_at")
	a.CommissionRate = field.NewField(table, "commission_rate")

	a.fillFieldMap()

	return a
}

func (a *agentReportConfigHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *agentReportConfigHistory) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 13)
	a.fieldMap["id"] = a.ID
	a.fieldMap["agent_id"] = a.AgentID
	a.fieldMap["agentname"] = a.Agentname
	a.fieldMap["agent_level"] = a.AgentLevel
	a.fieldMap["rate_sport"] = a.RateSport
	a.fieldMap["rate_slot"] = a.RateSlot
	a.fieldMap["rate_real"] = a.RateReal
	a.fieldMap["rate_machine"] = a.RateMachine
	a.fieldMap["rate_withdrawal"] = a.RateWithdrawal
	a.fieldMap["expired_at"] = a.ExpiredAt
	a.fieldMap["tag"] = a.Tag
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["commission_rate"] = a.CommissionRate
}

func (a agentReportConfigHistory) clone(db *gorm.DB) agentReportConfigHistory {
	a.agentReportConfigHistoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a agentReportConfigHistory) replaceDB(db *gorm.DB) agentReportConfigHistory {
	a.agentReportConfigHistoryDo.ReplaceDB(db)
	return a
}

type agentReportConfigHistoryDo struct{ gen.DO }

type IAgentReportConfigHistoryDo interface {
	gen.SubQuery
	Debug() IAgentReportConfigHistoryDo
	WithContext(ctx context.Context) IAgentReportConfigHistoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAgentReportConfigHistoryDo
	WriteDB() IAgentReportConfigHistoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAgentReportConfigHistoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAgentReportConfigHistoryDo
	Not(conds ...gen.Condition) IAgentReportConfigHistoryDo
	Or(conds ...gen.Condition) IAgentReportConfigHistoryDo
	Select(conds ...field.Expr) IAgentReportConfigHistoryDo
	Where(conds ...gen.Condition) IAgentReportConfigHistoryDo
	Order(conds ...field.Expr) IAgentReportConfigHistoryDo
	Distinct(cols ...field.Expr) IAgentReportConfigHistoryDo
	Omit(cols ...field.Expr) IAgentReportConfigHistoryDo
	Join(table schema.Tabler, on ...field.Expr) IAgentReportConfigHistoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAgentReportConfigHistoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAgentReportConfigHistoryDo
	Group(cols ...field.Expr) IAgentReportConfigHistoryDo
	Having(conds ...gen.Condition) IAgentReportConfigHistoryDo
	Limit(limit int) IAgentReportConfigHistoryDo
	Offset(offset int) IAgentReportConfigHistoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAgentReportConfigHistoryDo
	Unscoped() IAgentReportConfigHistoryDo
	Create(values ...*model.AgentReportConfigHistory) error
	CreateInBatches(values []*model.AgentReportConfigHistory, batchSize int) error
	Save(values ...*model.AgentReportConfigHistory) error
	First() (*model.AgentReportConfigHistory, error)
	Take() (*model.AgentReportConfigHistory, error)
	Last() (*model.AgentReportConfigHistory, error)
	Find() ([]*model.AgentReportConfigHistory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AgentReportConfigHistory, err error)
	FindInBatches(result *[]*model.AgentReportConfigHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AgentReportConfigHistory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAgentReportConfigHistoryDo
	Assign(attrs ...field.AssignExpr) IAgentReportConfigHistoryDo
	Joins(fields ...field.RelationField) IAgentReportConfigHistoryDo
	Preload(fields ...field.RelationField) IAgentReportConfigHistoryDo
	FirstOrInit() (*model.AgentReportConfigHistory, error)
	FirstOrCreate() (*model.AgentReportConfigHistory, error)
	FindByPage(offset int, limit int) (result []*model.AgentReportConfigHistory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAgentReportConfigHistoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a agentReportConfigHistoryDo) Debug() IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Debug())
}

func (a agentReportConfigHistoryDo) WithContext(ctx context.Context) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a agentReportConfigHistoryDo) ReadDB() IAgentReportConfigHistoryDo {
	return a.Clauses(dbresolver.Read)
}

func (a agentReportConfigHistoryDo) WriteDB() IAgentReportConfigHistoryDo {
	return a.Clauses(dbresolver.Write)
}

func (a agentReportConfigHistoryDo) Session(config *gorm.Session) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Session(config))
}

func (a agentReportConfigHistoryDo) Clauses(conds ...clause.Expression) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a agentReportConfigHistoryDo) Returning(value interface{}, columns ...string) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a agentReportConfigHistoryDo) Not(conds ...gen.Condition) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a agentReportConfigHistoryDo) Or(conds ...gen.Condition) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a agentReportConfigHistoryDo) Select(conds ...field.Expr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a agentReportConfigHistoryDo) Where(conds ...gen.Condition) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a agentReportConfigHistoryDo) Order(conds ...field.Expr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a agentReportConfigHistoryDo) Distinct(cols ...field.Expr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a agentReportConfigHistoryDo) Omit(cols ...field.Expr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a agentReportConfigHistoryDo) Join(table schema.Tabler, on ...field.Expr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a agentReportConfigHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a agentReportConfigHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a agentReportConfigHistoryDo) Group(cols ...field.Expr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a agentReportConfigHistoryDo) Having(conds ...gen.Condition) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a agentReportConfigHistoryDo) Limit(limit int) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a agentReportConfigHistoryDo) Offset(offset int) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a agentReportConfigHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a agentReportConfigHistoryDo) Unscoped() IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Unscoped())
}

func (a agentReportConfigHistoryDo) Create(values ...*model.AgentReportConfigHistory) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a agentReportConfigHistoryDo) CreateInBatches(values []*model.AgentReportConfigHistory, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a agentReportConfigHistoryDo) Save(values ...*model.AgentReportConfigHistory) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a agentReportConfigHistoryDo) First() (*model.AgentReportConfigHistory, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentReportConfigHistory), nil
	}
}

func (a agentReportConfigHistoryDo) Take() (*model.AgentReportConfigHistory, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentReportConfigHistory), nil
	}
}

func (a agentReportConfigHistoryDo) Last() (*model.AgentReportConfigHistory, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentReportConfigHistory), nil
	}
}

func (a agentReportConfigHistoryDo) Find() ([]*model.AgentReportConfigHistory, error) {
	result, err := a.DO.Find()
	return result.([]*model.AgentReportConfigHistory), err
}

func (a agentReportConfigHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AgentReportConfigHistory, err error) {
	buf := make([]*model.AgentReportConfigHistory, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a agentReportConfigHistoryDo) FindInBatches(result *[]*model.AgentReportConfigHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a agentReportConfigHistoryDo) Attrs(attrs ...field.AssignExpr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a agentReportConfigHistoryDo) Assign(attrs ...field.AssignExpr) IAgentReportConfigHistoryDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a agentReportConfigHistoryDo) Joins(fields ...field.RelationField) IAgentReportConfigHistoryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a agentReportConfigHistoryDo) Preload(fields ...field.RelationField) IAgentReportConfigHistoryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a agentReportConfigHistoryDo) FirstOrInit() (*model.AgentReportConfigHistory, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentReportConfigHistory), nil
	}
}

func (a agentReportConfigHistoryDo) FirstOrCreate() (*model.AgentReportConfigHistory, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentReportConfigHistory), nil
	}
}

func (a agentReportConfigHistoryDo) FindByPage(offset int, limit int) (result []*model.AgentReportConfigHistory, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a agentReportConfigHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a agentReportConfigHistoryDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a agentReportConfigHistoryDo) Delete(models ...*model.AgentReportConfigHistory) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *agentReportConfigHistoryDo) withDO(do gen.Dao) *agentReportConfigHistoryDo {
	a.DO = *do.(*gen.DO)
	return a
}
