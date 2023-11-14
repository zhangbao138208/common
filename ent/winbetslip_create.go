// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.skig.tech/zero-core/common/ent/winbetslip"
)

// WinBetslipCreate is the builder for creating a WinBetslip entity.
type WinBetslipCreate struct {
	config
	mutation *WinBetslipMutation
	hooks    []Hook
}

// SetRoundID sets the "round_id" field.
func (wbc *WinBetslipCreate) SetRoundID(s string) *WinBetslipCreate {
	wbc.mutation.SetRoundID(s)
	return wbc
}

// SetTransactionID sets the "transaction_id" field.
func (wbc *WinBetslipCreate) SetTransactionID(s string) *WinBetslipCreate {
	wbc.mutation.SetTransactionID(s)
	return wbc
}

// SetGamePlatID sets the "game_plat_id" field.
func (wbc *WinBetslipCreate) SetGamePlatID(i int32) *WinBetslipCreate {
	wbc.mutation.SetGamePlatID(i)
	return wbc
}

// SetXbStatus sets the "xb_status" field.
func (wbc *WinBetslipCreate) SetXbStatus(i int8) *WinBetslipCreate {
	wbc.mutation.SetXbStatus(i)
	return wbc
}

// SetXbUID sets the "xb_uid" field.
func (wbc *WinBetslipCreate) SetXbUID(u uint32) *WinBetslipCreate {
	wbc.mutation.SetXbUID(u)
	return wbc
}

// SetXbUsername sets the "xb_username" field.
func (wbc *WinBetslipCreate) SetXbUsername(s string) *WinBetslipCreate {
	wbc.mutation.SetXbUsername(s)
	return wbc
}

// SetXbProfit sets the "xb_profit" field.
func (wbc *WinBetslipCreate) SetXbProfit(f float64) *WinBetslipCreate {
	wbc.mutation.SetXbProfit(f)
	return wbc
}

// SetNillableXbProfit sets the "xb_profit" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableXbProfit(f *float64) *WinBetslipCreate {
	if f != nil {
		wbc.SetXbProfit(*f)
	}
	return wbc
}

// SetStake sets the "stake" field.
func (wbc *WinBetslipCreate) SetStake(f float64) *WinBetslipCreate {
	wbc.mutation.SetStake(f)
	return wbc
}

// SetValidStake sets the "valid_stake" field.
func (wbc *WinBetslipCreate) SetValidStake(f float64) *WinBetslipCreate {
	wbc.mutation.SetValidStake(f)
	return wbc
}

// SetPayout sets the "payout" field.
func (wbc *WinBetslipCreate) SetPayout(f float64) *WinBetslipCreate {
	wbc.mutation.SetPayout(f)
	return wbc
}

// SetNillablePayout sets the "payout" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillablePayout(f *float64) *WinBetslipCreate {
	if f != nil {
		wbc.SetPayout(*f)
	}
	return wbc
}

// SetCoinRefund sets the "coin_refund" field.
func (wbc *WinBetslipCreate) SetCoinRefund(f float64) *WinBetslipCreate {
	wbc.mutation.SetCoinRefund(f)
	return wbc
}

// SetCoinBefore sets the "coin_before" field.
func (wbc *WinBetslipCreate) SetCoinBefore(f float64) *WinBetslipCreate {
	wbc.mutation.SetCoinBefore(f)
	return wbc
}

// SetGameID sets the "game_id" field.
func (wbc *WinBetslipCreate) SetGameID(s string) *WinBetslipCreate {
	wbc.mutation.SetGameID(s)
	return wbc
}

// SetGameName sets the "game_name" field.
func (wbc *WinBetslipCreate) SetGameName(s string) *WinBetslipCreate {
	wbc.mutation.SetGameName(s)
	return wbc
}

// SetNillableGameName sets the "game_name" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableGameName(s *string) *WinBetslipCreate {
	if s != nil {
		wbc.SetGameName(*s)
	}
	return wbc
}

// SetAmountType sets the "amount_type" field.
func (wbc *WinBetslipCreate) SetAmountType(i int32) *WinBetslipCreate {
	wbc.mutation.SetAmountType(i)
	return wbc
}

// SetGameTypeID sets the "game_type_id" field.
func (wbc *WinBetslipCreate) SetGameTypeID(s string) *WinBetslipCreate {
	wbc.mutation.SetGameTypeID(s)
	return wbc
}

// SetGameGroupID sets the "game_group_id" field.
func (wbc *WinBetslipCreate) SetGameGroupID(i int32) *WinBetslipCreate {
	wbc.mutation.SetGameGroupID(i)
	return wbc
}

// SetNillableGameGroupID sets the "game_group_id" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableGameGroupID(i *int32) *WinBetslipCreate {
	if i != nil {
		wbc.SetGameGroupID(*i)
	}
	return wbc
}

// SetSportType sets the "sport_type" field.
func (wbc *WinBetslipCreate) SetSportType(s string) *WinBetslipCreate {
	wbc.mutation.SetSportType(s)
	return wbc
}

// SetNillableSportType sets the "sport_type" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableSportType(s *string) *WinBetslipCreate {
	if s != nil {
		wbc.SetSportType(*s)
	}
	return wbc
}

// SetDtStarted sets the "dt_started" field.
func (wbc *WinBetslipCreate) SetDtStarted(i int) *WinBetslipCreate {
	wbc.mutation.SetDtStarted(i)
	return wbc
}

// SetDtCompleted sets the "dt_completed" field.
func (wbc *WinBetslipCreate) SetDtCompleted(i int) *WinBetslipCreate {
	wbc.mutation.SetDtCompleted(i)
	return wbc
}

// SetNillableDtCompleted sets the "dt_completed" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableDtCompleted(i *int) *WinBetslipCreate {
	if i != nil {
		wbc.SetDtCompleted(*i)
	}
	return wbc
}

// SetWinTransactionID sets the "win_transaction_id" field.
func (wbc *WinBetslipCreate) SetWinTransactionID(s string) *WinBetslipCreate {
	wbc.mutation.SetWinTransactionID(s)
	return wbc
}

// SetNillableWinTransactionID sets the "win_transaction_id" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableWinTransactionID(s *string) *WinBetslipCreate {
	if s != nil {
		wbc.SetWinTransactionID(*s)
	}
	return wbc
}

// SetBetJSON sets the "bet_json" field.
func (wbc *WinBetslipCreate) SetBetJSON(s string) *WinBetslipCreate {
	wbc.mutation.SetBetJSON(s)
	return wbc
}

// SetNillableBetJSON sets the "bet_json" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableBetJSON(s *string) *WinBetslipCreate {
	if s != nil {
		wbc.SetBetJSON(*s)
	}
	return wbc
}

// SetRewardJSON sets the "reward_json" field.
func (wbc *WinBetslipCreate) SetRewardJSON(s string) *WinBetslipCreate {
	wbc.mutation.SetRewardJSON(s)
	return wbc
}

// SetNillableRewardJSON sets the "reward_json" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableRewardJSON(s *string) *WinBetslipCreate {
	if s != nil {
		wbc.SetRewardJSON(*s)
	}
	return wbc
}

// SetRefundJSON sets the "refund_json" field.
func (wbc *WinBetslipCreate) SetRefundJSON(s string) *WinBetslipCreate {
	wbc.mutation.SetRefundJSON(s)
	return wbc
}

// SetNillableRefundJSON sets the "refund_json" field if the given value is not nil.
func (wbc *WinBetslipCreate) SetNillableRefundJSON(s *string) *WinBetslipCreate {
	if s != nil {
		wbc.SetRefundJSON(*s)
	}
	return wbc
}

// SetCreateTimeStr sets the "create_time_str" field.
func (wbc *WinBetslipCreate) SetCreateTimeStr(s string) *WinBetslipCreate {
	wbc.mutation.SetCreateTimeStr(s)
	return wbc
}

// SetCreatedAt sets the "created_at" field.
func (wbc *WinBetslipCreate) SetCreatedAt(i int32) *WinBetslipCreate {
	wbc.mutation.SetCreatedAt(i)
	return wbc
}

// SetUpdatedAt sets the "updated_at" field.
func (wbc *WinBetslipCreate) SetUpdatedAt(i int32) *WinBetslipCreate {
	wbc.mutation.SetUpdatedAt(i)
	return wbc
}

// SetID sets the "id" field.
func (wbc *WinBetslipCreate) SetID(i int) *WinBetslipCreate {
	wbc.mutation.SetID(i)
	return wbc
}

// Mutation returns the WinBetslipMutation object of the builder.
func (wbc *WinBetslipCreate) Mutation() *WinBetslipMutation {
	return wbc.mutation
}

// Save creates the WinBetslip in the database.
func (wbc *WinBetslipCreate) Save(ctx context.Context) (*WinBetslip, error) {
	return withHooks(ctx, wbc.sqlSave, wbc.mutation, wbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wbc *WinBetslipCreate) SaveX(ctx context.Context) *WinBetslip {
	v, err := wbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wbc *WinBetslipCreate) Exec(ctx context.Context) error {
	_, err := wbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wbc *WinBetslipCreate) ExecX(ctx context.Context) {
	if err := wbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wbc *WinBetslipCreate) check() error {
	if _, ok := wbc.mutation.RoundID(); !ok {
		return &ValidationError{Name: "round_id", err: errors.New(`ent: missing required field "WinBetslip.round_id"`)}
	}
	if _, ok := wbc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "transaction_id", err: errors.New(`ent: missing required field "WinBetslip.transaction_id"`)}
	}
	if _, ok := wbc.mutation.GamePlatID(); !ok {
		return &ValidationError{Name: "game_plat_id", err: errors.New(`ent: missing required field "WinBetslip.game_plat_id"`)}
	}
	if _, ok := wbc.mutation.XbStatus(); !ok {
		return &ValidationError{Name: "xb_status", err: errors.New(`ent: missing required field "WinBetslip.xb_status"`)}
	}
	if _, ok := wbc.mutation.XbUID(); !ok {
		return &ValidationError{Name: "xb_uid", err: errors.New(`ent: missing required field "WinBetslip.xb_uid"`)}
	}
	if _, ok := wbc.mutation.XbUsername(); !ok {
		return &ValidationError{Name: "xb_username", err: errors.New(`ent: missing required field "WinBetslip.xb_username"`)}
	}
	if _, ok := wbc.mutation.Stake(); !ok {
		return &ValidationError{Name: "stake", err: errors.New(`ent: missing required field "WinBetslip.stake"`)}
	}
	if _, ok := wbc.mutation.ValidStake(); !ok {
		return &ValidationError{Name: "valid_stake", err: errors.New(`ent: missing required field "WinBetslip.valid_stake"`)}
	}
	if _, ok := wbc.mutation.CoinRefund(); !ok {
		return &ValidationError{Name: "coin_refund", err: errors.New(`ent: missing required field "WinBetslip.coin_refund"`)}
	}
	if _, ok := wbc.mutation.CoinBefore(); !ok {
		return &ValidationError{Name: "coin_before", err: errors.New(`ent: missing required field "WinBetslip.coin_before"`)}
	}
	if _, ok := wbc.mutation.GameID(); !ok {
		return &ValidationError{Name: "game_id", err: errors.New(`ent: missing required field "WinBetslip.game_id"`)}
	}
	if _, ok := wbc.mutation.AmountType(); !ok {
		return &ValidationError{Name: "amount_type", err: errors.New(`ent: missing required field "WinBetslip.amount_type"`)}
	}
	if _, ok := wbc.mutation.GameTypeID(); !ok {
		return &ValidationError{Name: "game_type_id", err: errors.New(`ent: missing required field "WinBetslip.game_type_id"`)}
	}
	if _, ok := wbc.mutation.DtStarted(); !ok {
		return &ValidationError{Name: "dt_started", err: errors.New(`ent: missing required field "WinBetslip.dt_started"`)}
	}
	if _, ok := wbc.mutation.CreateTimeStr(); !ok {
		return &ValidationError{Name: "create_time_str", err: errors.New(`ent: missing required field "WinBetslip.create_time_str"`)}
	}
	if _, ok := wbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WinBetslip.created_at"`)}
	}
	if _, ok := wbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WinBetslip.updated_at"`)}
	}
	return nil
}

func (wbc *WinBetslipCreate) sqlSave(ctx context.Context) (*WinBetslip, error) {
	if err := wbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	wbc.mutation.id = &_node.ID
	wbc.mutation.done = true
	return _node, nil
}

func (wbc *WinBetslipCreate) createSpec() (*WinBetslip, *sqlgraph.CreateSpec) {
	var (
		_node = &WinBetslip{config: wbc.config}
		_spec = sqlgraph.NewCreateSpec(winbetslip.Table, sqlgraph.NewFieldSpec(winbetslip.FieldID, field.TypeInt))
	)
	if id, ok := wbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wbc.mutation.RoundID(); ok {
		_spec.SetField(winbetslip.FieldRoundID, field.TypeString, value)
		_node.RoundID = value
	}
	if value, ok := wbc.mutation.TransactionID(); ok {
		_spec.SetField(winbetslip.FieldTransactionID, field.TypeString, value)
		_node.TransactionID = value
	}
	if value, ok := wbc.mutation.GamePlatID(); ok {
		_spec.SetField(winbetslip.FieldGamePlatID, field.TypeInt32, value)
		_node.GamePlatID = value
	}
	if value, ok := wbc.mutation.XbStatus(); ok {
		_spec.SetField(winbetslip.FieldXbStatus, field.TypeInt8, value)
		_node.XbStatus = value
	}
	if value, ok := wbc.mutation.XbUID(); ok {
		_spec.SetField(winbetslip.FieldXbUID, field.TypeUint32, value)
		_node.XbUID = value
	}
	if value, ok := wbc.mutation.XbUsername(); ok {
		_spec.SetField(winbetslip.FieldXbUsername, field.TypeString, value)
		_node.XbUsername = value
	}
	if value, ok := wbc.mutation.XbProfit(); ok {
		_spec.SetField(winbetslip.FieldXbProfit, field.TypeFloat64, value)
		_node.XbProfit = value
	}
	if value, ok := wbc.mutation.Stake(); ok {
		_spec.SetField(winbetslip.FieldStake, field.TypeFloat64, value)
		_node.Stake = value
	}
	if value, ok := wbc.mutation.ValidStake(); ok {
		_spec.SetField(winbetslip.FieldValidStake, field.TypeFloat64, value)
		_node.ValidStake = value
	}
	if value, ok := wbc.mutation.Payout(); ok {
		_spec.SetField(winbetslip.FieldPayout, field.TypeFloat64, value)
		_node.Payout = value
	}
	if value, ok := wbc.mutation.CoinRefund(); ok {
		_spec.SetField(winbetslip.FieldCoinRefund, field.TypeFloat64, value)
		_node.CoinRefund = value
	}
	if value, ok := wbc.mutation.CoinBefore(); ok {
		_spec.SetField(winbetslip.FieldCoinBefore, field.TypeFloat64, value)
		_node.CoinBefore = value
	}
	if value, ok := wbc.mutation.GameID(); ok {
		_spec.SetField(winbetslip.FieldGameID, field.TypeString, value)
		_node.GameID = value
	}
	if value, ok := wbc.mutation.GameName(); ok {
		_spec.SetField(winbetslip.FieldGameName, field.TypeString, value)
		_node.GameName = value
	}
	if value, ok := wbc.mutation.AmountType(); ok {
		_spec.SetField(winbetslip.FieldAmountType, field.TypeInt32, value)
		_node.AmountType = value
	}
	if value, ok := wbc.mutation.GameTypeID(); ok {
		_spec.SetField(winbetslip.FieldGameTypeID, field.TypeString, value)
		_node.GameTypeID = value
	}
	if value, ok := wbc.mutation.GameGroupID(); ok {
		_spec.SetField(winbetslip.FieldGameGroupID, field.TypeInt32, value)
		_node.GameGroupID = value
	}
	if value, ok := wbc.mutation.SportType(); ok {
		_spec.SetField(winbetslip.FieldSportType, field.TypeString, value)
		_node.SportType = value
	}
	if value, ok := wbc.mutation.DtStarted(); ok {
		_spec.SetField(winbetslip.FieldDtStarted, field.TypeInt, value)
		_node.DtStarted = value
	}
	if value, ok := wbc.mutation.DtCompleted(); ok {
		_spec.SetField(winbetslip.FieldDtCompleted, field.TypeInt, value)
		_node.DtCompleted = value
	}
	if value, ok := wbc.mutation.WinTransactionID(); ok {
		_spec.SetField(winbetslip.FieldWinTransactionID, field.TypeString, value)
		_node.WinTransactionID = value
	}
	if value, ok := wbc.mutation.BetJSON(); ok {
		_spec.SetField(winbetslip.FieldBetJSON, field.TypeString, value)
		_node.BetJSON = value
	}
	if value, ok := wbc.mutation.RewardJSON(); ok {
		_spec.SetField(winbetslip.FieldRewardJSON, field.TypeString, value)
		_node.RewardJSON = value
	}
	if value, ok := wbc.mutation.RefundJSON(); ok {
		_spec.SetField(winbetslip.FieldRefundJSON, field.TypeString, value)
		_node.RefundJSON = value
	}
	if value, ok := wbc.mutation.CreateTimeStr(); ok {
		_spec.SetField(winbetslip.FieldCreateTimeStr, field.TypeString, value)
		_node.CreateTimeStr = value
	}
	if value, ok := wbc.mutation.CreatedAt(); ok {
		_spec.SetField(winbetslip.FieldCreatedAt, field.TypeInt32, value)
		_node.CreatedAt = value
	}
	if value, ok := wbc.mutation.UpdatedAt(); ok {
		_spec.SetField(winbetslip.FieldUpdatedAt, field.TypeInt32, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WinBetslipCreateBulk is the builder for creating many WinBetslip entities in bulk.
type WinBetslipCreateBulk struct {
	config
	builders []*WinBetslipCreate
}

// Save creates the WinBetslip entities in the database.
func (wbcb *WinBetslipCreateBulk) Save(ctx context.Context) ([]*WinBetslip, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wbcb.builders))
	nodes := make([]*WinBetslip, len(wbcb.builders))
	mutators := make([]Mutator, len(wbcb.builders))
	for i := range wbcb.builders {
		func(i int, root context.Context) {
			builder := wbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WinBetslipMutation)
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
					_, err = mutators[i+1].Mutate(root, wbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wbcb.driver, spec); err != nil {
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
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, wbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wbcb *WinBetslipCreateBulk) SaveX(ctx context.Context) []*WinBetslip {
	v, err := wbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wbcb *WinBetslipCreateBulk) Exec(ctx context.Context) error {
	_, err := wbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wbcb *WinBetslipCreateBulk) ExecX(ctx context.Context) {
	if err := wbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
