// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"wynn-member-api/ent/game"
	"wynn-member-api/ent/gameaccount"
	"wynn-member-api/ent/predicate"
	"wynn-member-api/ent/transfertransaction"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GameUpdate is the builder for updating Game entities.
type GameUpdate struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameUpdate builder.
func (gu *GameUpdate) Where(ps ...predicate.Game) *GameUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUUID sets the "uuid" field.
func (gu *GameUpdate) SetUUID(u uuid.UUID) *GameUpdate {
	gu.mutation.SetUUID(u)
	return gu
}

// SetName sets the "name" field.
func (gu *GameUpdate) SetName(s string) *GameUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetBanner sets the "banner" field.
func (gu *GameUpdate) SetBanner(s string) *GameUpdate {
	gu.mutation.SetBanner(s)
	return gu
}

// SetNillableBanner sets the "banner" field if the given value is not nil.
func (gu *GameUpdate) SetNillableBanner(s *string) *GameUpdate {
	if s != nil {
		gu.SetBanner(*s)
	}
	return gu
}

// ClearBanner clears the value of the "banner" field.
func (gu *GameUpdate) ClearBanner() *GameUpdate {
	gu.mutation.ClearBanner()
	return gu
}

// SetDesc sets the "desc" field.
func (gu *GameUpdate) SetDesc(s string) *GameUpdate {
	gu.mutation.SetDesc(s)
	return gu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (gu *GameUpdate) SetNillableDesc(s *string) *GameUpdate {
	if s != nil {
		gu.SetDesc(*s)
	}
	return gu
}

// ClearDesc clears the value of the "desc" field.
func (gu *GameUpdate) ClearDesc() *GameUpdate {
	gu.mutation.ClearDesc()
	return gu
}

// SetStatus sets the "status" field.
func (gu *GameUpdate) SetStatus(ga game.Status) *GameUpdate {
	gu.mutation.SetStatus(ga)
	return gu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (gu *GameUpdate) SetNillableStatus(ga *game.Status) *GameUpdate {
	if ga != nil {
		gu.SetStatus(*ga)
	}
	return gu
}

// SetCreatedAt sets the "created_at" field.
func (gu *GameUpdate) SetCreatedAt(t time.Time) *GameUpdate {
	gu.mutation.SetCreatedAt(t)
	return gu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gu *GameUpdate) SetNillableCreatedAt(t *time.Time) *GameUpdate {
	if t != nil {
		gu.SetCreatedAt(*t)
	}
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GameUpdate) SetUpdatedAt(t time.Time) *GameUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gu *GameUpdate) SetNillableUpdatedAt(t *time.Time) *GameUpdate {
	if t != nil {
		gu.SetUpdatedAt(*t)
	}
	return gu
}

// AddAccountIDs adds the "accounts" edge to the GameAccount entity by IDs.
func (gu *GameUpdate) AddAccountIDs(ids ...int) *GameUpdate {
	gu.mutation.AddAccountIDs(ids...)
	return gu
}

// AddAccounts adds the "accounts" edges to the GameAccount entity.
func (gu *GameUpdate) AddAccounts(g ...*GameAccount) *GameUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddAccountIDs(ids...)
}

// AddTransferIDs adds the "transfers" edge to the TransferTransaction entity by IDs.
func (gu *GameUpdate) AddTransferIDs(ids ...int) *GameUpdate {
	gu.mutation.AddTransferIDs(ids...)
	return gu
}

// AddTransfers adds the "transfers" edges to the TransferTransaction entity.
func (gu *GameUpdate) AddTransfers(t ...*TransferTransaction) *GameUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return gu.AddTransferIDs(ids...)
}

// Mutation returns the GameMutation object of the builder.
func (gu *GameUpdate) Mutation() *GameMutation {
	return gu.mutation
}

// ClearAccounts clears all "accounts" edges to the GameAccount entity.
func (gu *GameUpdate) ClearAccounts() *GameUpdate {
	gu.mutation.ClearAccounts()
	return gu
}

// RemoveAccountIDs removes the "accounts" edge to GameAccount entities by IDs.
func (gu *GameUpdate) RemoveAccountIDs(ids ...int) *GameUpdate {
	gu.mutation.RemoveAccountIDs(ids...)
	return gu
}

// RemoveAccounts removes "accounts" edges to GameAccount entities.
func (gu *GameUpdate) RemoveAccounts(g ...*GameAccount) *GameUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveAccountIDs(ids...)
}

// ClearTransfers clears all "transfers" edges to the TransferTransaction entity.
func (gu *GameUpdate) ClearTransfers() *GameUpdate {
	gu.mutation.ClearTransfers()
	return gu
}

// RemoveTransferIDs removes the "transfers" edge to TransferTransaction entities by IDs.
func (gu *GameUpdate) RemoveTransferIDs(ids ...int) *GameUpdate {
	gu.mutation.RemoveTransferIDs(ids...)
	return gu
}

// RemoveTransfers removes "transfers" edges to TransferTransaction entities.
func (gu *GameUpdate) RemoveTransfers(t ...*TransferTransaction) *GameUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return gu.RemoveTransferIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GameUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GameUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GameUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GameUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GameUpdate) check() error {
	if v, ok := gu.mutation.Status(); ok {
		if err := game.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (gu *GameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: game.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: game.FieldUUID,
		})
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldName,
		})
	}
	if value, ok := gu.mutation.Banner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldBanner,
		})
	}
	if gu.mutation.BannerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: game.FieldBanner,
		})
	}
	if value, ok := gu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldDesc,
		})
	}
	if gu.mutation.DescCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: game.FieldDesc,
		})
	}
	if value, ok := gu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: game.FieldStatus,
		})
	}
	if value, ok := gu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: game.FieldCreatedAt,
		})
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: game.FieldUpdatedAt,
		})
	}
	if gu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.AccountsTable,
			Columns: []string{game.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !gu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.AccountsTable,
			Columns: []string{game.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.AccountsTable,
			Columns: []string{game.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.TransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.TransfersTable,
			Columns: []string{game.TransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfertransaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedTransfersIDs(); len(nodes) > 0 && !gu.mutation.TransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.TransfersTable,
			Columns: []string{game.TransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfertransaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.TransfersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.TransfersTable,
			Columns: []string{game.TransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfertransaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GameUpdateOne is the builder for updating a single Game entity.
type GameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameMutation
}

// SetUUID sets the "uuid" field.
func (guo *GameUpdateOne) SetUUID(u uuid.UUID) *GameUpdateOne {
	guo.mutation.SetUUID(u)
	return guo
}

// SetName sets the "name" field.
func (guo *GameUpdateOne) SetName(s string) *GameUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetBanner sets the "banner" field.
func (guo *GameUpdateOne) SetBanner(s string) *GameUpdateOne {
	guo.mutation.SetBanner(s)
	return guo
}

// SetNillableBanner sets the "banner" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableBanner(s *string) *GameUpdateOne {
	if s != nil {
		guo.SetBanner(*s)
	}
	return guo
}

// ClearBanner clears the value of the "banner" field.
func (guo *GameUpdateOne) ClearBanner() *GameUpdateOne {
	guo.mutation.ClearBanner()
	return guo
}

// SetDesc sets the "desc" field.
func (guo *GameUpdateOne) SetDesc(s string) *GameUpdateOne {
	guo.mutation.SetDesc(s)
	return guo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableDesc(s *string) *GameUpdateOne {
	if s != nil {
		guo.SetDesc(*s)
	}
	return guo
}

// ClearDesc clears the value of the "desc" field.
func (guo *GameUpdateOne) ClearDesc() *GameUpdateOne {
	guo.mutation.ClearDesc()
	return guo
}

// SetStatus sets the "status" field.
func (guo *GameUpdateOne) SetStatus(ga game.Status) *GameUpdateOne {
	guo.mutation.SetStatus(ga)
	return guo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableStatus(ga *game.Status) *GameUpdateOne {
	if ga != nil {
		guo.SetStatus(*ga)
	}
	return guo
}

// SetCreatedAt sets the "created_at" field.
func (guo *GameUpdateOne) SetCreatedAt(t time.Time) *GameUpdateOne {
	guo.mutation.SetCreatedAt(t)
	return guo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableCreatedAt(t *time.Time) *GameUpdateOne {
	if t != nil {
		guo.SetCreatedAt(*t)
	}
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GameUpdateOne) SetUpdatedAt(t time.Time) *GameUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableUpdatedAt(t *time.Time) *GameUpdateOne {
	if t != nil {
		guo.SetUpdatedAt(*t)
	}
	return guo
}

// AddAccountIDs adds the "accounts" edge to the GameAccount entity by IDs.
func (guo *GameUpdateOne) AddAccountIDs(ids ...int) *GameUpdateOne {
	guo.mutation.AddAccountIDs(ids...)
	return guo
}

// AddAccounts adds the "accounts" edges to the GameAccount entity.
func (guo *GameUpdateOne) AddAccounts(g ...*GameAccount) *GameUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddAccountIDs(ids...)
}

// AddTransferIDs adds the "transfers" edge to the TransferTransaction entity by IDs.
func (guo *GameUpdateOne) AddTransferIDs(ids ...int) *GameUpdateOne {
	guo.mutation.AddTransferIDs(ids...)
	return guo
}

// AddTransfers adds the "transfers" edges to the TransferTransaction entity.
func (guo *GameUpdateOne) AddTransfers(t ...*TransferTransaction) *GameUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return guo.AddTransferIDs(ids...)
}

// Mutation returns the GameMutation object of the builder.
func (guo *GameUpdateOne) Mutation() *GameMutation {
	return guo.mutation
}

// ClearAccounts clears all "accounts" edges to the GameAccount entity.
func (guo *GameUpdateOne) ClearAccounts() *GameUpdateOne {
	guo.mutation.ClearAccounts()
	return guo
}

// RemoveAccountIDs removes the "accounts" edge to GameAccount entities by IDs.
func (guo *GameUpdateOne) RemoveAccountIDs(ids ...int) *GameUpdateOne {
	guo.mutation.RemoveAccountIDs(ids...)
	return guo
}

// RemoveAccounts removes "accounts" edges to GameAccount entities.
func (guo *GameUpdateOne) RemoveAccounts(g ...*GameAccount) *GameUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveAccountIDs(ids...)
}

// ClearTransfers clears all "transfers" edges to the TransferTransaction entity.
func (guo *GameUpdateOne) ClearTransfers() *GameUpdateOne {
	guo.mutation.ClearTransfers()
	return guo
}

// RemoveTransferIDs removes the "transfers" edge to TransferTransaction entities by IDs.
func (guo *GameUpdateOne) RemoveTransferIDs(ids ...int) *GameUpdateOne {
	guo.mutation.RemoveTransferIDs(ids...)
	return guo
}

// RemoveTransfers removes "transfers" edges to TransferTransaction entities.
func (guo *GameUpdateOne) RemoveTransfers(t ...*TransferTransaction) *GameUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return guo.RemoveTransferIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GameUpdateOne) Select(field string, fields ...string) *GameUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Game entity.
func (guo *GameUpdateOne) Save(ctx context.Context) (*Game, error) {
	var (
		err  error
		node *Game
	)
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GameUpdateOne) SaveX(ctx context.Context) *Game {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GameUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GameUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GameUpdateOne) check() error {
	if v, ok := guo.mutation.Status(); ok {
		if err := game.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (guo *GameUpdateOne) sqlSave(ctx context.Context) (_node *Game, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: game.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Game.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for _, f := range fields {
			if !game.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: game.FieldUUID,
		})
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldName,
		})
	}
	if value, ok := guo.mutation.Banner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldBanner,
		})
	}
	if guo.mutation.BannerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: game.FieldBanner,
		})
	}
	if value, ok := guo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldDesc,
		})
	}
	if guo.mutation.DescCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: game.FieldDesc,
		})
	}
	if value, ok := guo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: game.FieldStatus,
		})
	}
	if value, ok := guo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: game.FieldCreatedAt,
		})
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: game.FieldUpdatedAt,
		})
	}
	if guo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.AccountsTable,
			Columns: []string{game.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !guo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.AccountsTable,
			Columns: []string{game.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.AccountsTable,
			Columns: []string{game.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.TransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.TransfersTable,
			Columns: []string{game.TransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfertransaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedTransfersIDs(); len(nodes) > 0 && !guo.mutation.TransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.TransfersTable,
			Columns: []string{game.TransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfertransaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.TransfersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.TransfersTable,
			Columns: []string{game.TransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfertransaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Game{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
