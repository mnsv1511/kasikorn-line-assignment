// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcards"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcardstatus"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

// DebitCardStatusCreate is the builder for creating a DebitCardStatus entity.
type DebitCardStatusCreate struct {
	config
	mutation *DebitCardStatusMutation
	hooks    []Hook
}

// SetCardID sets the "card_id" field.
func (dcsc *DebitCardStatusCreate) SetCardID(i int) *DebitCardStatusCreate {
	dcsc.mutation.SetCardID(i)
	return dcsc
}

// SetNillableCardID sets the "card_id" field if the given value is not nil.
func (dcsc *DebitCardStatusCreate) SetNillableCardID(i *int) *DebitCardStatusCreate {
	if i != nil {
		dcsc.SetCardID(*i)
	}
	return dcsc
}

// SetUserID sets the "user_id" field.
func (dcsc *DebitCardStatusCreate) SetUserID(i int) *DebitCardStatusCreate {
	dcsc.mutation.SetUserID(i)
	return dcsc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (dcsc *DebitCardStatusCreate) SetNillableUserID(i *int) *DebitCardStatusCreate {
	if i != nil {
		dcsc.SetUserID(*i)
	}
	return dcsc
}

// SetStatus sets the "status" field.
func (dcsc *DebitCardStatusCreate) SetStatus(s string) *DebitCardStatusCreate {
	dcsc.mutation.SetStatus(s)
	return dcsc
}

// SetUsersID sets the "users" edge to the Users entity by ID.
func (dcsc *DebitCardStatusCreate) SetUsersID(id int) *DebitCardStatusCreate {
	dcsc.mutation.SetUsersID(id)
	return dcsc
}

// SetNillableUsersID sets the "users" edge to the Users entity by ID if the given value is not nil.
func (dcsc *DebitCardStatusCreate) SetNillableUsersID(id *int) *DebitCardStatusCreate {
	if id != nil {
		dcsc = dcsc.SetUsersID(*id)
	}
	return dcsc
}

// SetUsers sets the "users" edge to the Users entity.
func (dcsc *DebitCardStatusCreate) SetUsers(u *Users) *DebitCardStatusCreate {
	return dcsc.SetUsersID(u.ID)
}

// SetDebitCardsID sets the "debit_cards" edge to the DebitCards entity by ID.
func (dcsc *DebitCardStatusCreate) SetDebitCardsID(id int) *DebitCardStatusCreate {
	dcsc.mutation.SetDebitCardsID(id)
	return dcsc
}

// SetNillableDebitCardsID sets the "debit_cards" edge to the DebitCards entity by ID if the given value is not nil.
func (dcsc *DebitCardStatusCreate) SetNillableDebitCardsID(id *int) *DebitCardStatusCreate {
	if id != nil {
		dcsc = dcsc.SetDebitCardsID(*id)
	}
	return dcsc
}

// SetDebitCards sets the "debit_cards" edge to the DebitCards entity.
func (dcsc *DebitCardStatusCreate) SetDebitCards(d *DebitCards) *DebitCardStatusCreate {
	return dcsc.SetDebitCardsID(d.ID)
}

// Mutation returns the DebitCardStatusMutation object of the builder.
func (dcsc *DebitCardStatusCreate) Mutation() *DebitCardStatusMutation {
	return dcsc.mutation
}

// Save creates the DebitCardStatus in the database.
func (dcsc *DebitCardStatusCreate) Save(ctx context.Context) (*DebitCardStatus, error) {
	return withHooks(ctx, dcsc.sqlSave, dcsc.mutation, dcsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dcsc *DebitCardStatusCreate) SaveX(ctx context.Context) *DebitCardStatus {
	v, err := dcsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcsc *DebitCardStatusCreate) Exec(ctx context.Context) error {
	_, err := dcsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcsc *DebitCardStatusCreate) ExecX(ctx context.Context) {
	if err := dcsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcsc *DebitCardStatusCreate) check() error {
	if _, ok := dcsc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "DebitCardStatus.status"`)}
	}
	return nil
}

func (dcsc *DebitCardStatusCreate) sqlSave(ctx context.Context) (*DebitCardStatus, error) {
	if err := dcsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dcsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dcsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dcsc.mutation.id = &_node.ID
	dcsc.mutation.done = true
	return _node, nil
}

func (dcsc *DebitCardStatusCreate) createSpec() (*DebitCardStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &DebitCardStatus{config: dcsc.config}
		_spec = sqlgraph.NewCreateSpec(debitcardstatus.Table, sqlgraph.NewFieldSpec(debitcardstatus.FieldID, field.TypeInt))
	)
	if value, ok := dcsc.mutation.Status(); ok {
		_spec.SetField(debitcardstatus.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if nodes := dcsc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   debitcardstatus.UsersTable,
			Columns: []string{debitcardstatus.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dcsc.mutation.DebitCardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   debitcardstatus.DebitCardsTable,
			Columns: []string{debitcardstatus.DebitCardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(debitcards.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CardID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DebitCardStatusCreateBulk is the builder for creating many DebitCardStatus entities in bulk.
type DebitCardStatusCreateBulk struct {
	config
	err      error
	builders []*DebitCardStatusCreate
}

// Save creates the DebitCardStatus entities in the database.
func (dcscb *DebitCardStatusCreateBulk) Save(ctx context.Context) ([]*DebitCardStatus, error) {
	if dcscb.err != nil {
		return nil, dcscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcscb.builders))
	nodes := make([]*DebitCardStatus, len(dcscb.builders))
	mutators := make([]Mutator, len(dcscb.builders))
	for i := range dcscb.builders {
		func(i int, root context.Context) {
			builder := dcscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DebitCardStatusMutation)
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
					_, err = mutators[i+1].Mutate(root, dcscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcscb *DebitCardStatusCreateBulk) SaveX(ctx context.Context) []*DebitCardStatus {
	v, err := dcscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcscb *DebitCardStatusCreateBulk) Exec(ctx context.Context) error {
	_, err := dcscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcscb *DebitCardStatusCreateBulk) ExecX(ctx context.Context) {
	if err := dcscb.Exec(ctx); err != nil {
		panic(err)
	}
}
