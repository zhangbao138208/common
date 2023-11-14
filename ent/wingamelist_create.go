// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.skig.tech/zero-core/common/ent/wingamelist"
)

// WinGameListCreate is the builder for creating a WinGameList entity.
type WinGameListCreate struct {
	config
	mutation *WinGameListMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (wglc *WinGameListCreate) SetCode(s string) *WinGameListCreate {
	wglc.mutation.SetCode(s)
	return wglc
}

// SetName sets the "name" field.
func (wglc *WinGameListCreate) SetName(s string) *WinGameListCreate {
	wglc.mutation.SetName(s)
	return wglc
}

// SetIcon sets the "icon" field.
func (wglc *WinGameListCreate) SetIcon(s string) *WinGameListCreate {
	wglc.mutation.SetIcon(s)
	return wglc
}

// SetGroupID sets the "group_id" field.
func (wglc *WinGameListCreate) SetGroupID(i int8) *WinGameListCreate {
	wglc.mutation.SetGroupID(i)
	return wglc
}

// SetPlatListID sets the "plat_list_id" field.
func (wglc *WinGameListCreate) SetPlatListID(i int32) *WinGameListCreate {
	wglc.mutation.SetPlatListID(i)
	return wglc
}

// SetRevenueRate sets the "revenue_rate" field.
func (wglc *WinGameListCreate) SetRevenueRate(f float64) *WinGameListCreate {
	wglc.mutation.SetRevenueRate(f)
	return wglc
}

// SetMaintenance sets the "maintenance" field.
func (wglc *WinGameListCreate) SetMaintenance(s string) *WinGameListCreate {
	wglc.mutation.SetMaintenance(s)
	return wglc
}

// SetNillableMaintenance sets the "maintenance" field if the given value is not nil.
func (wglc *WinGameListCreate) SetNillableMaintenance(s *string) *WinGameListCreate {
	if s != nil {
		wglc.SetMaintenance(*s)
	}
	return wglc
}

// SetGameCount sets the "game_count" field.
func (wglc *WinGameListCreate) SetGameCount(i int32) *WinGameListCreate {
	wglc.mutation.SetGameCount(i)
	return wglc
}

// SetNillableGameCount sets the "game_count" field if the given value is not nil.
func (wglc *WinGameListCreate) SetNillableGameCount(i *int32) *WinGameListCreate {
	if i != nil {
		wglc.SetGameCount(*i)
	}
	return wglc
}

// SetRemark sets the "remark" field.
func (wglc *WinGameListCreate) SetRemark(s string) *WinGameListCreate {
	wglc.mutation.SetRemark(s)
	return wglc
}

// SetSort sets the "sort" field.
func (wglc *WinGameListCreate) SetSort(i int8) *WinGameListCreate {
	wglc.mutation.SetSort(i)
	return wglc
}

// SetStatus sets the "status" field.
func (wglc *WinGameListCreate) SetStatus(i int8) *WinGameListCreate {
	wglc.mutation.SetStatus(i)
	return wglc
}

// SetCreatedAt sets the "created_at" field.
func (wglc *WinGameListCreate) SetCreatedAt(i int32) *WinGameListCreate {
	wglc.mutation.SetCreatedAt(i)
	return wglc
}

// SetUpdatedAt sets the "updated_at" field.
func (wglc *WinGameListCreate) SetUpdatedAt(i int32) *WinGameListCreate {
	wglc.mutation.SetUpdatedAt(i)
	return wglc
}

// SetUpdatedUser sets the "updated_user" field.
func (wglc *WinGameListCreate) SetUpdatedUser(s string) *WinGameListCreate {
	wglc.mutation.SetUpdatedUser(s)
	return wglc
}

// SetNillableUpdatedUser sets the "updated_user" field if the given value is not nil.
func (wglc *WinGameListCreate) SetNillableUpdatedUser(s *string) *WinGameListCreate {
	if s != nil {
		wglc.SetUpdatedUser(*s)
	}
	return wglc
}

// SetOperatorName sets the "operator_name" field.
func (wglc *WinGameListCreate) SetOperatorName(s string) *WinGameListCreate {
	wglc.mutation.SetOperatorName(s)
	return wglc
}

// SetNillableOperatorName sets the "operator_name" field if the given value is not nil.
func (wglc *WinGameListCreate) SetNillableOperatorName(s *string) *WinGameListCreate {
	if s != nil {
		wglc.SetOperatorName(*s)
	}
	return wglc
}

// SetCategory sets the "category" field.
func (wglc *WinGameListCreate) SetCategory(s string) *WinGameListCreate {
	wglc.mutation.SetCategory(s)
	return wglc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (wglc *WinGameListCreate) SetNillableCategory(s *string) *WinGameListCreate {
	if s != nil {
		wglc.SetCategory(*s)
	}
	return wglc
}

// SetID sets the "id" field.
func (wglc *WinGameListCreate) SetID(i int32) *WinGameListCreate {
	wglc.mutation.SetID(i)
	return wglc
}

// Mutation returns the WinGameListMutation object of the builder.
func (wglc *WinGameListCreate) Mutation() *WinGameListMutation {
	return wglc.mutation
}

// Save creates the WinGameList in the database.
func (wglc *WinGameListCreate) Save(ctx context.Context) (*WinGameList, error) {
	return withHooks(ctx, wglc.sqlSave, wglc.mutation, wglc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wglc *WinGameListCreate) SaveX(ctx context.Context) *WinGameList {
	v, err := wglc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wglc *WinGameListCreate) Exec(ctx context.Context) error {
	_, err := wglc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wglc *WinGameListCreate) ExecX(ctx context.Context) {
	if err := wglc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wglc *WinGameListCreate) check() error {
	if _, ok := wglc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "WinGameList.code"`)}
	}
	if _, ok := wglc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WinGameList.name"`)}
	}
	if _, ok := wglc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "WinGameList.icon"`)}
	}
	if _, ok := wglc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "WinGameList.group_id"`)}
	}
	if _, ok := wglc.mutation.PlatListID(); !ok {
		return &ValidationError{Name: "plat_list_id", err: errors.New(`ent: missing required field "WinGameList.plat_list_id"`)}
	}
	if _, ok := wglc.mutation.RevenueRate(); !ok {
		return &ValidationError{Name: "revenue_rate", err: errors.New(`ent: missing required field "WinGameList.revenue_rate"`)}
	}
	if _, ok := wglc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "WinGameList.remark"`)}
	}
	if _, ok := wglc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "WinGameList.sort"`)}
	}
	if _, ok := wglc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "WinGameList.status"`)}
	}
	if _, ok := wglc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WinGameList.created_at"`)}
	}
	if _, ok := wglc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WinGameList.updated_at"`)}
	}
	return nil
}

func (wglc *WinGameListCreate) sqlSave(ctx context.Context) (*WinGameList, error) {
	if err := wglc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wglc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wglc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	wglc.mutation.id = &_node.ID
	wglc.mutation.done = true
	return _node, nil
}

func (wglc *WinGameListCreate) createSpec() (*WinGameList, *sqlgraph.CreateSpec) {
	var (
		_node = &WinGameList{config: wglc.config}
		_spec = sqlgraph.NewCreateSpec(wingamelist.Table, sqlgraph.NewFieldSpec(wingamelist.FieldID, field.TypeInt32))
	)
	if id, ok := wglc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wglc.mutation.Code(); ok {
		_spec.SetField(wingamelist.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := wglc.mutation.Name(); ok {
		_spec.SetField(wingamelist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wglc.mutation.Icon(); ok {
		_spec.SetField(wingamelist.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := wglc.mutation.GroupID(); ok {
		_spec.SetField(wingamelist.FieldGroupID, field.TypeInt8, value)
		_node.GroupID = value
	}
	if value, ok := wglc.mutation.PlatListID(); ok {
		_spec.SetField(wingamelist.FieldPlatListID, field.TypeInt32, value)
		_node.PlatListID = value
	}
	if value, ok := wglc.mutation.RevenueRate(); ok {
		_spec.SetField(wingamelist.FieldRevenueRate, field.TypeFloat64, value)
		_node.RevenueRate = value
	}
	if value, ok := wglc.mutation.Maintenance(); ok {
		_spec.SetField(wingamelist.FieldMaintenance, field.TypeString, value)
		_node.Maintenance = value
	}
	if value, ok := wglc.mutation.GameCount(); ok {
		_spec.SetField(wingamelist.FieldGameCount, field.TypeInt32, value)
		_node.GameCount = value
	}
	if value, ok := wglc.mutation.Remark(); ok {
		_spec.SetField(wingamelist.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := wglc.mutation.Sort(); ok {
		_spec.SetField(wingamelist.FieldSort, field.TypeInt8, value)
		_node.Sort = value
	}
	if value, ok := wglc.mutation.Status(); ok {
		_spec.SetField(wingamelist.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := wglc.mutation.CreatedAt(); ok {
		_spec.SetField(wingamelist.FieldCreatedAt, field.TypeInt32, value)
		_node.CreatedAt = value
	}
	if value, ok := wglc.mutation.UpdatedAt(); ok {
		_spec.SetField(wingamelist.FieldUpdatedAt, field.TypeInt32, value)
		_node.UpdatedAt = value
	}
	if value, ok := wglc.mutation.UpdatedUser(); ok {
		_spec.SetField(wingamelist.FieldUpdatedUser, field.TypeString, value)
		_node.UpdatedUser = value
	}
	if value, ok := wglc.mutation.OperatorName(); ok {
		_spec.SetField(wingamelist.FieldOperatorName, field.TypeString, value)
		_node.OperatorName = value
	}
	if value, ok := wglc.mutation.Category(); ok {
		_spec.SetField(wingamelist.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	return _node, _spec
}

// WinGameListCreateBulk is the builder for creating many WinGameList entities in bulk.
type WinGameListCreateBulk struct {
	config
	builders []*WinGameListCreate
}

// Save creates the WinGameList entities in the database.
func (wglcb *WinGameListCreateBulk) Save(ctx context.Context) ([]*WinGameList, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wglcb.builders))
	nodes := make([]*WinGameList, len(wglcb.builders))
	mutators := make([]Mutator, len(wglcb.builders))
	for i := range wglcb.builders {
		func(i int, root context.Context) {
			builder := wglcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WinGameListMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wglcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wglcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wglcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wglcb *WinGameListCreateBulk) SaveX(ctx context.Context) []*WinGameList {
	v, err := wglcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wglcb *WinGameListCreateBulk) Exec(ctx context.Context) error {
	_, err := wglcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wglcb *WinGameListCreateBulk) ExecX(ctx context.Context) {
	if err := wglcb.Exec(ctx); err != nil {
		panic(err)
	}
}