// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"drive/ent/predicate"
	"drive/ent/social"
	"drive/ent/wechat"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WechatQuery is the builder for querying Wechat entities.
type WechatQuery struct {
	config
	limit            *int
	offset           *int
	unique           *bool
	order            []OrderFunc
	fields           []string
	predicates       []predicate.Wechat
	withSocials      *SocialQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*Wechat) error
	withNamedSocials map[string]*SocialQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WechatQuery builder.
func (wq *WechatQuery) Where(ps ...predicate.Wechat) *WechatQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *WechatQuery) Limit(limit int) *WechatQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *WechatQuery) Offset(offset int) *WechatQuery {
	wq.offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WechatQuery) Unique(unique bool) *WechatQuery {
	wq.unique = &unique
	return wq
}

// Order adds an order step to the query.
func (wq *WechatQuery) Order(o ...OrderFunc) *WechatQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QuerySocials chains the current query on the "socials" edge.
func (wq *WechatQuery) QuerySocials() *SocialQuery {
	query := &SocialQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(wechat.Table, wechat.FieldID, selector),
			sqlgraph.To(social.Table, social.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wechat.SocialsTable, wechat.SocialsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Wechat entity from the query.
// Returns a *NotFoundError when no Wechat was found.
func (wq *WechatQuery) First(ctx context.Context) (*Wechat, error) {
	nodes, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wechat.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WechatQuery) FirstX(ctx context.Context) *Wechat {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Wechat ID from the query.
// Returns a *NotFoundError when no Wechat ID was found.
func (wq *WechatQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wechat.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WechatQuery) FirstIDX(ctx context.Context) int64 {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Wechat entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Wechat entity is found.
// Returns a *NotFoundError when no Wechat entities are found.
func (wq *WechatQuery) Only(ctx context.Context) (*Wechat, error) {
	nodes, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wechat.Label}
	default:
		return nil, &NotSingularError{wechat.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WechatQuery) OnlyX(ctx context.Context) *Wechat {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Wechat ID in the query.
// Returns a *NotSingularError when more than one Wechat ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WechatQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wechat.Label}
	default:
		err = &NotSingularError{wechat.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WechatQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Wechats.
func (wq *WechatQuery) All(ctx context.Context) ([]*Wechat, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *WechatQuery) AllX(ctx context.Context) []*Wechat {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Wechat IDs.
func (wq *WechatQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := wq.Select(wechat.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WechatQuery) IDsX(ctx context.Context) []int64 {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WechatQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WechatQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WechatQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WechatQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WechatQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WechatQuery) Clone() *WechatQuery {
	if wq == nil {
		return nil
	}
	return &WechatQuery{
		config:      wq.config,
		limit:       wq.limit,
		offset:      wq.offset,
		order:       append([]OrderFunc{}, wq.order...),
		predicates:  append([]predicate.Wechat{}, wq.predicates...),
		withSocials: wq.withSocials.Clone(),
		// clone intermediate query.
		sql:    wq.sql.Clone(),
		path:   wq.path,
		unique: wq.unique,
	}
}

// WithSocials tells the query-builder to eager-load the nodes that are connected to
// the "socials" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WechatQuery) WithSocials(opts ...func(*SocialQuery)) *WechatQuery {
	query := &SocialQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withSocials = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Wechat.Query().
//		GroupBy(wechat.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WechatQuery) GroupBy(field string, fields ...string) *WechatGroupBy {
	grbuild := &WechatGroupBy{config: wq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(ctx), nil
	}
	grbuild.label = wechat.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//	}
//
//	client.Wechat.Query().
//		Select(wechat.FieldCreatedBy).
//		Scan(ctx, &v)
func (wq *WechatQuery) Select(fields ...string) *WechatSelect {
	wq.fields = append(wq.fields, fields...)
	selbuild := &WechatSelect{WechatQuery: wq}
	selbuild.label = wechat.Label
	selbuild.flds, selbuild.scan = &wq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a WechatSelect configured with the given aggregations.
func (wq *WechatQuery) Aggregate(fns ...AggregateFunc) *WechatSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WechatQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wq.fields {
		if !wechat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WechatQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Wechat, error) {
	var (
		nodes       = []*Wechat{}
		_spec       = wq.querySpec()
		loadedTypes = [1]bool{
			wq.withSocials != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Wechat).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Wechat{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withSocials; query != nil {
		if err := wq.loadSocials(ctx, query, nodes,
			func(n *Wechat) { n.Edges.Socials = []*Social{} },
			func(n *Wechat, e *Social) { n.Edges.Socials = append(n.Edges.Socials, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedSocials {
		if err := wq.loadSocials(ctx, query, nodes,
			func(n *Wechat) { n.appendNamedSocials(name) },
			func(n *Wechat, e *Social) { n.appendNamedSocials(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range wq.loadTotal {
		if err := wq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WechatQuery) loadSocials(ctx context.Context, query *SocialQuery, nodes []*Wechat, init func(*Wechat), assign func(*Wechat, *Social)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Wechat)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Social(func(s *sql.Selector) {
		s.Where(sql.InValues(wechat.SocialsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RelationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wq *WechatQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.fields
	if len(wq.fields) > 0 {
		_spec.Unique = wq.unique != nil && *wq.unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WechatQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (wq *WechatQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wechat.Table,
			Columns: wechat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: wechat.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if unique := wq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wechat.FieldID)
		for i := range fields {
			if fields[i] != wechat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WechatQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(wechat.Table)
	columns := wq.fields
	if len(columns) == 0 {
		columns = wechat.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.unique != nil && *wq.unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedSocials tells the query-builder to eager-load the nodes that are connected to the "socials"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WechatQuery) WithNamedSocials(name string, opts ...func(*SocialQuery)) *WechatQuery {
	query := &SocialQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedSocials == nil {
		wq.withNamedSocials = make(map[string]*SocialQuery)
	}
	wq.withNamedSocials[name] = query
	return wq
}

// WechatGroupBy is the group-by builder for Wechat entities.
type WechatGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WechatGroupBy) Aggregate(fns ...AggregateFunc) *WechatGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wgb *WechatGroupBy) Scan(ctx context.Context, v any) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

func (wgb *WechatGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range wgb.fields {
		if !wechat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *WechatGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql.Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
		for _, f := range wgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wgb.fields...)...)
}

// WechatSelect is the builder for selecting fields of Wechat entities.
type WechatSelect struct {
	*WechatQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WechatSelect) Aggregate(fns ...AggregateFunc) *WechatSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WechatSelect) Scan(ctx context.Context, v any) error {
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	ws.sql = ws.WechatQuery.sqlQuery(ctx)
	return ws.sqlScan(ctx, v)
}

func (ws *WechatSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(ws.sql))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ws.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ws.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ws.sql.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
