// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/mnsv1511/kasikorn-line-assignment/ent"
)

// The AccountBalancesFunc type is an adapter to allow the use of ordinary
// function as AccountBalances mutator.
type AccountBalancesFunc func(context.Context, *ent.AccountBalancesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountBalancesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AccountBalancesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountBalancesMutation", m)
}

// The AccountDetailsFunc type is an adapter to allow the use of ordinary
// function as AccountDetails mutator.
type AccountDetailsFunc func(context.Context, *ent.AccountDetailsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountDetailsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AccountDetailsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountDetailsMutation", m)
}

// The AccountFlagsFunc type is an adapter to allow the use of ordinary
// function as AccountFlags mutator.
type AccountFlagsFunc func(context.Context, *ent.AccountFlagsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountFlagsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AccountFlagsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountFlagsMutation", m)
}

// The AccountsFunc type is an adapter to allow the use of ordinary
// function as Accounts mutator.
type AccountsFunc func(context.Context, *ent.AccountsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AccountsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountsMutation", m)
}

// The BannersFunc type is an adapter to allow the use of ordinary
// function as Banners mutator.
type BannersFunc func(context.Context, *ent.BannersMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BannersFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BannersMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BannersMutation", m)
}

// The DebitCardDesignFunc type is an adapter to allow the use of ordinary
// function as DebitCardDesign mutator.
type DebitCardDesignFunc func(context.Context, *ent.DebitCardDesignMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DebitCardDesignFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DebitCardDesignMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DebitCardDesignMutation", m)
}

// The DebitCardDetailsFunc type is an adapter to allow the use of ordinary
// function as DebitCardDetails mutator.
type DebitCardDetailsFunc func(context.Context, *ent.DebitCardDetailsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DebitCardDetailsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DebitCardDetailsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DebitCardDetailsMutation", m)
}

// The DebitCardStatusFunc type is an adapter to allow the use of ordinary
// function as DebitCardStatus mutator.
type DebitCardStatusFunc func(context.Context, *ent.DebitCardStatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DebitCardStatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DebitCardStatusMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DebitCardStatusMutation", m)
}

// The DebitCardsFunc type is an adapter to allow the use of ordinary
// function as DebitCards mutator.
type DebitCardsFunc func(context.Context, *ent.DebitCardsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DebitCardsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DebitCardsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DebitCardsMutation", m)
}

// The TransactionsFunc type is an adapter to allow the use of ordinary
// function as Transactions mutator.
type TransactionsFunc func(context.Context, *ent.TransactionsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TransactionsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TransactionsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TransactionsMutation", m)
}

// The UserGreetingsFunc type is an adapter to allow the use of ordinary
// function as UserGreetings mutator.
type UserGreetingsFunc func(context.Context, *ent.UserGreetingsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserGreetingsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserGreetingsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserGreetingsMutation", m)
}

// The UsersFunc type is an adapter to allow the use of ordinary
// function as Users mutator.
type UsersFunc func(context.Context, *ent.UsersMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UsersFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UsersMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UsersMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
