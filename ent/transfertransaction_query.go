// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"wynn-member-api/ent/game"
	"wynn-member-api/ent/predicate"
	"wynn-member-api/ent/transfertransaction"
	"wynn-member-api/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransferTransactionQuery is the builder for querying TransferTransaction entities.
type TransferTransactionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TransferTransaction
	// eager-loading edges.
	withOwner *UserQuery
	withGame  *GameQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransferTransactionQuery builder.
func (ttq *TransferTransactionQuery) Where(ps ...predicate.TransferTransaction) *TransferTransactionQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit adds a limit step to the query.
func (ttq *TransferTransactionQuery) Limit(limit int) *TransferTransactionQuery {
	ttq.limit = &limit
	return ttq
}

// Offset adds an offset step to the query.
func (ttq *TransferTransactionQuery) Offset(offset int) *TransferTransactionQuery {
	ttq.offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TransferTransactionQuery) Unique(unique bool) *TransferTransactionQuery {
	ttq.unique = &unique
	return ttq
}

// Order adds an order step to the query.
func (ttq *TransferTransactionQuery) Order(o ...OrderFunc) *TransferTransactionQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryOwner chains the current query on the "owner" edge.
func (ttq *TransferTransactionQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transfertransaction.Table, transfertransaction.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transfertransaction.OwnerTable, transfertransaction.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGame chains the current query on the "game" edge.
func (ttq *TransferTransactionQuery) QueryGame() *GameQuery {
	query := &GameQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transfertransaction.Table, transfertransaction.FieldID, selector),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transfertransaction.GameTable, transfertransaction.GameColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TransferTransaction entity from the query.
// Returns a *NotFoundError when no TransferTransaction was found.
func (ttq *TransferTransactionQuery) First(ctx context.Context) (*TransferTransaction, error) {
	nodes, err := ttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transfertransaction.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TransferTransactionQuery) FirstX(ctx context.Context) *TransferTransaction {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TransferTransaction ID from the query.
// Returns a *NotFoundError when no TransferTransaction ID was found.
func (ttq *TransferTransactionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transfertransaction.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TransferTransactionQuery) FirstIDX(ctx context.Context) int {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TransferTransaction entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TransferTransaction entity is not found.
// Returns a *NotFoundError when no TransferTransaction entities are found.
func (ttq *TransferTransactionQuery) Only(ctx context.Context) (*TransferTransaction, error) {
	nodes, err := ttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transfertransaction.Label}
	default:
		return nil, &NotSingularError{transfertransaction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TransferTransactionQuery) OnlyX(ctx context.Context) *TransferTransaction {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TransferTransaction ID in the query.
// Returns a *NotSingularError when exactly one TransferTransaction ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TransferTransactionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = &NotSingularError{transfertransaction.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TransferTransactionQuery) OnlyIDX(ctx context.Context) int {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TransferTransactions.
func (ttq *TransferTransactionQuery) All(ctx context.Context) ([]*TransferTransaction, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TransferTransactionQuery) AllX(ctx context.Context) []*TransferTransaction {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TransferTransaction IDs.
func (ttq *TransferTransactionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ttq.Select(transfertransaction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TransferTransactionQuery) IDsX(ctx context.Context) []int {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TransferTransactionQuery) Count(ctx context.Context) (int, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TransferTransactionQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TransferTransactionQuery) Exist(ctx context.Context) (bool, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TransferTransactionQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransferTransactionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TransferTransactionQuery) Clone() *TransferTransactionQuery {
	if ttq == nil {
		return nil
	}
	return &TransferTransactionQuery{
		config:     ttq.config,
		limit:      ttq.limit,
		offset:     ttq.offset,
		order:      append([]OrderFunc{}, ttq.order...),
		predicates: append([]predicate.TransferTransaction{}, ttq.predicates...),
		withOwner:  ttq.withOwner.Clone(),
		withGame:   ttq.withGame.Clone(),
		// clone intermediate query.
		sql:  ttq.sql.Clone(),
		path: ttq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TransferTransactionQuery) WithOwner(opts ...func(*UserQuery)) *TransferTransactionQuery {
	query := &UserQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withOwner = query
	return ttq
}

// WithGame tells the query-builder to eager-load the nodes that are connected to
// the "game" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TransferTransactionQuery) WithGame(opts ...func(*GameQuery)) *TransferTransactionQuery {
	query := &GameQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withGame = query
	return ttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TransferTransaction.Query().
//		GroupBy(transfertransaction.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ttq *TransferTransactionQuery) GroupBy(field string, fields ...string) *TransferTransactionGroupBy {
	group := &TransferTransactionGroupBy{config: ttq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ttq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//	}
//
//	client.TransferTransaction.Query().
//		Select(transfertransaction.FieldUUID).
//		Scan(ctx, &v)
//
func (ttq *TransferTransactionQuery) Select(fields ...string) *TransferTransactionSelect {
	ttq.fields = append(ttq.fields, fields...)
	return &TransferTransactionSelect{TransferTransactionQuery: ttq}
}

func (ttq *TransferTransactionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ttq.fields {
		if !transfertransaction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TransferTransactionQuery) sqlAll(ctx context.Context) ([]*TransferTransaction, error) {
	var (
		nodes       = []*TransferTransaction{}
		withFKs     = ttq.withFKs
		_spec       = ttq.querySpec()
		loadedTypes = [2]bool{
			ttq.withOwner != nil,
			ttq.withGame != nil,
		}
	)
	if ttq.withOwner != nil || ttq.withGame != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, transfertransaction.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TransferTransaction{config: ttq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ttq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TransferTransaction)
		for i := range nodes {
			if nodes[i].user_transfers == nil {
				continue
			}
			fk := *nodes[i].user_transfers
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_transfers" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	if query := ttq.withGame; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TransferTransaction)
		for i := range nodes {
			if nodes[i].game_transfers == nil {
				continue
			}
			fk := *nodes[i].game_transfers
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(game.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "game_transfers" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Game = n
			}
		}
	}

	return nodes, nil
}

func (ttq *TransferTransactionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TransferTransactionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ttq *TransferTransactionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transfertransaction.Table,
			Columns: transfertransaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transfertransaction.FieldID,
			},
		},
		From:   ttq.sql,
		Unique: true,
	}
	if unique := ttq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transfertransaction.FieldID)
		for i := range fields {
			if fields[i] != transfertransaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ttq *TransferTransactionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(transfertransaction.Table)
	columns := ttq.fields
	if len(columns) == 0 {
		columns = transfertransaction.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector)
	}
	if offset := ttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TransferTransactionGroupBy is the group-by builder for TransferTransaction entities.
type TransferTransactionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TransferTransactionGroupBy) Aggregate(fns ...AggregateFunc) *TransferTransactionGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ttgb *TransferTransactionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ttgb.path(ctx)
	if err != nil {
		return err
	}
	ttgb.sql = query
	return ttgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ttgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TransferTransactionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TransferTransactionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) StringsX(ctx context.Context) []string {
	v, err := ttgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TransferTransactionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ttgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = fmt.Errorf("ent: TransferTransactionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) StringX(ctx context.Context) string {
	v, err := ttgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TransferTransactionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TransferTransactionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) IntsX(ctx context.Context) []int {
	v, err := ttgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TransferTransactionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ttgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = fmt.Errorf("ent: TransferTransactionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) IntX(ctx context.Context) int {
	v, err := ttgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TransferTransactionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TransferTransactionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ttgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TransferTransactionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ttgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = fmt.Errorf("ent: TransferTransactionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ttgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TransferTransactionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TransferTransactionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ttgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TransferTransactionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ttgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = fmt.Errorf("ent: TransferTransactionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ttgb *TransferTransactionGroupBy) BoolX(ctx context.Context) bool {
	v, err := ttgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ttgb *TransferTransactionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ttgb.fields {
		if !transfertransaction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ttgb *TransferTransactionGroupBy) sqlQuery() *sql.Selector {
	selector := ttgb.sql.Select()
	aggregation := make([]string, 0, len(ttgb.fns))
	for _, fn := range ttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ttgb.fields)+len(ttgb.fns))
		for _, f := range ttgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ttgb.fields...)...)
}

// TransferTransactionSelect is the builder for selecting fields of TransferTransaction entities.
type TransferTransactionSelect struct {
	*TransferTransactionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TransferTransactionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	tts.sql = tts.TransferTransactionQuery.sqlQuery(ctx)
	return tts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tts *TransferTransactionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tts *TransferTransactionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TransferTransactionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tts *TransferTransactionSelect) StringsX(ctx context.Context) []string {
	v, err := tts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tts *TransferTransactionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = fmt.Errorf("ent: TransferTransactionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tts *TransferTransactionSelect) StringX(ctx context.Context) string {
	v, err := tts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tts *TransferTransactionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TransferTransactionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tts *TransferTransactionSelect) IntsX(ctx context.Context) []int {
	v, err := tts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tts *TransferTransactionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = fmt.Errorf("ent: TransferTransactionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tts *TransferTransactionSelect) IntX(ctx context.Context) int {
	v, err := tts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tts *TransferTransactionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TransferTransactionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tts *TransferTransactionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tts *TransferTransactionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = fmt.Errorf("ent: TransferTransactionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tts *TransferTransactionSelect) Float64X(ctx context.Context) float64 {
	v, err := tts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tts *TransferTransactionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TransferTransactionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tts *TransferTransactionSelect) BoolsX(ctx context.Context) []bool {
	v, err := tts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tts *TransferTransactionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfertransaction.Label}
	default:
		err = fmt.Errorf("ent: TransferTransactionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tts *TransferTransactionSelect) BoolX(ctx context.Context) bool {
	v, err := tts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tts *TransferTransactionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tts.sql.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
