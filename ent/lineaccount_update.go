// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"wynn-member-api/ent/lineaccount"
	"wynn-member-api/ent/predicate"
	"wynn-member-api/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LineAccountUpdate is the builder for updating LineAccount entities.
type LineAccountUpdate struct {
	config
	hooks    []Hook
	mutation *LineAccountMutation
}

// Where appends a list predicates to the LineAccountUpdate builder.
func (lau *LineAccountUpdate) Where(ps ...predicate.LineAccount) *LineAccountUpdate {
	lau.mutation.Where(ps...)
	return lau
}

// SetUUID sets the "uuid" field.
func (lau *LineAccountUpdate) SetUUID(u uuid.UUID) *LineAccountUpdate {
	lau.mutation.SetUUID(u)
	return lau
}

// SetLineID sets the "line_id" field.
func (lau *LineAccountUpdate) SetLineID(s string) *LineAccountUpdate {
	lau.mutation.SetLineID(s)
	return lau
}

// SetNillableLineID sets the "line_id" field if the given value is not nil.
func (lau *LineAccountUpdate) SetNillableLineID(s *string) *LineAccountUpdate {
	if s != nil {
		lau.SetLineID(*s)
	}
	return lau
}

// SetLineClientID sets the "line_client_id" field.
func (lau *LineAccountUpdate) SetLineClientID(s string) *LineAccountUpdate {
	lau.mutation.SetLineClientID(s)
	return lau
}

// SetNillableLineClientID sets the "line_client_id" field if the given value is not nil.
func (lau *LineAccountUpdate) SetNillableLineClientID(s *string) *LineAccountUpdate {
	if s != nil {
		lau.SetLineClientID(*s)
	}
	return lau
}

// ClearLineClientID clears the value of the "line_client_id" field.
func (lau *LineAccountUpdate) ClearLineClientID() *LineAccountUpdate {
	lau.mutation.ClearLineClientID()
	return lau
}

// SetCreatedAt sets the "created_at" field.
func (lau *LineAccountUpdate) SetCreatedAt(t time.Time) *LineAccountUpdate {
	lau.mutation.SetCreatedAt(t)
	return lau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lau *LineAccountUpdate) SetNillableCreatedAt(t *time.Time) *LineAccountUpdate {
	if t != nil {
		lau.SetCreatedAt(*t)
	}
	return lau
}

// SetUpdatedAt sets the "updated_at" field.
func (lau *LineAccountUpdate) SetUpdatedAt(t time.Time) *LineAccountUpdate {
	lau.mutation.SetUpdatedAt(t)
	return lau
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lau *LineAccountUpdate) SetNillableUpdatedAt(t *time.Time) *LineAccountUpdate {
	if t != nil {
		lau.SetUpdatedAt(*t)
	}
	return lau
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (lau *LineAccountUpdate) SetOwnerID(id int) *LineAccountUpdate {
	lau.mutation.SetOwnerID(id)
	return lau
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (lau *LineAccountUpdate) SetNillableOwnerID(id *int) *LineAccountUpdate {
	if id != nil {
		lau = lau.SetOwnerID(*id)
	}
	return lau
}

// SetOwner sets the "owner" edge to the User entity.
func (lau *LineAccountUpdate) SetOwner(u *User) *LineAccountUpdate {
	return lau.SetOwnerID(u.ID)
}

// Mutation returns the LineAccountMutation object of the builder.
func (lau *LineAccountUpdate) Mutation() *LineAccountMutation {
	return lau.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (lau *LineAccountUpdate) ClearOwner() *LineAccountUpdate {
	lau.mutation.ClearOwner()
	return lau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lau *LineAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lau.hooks) == 0 {
		affected, err = lau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LineAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lau.mutation = mutation
			affected, err = lau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lau.hooks) - 1; i >= 0; i-- {
			if lau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lau *LineAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := lau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lau *LineAccountUpdate) Exec(ctx context.Context) error {
	_, err := lau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lau *LineAccountUpdate) ExecX(ctx context.Context) {
	if err := lau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lau *LineAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lineaccount.Table,
			Columns: lineaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lineaccount.FieldID,
			},
		},
	}
	if ps := lau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lau.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: lineaccount.FieldUUID,
		})
	}
	if value, ok := lau.mutation.LineID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lineaccount.FieldLineID,
		})
	}
	if value, ok := lau.mutation.LineClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lineaccount.FieldLineClientID,
		})
	}
	if lau.mutation.LineClientIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: lineaccount.FieldLineClientID,
		})
	}
	if value, ok := lau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lineaccount.FieldCreatedAt,
		})
	}
	if value, ok := lau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lineaccount.FieldUpdatedAt,
		})
	}
	if lau.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lau.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lineaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LineAccountUpdateOne is the builder for updating a single LineAccount entity.
type LineAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LineAccountMutation
}

// SetUUID sets the "uuid" field.
func (lauo *LineAccountUpdateOne) SetUUID(u uuid.UUID) *LineAccountUpdateOne {
	lauo.mutation.SetUUID(u)
	return lauo
}

// SetLineID sets the "line_id" field.
func (lauo *LineAccountUpdateOne) SetLineID(s string) *LineAccountUpdateOne {
	lauo.mutation.SetLineID(s)
	return lauo
}

// SetNillableLineID sets the "line_id" field if the given value is not nil.
func (lauo *LineAccountUpdateOne) SetNillableLineID(s *string) *LineAccountUpdateOne {
	if s != nil {
		lauo.SetLineID(*s)
	}
	return lauo
}

// SetLineClientID sets the "line_client_id" field.
func (lauo *LineAccountUpdateOne) SetLineClientID(s string) *LineAccountUpdateOne {
	lauo.mutation.SetLineClientID(s)
	return lauo
}

// SetNillableLineClientID sets the "line_client_id" field if the given value is not nil.
func (lauo *LineAccountUpdateOne) SetNillableLineClientID(s *string) *LineAccountUpdateOne {
	if s != nil {
		lauo.SetLineClientID(*s)
	}
	return lauo
}

// ClearLineClientID clears the value of the "line_client_id" field.
func (lauo *LineAccountUpdateOne) ClearLineClientID() *LineAccountUpdateOne {
	lauo.mutation.ClearLineClientID()
	return lauo
}

// SetCreatedAt sets the "created_at" field.
func (lauo *LineAccountUpdateOne) SetCreatedAt(t time.Time) *LineAccountUpdateOne {
	lauo.mutation.SetCreatedAt(t)
	return lauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lauo *LineAccountUpdateOne) SetNillableCreatedAt(t *time.Time) *LineAccountUpdateOne {
	if t != nil {
		lauo.SetCreatedAt(*t)
	}
	return lauo
}

// SetUpdatedAt sets the "updated_at" field.
func (lauo *LineAccountUpdateOne) SetUpdatedAt(t time.Time) *LineAccountUpdateOne {
	lauo.mutation.SetUpdatedAt(t)
	return lauo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lauo *LineAccountUpdateOne) SetNillableUpdatedAt(t *time.Time) *LineAccountUpdateOne {
	if t != nil {
		lauo.SetUpdatedAt(*t)
	}
	return lauo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (lauo *LineAccountUpdateOne) SetOwnerID(id int) *LineAccountUpdateOne {
	lauo.mutation.SetOwnerID(id)
	return lauo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (lauo *LineAccountUpdateOne) SetNillableOwnerID(id *int) *LineAccountUpdateOne {
	if id != nil {
		lauo = lauo.SetOwnerID(*id)
	}
	return lauo
}

// SetOwner sets the "owner" edge to the User entity.
func (lauo *LineAccountUpdateOne) SetOwner(u *User) *LineAccountUpdateOne {
	return lauo.SetOwnerID(u.ID)
}

// Mutation returns the LineAccountMutation object of the builder.
func (lauo *LineAccountUpdateOne) Mutation() *LineAccountMutation {
	return lauo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (lauo *LineAccountUpdateOne) ClearOwner() *LineAccountUpdateOne {
	lauo.mutation.ClearOwner()
	return lauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lauo *LineAccountUpdateOne) Select(field string, fields ...string) *LineAccountUpdateOne {
	lauo.fields = append([]string{field}, fields...)
	return lauo
}

// Save executes the query and returns the updated LineAccount entity.
func (lauo *LineAccountUpdateOne) Save(ctx context.Context) (*LineAccount, error) {
	var (
		err  error
		node *LineAccount
	)
	if len(lauo.hooks) == 0 {
		node, err = lauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LineAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lauo.mutation = mutation
			node, err = lauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lauo.hooks) - 1; i >= 0; i-- {
			if lauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lauo *LineAccountUpdateOne) SaveX(ctx context.Context) *LineAccount {
	node, err := lauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lauo *LineAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := lauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lauo *LineAccountUpdateOne) ExecX(ctx context.Context) {
	if err := lauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lauo *LineAccountUpdateOne) sqlSave(ctx context.Context) (_node *LineAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lineaccount.Table,
			Columns: lineaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lineaccount.FieldID,
			},
		},
	}
	id, ok := lauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LineAccount.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := lauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lineaccount.FieldID)
		for _, f := range fields {
			if !lineaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lineaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lauo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: lineaccount.FieldUUID,
		})
	}
	if value, ok := lauo.mutation.LineID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lineaccount.FieldLineID,
		})
	}
	if value, ok := lauo.mutation.LineClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lineaccount.FieldLineClientID,
		})
	}
	if lauo.mutation.LineClientIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: lineaccount.FieldLineClientID,
		})
	}
	if value, ok := lauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lineaccount.FieldCreatedAt,
		})
	}
	if value, ok := lauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lineaccount.FieldUpdatedAt,
		})
	}
	if lauo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lauo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LineAccount{config: lauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lineaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
