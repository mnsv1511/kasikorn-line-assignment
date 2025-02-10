// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountflags"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
)

// AccountFlagsDelete is the builder for deleting a AccountFlags entity.
type AccountFlagsDelete struct {
	config
	hooks    []Hook
	mutation *AccountFlagsMutation
}

// Where appends a list predicates to the AccountFlagsDelete builder.
func (afd *AccountFlagsDelete) Where(ps ...predicate.AccountFlags) *AccountFlagsDelete {
	afd.mutation.Where(ps...)
	return afd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (afd *AccountFlagsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, afd.sqlExec, afd.mutation, afd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (afd *AccountFlagsDelete) ExecX(ctx context.Context) int {
	n, err := afd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (afd *AccountFlagsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(accountflags.Table, sqlgraph.NewFieldSpec(accountflags.FieldID, field.TypeInt))
	if ps := afd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, afd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	afd.mutation.done = true
	return affected, err
}

// AccountFlagsDeleteOne is the builder for deleting a single AccountFlags entity.
type AccountFlagsDeleteOne struct {
	afd *AccountFlagsDelete
}

// Where appends a list predicates to the AccountFlagsDelete builder.
func (afdo *AccountFlagsDeleteOne) Where(ps ...predicate.AccountFlags) *AccountFlagsDeleteOne {
	afdo.afd.mutation.Where(ps...)
	return afdo
}

// Exec executes the deletion query.
func (afdo *AccountFlagsDeleteOne) Exec(ctx context.Context) error {
	n, err := afdo.afd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountflags.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (afdo *AccountFlagsDeleteOne) ExecX(ctx context.Context) {
	if err := afdo.Exec(ctx); err != nil {
		panic(err)
	}
}
