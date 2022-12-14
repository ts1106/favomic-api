// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ts1106/favomic-api/ent/author"
	"github.com/ts1106/favomic-api/ent/comic"
	"github.com/ts1106/favomic-api/ent/episode"
	"github.com/ts1106/favomic-api/ent/magazine"
	"github.com/ts1106/favomic-api/ent/predicate"
	"github.com/ts1106/favomic-api/ent/tag"
)

// ComicQuery is the builder for querying Comic entities.
type ComicQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.Comic
	withAuthor   *AuthorQuery
	withMagazine *MagazineQuery
	withTags     *TagQuery
	withEpisodes *EpisodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ComicQuery builder.
func (cq *ComicQuery) Where(ps ...predicate.Comic) *ComicQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *ComicQuery) Limit(limit int) *ComicQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *ComicQuery) Offset(offset int) *ComicQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ComicQuery) Unique(unique bool) *ComicQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *ComicQuery) Order(o ...OrderFunc) *ComicQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryAuthor chains the current query on the "author" edge.
func (cq *ComicQuery) QueryAuthor() *AuthorQuery {
	query := &AuthorQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(comic.Table, comic.FieldID, selector),
			sqlgraph.To(author.Table, author.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comic.AuthorTable, comic.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMagazine chains the current query on the "magazine" edge.
func (cq *ComicQuery) QueryMagazine() *MagazineQuery {
	query := &MagazineQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(comic.Table, comic.FieldID, selector),
			sqlgraph.To(magazine.Table, magazine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comic.MagazineTable, comic.MagazineColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (cq *ComicQuery) QueryTags() *TagQuery {
	query := &TagQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(comic.Table, comic.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, comic.TagsTable, comic.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEpisodes chains the current query on the "episodes" edge.
func (cq *ComicQuery) QueryEpisodes() *EpisodeQuery {
	query := &EpisodeQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(comic.Table, comic.FieldID, selector),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, comic.EpisodesTable, comic.EpisodesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Comic entity from the query.
// Returns a *NotFoundError when no Comic was found.
func (cq *ComicQuery) First(ctx context.Context) (*Comic, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{comic.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ComicQuery) FirstX(ctx context.Context) *Comic {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Comic ID from the query.
// Returns a *NotFoundError when no Comic ID was found.
func (cq *ComicQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{comic.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ComicQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Comic entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Comic entity is found.
// Returns a *NotFoundError when no Comic entities are found.
func (cq *ComicQuery) Only(ctx context.Context) (*Comic, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{comic.Label}
	default:
		return nil, &NotSingularError{comic.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ComicQuery) OnlyX(ctx context.Context) *Comic {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Comic ID in the query.
// Returns a *NotSingularError when more than one Comic ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ComicQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{comic.Label}
	default:
		err = &NotSingularError{comic.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ComicQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Comics.
func (cq *ComicQuery) All(ctx context.Context) ([]*Comic, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *ComicQuery) AllX(ctx context.Context) []*Comic {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Comic IDs.
func (cq *ComicQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cq.Select(comic.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ComicQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ComicQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ComicQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ComicQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ComicQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ComicQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ComicQuery) Clone() *ComicQuery {
	if cq == nil {
		return nil
	}
	return &ComicQuery{
		config:       cq.config,
		limit:        cq.limit,
		offset:       cq.offset,
		order:        append([]OrderFunc{}, cq.order...),
		predicates:   append([]predicate.Comic{}, cq.predicates...),
		withAuthor:   cq.withAuthor.Clone(),
		withMagazine: cq.withMagazine.Clone(),
		withTags:     cq.withTags.Clone(),
		withEpisodes: cq.withEpisodes.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ComicQuery) WithAuthor(opts ...func(*AuthorQuery)) *ComicQuery {
	query := &AuthorQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withAuthor = query
	return cq
}

// WithMagazine tells the query-builder to eager-load the nodes that are connected to
// the "magazine" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ComicQuery) WithMagazine(opts ...func(*MagazineQuery)) *ComicQuery {
	query := &MagazineQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withMagazine = query
	return cq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ComicQuery) WithTags(opts ...func(*TagQuery)) *ComicQuery {
	query := &TagQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withTags = query
	return cq
}

// WithEpisodes tells the query-builder to eager-load the nodes that are connected to
// the "episodes" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ComicQuery) WithEpisodes(opts ...func(*EpisodeQuery)) *ComicQuery {
	query := &EpisodeQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withEpisodes = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Comic.Query().
//		GroupBy(comic.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ComicQuery) GroupBy(field string, fields ...string) *ComicGroupBy {
	grbuild := &ComicGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = comic.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Comic.Query().
//		Select(comic.FieldTitle).
//		Scan(ctx, &v)
func (cq *ComicQuery) Select(fields ...string) *ComicSelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &ComicSelect{ComicQuery: cq}
	selbuild.label = comic.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ComicSelect configured with the given aggregations.
func (cq *ComicQuery) Aggregate(fns ...AggregateFunc) *ComicSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ComicQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !comic.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ComicQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Comic, error) {
	var (
		nodes       = []*Comic{}
		_spec       = cq.querySpec()
		loadedTypes = [4]bool{
			cq.withAuthor != nil,
			cq.withMagazine != nil,
			cq.withTags != nil,
			cq.withEpisodes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Comic).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Comic{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withAuthor; query != nil {
		if err := cq.loadAuthor(ctx, query, nodes, nil,
			func(n *Comic, e *Author) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withMagazine; query != nil {
		if err := cq.loadMagazine(ctx, query, nodes, nil,
			func(n *Comic, e *Magazine) { n.Edges.Magazine = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withTags; query != nil {
		if err := cq.loadTags(ctx, query, nodes,
			func(n *Comic) { n.Edges.Tags = []*Tag{} },
			func(n *Comic, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withEpisodes; query != nil {
		if err := cq.loadEpisodes(ctx, query, nodes,
			func(n *Comic) { n.Edges.Episodes = []*Episode{} },
			func(n *Comic, e *Episode) { n.Edges.Episodes = append(n.Edges.Episodes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ComicQuery) loadAuthor(ctx context.Context, query *AuthorQuery, nodes []*Comic, init func(*Comic), assign func(*Comic, *Author)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Comic)
	for i := range nodes {
		fk := nodes[i].AuthorID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(author.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "author_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ComicQuery) loadMagazine(ctx context.Context, query *MagazineQuery, nodes []*Comic, init func(*Comic), assign func(*Comic, *Magazine)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Comic)
	for i := range nodes {
		fk := nodes[i].MagazineID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(magazine.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "magazine_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ComicQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*Comic, init func(*Comic), assign func(*Comic, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Comic)
	nids := make(map[uuid.UUID]map[*Comic]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(comic.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(comic.TagsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(comic.TagsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(comic.TagsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := *values[1].(*uuid.UUID)
			if nids[inValue] == nil {
				nids[inValue] = map[*Comic]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *ComicQuery) loadEpisodes(ctx context.Context, query *EpisodeQuery, nodes []*Comic, init func(*Comic), assign func(*Comic, *Episode)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Comic)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.InValues(comic.EpisodesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ComicID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "comic_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ComicQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ComicQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (cq *ComicQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comic.Table,
			Columns: comic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: comic.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comic.FieldID)
		for i := range fields {
			if fields[i] != comic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ComicQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(comic.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = comic.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ComicGroupBy is the group-by builder for Comic entities.
type ComicGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ComicGroupBy) Aggregate(fns ...AggregateFunc) *ComicGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *ComicGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *ComicGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cgb.fields {
		if !comic.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *ComicGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// ComicSelect is the builder for selecting fields of Comic entities.
type ComicSelect struct {
	*ComicQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ComicSelect) Aggregate(fns ...AggregateFunc) *ComicSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ComicSelect) Scan(ctx context.Context, v any) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.ComicQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *ComicSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(cs.sql))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		cs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		cs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
