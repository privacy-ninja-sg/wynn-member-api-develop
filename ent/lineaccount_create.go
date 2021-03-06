// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wynn-member-api/ent/lineaccount"
	"wynn-member-api/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LineAccountCreate is the builder for creating a LineAccount entity.
type LineAccountCreate struct {
	config
	mutation *LineAccountMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (lac *LineAccountCreate) SetUUID(u uuid.UUID) *LineAccountCreate {
	lac.mutation.SetUUID(u)
	return lac
}

// SetLineID sets the "line_id" field.
func (lac *LineAccountCreate) SetLineID(s string) *LineAccountCreate {
	lac.mutation.SetLineID(s)
	return lac
}

// SetNillableLineID sets the "line_id" field if the given value is not nil.
func (lac *LineAccountCreate) SetNillableLineID(s *string) *LineAccountCreate {
	if s != nil {
		lac.SetLineID(*s)
	}
	return lac
}

// SetLineClientID sets the "line_client_id" field.
func (lac *LineAccountCreate) SetLineClientID(s string) *LineAccountCreate {
	lac.mutation.SetLineClientID(s)
	return lac
}

// SetNillableLineClientID sets the "line_client_id" field if the given value is not nil.
func (lac *LineAccountCreate) SetNillableLineClientID(s *string) *LineAccountCreate {
	if s != nil {
		lac.SetLineClientID(*s)
	}
	return lac
}

// SetCreatedAt sets the "created_at" field.
func (lac *LineAccountCreate) SetCreatedAt(t time.Time) *LineAccountCreate {
	lac.mutation.SetCreatedAt(t)
	return lac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lac *LineAccountCreate) SetNillableCreatedAt(t *time.Time) *LineAccountCreate {
	if t != nil {
		lac.SetCreatedAt(*t)
	}
	return lac
}

// SetUpdatedAt sets the "updated_at" field.
func (lac *LineAccountCreate) SetUpdatedAt(t time.Time) *LineAccountCreate {
	lac.mutation.SetUpdatedAt(t)
	return lac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lac *LineAccountCreate) SetNillableUpdatedAt(t *time.Time) *LineAccountCreate {
	if t != nil {
		lac.SetUpdatedAt(*t)
	}
	return lac
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (lac *LineAccountCreate) SetOwnerID(id int) *LineAccountCreate {
	lac.mutation.SetOwnerID(id)
	return lac
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (lac *LineAccountCreate) SetNillableOwnerID(id *int) *LineAccountCreate {
	if id != nil {
		lac = lac.SetOwnerID(*id)
	}
	return lac
}

// SetOwner sets the "owner" edge to the User entity.
func (lac *LineAccountCreate) SetOwner(u *User) *LineAccountCreate {
	return lac.SetOwnerID(u.ID)
}

// Mutation returns the LineAccountMutation object of the builder.
func (lac *LineAccountCreate) Mutation() *LineAccountMutation {
	return lac.mutation
}

// Save creates the LineAccount in the database.
func (lac *LineAccountCreate) Save(ctx context.Context) (*LineAccount, error) {
	var (
		err  error
		node *LineAccount
	)
	lac.defaults()
	if len(lac.hooks) == 0 {
		if err = lac.check(); err != nil {
			return nil, err
		}
		node, err = lac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LineAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lac.check(); err != nil {
				return nil, err
			}
			lac.mutation = mutation
			if node, err = lac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lac.hooks) - 1; i >= 0; i-- {
			if lac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lac *LineAccountCreate) SaveX(ctx context.Context) *LineAccount {
	v, err := lac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lac *LineAccountCreate) Exec(ctx context.Context) error {
	_, err := lac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lac *LineAccountCreate) ExecX(ctx context.Context) {
	if err := lac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lac *LineAccountCreate) defaults() {
	if _, ok := lac.mutation.UUID(); !ok {
		v := lineaccount.DefaultUUID()
		lac.mutation.SetUUID(v)
	}
	if _, ok := lac.mutation.LineID(); !ok {
		v := lineaccount.DefaultLineID
		lac.mutation.SetLineID(v)
	}
	if _, ok := lac.mutation.CreatedAt(); !ok {
		v := lineaccount.DefaultCreatedAt()
		lac.mutation.SetCreatedAt(v)
	}
	if _, ok := lac.mutation.UpdatedAt(); !ok {
		v := lineaccount.DefaultUpdatedAt()
		lac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lac *LineAccountCreate) check() error {
	if _, ok := lac.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "uuid"`)}
	}
	if _, ok := lac.mutation.LineID(); !ok {
		return &ValidationError{Name: "line_id", err: errors.New(`ent: missing required field "line_id"`)}
	}
	if _, ok := lac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := lac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (lac *LineAccountCreate) sqlSave(ctx context.Context) (*LineAccount, error) {
	_node, _spec := lac.createSpec()
	if err := sqlgraph.CreateNode(ctx, lac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lac *LineAccountCreate) createSpec() (*LineAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &LineAccount{config: lac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lineaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lineaccount.FieldID,
			},
		}
	)
	if value, ok := lac.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: lineaccount.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := lac.mutation.LineID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lineaccount.FieldLineID,
		})
		_node.LineID = value
	}
	if value, ok := lac.mutation.LineClientID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lineaccount.FieldLineClientID,
		})
		_node.LineClientID = value
	}
	if value, ok := lac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lineaccount.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := lac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lineaccount.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := lac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lineaccount.OwnerTable,
			Columns: []string{lineaccount.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_line = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LineAccountCreateBulk is the builder for creating many LineAccount entities in bulk.
type LineAccountCreateBulk struct {
	config
	builders []*LineAccountCreate
}

// Save creates the LineAccount entities in the database.
func (lacb *LineAccountCreateBulk) Save(ctx context.Context) ([]*LineAccount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lacb.builders))
	nodes := make([]*LineAccount, len(lacb.builders))
	mutators := make([]Mutator, len(lacb.builders))
	for i := range lacb.builders {
		func(i int, root context.Context) {
			builder := lacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LineAccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lacb *LineAccountCreateBulk) SaveX(ctx context.Context) []*LineAccount {
	v, err := lacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lacb *LineAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := lacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lacb *LineAccountCreateBulk) ExecX(ctx context.Context) {
	if err := lacb.Exec(ctx); err != nil {
		panic(err)
	}
}
