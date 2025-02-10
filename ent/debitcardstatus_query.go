// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcards"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcardstatus"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

// DebitCardStatusQuery is the builder for querying DebitCardStatus entities.
type DebitCardStatusQuery struct {
	config
	ctx            *QueryContext
	order          []debitcardstatus.OrderOption
	inters         []Interceptor
	predicates     []predicate.DebitCardStatus
	withUsers      *UsersQuery
	withDebitCards *DebitCardsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DebitCardStatusQuery builder.
func (dcsq *DebitCardStatusQuery) Where(ps ...predicate.DebitCardStatus) *DebitCardStatusQuery {
	dcsq.predicates = append(dcsq.predicates, ps...)
	return dcsq
}

// Limit the number of records to be returned by this query.
func (dcsq *DebitCardStatusQuery) Limit(limit int) *DebitCardStatusQuery {
	dcsq.ctx.Limit = &limit
	return dcsq
}

// Offset to start from.
func (dcsq *DebitCardStatusQuery) Offset(offset int) *DebitCardStatusQuery {
	dcsq.ctx.Offset = &offset
	return dcsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dcsq *DebitCardStatusQuery) Unique(unique bool) *DebitCardStatusQuery {
	dcsq.ctx.Unique = &unique
	return dcsq
}

// Order specifies how the records should be ordered.
func (dcsq *DebitCardStatusQuery) Order(o ...debitcardstatus.OrderOption) *DebitCardStatusQuery {
	dcsq.order = append(dcsq.order, o...)
	return dcsq
}

// QueryUsers chains the current query on the "users" edge.
func (dcsq *DebitCardStatusQuery) QueryUsers() *UsersQuery {
	query := (&UsersClient{config: dcsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(debitcardstatus.Table, debitcardstatus.FieldID, selector),
			sqlgraph.To(users.Table, users.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, debitcardstatus.UsersTable, debitcardstatus.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDebitCards chains the current query on the "debit_cards" edge.
func (dcsq *DebitCardStatusQuery) QueryDebitCards() *DebitCardsQuery {
	query := (&DebitCardsClient{config: dcsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(debitcardstatus.Table, debitcardstatus.FieldID, selector),
			sqlgraph.To(debitcards.Table, debitcards.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, debitcardstatus.DebitCardsTable, debitcardstatus.DebitCardsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DebitCardStatus entity from the query.
// Returns a *NotFoundError when no DebitCardStatus was found.
func (dcsq *DebitCardStatusQuery) First(ctx context.Context) (*DebitCardStatus, error) {
	nodes, err := dcsq.Limit(1).All(setContextOp(ctx, dcsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{debitcardstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dcsq *DebitCardStatusQuery) FirstX(ctx context.Context) *DebitCardStatus {
	node, err := dcsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DebitCardStatus ID from the query.
// Returns a *NotFoundError when no DebitCardStatus ID was found.
func (dcsq *DebitCardStatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dcsq.Limit(1).IDs(setContextOp(ctx, dcsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{debitcardstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dcsq *DebitCardStatusQuery) FirstIDX(ctx context.Context) int {
	id, err := dcsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DebitCardStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DebitCardStatus entity is found.
// Returns a *NotFoundError when no DebitCardStatus entities are found.
func (dcsq *DebitCardStatusQuery) Only(ctx context.Context) (*DebitCardStatus, error) {
	nodes, err := dcsq.Limit(2).All(setContextOp(ctx, dcsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{debitcardstatus.Label}
	default:
		return nil, &NotSingularError{debitcardstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dcsq *DebitCardStatusQuery) OnlyX(ctx context.Context) *DebitCardStatus {
	node, err := dcsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DebitCardStatus ID in the query.
// Returns a *NotSingularError when more than one DebitCardStatus ID is found.
// Returns a *NotFoundError when no entities are found.
func (dcsq *DebitCardStatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dcsq.Limit(2).IDs(setContextOp(ctx, dcsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{debitcardstatus.Label}
	default:
		err = &NotSingularError{debitcardstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dcsq *DebitCardStatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := dcsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DebitCardStatusSlice.
func (dcsq *DebitCardStatusQuery) All(ctx context.Context) ([]*DebitCardStatus, error) {
	ctx = setContextOp(ctx, dcsq.ctx, ent.OpQueryAll)
	if err := dcsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DebitCardStatus, *DebitCardStatusQuery]()
	return withInterceptors[[]*DebitCardStatus](ctx, dcsq, qr, dcsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dcsq *DebitCardStatusQuery) AllX(ctx context.Context) []*DebitCardStatus {
	nodes, err := dcsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DebitCardStatus IDs.
func (dcsq *DebitCardStatusQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dcsq.ctx.Unique == nil && dcsq.path != nil {
		dcsq.Unique(true)
	}
	ctx = setContextOp(ctx, dcsq.ctx, ent.OpQueryIDs)
	if err = dcsq.Select(debitcardstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dcsq *DebitCardStatusQuery) IDsX(ctx context.Context) []int {
	ids, err := dcsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dcsq *DebitCardStatusQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dcsq.ctx, ent.OpQueryCount)
	if err := dcsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dcsq, querierCount[*DebitCardStatusQuery](), dcsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dcsq *DebitCardStatusQuery) CountX(ctx context.Context) int {
	count, err := dcsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dcsq *DebitCardStatusQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dcsq.ctx, ent.OpQueryExist)
	switch _, err := dcsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dcsq *DebitCardStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := dcsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DebitCardStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dcsq *DebitCardStatusQuery) Clone() *DebitCardStatusQuery {
	if dcsq == nil {
		return nil
	}
	return &DebitCardStatusQuery{
		config:         dcsq.config,
		ctx:            dcsq.ctx.Clone(),
		order:          append([]debitcardstatus.OrderOption{}, dcsq.order...),
		inters:         append([]Interceptor{}, dcsq.inters...),
		predicates:     append([]predicate.DebitCardStatus{}, dcsq.predicates...),
		withUsers:      dcsq.withUsers.Clone(),
		withDebitCards: dcsq.withDebitCards.Clone(),
		// clone intermediate query.
		sql:  dcsq.sql.Clone(),
		path: dcsq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (dcsq *DebitCardStatusQuery) WithUsers(opts ...func(*UsersQuery)) *DebitCardStatusQuery {
	query := (&UsersClient{config: dcsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcsq.withUsers = query
	return dcsq
}

// WithDebitCards tells the query-builder to eager-load the nodes that are connected to
// the "debit_cards" edge. The optional arguments are used to configure the query builder of the edge.
func (dcsq *DebitCardStatusQuery) WithDebitCards(opts ...func(*DebitCardsQuery)) *DebitCardStatusQuery {
	query := (&DebitCardsClient{config: dcsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcsq.withDebitCards = query
	return dcsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CardID int `json:"card_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DebitCardStatus.Query().
//		GroupBy(debitcardstatus.FieldCardID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dcsq *DebitCardStatusQuery) GroupBy(field string, fields ...string) *DebitCardStatusGroupBy {
	dcsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DebitCardStatusGroupBy{build: dcsq}
	grbuild.flds = &dcsq.ctx.Fields
	grbuild.label = debitcardstatus.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CardID int `json:"card_id,omitempty"`
//	}
//
//	client.DebitCardStatus.Query().
//		Select(debitcardstatus.FieldCardID).
//		Scan(ctx, &v)
func (dcsq *DebitCardStatusQuery) Select(fields ...string) *DebitCardStatusSelect {
	dcsq.ctx.Fields = append(dcsq.ctx.Fields, fields...)
	sbuild := &DebitCardStatusSelect{DebitCardStatusQuery: dcsq}
	sbuild.label = debitcardstatus.Label
	sbuild.flds, sbuild.scan = &dcsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DebitCardStatusSelect configured with the given aggregations.
func (dcsq *DebitCardStatusQuery) Aggregate(fns ...AggregateFunc) *DebitCardStatusSelect {
	return dcsq.Select().Aggregate(fns...)
}

func (dcsq *DebitCardStatusQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dcsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dcsq); err != nil {
				return err
			}
		}
	}
	for _, f := range dcsq.ctx.Fields {
		if !debitcardstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dcsq.path != nil {
		prev, err := dcsq.path(ctx)
		if err != nil {
			return err
		}
		dcsq.sql = prev
	}
	return nil
}

func (dcsq *DebitCardStatusQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DebitCardStatus, error) {
	var (
		nodes       = []*DebitCardStatus{}
		_spec       = dcsq.querySpec()
		loadedTypes = [2]bool{
			dcsq.withUsers != nil,
			dcsq.withDebitCards != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DebitCardStatus).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DebitCardStatus{config: dcsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dcsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dcsq.withUsers; query != nil {
		if err := dcsq.loadUsers(ctx, query, nodes, nil,
			func(n *DebitCardStatus, e *Users) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := dcsq.withDebitCards; query != nil {
		if err := dcsq.loadDebitCards(ctx, query, nodes, nil,
			func(n *DebitCardStatus, e *DebitCards) { n.Edges.DebitCards = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dcsq *DebitCardStatusQuery) loadUsers(ctx context.Context, query *UsersQuery, nodes []*DebitCardStatus, init func(*DebitCardStatus), assign func(*DebitCardStatus, *Users)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DebitCardStatus)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(users.IDIn(ids...))
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
func (dcsq *DebitCardStatusQuery) loadDebitCards(ctx context.Context, query *DebitCardsQuery, nodes []*DebitCardStatus, init func(*DebitCardStatus), assign func(*DebitCardStatus, *DebitCards)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DebitCardStatus)
	for i := range nodes {
		fk := nodes[i].CardID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(debitcards.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "card_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dcsq *DebitCardStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dcsq.querySpec()
	_spec.Node.Columns = dcsq.ctx.Fields
	if len(dcsq.ctx.Fields) > 0 {
		_spec.Unique = dcsq.ctx.Unique != nil && *dcsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dcsq.driver, _spec)
}

func (dcsq *DebitCardStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(debitcardstatus.Table, debitcardstatus.Columns, sqlgraph.NewFieldSpec(debitcardstatus.FieldID, field.TypeInt))
	_spec.From = dcsq.sql
	if unique := dcsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dcsq.path != nil {
		_spec.Unique = true
	}
	if fields := dcsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, debitcardstatus.FieldID)
		for i := range fields {
			if fields[i] != debitcardstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dcsq.withUsers != nil {
			_spec.Node.AddColumnOnce(debitcardstatus.FieldUserID)
		}
		if dcsq.withDebitCards != nil {
			_spec.Node.AddColumnOnce(debitcardstatus.FieldCardID)
		}
	}
	if ps := dcsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dcsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dcsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dcsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dcsq *DebitCardStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dcsq.driver.Dialect())
	t1 := builder.Table(debitcardstatus.Table)
	columns := dcsq.ctx.Fields
	if len(columns) == 0 {
		columns = debitcardstatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dcsq.sql != nil {
		selector = dcsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dcsq.ctx.Unique != nil && *dcsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dcsq.predicates {
		p(selector)
	}
	for _, p := range dcsq.order {
		p(selector)
	}
	if offset := dcsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dcsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DebitCardStatusGroupBy is the group-by builder for DebitCardStatus entities.
type DebitCardStatusGroupBy struct {
	selector
	build *DebitCardStatusQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dcsgb *DebitCardStatusGroupBy) Aggregate(fns ...AggregateFunc) *DebitCardStatusGroupBy {
	dcsgb.fns = append(dcsgb.fns, fns...)
	return dcsgb
}

// Scan applies the selector query and scans the result into the given value.
func (dcsgb *DebitCardStatusGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcsgb.build.ctx, ent.OpQueryGroupBy)
	if err := dcsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DebitCardStatusQuery, *DebitCardStatusGroupBy](ctx, dcsgb.build, dcsgb, dcsgb.build.inters, v)
}

func (dcsgb *DebitCardStatusGroupBy) sqlScan(ctx context.Context, root *DebitCardStatusQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dcsgb.fns))
	for _, fn := range dcsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dcsgb.flds)+len(dcsgb.fns))
		for _, f := range *dcsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dcsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DebitCardStatusSelect is the builder for selecting fields of DebitCardStatus entities.
type DebitCardStatusSelect struct {
	*DebitCardStatusQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dcss *DebitCardStatusSelect) Aggregate(fns ...AggregateFunc) *DebitCardStatusSelect {
	dcss.fns = append(dcss.fns, fns...)
	return dcss
}

// Scan applies the selector query and scans the result into the given value.
func (dcss *DebitCardStatusSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcss.ctx, ent.OpQuerySelect)
	if err := dcss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DebitCardStatusQuery, *DebitCardStatusSelect](ctx, dcss.DebitCardStatusQuery, dcss, dcss.inters, v)
}

func (dcss *DebitCardStatusSelect) sqlScan(ctx context.Context, root *DebitCardStatusQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dcss.fns))
	for _, fn := range dcss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dcss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
