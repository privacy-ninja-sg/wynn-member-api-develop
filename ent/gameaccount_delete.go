// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"wynn-member-api/ent/gameaccount"
	"wynn-member-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GameAccountDelete is the builder for deleting a GameAccount entity.
type GameAccountDelete struct {
	config
	hooks    []Hook
	mutation *GameAccountMutation
}

// Where appends a list predicates to the GameAccountDelete builder.
func (gad *GameAccountDelete) Where(ps ...predicate.GameAccount) *GameAccountDelete {
	gad.mutation.Where(ps...)
	return gad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gad *GameAccountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gad.hooks) == 0 {
		affected, err = gad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gad.mutation = mutation
			affected, err = gad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gad.hooks) - 1; i >= 0; i-- {
			if gad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gad *GameAccountDelete) ExecX(ctx context.Context) int {
	n, err := gad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gad *GameAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: gameaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gameaccount.FieldID,
			},
		},
	}
	if ps := gad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gad.driver, _spec)
}

// GameAccountDeleteOne is the builder for deleting a single GameAccount entity.
type GameAccountDeleteOne struct {
	gad *GameAccountDelete
}

// Exec executes the deletion query.
func (gado *GameAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := gado.gad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gameaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gado *GameAccountDeleteOne) ExecX(ctx context.Context) {
	gado.gad.ExecX(ctx)
}
