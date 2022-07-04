// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"wynn-member-api/ent/predicate"
	"wynn-member-api/ent/sagameaccount"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SAGameAccountDelete is the builder for deleting a SAGameAccount entity.
type SAGameAccountDelete struct {
	config
	hooks    []Hook
	mutation *SAGameAccountMutation
}

// Where appends a list predicates to the SAGameAccountDelete builder.
func (sgad *SAGameAccountDelete) Where(ps ...predicate.SAGameAccount) *SAGameAccountDelete {
	sgad.mutation.Where(ps...)
	return sgad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sgad *SAGameAccountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sgad.hooks) == 0 {
		affected, err = sgad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SAGameAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sgad.mutation = mutation
			affected, err = sgad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sgad.hooks) - 1; i >= 0; i-- {
			if sgad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sgad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sgad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sgad *SAGameAccountDelete) ExecX(ctx context.Context) int {
	n, err := sgad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sgad *SAGameAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sagameaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sagameaccount.FieldID,
			},
		},
	}
	if ps := sgad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sgad.driver, _spec)
}

// SAGameAccountDeleteOne is the builder for deleting a single SAGameAccount entity.
type SAGameAccountDeleteOne struct {
	sgad *SAGameAccountDelete
}

// Exec executes the deletion query.
func (sgado *SAGameAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := sgado.sgad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sagameaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sgado *SAGameAccountDeleteOne) ExecX(ctx context.Context) {
	sgado.sgad.ExecX(ctx)
}
