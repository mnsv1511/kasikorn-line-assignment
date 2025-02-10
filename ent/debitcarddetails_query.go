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
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcarddetails"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcards"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

// DebitCardDetailsQuery is the builder for querying DebitCardDetails entities.
type DebitCardDetailsQuery struct {
	config
	ctx            *QueryContext
	order          []debitcarddetails.OrderOption
	inters         []Interceptor
	predicates     []predicate.DebitCardDetails
	withUsers      *UsersQuery
	withDebitCards *DebitCardsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DebitCardDetailsQuery builder.
func (dcdq *DebitCardDetailsQuery) Where(ps ...predicate.DebitCardDetails) *DebitCardDetailsQuery {
	dcdq.predicates = append(dcdq.predicates, ps...)
	return dcdq
}

// Limit the number of records to be returned by this query.
func (dcdq *DebitCardDetailsQuery) Limit(limit int) *DebitCardDetailsQuery {
	dcdq.ctx.Limit = &limit
	return dcdq
}

// Offset to start from.
func (dcdq *DebitCardDetailsQuery) Offset(offset int) *DebitCardDetailsQuery {
	dcdq.ctx.Offset = &offset
	return dcdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dcdq *DebitCardDetailsQuery) Unique(unique bool) *DebitCardDetailsQuery {
	dcdq.ctx.Unique = &unique
	return dcdq
}

// Order specifies how the records should be ordered.
func (dcdq *DebitCardDetailsQuery) Order(o ...debitcarddetails.OrderOption) *DebitCardDetailsQuery {
	dcdq.order = append(dcdq.order, o...)
	return dcdq
}

// QueryUsers chains the current query on the "users" edge.
func (dcdq *DebitCardDetailsQuery) QueryUsers() *UsersQuery {
	query := (&UsersClient{config: dcdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(debitcarddetails.Table, debitcarddetails.FieldID, selector),
			sqlgraph.To(users.Table, users.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, debitcarddetails.UsersTable, debitcarddetails.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDebitCards chains the current query on the "debit_cards" edge.
func (dcdq *DebitCardDetailsQuery) QueryDebitCards() *DebitCardsQuery {
	query := (&DebitCardsClient{config: dcdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(debitcarddetails.Table, debitcarddetails.FieldID, selector),
			sqlgraph.To(debitcards.Table, debitcards.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, debitcarddetails.DebitCardsTable, debitcarddetails.DebitCardsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DebitCardDetails entity from the query.
// Returns a *NotFoundError when no DebitCardDetails was found.
func (dcdq *DebitCardDetailsQuery) First(ctx context.Context) (*DebitCardDetails, error) {
	nodes, err := dcdq.Limit(1).All(setContextOp(ctx, dcdq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{debitcarddetails.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dcdq *DebitCardDetailsQuery) FirstX(ctx context.Context) *DebitCardDetails {
	node, err := dcdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DebitCardDetails ID from the query.
// Returns a *NotFoundError when no DebitCardDetails ID was found.
func (dcdq *DebitCardDetailsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dcdq.Limit(1).IDs(setContextOp(ctx, dcdq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{debitcarddetails.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dcdq *DebitCardDetailsQuery) FirstIDX(ctx context.Context) int {
	id, err := dcdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DebitCardDetails entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DebitCardDetails entity is found.
// Returns a *NotFoundError when no DebitCardDetails entities are found.
func (dcdq *DebitCardDetailsQuery) Only(ctx context.Context) (*DebitCardDetails, error) {
	nodes, err := dcdq.Limit(2).All(setContextOp(ctx, dcdq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{debitcarddetails.Label}
	default:
		return nil, &NotSingularError{debitcarddetails.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dcdq *DebitCardDetailsQuery) OnlyX(ctx context.Context) *DebitCardDetails {
	node, err := dcdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DebitCardDetails ID in the query.
// Returns a *NotSingularError when more than one DebitCardDetails ID is found.
// Returns a *NotFoundError when no entities are found.
func (dcdq *DebitCardDetailsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dcdq.Limit(2).IDs(setContextOp(ctx, dcdq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{debitcarddetails.Label}
	default:
		err = &NotSingularError{debitcarddetails.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dcdq *DebitCardDetailsQuery) OnlyIDX(ctx context.Context) int {
	id, err := dcdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DebitCardDetailsSlice.
func (dcdq *DebitCardDetailsQuery) All(ctx context.Context) ([]*DebitCardDetails, error) {
	ctx = setContextOp(ctx, dcdq.ctx, ent.OpQueryAll)
	if err := dcdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DebitCardDetails, *DebitCardDetailsQuery]()
	return withInterceptors[[]*DebitCardDetails](ctx, dcdq, qr, dcdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dcdq *DebitCardDetailsQuery) AllX(ctx context.Context) []*DebitCardDetails {
	nodes, err := dcdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DebitCardDetails IDs.
func (dcdq *DebitCardDetailsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dcdq.ctx.Unique == nil && dcdq.path != nil {
		dcdq.Unique(true)
	}
	ctx = setContextOp(ctx, dcdq.ctx, ent.OpQueryIDs)
	if err = dcdq.Select(debitcarddetails.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dcdq *DebitCardDetailsQuery) IDsX(ctx context.Context) []int {
	ids, err := dcdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dcdq *DebitCardDetailsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dcdq.ctx, ent.OpQueryCount)
	if err := dcdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dcdq, querierCount[*DebitCardDetailsQuery](), dcdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dcdq *DebitCardDetailsQuery) CountX(ctx context.Context) int {
	count, err := dcdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dcdq *DebitCardDetailsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dcdq.ctx, ent.OpQueryExist)
	switch _, err := dcdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dcdq *DebitCardDetailsQuery) ExistX(ctx context.Context) bool {
	exist, err := dcdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DebitCardDetailsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dcdq *DebitCardDetailsQuery) Clone() *DebitCardDetailsQuery {
	if dcdq == nil {
		return nil
	}
	return &DebitCardDetailsQuery{
		config:         dcdq.config,
		ctx:            dcdq.ctx.Clone(),
		order:          append([]debitcarddetails.OrderOption{}, dcdq.order...),
		inters:         append([]Interceptor{}, dcdq.inters...),
		predicates:     append([]predicate.DebitCardDetails{}, dcdq.predicates...),
		withUsers:      dcdq.withUsers.Clone(),
		withDebitCards: dcdq.withDebitCards.Clone(),
		// clone intermediate query.
		sql:  dcdq.sql.Clone(),
		path: dcdq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (dcdq *DebitCardDetailsQuery) WithUsers(opts ...func(*UsersQuery)) *DebitCardDetailsQuery {
	query := (&UsersClient{config: dcdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcdq.withUsers = query
	return dcdq
}

// WithDebitCards tells the query-builder to eager-load the nodes that are connected to
// the "debit_cards" edge. The optional arguments are used to configure the query builder of the edge.
func (dcdq *DebitCardDetailsQuery) WithDebitCards(opts ...func(*DebitCardsQuery)) *DebitCardDetailsQuery {
	query := (&DebitCardsClient{config: dcdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcdq.withDebitCards = query
	return dcdq
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
//	client.DebitCardDetails.Query().
//		GroupBy(debitcarddetails.FieldCardID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dcdq *DebitCardDetailsQuery) GroupBy(field string, fields ...string) *DebitCardDetailsGroupBy {
	dcdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DebitCardDetailsGroupBy{build: dcdq}
	grbuild.flds = &dcdq.ctx.Fields
	grbuild.label = debitcarddetails.Label
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
//	client.DebitCardDetails.Query().
//		Select(debitcarddetails.FieldCardID).
//		Scan(ctx, &v)
func (dcdq *DebitCardDetailsQuery) Select(fields ...string) *DebitCardDetailsSelect {
	dcdq.ctx.Fields = append(dcdq.ctx.Fields, fields...)
	sbuild := &DebitCardDetailsSelect{DebitCardDetailsQuery: dcdq}
	sbuild.label = debitcarddetails.Label
	sbuild.flds, sbuild.scan = &dcdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DebitCardDetailsSelect configured with the given aggregations.
func (dcdq *DebitCardDetailsQuery) Aggregate(fns ...AggregateFunc) *DebitCardDetailsSelect {
	return dcdq.Select().Aggregate(fns...)
}

func (dcdq *DebitCardDetailsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dcdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dcdq); err != nil {
				return err
			}
		}
	}
	for _, f := range dcdq.ctx.Fields {
		if !debitcarddetails.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dcdq.path != nil {
		prev, err := dcdq.path(ctx)
		if err != nil {
			return err
		}
		dcdq.sql = prev
	}
	return nil
}

func (dcdq *DebitCardDetailsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DebitCardDetails, error) {
	var (
		nodes       = []*DebitCardDetails{}
		_spec       = dcdq.querySpec()
		loadedTypes = [2]bool{
			dcdq.withUsers != nil,
			dcdq.withDebitCards != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DebitCardDetails).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DebitCardDetails{config: dcdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dcdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dcdq.withUsers; query != nil {
		if err := dcdq.loadUsers(ctx, query, nodes, nil,
			func(n *DebitCardDetails, e *Users) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := dcdq.withDebitCards; query != nil {
		if err := dcdq.loadDebitCards(ctx, query, nodes, nil,
			func(n *DebitCardDetails, e *DebitCards) { n.Edges.DebitCards = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dcdq *DebitCardDetailsQuery) loadUsers(ctx context.Context, query *UsersQuery, nodes []*DebitCardDetails, init func(*DebitCardDetails), assign func(*DebitCardDetails, *Users)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DebitCardDetails)
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
func (dcdq *DebitCardDetailsQuery) loadDebitCards(ctx context.Context, query *DebitCardsQuery, nodes []*DebitCardDetails, init func(*DebitCardDetails), assign func(*DebitCardDetails, *DebitCards)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DebitCardDetails)
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

func (dcdq *DebitCardDetailsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dcdq.querySpec()
	_spec.Node.Columns = dcdq.ctx.Fields
	if len(dcdq.ctx.Fields) > 0 {
		_spec.Unique = dcdq.ctx.Unique != nil && *dcdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dcdq.driver, _spec)
}

func (dcdq *DebitCardDetailsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(debitcarddetails.Table, debitcarddetails.Columns, sqlgraph.NewFieldSpec(debitcarddetails.FieldID, field.TypeInt))
	_spec.From = dcdq.sql
	if unique := dcdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dcdq.path != nil {
		_spec.Unique = true
	}
	if fields := dcdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, debitcarddetails.FieldID)
		for i := range fields {
			if fields[i] != debitcarddetails.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dcdq.withUsers != nil {
			_spec.Node.AddColumnOnce(debitcarddetails.FieldUserID)
		}
		if dcdq.withDebitCards != nil {
			_spec.Node.AddColumnOnce(debitcarddetails.FieldCardID)
		}
	}
	if ps := dcdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dcdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dcdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dcdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dcdq *DebitCardDetailsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dcdq.driver.Dialect())
	t1 := builder.Table(debitcarddetails.Table)
	columns := dcdq.ctx.Fields
	if len(columns) == 0 {
		columns = debitcarddetails.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dcdq.sql != nil {
		selector = dcdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dcdq.ctx.Unique != nil && *dcdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dcdq.predicates {
		p(selector)
	}
	for _, p := range dcdq.order {
		p(selector)
	}
	if offset := dcdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dcdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DebitCardDetailsGroupBy is the group-by builder for DebitCardDetails entities.
type DebitCardDetailsGroupBy struct {
	selector
	build *DebitCardDetailsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dcdgb *DebitCardDetailsGroupBy) Aggregate(fns ...AggregateFunc) *DebitCardDetailsGroupBy {
	dcdgb.fns = append(dcdgb.fns, fns...)
	return dcdgb
}

// Scan applies the selector query and scans the result into the given value.
func (dcdgb *DebitCardDetailsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcdgb.build.ctx, ent.OpQueryGroupBy)
	if err := dcdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DebitCardDetailsQuery, *DebitCardDetailsGroupBy](ctx, dcdgb.build, dcdgb, dcdgb.build.inters, v)
}

func (dcdgb *DebitCardDetailsGroupBy) sqlScan(ctx context.Context, root *DebitCardDetailsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dcdgb.fns))
	for _, fn := range dcdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dcdgb.flds)+len(dcdgb.fns))
		for _, f := range *dcdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dcdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DebitCardDetailsSelect is the builder for selecting fields of DebitCardDetails entities.
type DebitCardDetailsSelect struct {
	*DebitCardDetailsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dcds *DebitCardDetailsSelect) Aggregate(fns ...AggregateFunc) *DebitCardDetailsSelect {
	dcds.fns = append(dcds.fns, fns...)
	return dcds
}

// Scan applies the selector query and scans the result into the given value.
func (dcds *DebitCardDetailsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcds.ctx, ent.OpQuerySelect)
	if err := dcds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DebitCardDetailsQuery, *DebitCardDetailsSelect](ctx, dcds.DebitCardDetailsQuery, dcds, dcds.inters, v)
}

func (dcds *DebitCardDetailsSelect) sqlScan(ctx context.Context, root *DebitCardDetailsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dcds.fns))
	for _, fn := range dcds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dcds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
