// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcardstatus"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
)

// DebitCardStatusDelete is the builder for deleting a DebitCardStatus entity.
type DebitCardStatusDelete struct {
	config
	hooks    []Hook
	mutation *DebitCardStatusMutation
}

// Where appends a list predicates to the DebitCardStatusDelete builder.
func (dcsd *DebitCardStatusDelete) Where(ps ...predicate.DebitCardStatus) *DebitCardStatusDelete {
	dcsd.mutation.Where(ps...)
	return dcsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dcsd *DebitCardStatusDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dcsd.sqlExec, dcsd.mutation, dcsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dcsd *DebitCardStatusDelete) ExecX(ctx context.Context) int {
	n, err := dcsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dcsd *DebitCardStatusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(debitcardstatus.Table, sqlgraph.NewFieldSpec(debitcardstatus.FieldID, field.TypeInt))
	if ps := dcsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dcsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dcsd.mutation.done = true
	return affected, err
}

// DebitCardStatusDeleteOne is the builder for deleting a single DebitCardStatus entity.
type DebitCardStatusDeleteOne struct {
	dcsd *DebitCardStatusDelete
}

// Where appends a list predicates to the DebitCardStatusDelete builder.
func (dcsdo *DebitCardStatusDeleteOne) Where(ps ...predicate.DebitCardStatus) *DebitCardStatusDeleteOne {
	dcsdo.dcsd.mutation.Where(ps...)
	return dcsdo
}

// Exec executes the deletion query.
func (dcsdo *DebitCardStatusDeleteOne) Exec(ctx context.Context) error {
	n, err := dcsdo.dcsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{debitcardstatus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dcsdo *DebitCardStatusDeleteOne) ExecX(ctx context.Context) {
	if err := dcsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
