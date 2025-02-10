// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountbalances"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
)

// AccountBalancesDelete is the builder for deleting a AccountBalances entity.
type AccountBalancesDelete struct {
	config
	hooks    []Hook
	mutation *AccountBalancesMutation
}

// Where appends a list predicates to the AccountBalancesDelete builder.
func (abd *AccountBalancesDelete) Where(ps ...predicate.AccountBalances) *AccountBalancesDelete {
	abd.mutation.Where(ps...)
	return abd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (abd *AccountBalancesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, abd.sqlExec, abd.mutation, abd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (abd *AccountBalancesDelete) ExecX(ctx context.Context) int {
	n, err := abd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (abd *AccountBalancesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(accountbalances.Table, sqlgraph.NewFieldSpec(accountbalances.FieldID, field.TypeInt))
	if ps := abd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, abd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	abd.mutation.done = true
	return affected, err
}

// AccountBalancesDeleteOne is the builder for deleting a single AccountBalances entity.
type AccountBalancesDeleteOne struct {
	abd *AccountBalancesDelete
}

// Where appends a list predicates to the AccountBalancesDelete builder.
func (abdo *AccountBalancesDeleteOne) Where(ps ...predicate.AccountBalances) *AccountBalancesDeleteOne {
	abdo.abd.mutation.Where(ps...)
	return abdo
}

// Exec executes the deletion query.
func (abdo *AccountBalancesDeleteOne) Exec(ctx context.Context) error {
	n, err := abdo.abd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountbalances.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (abdo *AccountBalancesDeleteOne) ExecX(ctx context.Context) {
	if err := abdo.Exec(ctx); err != nil {
		panic(err)
	}
}
