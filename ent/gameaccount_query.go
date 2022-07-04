// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"wynn-member-api/ent/game"
	"wynn-member-api/ent/gameaccount"
	"wynn-member-api/ent/pgslotaccount"
	"wynn-member-api/ent/predicate"
	"wynn-member-api/ent/prettygameaccount"
	"wynn-member-api/ent/sagameaccount"
	"wynn-member-api/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GameAccountQuery is the builder for querying GameAccount entities.
type GameAccountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GameAccount
	// eager-loading edges.
	withOwner  *UserQuery
	withGame   *GameQuery
	withPgslot *PgSlotAccountQuery
	withPretty *PrettyGameAccountQuery
	withSagame *SAGameAccountQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GameAccountQuery builder.
func (gaq *GameAccountQuery) Where(ps ...predicate.GameAccount) *GameAccountQuery {
	gaq.predicates = append(gaq.predicates, ps...)
	return gaq
}

// Limit adds a limit step to the query.
func (gaq *GameAccountQuery) Limit(limit int) *GameAccountQuery {
	gaq.limit = &limit
	return gaq
}

// Offset adds an offset step to the query.
func (gaq *GameAccountQuery) Offset(offset int) *GameAccountQuery {
	gaq.offset = &offset
	return gaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gaq *GameAccountQuery) Unique(unique bool) *GameAccountQuery {
	gaq.unique = &unique
	return gaq
}

// Order adds an order step to the query.
func (gaq *GameAccountQuery) Order(o ...OrderFunc) *GameAccountQuery {
	gaq.order = append(gaq.order, o...)
	return gaq
}

// QueryOwner chains the current query on the "owner" edge.
func (gaq *GameAccountQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: gaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gameaccount.Table, gameaccount.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, gameaccount.OwnerTable, gameaccount.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGame chains the current query on the "game" edge.
func (gaq *GameAccountQuery) QueryGame() *GameQuery {
	query := &GameQuery{config: gaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gameaccount.Table, gameaccount.FieldID, selector),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, gameaccount.GameTable, gameaccount.GameColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPgslot chains the current query on the "pgslot" edge.
func (gaq *GameAccountQuery) QueryPgslot() *PgSlotAccountQuery {
	query := &PgSlotAccountQuery{config: gaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gameaccount.Table, gameaccount.FieldID, selector),
			sqlgraph.To(pgslotaccount.Table, pgslotaccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gameaccount.PgslotTable, gameaccount.PgslotColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPretty chains the current query on the "pretty" edge.
func (gaq *GameAccountQuery) QueryPretty() *PrettyGameAccountQuery {
	query := &PrettyGameAccountQuery{config: gaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gameaccount.Table, gameaccount.FieldID, selector),
			sqlgraph.To(prettygameaccount.Table, prettygameaccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gameaccount.PrettyTable, gameaccount.PrettyColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySagame chains the current query on the "sagame" edge.
func (gaq *GameAccountQuery) QuerySagame() *SAGameAccountQuery {
	query := &SAGameAccountQuery{config: gaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gameaccount.Table, gameaccount.FieldID, selector),
			sqlgraph.To(sagameaccount.Table, sagameaccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gameaccount.SagameTable, gameaccount.SagameColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GameAccount entity from the query.
// Returns a *NotFoundError when no GameAccount was found.
func (gaq *GameAccountQuery) First(ctx context.Context) (*GameAccount, error) {
	nodes, err := gaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gameaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gaq *GameAccountQuery) FirstX(ctx context.Context) *GameAccount {
	node, err := gaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GameAccount ID from the query.
// Returns a *NotFoundError when no GameAccount ID was found.
func (gaq *GameAccountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gameaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gaq *GameAccountQuery) FirstIDX(ctx context.Context) int {
	id, err := gaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GameAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GameAccount entity is not found.
// Returns a *NotFoundError when no GameAccount entities are found.
func (gaq *GameAccountQuery) Only(ctx context.Context) (*GameAccount, error) {
	nodes, err := gaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gameaccount.Label}
	default:
		return nil, &NotSingularError{gameaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gaq *GameAccountQuery) OnlyX(ctx context.Context) *GameAccount {
	node, err := gaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GameAccount ID in the query.
// Returns a *NotSingularError when exactly one GameAccount ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gaq *GameAccountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = &NotSingularError{gameaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gaq *GameAccountQuery) OnlyIDX(ctx context.Context) int {
	id, err := gaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GameAccounts.
func (gaq *GameAccountQuery) All(ctx context.Context) ([]*GameAccount, error) {
	if err := gaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gaq *GameAccountQuery) AllX(ctx context.Context) []*GameAccount {
	nodes, err := gaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GameAccount IDs.
func (gaq *GameAccountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gaq.Select(gameaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gaq *GameAccountQuery) IDsX(ctx context.Context) []int {
	ids, err := gaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gaq *GameAccountQuery) Count(ctx context.Context) (int, error) {
	if err := gaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gaq *GameAccountQuery) CountX(ctx context.Context) int {
	count, err := gaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gaq *GameAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := gaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gaq *GameAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := gaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GameAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gaq *GameAccountQuery) Clone() *GameAccountQuery {
	if gaq == nil {
		return nil
	}
	return &GameAccountQuery{
		config:     gaq.config,
		limit:      gaq.limit,
		offset:     gaq.offset,
		order:      append([]OrderFunc{}, gaq.order...),
		predicates: append([]predicate.GameAccount{}, gaq.predicates...),
		withOwner:  gaq.withOwner.Clone(),
		withGame:   gaq.withGame.Clone(),
		withPgslot: gaq.withPgslot.Clone(),
		withPretty: gaq.withPretty.Clone(),
		withSagame: gaq.withSagame.Clone(),
		// clone intermediate query.
		sql:  gaq.sql.Clone(),
		path: gaq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GameAccountQuery) WithOwner(opts ...func(*UserQuery)) *GameAccountQuery {
	query := &UserQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	gaq.withOwner = query
	return gaq
}

// WithGame tells the query-builder to eager-load the nodes that are connected to
// the "game" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GameAccountQuery) WithGame(opts ...func(*GameQuery)) *GameAccountQuery {
	query := &GameQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	gaq.withGame = query
	return gaq
}

// WithPgslot tells the query-builder to eager-load the nodes that are connected to
// the "pgslot" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GameAccountQuery) WithPgslot(opts ...func(*PgSlotAccountQuery)) *GameAccountQuery {
	query := &PgSlotAccountQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	gaq.withPgslot = query
	return gaq
}

// WithPretty tells the query-builder to eager-load the nodes that are connected to
// the "pretty" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GameAccountQuery) WithPretty(opts ...func(*PrettyGameAccountQuery)) *GameAccountQuery {
	query := &PrettyGameAccountQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	gaq.withPretty = query
	return gaq
}

// WithSagame tells the query-builder to eager-load the nodes that are connected to
// the "sagame" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GameAccountQuery) WithSagame(opts ...func(*SAGameAccountQuery)) *GameAccountQuery {
	query := &SAGameAccountQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	gaq.withSagame = query
	return gaq
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
//	client.GameAccount.Query().
//		GroupBy(gameaccount.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gaq *GameAccountQuery) GroupBy(field string, fields ...string) *GameAccountGroupBy {
	group := &GameAccountGroupBy{config: gaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gaq.sqlQuery(ctx), nil
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
//	client.GameAccount.Query().
//		Select(gameaccount.FieldUUID).
//		Scan(ctx, &v)
//
func (gaq *GameAccountQuery) Select(fields ...string) *GameAccountSelect {
	gaq.fields = append(gaq.fields, fields...)
	return &GameAccountSelect{GameAccountQuery: gaq}
}

func (gaq *GameAccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gaq.fields {
		if !gameaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gaq.path != nil {
		prev, err := gaq.path(ctx)
		if err != nil {
			return err
		}
		gaq.sql = prev
	}
	return nil
}

func (gaq *GameAccountQuery) sqlAll(ctx context.Context) ([]*GameAccount, error) {
	var (
		nodes       = []*GameAccount{}
		withFKs     = gaq.withFKs
		_spec       = gaq.querySpec()
		loadedTypes = [5]bool{
			gaq.withOwner != nil,
			gaq.withGame != nil,
			gaq.withPgslot != nil,
			gaq.withPretty != nil,
			gaq.withSagame != nil,
		}
	)
	if gaq.withOwner != nil || gaq.withGame != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, gameaccount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GameAccount{config: gaq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gaq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*GameAccount)
		for i := range nodes {
			if nodes[i].user_games == nil {
				continue
			}
			fk := *nodes[i].user_games
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_games" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	if query := gaq.withGame; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*GameAccount)
		for i := range nodes {
			if nodes[i].game_accounts == nil {
				continue
			}
			fk := *nodes[i].game_accounts
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
				return nil, fmt.Errorf(`unexpected foreign-key "game_accounts" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Game = n
			}
		}
	}

	if query := gaq.withPgslot; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GameAccount)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Pgslot = []*PgSlotAccount{}
		}
		query.withFKs = true
		query.Where(predicate.PgSlotAccount(func(s *sql.Selector) {
			s.Where(sql.InValues(gameaccount.PgslotColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.game_account_pgslot
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "game_account_pgslot" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "game_account_pgslot" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Pgslot = append(node.Edges.Pgslot, n)
		}
	}

	if query := gaq.withPretty; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GameAccount)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Pretty = []*PrettyGameAccount{}
		}
		query.withFKs = true
		query.Where(predicate.PrettyGameAccount(func(s *sql.Selector) {
			s.Where(sql.InValues(gameaccount.PrettyColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.game_account_pretty
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "game_account_pretty" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "game_account_pretty" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Pretty = append(node.Edges.Pretty, n)
		}
	}

	if query := gaq.withSagame; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GameAccount)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Sagame = []*SAGameAccount{}
		}
		query.withFKs = true
		query.Where(predicate.SAGameAccount(func(s *sql.Selector) {
			s.Where(sql.InValues(gameaccount.SagameColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.game_account_sagame
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "game_account_sagame" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "game_account_sagame" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Sagame = append(node.Edges.Sagame, n)
		}
	}

	return nodes, nil
}

func (gaq *GameAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gaq.querySpec()
	return sqlgraph.CountNodes(ctx, gaq.driver, _spec)
}

func (gaq *GameAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gaq *GameAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gameaccount.Table,
			Columns: gameaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gameaccount.FieldID,
			},
		},
		From:   gaq.sql,
		Unique: true,
	}
	if unique := gaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gameaccount.FieldID)
		for i := range fields {
			if fields[i] != gameaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gaq *GameAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gaq.driver.Dialect())
	t1 := builder.Table(gameaccount.Table)
	columns := gaq.fields
	if len(columns) == 0 {
		columns = gameaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gaq.sql != nil {
		selector = gaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range gaq.predicates {
		p(selector)
	}
	for _, p := range gaq.order {
		p(selector)
	}
	if offset := gaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GameAccountGroupBy is the group-by builder for GameAccount entities.
type GameAccountGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gagb *GameAccountGroupBy) Aggregate(fns ...AggregateFunc) *GameAccountGroupBy {
	gagb.fns = append(gagb.fns, fns...)
	return gagb
}

// Scan applies the group-by query and scans the result into the given value.
func (gagb *GameAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gagb.path(ctx)
	if err != nil {
		return err
	}
	gagb.sql = query
	return gagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gagb *GameAccountGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gagb *GameAccountGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gagb.fields) > 1 {
		return nil, errors.New("ent: GameAccountGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gagb *GameAccountGroupBy) StringsX(ctx context.Context) []string {
	v, err := gagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gagb *GameAccountGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = fmt.Errorf("ent: GameAccountGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gagb *GameAccountGroupBy) StringX(ctx context.Context) string {
	v, err := gagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gagb *GameAccountGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gagb.fields) > 1 {
		return nil, errors.New("ent: GameAccountGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gagb *GameAccountGroupBy) IntsX(ctx context.Context) []int {
	v, err := gagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gagb *GameAccountGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = fmt.Errorf("ent: GameAccountGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gagb *GameAccountGroupBy) IntX(ctx context.Context) int {
	v, err := gagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gagb *GameAccountGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gagb.fields) > 1 {
		return nil, errors.New("ent: GameAccountGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gagb *GameAccountGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gagb *GameAccountGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = fmt.Errorf("ent: GameAccountGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gagb *GameAccountGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gagb *GameAccountGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gagb.fields) > 1 {
		return nil, errors.New("ent: GameAccountGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gagb *GameAccountGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gagb *GameAccountGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = fmt.Errorf("ent: GameAccountGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gagb *GameAccountGroupBy) BoolX(ctx context.Context) bool {
	v, err := gagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gagb *GameAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gagb.fields {
		if !gameaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gagb *GameAccountGroupBy) sqlQuery() *sql.Selector {
	selector := gagb.sql.Select()
	aggregation := make([]string, 0, len(gagb.fns))
	for _, fn := range gagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gagb.fields)+len(gagb.fns))
		for _, f := range gagb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gagb.fields...)...)
}

// GameAccountSelect is the builder for selecting fields of GameAccount entities.
type GameAccountSelect struct {
	*GameAccountQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gas *GameAccountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gas.prepareQuery(ctx); err != nil {
		return err
	}
	gas.sql = gas.GameAccountQuery.sqlQuery(ctx)
	return gas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gas *GameAccountSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gas *GameAccountSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gas.fields) > 1 {
		return nil, errors.New("ent: GameAccountSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gas *GameAccountSelect) StringsX(ctx context.Context) []string {
	v, err := gas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gas *GameAccountSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = fmt.Errorf("ent: GameAccountSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gas *GameAccountSelect) StringX(ctx context.Context) string {
	v, err := gas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gas *GameAccountSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gas.fields) > 1 {
		return nil, errors.New("ent: GameAccountSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gas *GameAccountSelect) IntsX(ctx context.Context) []int {
	v, err := gas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gas *GameAccountSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = fmt.Errorf("ent: GameAccountSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gas *GameAccountSelect) IntX(ctx context.Context) int {
	v, err := gas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gas *GameAccountSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gas.fields) > 1 {
		return nil, errors.New("ent: GameAccountSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gas *GameAccountSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gas *GameAccountSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = fmt.Errorf("ent: GameAccountSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gas *GameAccountSelect) Float64X(ctx context.Context) float64 {
	v, err := gas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gas *GameAccountSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gas.fields) > 1 {
		return nil, errors.New("ent: GameAccountSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gas *GameAccountSelect) BoolsX(ctx context.Context) []bool {
	v, err := gas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gas *GameAccountSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gameaccount.Label}
	default:
		err = fmt.Errorf("ent: GameAccountSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gas *GameAccountSelect) BoolX(ctx context.Context) bool {
	v, err := gas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gas *GameAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gas.sql.Query()
	if err := gas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
