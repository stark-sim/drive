// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"drive/ent/directory"
	"drive/ent/object"
	"drive/ent/predicate"
	"drive/ent/user"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ObjectQuery is the builder for querying Object entities.
type ObjectQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.Object
	withUser      *UserQuery
	withDirectory *DirectoryQuery
	withFKs       bool
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*Object) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ObjectQuery builder.
func (oq *ObjectQuery) Where(ps ...predicate.Object) *ObjectQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *ObjectQuery) Limit(limit int) *ObjectQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *ObjectQuery) Offset(offset int) *ObjectQuery {
	oq.offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *ObjectQuery) Unique(unique bool) *ObjectQuery {
	oq.unique = &unique
	return oq
}

// Order adds an order step to the query.
func (oq *ObjectQuery) Order(o ...OrderFunc) *ObjectQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryUser chains the current query on the "user" edge.
func (oq *ObjectQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(object.Table, object.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, object.UserTable, object.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDirectory chains the current query on the "directory" edge.
func (oq *ObjectQuery) QueryDirectory() *DirectoryQuery {
	query := &DirectoryQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(object.Table, object.FieldID, selector),
			sqlgraph.To(directory.Table, directory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, object.DirectoryTable, object.DirectoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Object entity from the query.
// Returns a *NotFoundError when no Object was found.
func (oq *ObjectQuery) First(ctx context.Context) (*Object, error) {
	nodes, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{object.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *ObjectQuery) FirstX(ctx context.Context) *Object {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Object ID from the query.
// Returns a *NotFoundError when no Object ID was found.
func (oq *ObjectQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{object.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *ObjectQuery) FirstIDX(ctx context.Context) int64 {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Object entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Object entity is found.
// Returns a *NotFoundError when no Object entities are found.
func (oq *ObjectQuery) Only(ctx context.Context) (*Object, error) {
	nodes, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{object.Label}
	default:
		return nil, &NotSingularError{object.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *ObjectQuery) OnlyX(ctx context.Context) *Object {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Object ID in the query.
// Returns a *NotSingularError when more than one Object ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *ObjectQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{object.Label}
	default:
		err = &NotSingularError{object.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *ObjectQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Objects.
func (oq *ObjectQuery) All(ctx context.Context) ([]*Object, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *ObjectQuery) AllX(ctx context.Context) []*Object {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Object IDs.
func (oq *ObjectQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := oq.Select(object.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *ObjectQuery) IDsX(ctx context.Context) []int64 {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *ObjectQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *ObjectQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *ObjectQuery) Exist(ctx context.Context) (bool, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *ObjectQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ObjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *ObjectQuery) Clone() *ObjectQuery {
	if oq == nil {
		return nil
	}
	return &ObjectQuery{
		config:        oq.config,
		limit:         oq.limit,
		offset:        oq.offset,
		order:         append([]OrderFunc{}, oq.order...),
		predicates:    append([]predicate.Object{}, oq.predicates...),
		withUser:      oq.withUser.Clone(),
		withDirectory: oq.withDirectory.Clone(),
		// clone intermediate query.
		sql:    oq.sql.Clone(),
		path:   oq.path,
		unique: oq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *ObjectQuery) WithUser(opts ...func(*UserQuery)) *ObjectQuery {
	query := &UserQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withUser = query
	return oq
}

// WithDirectory tells the query-builder to eager-load the nodes that are connected to
// the "directory" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *ObjectQuery) WithDirectory(opts ...func(*DirectoryQuery)) *ObjectQuery {
	query := &DirectoryQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withDirectory = query
	return oq
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
//	client.Object.Query().
//		GroupBy(object.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *ObjectQuery) GroupBy(field string, fields ...string) *ObjectGroupBy {
	grbuild := &ObjectGroupBy{config: oq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(ctx), nil
	}
	grbuild.label = object.Label
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
//	client.Object.Query().
//		Select(object.FieldCreatedBy).
//		Scan(ctx, &v)
func (oq *ObjectQuery) Select(fields ...string) *ObjectSelect {
	oq.fields = append(oq.fields, fields...)
	selbuild := &ObjectSelect{ObjectQuery: oq}
	selbuild.label = object.Label
	selbuild.flds, selbuild.scan = &oq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ObjectSelect configured with the given aggregations.
func (oq *ObjectQuery) Aggregate(fns ...AggregateFunc) *ObjectSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *ObjectQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oq.fields {
		if !object.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *ObjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Object, error) {
	var (
		nodes       = []*Object{}
		withFKs     = oq.withFKs
		_spec       = oq.querySpec()
		loadedTypes = [2]bool{
			oq.withUser != nil,
			oq.withDirectory != nil,
		}
	)
	if oq.withUser != nil || oq.withDirectory != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, object.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Object).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Object{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withUser; query != nil {
		if err := oq.loadUser(ctx, query, nodes, nil,
			func(n *Object, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := oq.withDirectory; query != nil {
		if err := oq.loadDirectory(ctx, query, nodes, nil,
			func(n *Object, e *Directory) { n.Edges.Directory = e }); err != nil {
			return nil, err
		}
	}
	for i := range oq.loadTotal {
		if err := oq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *ObjectQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Object, init func(*Object), assign func(*Object, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Object)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oq *ObjectQuery) loadDirectory(ctx context.Context, query *DirectoryQuery, nodes []*Object, init func(*Object), assign func(*Object, *Directory)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Object)
	for i := range nodes {
		if nodes[i].directory_objects == nil {
			continue
		}
		fk := *nodes[i].directory_objects
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(directory.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "directory_objects" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (oq *ObjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	_spec.Node.Columns = oq.fields
	if len(oq.fields) > 0 {
		_spec.Unique = oq.unique != nil && *oq.unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *ObjectQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (oq *ObjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   object.Table,
			Columns: object.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: object.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if unique := oq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, object.FieldID)
		for i := range fields {
			if fields[i] != object.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *ObjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(object.Table)
	columns := oq.fields
	if len(columns) == 0 {
		columns = object.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.unique != nil && *oq.unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ObjectGroupBy is the group-by builder for Object entities.
type ObjectGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *ObjectGroupBy) Aggregate(fns ...AggregateFunc) *ObjectGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scans the result into the given value.
func (ogb *ObjectGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ogb.path(ctx)
	if err != nil {
		return err
	}
	ogb.sql = query
	return ogb.sqlScan(ctx, v)
}

func (ogb *ObjectGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ogb.fields {
		if !object.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *ObjectGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql.Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
		for _, f := range ogb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ogb.fields...)...)
}

// ObjectSelect is the builder for selecting fields of Object entities.
type ObjectSelect struct {
	*ObjectQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *ObjectSelect) Aggregate(fns ...AggregateFunc) *ObjectSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *ObjectSelect) Scan(ctx context.Context, v any) error {
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	os.sql = os.ObjectQuery.sqlQuery(ctx)
	return os.sqlScan(ctx, v)
}

func (os *ObjectSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(os.sql))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		os.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		os.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := os.sql.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
