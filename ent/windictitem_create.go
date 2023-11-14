// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.skig.tech/zero-core/common/ent/windictitem"
)

// WinDictItemCreate is the builder for creating a WinDictItem entity.
type WinDictItemCreate struct {
	config
	mutation *WinDictItemMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (wdic *WinDictItemCreate) SetCode(s string) *WinDictItemCreate {
	wdic.mutation.SetCode(s)
	return wdic
}

// SetTitle sets the "title" field.
func (wdic *WinDictItemCreate) SetTitle(s string) *WinDictItemCreate {
	wdic.mutation.SetTitle(s)
	return wdic
}

// SetRemark sets the "remark" field.
func (wdic *WinDictItemCreate) SetRemark(s string) *WinDictItemCreate {
	wdic.mutation.SetRemark(s)
	return wdic
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (wdic *WinDictItemCreate) SetNillableRemark(s *string) *WinDictItemCreate {
	if s != nil {
		wdic.SetRemark(*s)
	}
	return wdic
}

// SetSort sets the "sort" field.
func (wdic *WinDictItemCreate) SetSort(i int32) *WinDictItemCreate {
	wdic.mutation.SetSort(i)
	return wdic
}

// SetReferID sets the "refer_id" field.
func (wdic *WinDictItemCreate) SetReferID(i int32) *WinDictItemCreate {
	wdic.mutation.SetReferID(i)
	return wdic
}

// SetStatus sets the "status" field.
func (wdic *WinDictItemCreate) SetStatus(b bool) *WinDictItemCreate {
	wdic.mutation.SetStatus(b)
	return wdic
}

// SetIsShow sets the "is_show" field.
func (wdic *WinDictItemCreate) SetIsShow(i int8) *WinDictItemCreate {
	wdic.mutation.SetIsShow(i)
	return wdic
}

// SetCreatedAt sets the "created_at" field.
func (wdic *WinDictItemCreate) SetCreatedAt(i int32) *WinDictItemCreate {
	wdic.mutation.SetCreatedAt(i)
	return wdic
}

// SetUpdatedAt sets the "updated_at" field.
func (wdic *WinDictItemCreate) SetUpdatedAt(i int32) *WinDictItemCreate {
	wdic.mutation.SetUpdatedAt(i)
	return wdic
}

// SetID sets the "id" field.
func (wdic *WinDictItemCreate) SetID(i int32) *WinDictItemCreate {
	wdic.mutation.SetID(i)
	return wdic
}

// Mutation returns the WinDictItemMutation object of the builder.
func (wdic *WinDictItemCreate) Mutation() *WinDictItemMutation {
	return wdic.mutation
}

// Save creates the WinDictItem in the database.
func (wdic *WinDictItemCreate) Save(ctx context.Context) (*WinDictItem, error) {
	return withHooks(ctx, wdic.sqlSave, wdic.mutation, wdic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wdic *WinDictItemCreate) SaveX(ctx context.Context) *WinDictItem {
	v, err := wdic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wdic *WinDictItemCreate) Exec(ctx context.Context) error {
	_, err := wdic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdic *WinDictItemCreate) ExecX(ctx context.Context) {
	if err := wdic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wdic *WinDictItemCreate) check() error {
	if _, ok := wdic.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "WinDictItem.code"`)}
	}
	if _, ok := wdic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "WinDictItem.title"`)}
	}
	if _, ok := wdic.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "WinDictItem.sort"`)}
	}
	if _, ok := wdic.mutation.ReferID(); !ok {
		return &ValidationError{Name: "refer_id", err: errors.New(`ent: missing required field "WinDictItem.refer_id"`)}
	}
	if _, ok := wdic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "WinDictItem.status"`)}
	}
	if _, ok := wdic.mutation.IsShow(); !ok {
		return &ValidationError{Name: "is_show", err: errors.New(`ent: missing required field "WinDictItem.is_show"`)}
	}
	if _, ok := wdic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WinDictItem.created_at"`)}
	}
	if _, ok := wdic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WinDictItem.updated_at"`)}
	}
	return nil
}

func (wdic *WinDictItemCreate) sqlSave(ctx context.Context) (*WinDictItem, error) {
	if err := wdic.check(); err != nil {
		return nil, err
	}
	_node, _spec := wdic.createSpec()
	if err := sqlgraph.CreateNode(ctx, wdic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	wdic.mutation.id = &_node.ID
	wdic.mutation.done = true
	return _node, nil
}

func (wdic *WinDictItemCreate) createSpec() (*WinDictItem, *sqlgraph.CreateSpec) {
	var (
		_node = &WinDictItem{config: wdic.config}
		_spec = sqlgraph.NewCreateSpec(windictitem.Table, sqlgraph.NewFieldSpec(windictitem.FieldID, field.TypeInt32))
	)
	if id, ok := wdic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wdic.mutation.Code(); ok {
		_spec.SetField(windictitem.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := wdic.mutation.Title(); ok {
		_spec.SetField(windictitem.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := wdic.mutation.Remark(); ok {
		_spec.SetField(windictitem.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := wdic.mutation.Sort(); ok {
		_spec.SetField(windictitem.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := wdic.mutation.ReferID(); ok {
		_spec.SetField(windictitem.FieldReferID, field.TypeInt32, value)
		_node.ReferID = value
	}
	if value, ok := wdic.mutation.Status(); ok {
		_spec.SetField(windictitem.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	if value, ok := wdic.mutation.IsShow(); ok {
		_spec.SetField(windictitem.FieldIsShow, field.TypeInt8, value)
		_node.IsShow = value
	}
	if value, ok := wdic.mutation.CreatedAt(); ok {
		_spec.SetField(windictitem.FieldCreatedAt, field.TypeInt32, value)
		_node.CreatedAt = value
	}
	if value, ok := wdic.mutation.UpdatedAt(); ok {
		_spec.SetField(windictitem.FieldUpdatedAt, field.TypeInt32, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WinDictItemCreateBulk is the builder for creating many WinDictItem entities in bulk.
type WinDictItemCreateBulk struct {
	config
	builders []*WinDictItemCreate
}

// Save creates the WinDictItem entities in the database.
func (wdicb *WinDictItemCreateBulk) Save(ctx context.Context) ([]*WinDictItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wdicb.builders))
	nodes := make([]*WinDictItem, len(wdicb.builders))
	mutators := make([]Mutator, len(wdicb.builders))
	for i := range wdicb.builders {
		func(i int, root context.Context) {
			builder := wdicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WinDictItemMutation)
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
					_, err = mutators[i+1].Mutate(root, wdicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wdicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wdicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wdicb *WinDictItemCreateBulk) SaveX(ctx context.Context) []*WinDictItem {
	v, err := wdicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wdicb *WinDictItemCreateBulk) Exec(ctx context.Context) error {
	_, err := wdicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdicb *WinDictItemCreateBulk) ExecX(ctx context.Context) {
	if err := wdicb.Exec(ctx); err != nil {
		panic(err)
	}
}
