// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountbalances"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountdetails"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountflags"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accounts"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/banners"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcarddesign"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcarddetails"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcards"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcardstatus"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/transactions"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/usergreetings"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	mutation *UsersMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (uc *UsersCreate) SetName(s string) *UsersCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetImageURL sets the "image_url" field.
func (uc *UsersCreate) SetImageURL(s string) *UsersCreate {
	uc.mutation.SetImageURL(s)
	return uc
}

// AddAccountIDs adds the "accounts" edge to the Accounts entity by IDs.
func (uc *UsersCreate) AddAccountIDs(ids ...int) *UsersCreate {
	uc.mutation.AddAccountIDs(ids...)
	return uc
}

// AddAccounts adds the "accounts" edges to the Accounts entity.
func (uc *UsersCreate) AddAccounts(a ...*Accounts) *UsersCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAccountIDs(ids...)
}

// AddBannerIDs adds the "banners" edge to the Banners entity by IDs.
func (uc *UsersCreate) AddBannerIDs(ids ...int) *UsersCreate {
	uc.mutation.AddBannerIDs(ids...)
	return uc
}

// AddBanners adds the "banners" edges to the Banners entity.
func (uc *UsersCreate) AddBanners(b ...*Banners) *UsersCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uc.AddBannerIDs(ids...)
}

// AddDebitCardIDs adds the "debit_cards" edge to the DebitCards entity by IDs.
func (uc *UsersCreate) AddDebitCardIDs(ids ...int) *UsersCreate {
	uc.mutation.AddDebitCardIDs(ids...)
	return uc
}

// AddDebitCards adds the "debit_cards" edges to the DebitCards entity.
func (uc *UsersCreate) AddDebitCards(d ...*DebitCards) *UsersCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDebitCardIDs(ids...)
}

// AddTransactionIDs adds the "transactions" edge to the Transactions entity by IDs.
func (uc *UsersCreate) AddTransactionIDs(ids ...int) *UsersCreate {
	uc.mutation.AddTransactionIDs(ids...)
	return uc
}

// AddTransactions adds the "transactions" edges to the Transactions entity.
func (uc *UsersCreate) AddTransactions(t ...*Transactions) *UsersCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTransactionIDs(ids...)
}

// AddUserGreetingIDs adds the "user_greetings" edge to the UserGreetings entity by IDs.
func (uc *UsersCreate) AddUserGreetingIDs(ids ...int) *UsersCreate {
	uc.mutation.AddUserGreetingIDs(ids...)
	return uc
}

// AddUserGreetings adds the "user_greetings" edges to the UserGreetings entity.
func (uc *UsersCreate) AddUserGreetings(u ...*UserGreetings) *UsersCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserGreetingIDs(ids...)
}

// AddAccountBalanceIDs adds the "account_balances" edge to the AccountBalances entity by IDs.
func (uc *UsersCreate) AddAccountBalanceIDs(ids ...int) *UsersCreate {
	uc.mutation.AddAccountBalanceIDs(ids...)
	return uc
}

// AddAccountBalances adds the "account_balances" edges to the AccountBalances entity.
func (uc *UsersCreate) AddAccountBalances(a ...*AccountBalances) *UsersCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAccountBalanceIDs(ids...)
}

// AddAccountDetailIDs adds the "account_details" edge to the AccountDetails entity by IDs.
func (uc *UsersCreate) AddAccountDetailIDs(ids ...int) *UsersCreate {
	uc.mutation.AddAccountDetailIDs(ids...)
	return uc
}

// AddAccountDetails adds the "account_details" edges to the AccountDetails entity.
func (uc *UsersCreate) AddAccountDetails(a ...*AccountDetails) *UsersCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAccountDetailIDs(ids...)
}

// AddAccountFlagIDs adds the "account_flags" edge to the AccountFlags entity by IDs.
func (uc *UsersCreate) AddAccountFlagIDs(ids ...int) *UsersCreate {
	uc.mutation.AddAccountFlagIDs(ids...)
	return uc
}

// AddAccountFlags adds the "account_flags" edges to the AccountFlags entity.
func (uc *UsersCreate) AddAccountFlags(a ...*AccountFlags) *UsersCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAccountFlagIDs(ids...)
}

// AddDebitCardDesignIDs adds the "debit_card_design" edge to the DebitCardDesign entity by IDs.
func (uc *UsersCreate) AddDebitCardDesignIDs(ids ...int) *UsersCreate {
	uc.mutation.AddDebitCardDesignIDs(ids...)
	return uc
}

// AddDebitCardDesign adds the "debit_card_design" edges to the DebitCardDesign entity.
func (uc *UsersCreate) AddDebitCardDesign(d ...*DebitCardDesign) *UsersCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDebitCardDesignIDs(ids...)
}

// AddDebitCardDetailIDs adds the "debit_card_details" edge to the DebitCardDetails entity by IDs.
func (uc *UsersCreate) AddDebitCardDetailIDs(ids ...int) *UsersCreate {
	uc.mutation.AddDebitCardDetailIDs(ids...)
	return uc
}

// AddDebitCardDetails adds the "debit_card_details" edges to the DebitCardDetails entity.
func (uc *UsersCreate) AddDebitCardDetails(d ...*DebitCardDetails) *UsersCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDebitCardDetailIDs(ids...)
}

// AddDebitCardStatuIDs adds the "debit_card_status" edge to the DebitCardStatus entity by IDs.
func (uc *UsersCreate) AddDebitCardStatuIDs(ids ...int) *UsersCreate {
	uc.mutation.AddDebitCardStatuIDs(ids...)
	return uc
}

// AddDebitCardStatus adds the "debit_card_status" edges to the DebitCardStatus entity.
func (uc *UsersCreate) AddDebitCardStatus(d ...*DebitCardStatus) *UsersCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDebitCardStatuIDs(ids...)
}

// Mutation returns the UsersMutation object of the builder.
func (uc *UsersCreate) Mutation() *UsersMutation {
	return uc.mutation
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UsersCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UsersCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UsersCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Users.name"`)}
	}
	if _, ok := uc.mutation.ImageURL(); !ok {
		return &ValidationError{Name: "image_url", err: errors.New(`ent: missing required field "Users.image_url"`)}
	}
	return nil
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UsersCreate) createSpec() (*Users, *sqlgraph.CreateSpec) {
	var (
		_node = &Users{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(users.Table, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(users.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.ImageURL(); ok {
		_spec.SetField(users.FieldImageURL, field.TypeString, value)
		_node.ImageURL = value
	}
	if nodes := uc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AccountsTable,
			Columns: []string{users.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accounts.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.BannersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.BannersTable,
			Columns: []string{users.BannersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banners.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DebitCardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.DebitCardsTable,
			Columns: []string{users.DebitCardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(debitcards.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.TransactionsTable,
			Columns: []string{users.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transactions.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserGreetingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.UserGreetingsTable,
			Columns: []string{users.UserGreetingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usergreetings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AccountBalancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AccountBalancesTable,
			Columns: []string{users.AccountBalancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accountbalances.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AccountDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AccountDetailsTable,
			Columns: []string{users.AccountDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accountdetails.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AccountFlagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AccountFlagsTable,
			Columns: []string{users.AccountFlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accountflags.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DebitCardDesignIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.DebitCardDesignTable,
			Columns: []string{users.DebitCardDesignColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(debitcarddesign.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DebitCardDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.DebitCardDetailsTable,
			Columns: []string{users.DebitCardDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(debitcarddetails.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DebitCardStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.DebitCardStatusTable,
			Columns: []string{users.DebitCardStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(debitcardstatus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UsersCreateBulk is the builder for creating many Users entities in bulk.
type UsersCreateBulk struct {
	config
	err      error
	builders []*UsersCreate
}

// Save creates the Users entities in the database.
func (ucb *UsersCreateBulk) Save(ctx context.Context) ([]*Users, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Users, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UsersCreateBulk) SaveX(ctx context.Context) []*Users {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UsersCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UsersCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
