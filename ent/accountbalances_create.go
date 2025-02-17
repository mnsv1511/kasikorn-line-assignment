// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountbalances"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accounts"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

// AccountBalancesCreate is the builder for creating a AccountBalances entity.
type AccountBalancesCreate struct {
	config
	mutation *AccountBalancesMutation
	hooks    []Hook
}

// SetAccountID sets the "account_id" field.
func (abc *AccountBalancesCreate) SetAccountID(i int) *AccountBalancesCreate {
	abc.mutation.SetAccountID(i)
	return abc
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (abc *AccountBalancesCreate) SetNillableAccountID(i *int) *AccountBalancesCreate {
	if i != nil {
		abc.SetAccountID(*i)
	}
	return abc
}

// SetUserID sets the "user_id" field.
func (abc *AccountBalancesCreate) SetUserID(i int) *AccountBalancesCreate {
	abc.mutation.SetUserID(i)
	return abc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (abc *AccountBalancesCreate) SetNillableUserID(i *int) *AccountBalancesCreate {
	if i != nil {
		abc.SetUserID(*i)
	}
	return abc
}

// SetAmount sets the "amount" field.
func (abc *AccountBalancesCreate) SetAmount(f float64) *AccountBalancesCreate {
	abc.mutation.SetAmount(f)
	return abc
}

// SetUsersID sets the "users" edge to the Users entity by ID.
func (abc *AccountBalancesCreate) SetUsersID(id int) *AccountBalancesCreate {
	abc.mutation.SetUsersID(id)
	return abc
}

// SetNillableUsersID sets the "users" edge to the Users entity by ID if the given value is not nil.
func (abc *AccountBalancesCreate) SetNillableUsersID(id *int) *AccountBalancesCreate {
	if id != nil {
		abc = abc.SetUsersID(*id)
	}
	return abc
}

// SetUsers sets the "users" edge to the Users entity.
func (abc *AccountBalancesCreate) SetUsers(u *Users) *AccountBalancesCreate {
	return abc.SetUsersID(u.ID)
}

// SetAccountsID sets the "accounts" edge to the Accounts entity by ID.
func (abc *AccountBalancesCreate) SetAccountsID(id int) *AccountBalancesCreate {
	abc.mutation.SetAccountsID(id)
	return abc
}

// SetNillableAccountsID sets the "accounts" edge to the Accounts entity by ID if the given value is not nil.
func (abc *AccountBalancesCreate) SetNillableAccountsID(id *int) *AccountBalancesCreate {
	if id != nil {
		abc = abc.SetAccountsID(*id)
	}
	return abc
}

// SetAccounts sets the "accounts" edge to the Accounts entity.
func (abc *AccountBalancesCreate) SetAccounts(a *Accounts) *AccountBalancesCreate {
	return abc.SetAccountsID(a.ID)
}

// Mutation returns the AccountBalancesMutation object of the builder.
func (abc *AccountBalancesCreate) Mutation() *AccountBalancesMutation {
	return abc.mutation
}

// Save creates the AccountBalances in the database.
func (abc *AccountBalancesCreate) Save(ctx context.Context) (*AccountBalances, error) {
	return withHooks(ctx, abc.sqlSave, abc.mutation, abc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (abc *AccountBalancesCreate) SaveX(ctx context.Context) *AccountBalances {
	v, err := abc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (abc *AccountBalancesCreate) Exec(ctx context.Context) error {
	_, err := abc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abc *AccountBalancesCreate) ExecX(ctx context.Context) {
	if err := abc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (abc *AccountBalancesCreate) check() error {
	if _, ok := abc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "AccountBalances.amount"`)}
	}
	return nil
}

func (abc *AccountBalancesCreate) sqlSave(ctx context.Context) (*AccountBalances, error) {
	if err := abc.check(); err != nil {
		return nil, err
	}
	_node, _spec := abc.createSpec()
	if err := sqlgraph.CreateNode(ctx, abc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	abc.mutation.id = &_node.ID
	abc.mutation.done = true
	return _node, nil
}

func (abc *AccountBalancesCreate) createSpec() (*AccountBalances, *sqlgraph.CreateSpec) {
	var (
		_node = &AccountBalances{config: abc.config}
		_spec = sqlgraph.NewCreateSpec(accountbalances.Table, sqlgraph.NewFieldSpec(accountbalances.FieldID, field.TypeInt))
	)
	if value, ok := abc.mutation.Amount(); ok {
		_spec.SetField(accountbalances.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if nodes := abc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountbalances.UsersTable,
			Columns: []string{accountbalances.UsersColumn},
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
	if nodes := abc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountbalances.AccountsTable,
			Columns: []string{accountbalances.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accounts.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountBalancesCreateBulk is the builder for creating many AccountBalances entities in bulk.
type AccountBalancesCreateBulk struct {
	config
	err      error
	builders []*AccountBalancesCreate
}

// Save creates the AccountBalances entities in the database.
func (abcb *AccountBalancesCreateBulk) Save(ctx context.Context) ([]*AccountBalances, error) {
	if abcb.err != nil {
		return nil, abcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(abcb.builders))
	nodes := make([]*AccountBalances, len(abcb.builders))
	mutators := make([]Mutator, len(abcb.builders))
	for i := range abcb.builders {
		func(i int, root context.Context) {
			builder := abcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountBalancesMutation)
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
					_, err = mutators[i+1].Mutate(root, abcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, abcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, abcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (abcb *AccountBalancesCreateBulk) SaveX(ctx context.Context) []*AccountBalances {
	v, err := abcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (abcb *AccountBalancesCreateBulk) Exec(ctx context.Context) error {
	_, err := abcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abcb *AccountBalancesCreateBulk) ExecX(ctx context.Context) {
	if err := abcb.Exec(ctx); err != nil {
		panic(err)
	}
}
