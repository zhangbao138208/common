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

func newWinPromotion(db *gorm.DB, opts ...gen.DOOption) winPromotion {
	_winPromotion := winPromotion{}

	_winPromotion.winPromotionDo.UseDB(db, opts...)
	_winPromotion.winPromotionDo.UseModel(&model.WinPromotion{})

	tableName := _winPromotion.winPromotionDo.TableName()
	_winPromotion.ALL = field.NewAsterisk(tableName)
	_winPromotion.ID = field.NewInt64(tableName, "id")
	_winPromotion.Code = field.NewString(tableName, "code")
	_winPromotion.CodeZh = field.NewString(tableName, "code_zh")
	_winPromotion.Img = field.NewString(tableName, "img")
	_winPromotion.Category = field.NewString(tableName, "category")
	_winPromotion.GameType = field.NewInt64(tableName, "game_type")
	_winPromotion.Info = field.NewString(tableName, "info")
	_winPromotion.Descript = field.NewString(tableName, "descript")
	_winPromotion.StartedAt = field.NewInt64(tableName, "started_at")
	_winPromotion.Ladder = field.NewString(tableName, "ladder")
	_winPromotion.PayoutCategory = field.NewInt64(tableName, "payout_category")
	_winPromotion.EndedAt = field.NewInt64(tableName, "ended_at")
	_winPromotion.Sort = field.NewInt64(tableName, "sort")
	_winPromotion.Status = field.NewInt64(tableName, "status")
	_winPromotion.CreatedAt = field.NewInt64(tableName, "created_at")
	_winPromotion.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winPromotion.OperatorName = field.NewString(tableName, "operator_name")

	_winPromotion.fillFieldMap()

	return _winPromotion
}

type winPromotion struct {
	winPromotionDo

	ALL            field.Asterisk
	ID             field.Int64
	Code           field.String // 活动标识:首充优惠-First Deposit Bonus 续充优惠-Second Deposit Bonus 首单包赔-Risk-Free Bet 快乐周末-Happy Weekend Bonus
	CodeZh         field.String // 名称中文
	Img            field.String // 图片
	Category       field.String // 类型:1-充值优惠 2-豪礼赠送 3-新活动
	GameType       field.Int64  // 活动游戏类型，见字典dic_promotion_game_type
	Info           field.String // 补充信息
	Descript       field.String // 详情描述
	StartedAt      field.Int64  // 开始时间
	Ladder         field.String // 新活动阶梯
	PayoutCategory field.Int64  // 派彩类型: 0-自动派彩 1-人工派彩 2-手动派彩
	EndedAt        field.Int64  // 结算时间
	Sort           field.Int64  // 排序(从高到底、ID降序)
	Status         field.Int64  // 状态:1-启用 0-停用
	CreatedAt      field.Int64
	UpdatedAt      field.Int64
	OperatorName   field.String // 操作人姓名

	fieldMap map[string]field.Expr
}

func (w winPromotion) Table(newTableName string) *winPromotion {
	w.winPromotionDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winPromotion) As(alias string) *winPromotion {
	w.winPromotionDo.DO = *(w.winPromotionDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winPromotion) updateTableName(table string) *winPromotion {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Code = field.NewString(table, "code")
	w.CodeZh = field.NewString(table, "code_zh")
	w.Img = field.NewString(table, "img")
	w.Category = field.NewString(table, "category")
	w.GameType = field.NewInt64(table, "game_type")
	w.Info = field.NewString(table, "info")
	w.Descript = field.NewString(table, "descript")
	w.StartedAt = field.NewInt64(table, "started_at")
	w.Ladder = field.NewString(table, "ladder")
	w.PayoutCategory = field.NewInt64(table, "payout_category")
	w.EndedAt = field.NewInt64(table, "ended_at")
	w.Sort = field.NewInt64(table, "sort")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.OperatorName = field.NewString(table, "operator_name")

	w.fillFieldMap()

	return w
}

func (w *winPromotion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winPromotion) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 17)
	w.fieldMap["id"] = w.ID
	w.fieldMap["code"] = w.Code
	w.fieldMap["code_zh"] = w.CodeZh
	w.fieldMap["img"] = w.Img
	w.fieldMap["category"] = w.Category
	w.fieldMap["game_type"] = w.GameType
	w.fieldMap["info"] = w.Info
	w.fieldMap["descript"] = w.Descript
	w.fieldMap["started_at"] = w.StartedAt
	w.fieldMap["ladder"] = w.Ladder
	w.fieldMap["payout_category"] = w.PayoutCategory
	w.fieldMap["ended_at"] = w.EndedAt
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["operator_name"] = w.OperatorName
}

func (w winPromotion) clone(db *gorm.DB) winPromotion {
	w.winPromotionDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winPromotion) replaceDB(db *gorm.DB) winPromotion {
	w.winPromotionDo.ReplaceDB(db)
	return w
}

type winPromotionDo struct{ gen.DO }

type IWinPromotionDo interface {
	gen.SubQuery
	Debug() IWinPromotionDo
	WithContext(ctx context.Context) IWinPromotionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinPromotionDo
	WriteDB() IWinPromotionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinPromotionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinPromotionDo
	Not(conds ...gen.Condition) IWinPromotionDo
	Or(conds ...gen.Condition) IWinPromotionDo
	Select(conds ...field.Expr) IWinPromotionDo
	Where(conds ...gen.Condition) IWinPromotionDo
	Order(conds ...field.Expr) IWinPromotionDo
	Distinct(cols ...field.Expr) IWinPromotionDo
	Omit(cols ...field.Expr) IWinPromotionDo
	Join(table schema.Tabler, on ...field.Expr) IWinPromotionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinPromotionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinPromotionDo
	Group(cols ...field.Expr) IWinPromotionDo
	Having(conds ...gen.Condition) IWinPromotionDo
	Limit(limit int) IWinPromotionDo
	Offset(offset int) IWinPromotionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPromotionDo
	Unscoped() IWinPromotionDo
	Create(values ...*model.WinPromotion) error
	CreateInBatches(values []*model.WinPromotion, batchSize int) error
	Save(values ...*model.WinPromotion) error
	First() (*model.WinPromotion, error)
	Take() (*model.WinPromotion, error)
	Last() (*model.WinPromotion, error)
	Find() ([]*model.WinPromotion, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPromotion, err error)
	FindInBatches(result *[]*model.WinPromotion, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinPromotion) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinPromotionDo
	Assign(attrs ...field.AssignExpr) IWinPromotionDo
	Joins(fields ...field.RelationField) IWinPromotionDo
	Preload(fields ...field.RelationField) IWinPromotionDo
	FirstOrInit() (*model.WinPromotion, error)
	FirstOrCreate() (*model.WinPromotion, error)
	FindByPage(offset int, limit int) (result []*model.WinPromotion, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinPromotionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winPromotionDo) Debug() IWinPromotionDo {
	return w.withDO(w.DO.Debug())
}

func (w winPromotionDo) WithContext(ctx context.Context) IWinPromotionDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winPromotionDo) ReadDB() IWinPromotionDo {
	return w.Clauses(dbresolver.Read)
}

func (w winPromotionDo) WriteDB() IWinPromotionDo {
	return w.Clauses(dbresolver.Write)
}

func (w winPromotionDo) Session(config *gorm.Session) IWinPromotionDo {
	return w.withDO(w.DO.Session(config))
}

func (w winPromotionDo) Clauses(conds ...clause.Expression) IWinPromotionDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winPromotionDo) Returning(value interface{}, columns ...string) IWinPromotionDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winPromotionDo) Not(conds ...gen.Condition) IWinPromotionDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winPromotionDo) Or(conds ...gen.Condition) IWinPromotionDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winPromotionDo) Select(conds ...field.Expr) IWinPromotionDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winPromotionDo) Where(conds ...gen.Condition) IWinPromotionDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winPromotionDo) Order(conds ...field.Expr) IWinPromotionDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winPromotionDo) Distinct(cols ...field.Expr) IWinPromotionDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winPromotionDo) Omit(cols ...field.Expr) IWinPromotionDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winPromotionDo) Join(table schema.Tabler, on ...field.Expr) IWinPromotionDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winPromotionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinPromotionDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winPromotionDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinPromotionDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winPromotionDo) Group(cols ...field.Expr) IWinPromotionDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winPromotionDo) Having(conds ...gen.Condition) IWinPromotionDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winPromotionDo) Limit(limit int) IWinPromotionDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winPromotionDo) Offset(offset int) IWinPromotionDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winPromotionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPromotionDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winPromotionDo) Unscoped() IWinPromotionDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winPromotionDo) Create(values ...*model.WinPromotion) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winPromotionDo) CreateInBatches(values []*model.WinPromotion, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winPromotionDo) Save(values ...*model.WinPromotion) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winPromotionDo) First() (*model.WinPromotion, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotion), nil
	}
}

func (w winPromotionDo) Take() (*model.WinPromotion, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotion), nil
	}
}

func (w winPromotionDo) Last() (*model.WinPromotion, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotion), nil
	}
}

func (w winPromotionDo) Find() ([]*model.WinPromotion, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinPromotion), err
}

func (w winPromotionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPromotion, err error) {
	buf := make([]*model.WinPromotion, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winPromotionDo) FindInBatches(result *[]*model.WinPromotion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winPromotionDo) Attrs(attrs ...field.AssignExpr) IWinPromotionDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winPromotionDo) Assign(attrs ...field.AssignExpr) IWinPromotionDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winPromotionDo) Joins(fields ...field.RelationField) IWinPromotionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winPromotionDo) Preload(fields ...field.RelationField) IWinPromotionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winPromotionDo) FirstOrInit() (*model.WinPromotion, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotion), nil
	}
}

func (w winPromotionDo) FirstOrCreate() (*model.WinPromotion, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotion), nil
	}
}

func (w winPromotionDo) FindByPage(offset int, limit int) (result []*model.WinPromotion, count int64, err error) {
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

func (w winPromotionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winPromotionDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winPromotionDo) Delete(models ...*model.WinPromotion) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winPromotionDo) withDO(do gen.Dao) *winPromotionDo {
	w.DO = *do.(*gen.DO)
	return w
}