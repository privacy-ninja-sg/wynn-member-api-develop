// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"wynn-member-api/ent/lineaccount"
	"wynn-member-api/ent/predicate"
	"wynn-member-api/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LineAccountQuery is the builder for querying LineAccount entities.
type LineAccountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LineAccount
	// eager-loading edges.
	withOwner *UserQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LineAccountQuery builder.
func (laq *LineAccountQuery) Where(ps ...predicate.LineAccount) *LineAccountQuery {
	laq.predicates = append(laq.predicates, ps...)
	return laq
}

// Limit adds a limit step to the query.
func (laq *LineAccountQuery) Limit(limit int) *LineAccountQuery {
	laq.limit = &limit
	return laq
}

// Offset adds an offset step to the query.
func (laq *LineAccountQuery) Offset(offset int) *LineAccountQuery {
	laq.offset = &offset
	return laq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (laq *LineAccountQuery) Unique(unique bool) *LineAccountQuery {
	laq.unique = &unique
	return laq
}

// Order adds an order step to the query.
func (laq *LineAccountQuery) Order(o ...OrderFunc) *LineAccountQuery {
	laq.order = append(laq.order, o...)
	return laq
}

// QueryOwner chains the current query on the "owner" edge.
func (laq *LineAccountQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: laq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := laq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := laq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lineaccount.Table, lineaccount.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, lineaccount.OwnerTable, lineaccount.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(laq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LineAccount entity from the query.
// Returns a *NotFoundError when no LineAccount was found.
func (laq *LineAccountQuery) First(ctx context.Context) (*LineAccount, error) {
	nodes, err := laq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lineaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (laq *LineAccountQuery) FirstX(ctx context.Context) *LineAccount {
	node, err := laq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LineAccount ID from the query.
// Returns a *NotFoundError when no LineAccount ID was found.
func (laq *LineAccountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = laq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lineaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (laq *LineAccountQuery) FirstIDX(ctx context.Context) int {
	id, err := laq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LineAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one LineAccount entity is not found.
// Returns a *NotFoundError when no LineAccount entities are found.
func (laq *LineAccountQuery) Only(ctx context.Context) (*LineAccount, error) {
	nodes, err := laq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lineaccount.Label}
	default:
		return nil, &NotSingularError{lineaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (laq *LineAccountQuery) OnlyX(ctx context.Context) *LineAccount {
	node, err := laq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LineAccount ID in the query.
// Returns a *NotSingularError when exactly one LineAccount ID is not found.
// Returns a *NotFoundError when no entities are found.
func (laq *LineAccountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = laq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = &NotSingularError{lineaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (laq *LineAccountQuery) OnlyIDX(ctx context.Context) int {
	id, err := laq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LineAccounts.
func (laq *LineAccountQuery) All(ctx context.Context) ([]*LineAccount, error) {
	if err := laq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return laq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (laq *LineAccountQuery) AllX(ctx context.Context) []*LineAccount {
	nodes, err := laq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LineAccount IDs.
func (laq *LineAccountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := laq.Select(lineaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (laq *LineAccountQuery) IDsX(ctx context.Context) []int {
	ids, err := laq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (laq *LineAccountQuery) Count(ctx context.Context) (int, error) {
	if err := laq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return laq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (laq *LineAccountQuery) CountX(ctx context.Context) int {
	count, err := laq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (laq *LineAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := laq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return laq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (laq *LineAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := laq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LineAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (laq *LineAccountQuery) Clone() *LineAccountQuery {
	if laq == nil {
		return nil
	}
	return &LineAccountQuery{
		config:     laq.config,
		limit:      laq.limit,
		offset:     laq.offset,
		order:      append([]OrderFunc{}, laq.order...),
		predicates: append([]predicate.LineAccount{}, laq.predicates...),
		withOwner:  laq.withOwner.Clone(),
		// clone intermediate query.
		sql:  laq.sql.Clone(),
		path: laq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (laq *LineAccountQuery) WithOwner(opts ...func(*UserQuery)) *LineAccountQuery {
	query := &UserQuery{config: laq.config}
	for _, opt := range opts {
		opt(query)
	}
	laq.withOwner = query
	return laq
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
//	client.LineAccount.Query().
//		GroupBy(lineaccount.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (laq *LineAccountQuery) GroupBy(field string, fields ...string) *LineAccountGroupBy {
	group := &LineAccountGroupBy{config: laq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := laq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return laq.sqlQuery(ctx), nil
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
//	client.LineAccount.Query().
//		Select(lineaccount.FieldUUID).
//		Scan(ctx, &v)
//
func (laq *LineAccountQuery) Select(fields ...string) *LineAccountSelect {
	laq.fields = append(laq.fields, fields...)
	return &LineAccountSelect{LineAccountQuery: laq}
}

func (laq *LineAccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range laq.fields {
		if !lineaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if laq.path != nil {
		prev, err := laq.path(ctx)
		if err != nil {
			return err
		}
		laq.sql = prev
	}
	return nil
}

func (laq *LineAccountQuery) sqlAll(ctx context.Context) ([]*LineAccount, error) {
	var (
		nodes       = []*LineAccount{}
		withFKs     = laq.withFKs
		_spec       = laq.querySpec()
		loadedTypes = [1]bool{
			laq.withOwner != nil,
		}
	)
	if laq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, lineaccount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &LineAccount{config: laq.config}
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
	if err := sqlgraph.QueryNodes(ctx, laq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := laq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*LineAccount)
		for i := range nodes {
			if nodes[i].user_line == nil {
				continue
			}
			fk := *nodes[i].user_line
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_line" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	return nodes, nil
}

func (laq *LineAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := laq.querySpec()
	return sqlgraph.CountNodes(ctx, laq.driver, _spec)
}

func (laq *LineAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := laq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (laq *LineAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lineaccount.Table,
			Columns: lineaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lineaccount.FieldID,
			},
		},
		From:   laq.sql,
		Unique: true,
	}
	if unique := laq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := laq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lineaccount.FieldID)
		for i := range fields {
			if fields[i] != lineaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := laq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := laq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := laq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := laq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (laq *LineAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(laq.driver.Dialect())
	t1 := builder.Table(lineaccount.Table)
	columns := laq.fields
	if len(columns) == 0 {
		columns = lineaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if laq.sql != nil {
		selector = laq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range laq.predicates {
		p(selector)
	}
	for _, p := range laq.order {
		p(selector)
	}
	if offset := laq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := laq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LineAccountGroupBy is the group-by builder for LineAccount entities.
type LineAccountGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lagb *LineAccountGroupBy) Aggregate(fns ...AggregateFunc) *LineAccountGroupBy {
	lagb.fns = append(lagb.fns, fns...)
	return lagb
}

// Scan applies the group-by query and scans the result into the given value.
func (lagb *LineAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lagb.path(ctx)
	if err != nil {
		return err
	}
	lagb.sql = query
	return lagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lagb *LineAccountGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := lagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (lagb *LineAccountGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(lagb.fields) > 1 {
		return nil, errors.New("ent: LineAccountGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := lagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lagb *LineAccountGroupBy) StringsX(ctx context.Context) []string {
	v, err := lagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lagb *LineAccountGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = fmt.Errorf("ent: LineAccountGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lagb *LineAccountGroupBy) StringX(ctx context.Context) string {
	v, err := lagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (lagb *LineAccountGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(lagb.fields) > 1 {
		return nil, errors.New("ent: LineAccountGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := lagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lagb *LineAccountGroupBy) IntsX(ctx context.Context) []int {
	v, err := lagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lagb *LineAccountGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = fmt.Errorf("ent: LineAccountGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lagb *LineAccountGroupBy) IntX(ctx context.Context) int {
	v, err := lagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (lagb *LineAccountGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(lagb.fields) > 1 {
		return nil, errors.New("ent: LineAccountGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := lagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lagb *LineAccountGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := lagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lagb *LineAccountGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = fmt.Errorf("ent: LineAccountGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lagb *LineAccountGroupBy) Float64X(ctx context.Context) float64 {
	v, err := lagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (lagb *LineAccountGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(lagb.fields) > 1 {
		return nil, errors.New("ent: LineAccountGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := lagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lagb *LineAccountGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := lagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lagb *LineAccountGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = fmt.Errorf("ent: LineAccountGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lagb *LineAccountGroupBy) BoolX(ctx context.Context) bool {
	v, err := lagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lagb *LineAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lagb.fields {
		if !lineaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lagb *LineAccountGroupBy) sqlQuery() *sql.Selector {
	selector := lagb.sql.Select()
	aggregation := make([]string, 0, len(lagb.fns))
	for _, fn := range lagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lagb.fields)+len(lagb.fns))
		for _, f := range lagb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lagb.fields...)...)
}

// LineAccountSelect is the builder for selecting fields of LineAccount entities.
type LineAccountSelect struct {
	*LineAccountQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (las *LineAccountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := las.prepareQuery(ctx); err != nil {
		return err
	}
	las.sql = las.LineAccountQuery.sqlQuery(ctx)
	return las.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (las *LineAccountSelect) ScanX(ctx context.Context, v interface{}) {
	if err := las.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (las *LineAccountSelect) Strings(ctx context.Context) ([]string, error) {
	if len(las.fields) > 1 {
		return nil, errors.New("ent: LineAccountSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := las.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (las *LineAccountSelect) StringsX(ctx context.Context) []string {
	v, err := las.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (las *LineAccountSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = las.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = fmt.Errorf("ent: LineAccountSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (las *LineAccountSelect) StringX(ctx context.Context) string {
	v, err := las.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (las *LineAccountSelect) Ints(ctx context.Context) ([]int, error) {
	if len(las.fields) > 1 {
		return nil, errors.New("ent: LineAccountSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := las.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (las *LineAccountSelect) IntsX(ctx context.Context) []int {
	v, err := las.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (las *LineAccountSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = las.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = fmt.Errorf("ent: LineAccountSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (las *LineAccountSelect) IntX(ctx context.Context) int {
	v, err := las.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (las *LineAccountSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(las.fields) > 1 {
		return nil, errors.New("ent: LineAccountSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := las.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (las *LineAccountSelect) Float64sX(ctx context.Context) []float64 {
	v, err := las.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (las *LineAccountSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = las.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = fmt.Errorf("ent: LineAccountSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (las *LineAccountSelect) Float64X(ctx context.Context) float64 {
	v, err := las.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (las *LineAccountSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(las.fields) > 1 {
		return nil, errors.New("ent: LineAccountSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := las.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (las *LineAccountSelect) BoolsX(ctx context.Context) []bool {
	v, err := las.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (las *LineAccountSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = las.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lineaccount.Label}
	default:
		err = fmt.Errorf("ent: LineAccountSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (las *LineAccountSelect) BoolX(ctx context.Context) bool {
	v, err := las.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (las *LineAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := las.sql.Query()
	if err := las.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}