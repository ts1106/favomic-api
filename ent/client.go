// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/ts1106/favomic-api/ent/migrate"

	"github.com/ts1106/favomic-api/ent/author"
	"github.com/ts1106/favomic-api/ent/comic"
	"github.com/ts1106/favomic-api/ent/episode"
	"github.com/ts1106/favomic-api/ent/magazine"
	"github.com/ts1106/favomic-api/ent/tag"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Author is the client for interacting with the Author builders.
	Author *AuthorClient
	// Comic is the client for interacting with the Comic builders.
	Comic *ComicClient
	// Episode is the client for interacting with the Episode builders.
	Episode *EpisodeClient
	// Magazine is the client for interacting with the Magazine builders.
	Magazine *MagazineClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Author = NewAuthorClient(c.config)
	c.Comic = NewComicClient(c.config)
	c.Episode = NewEpisodeClient(c.config)
	c.Magazine = NewMagazineClient(c.config)
	c.Tag = NewTagClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Author:   NewAuthorClient(cfg),
		Comic:    NewComicClient(cfg),
		Episode:  NewEpisodeClient(cfg),
		Magazine: NewMagazineClient(cfg),
		Tag:      NewTagClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Author:   NewAuthorClient(cfg),
		Comic:    NewComicClient(cfg),
		Episode:  NewEpisodeClient(cfg),
		Magazine: NewMagazineClient(cfg),
		Tag:      NewTagClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Author.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Author.Use(hooks...)
	c.Comic.Use(hooks...)
	c.Episode.Use(hooks...)
	c.Magazine.Use(hooks...)
	c.Tag.Use(hooks...)
}

// AuthorClient is a client for the Author schema.
type AuthorClient struct {
	config
}

// NewAuthorClient returns a client for the Author from the given config.
func NewAuthorClient(c config) *AuthorClient {
	return &AuthorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `author.Hooks(f(g(h())))`.
func (c *AuthorClient) Use(hooks ...Hook) {
	c.hooks.Author = append(c.hooks.Author, hooks...)
}

// Create returns a builder for creating a Author entity.
func (c *AuthorClient) Create() *AuthorCreate {
	mutation := newAuthorMutation(c.config, OpCreate)
	return &AuthorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Author entities.
func (c *AuthorClient) CreateBulk(builders ...*AuthorCreate) *AuthorCreateBulk {
	return &AuthorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Author.
func (c *AuthorClient) Update() *AuthorUpdate {
	mutation := newAuthorMutation(c.config, OpUpdate)
	return &AuthorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthorClient) UpdateOne(a *Author) *AuthorUpdateOne {
	mutation := newAuthorMutation(c.config, OpUpdateOne, withAuthor(a))
	return &AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthorClient) UpdateOneID(id uuid.UUID) *AuthorUpdateOne {
	mutation := newAuthorMutation(c.config, OpUpdateOne, withAuthorID(id))
	return &AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Author.
func (c *AuthorClient) Delete() *AuthorDelete {
	mutation := newAuthorMutation(c.config, OpDelete)
	return &AuthorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AuthorClient) DeleteOne(a *Author) *AuthorDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AuthorClient) DeleteOneID(id uuid.UUID) *AuthorDeleteOne {
	builder := c.Delete().Where(author.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthorDeleteOne{builder}
}

// Query returns a query builder for Author.
func (c *AuthorClient) Query() *AuthorQuery {
	return &AuthorQuery{
		config: c.config,
	}
}

// Get returns a Author entity by its id.
func (c *AuthorClient) Get(ctx context.Context, id uuid.UUID) (*Author, error) {
	return c.Query().Where(author.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthorClient) GetX(ctx context.Context, id uuid.UUID) *Author {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComics queries the comics edge of a Author.
func (c *AuthorClient) QueryComics(a *Author) *ComicQuery {
	query := &ComicQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(author.Table, author.FieldID, id),
			sqlgraph.To(comic.Table, comic.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, author.ComicsTable, author.ComicsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AuthorClient) Hooks() []Hook {
	return c.hooks.Author
}

// ComicClient is a client for the Comic schema.
type ComicClient struct {
	config
}

// NewComicClient returns a client for the Comic from the given config.
func NewComicClient(c config) *ComicClient {
	return &ComicClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comic.Hooks(f(g(h())))`.
func (c *ComicClient) Use(hooks ...Hook) {
	c.hooks.Comic = append(c.hooks.Comic, hooks...)
}

// Create returns a builder for creating a Comic entity.
func (c *ComicClient) Create() *ComicCreate {
	mutation := newComicMutation(c.config, OpCreate)
	return &ComicCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comic entities.
func (c *ComicClient) CreateBulk(builders ...*ComicCreate) *ComicCreateBulk {
	return &ComicCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comic.
func (c *ComicClient) Update() *ComicUpdate {
	mutation := newComicMutation(c.config, OpUpdate)
	return &ComicUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ComicClient) UpdateOne(co *Comic) *ComicUpdateOne {
	mutation := newComicMutation(c.config, OpUpdateOne, withComic(co))
	return &ComicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ComicClient) UpdateOneID(id uuid.UUID) *ComicUpdateOne {
	mutation := newComicMutation(c.config, OpUpdateOne, withComicID(id))
	return &ComicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comic.
func (c *ComicClient) Delete() *ComicDelete {
	mutation := newComicMutation(c.config, OpDelete)
	return &ComicDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ComicClient) DeleteOne(co *Comic) *ComicDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ComicClient) DeleteOneID(id uuid.UUID) *ComicDeleteOne {
	builder := c.Delete().Where(comic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ComicDeleteOne{builder}
}

// Query returns a query builder for Comic.
func (c *ComicClient) Query() *ComicQuery {
	return &ComicQuery{
		config: c.config,
	}
}

// Get returns a Comic entity by its id.
func (c *ComicClient) Get(ctx context.Context, id uuid.UUID) (*Comic, error) {
	return c.Query().Where(comic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ComicClient) GetX(ctx context.Context, id uuid.UUID) *Comic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Comic.
func (c *ComicClient) QueryAuthor(co *Comic) *AuthorQuery {
	query := &AuthorQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comic.Table, comic.FieldID, id),
			sqlgraph.To(author.Table, author.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comic.AuthorTable, comic.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMagazine queries the magazine edge of a Comic.
func (c *ComicClient) QueryMagazine(co *Comic) *MagazineQuery {
	query := &MagazineQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comic.Table, comic.FieldID, id),
			sqlgraph.To(magazine.Table, magazine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comic.MagazineTable, comic.MagazineColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTags queries the tags edge of a Comic.
func (c *ComicClient) QueryTags(co *Comic) *TagQuery {
	query := &TagQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comic.Table, comic.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, comic.TagsTable, comic.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEpisodes queries the episodes edge of a Comic.
func (c *ComicClient) QueryEpisodes(co *Comic) *EpisodeQuery {
	query := &EpisodeQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comic.Table, comic.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, comic.EpisodesTable, comic.EpisodesColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ComicClient) Hooks() []Hook {
	return c.hooks.Comic
}

// EpisodeClient is a client for the Episode schema.
type EpisodeClient struct {
	config
}

// NewEpisodeClient returns a client for the Episode from the given config.
func NewEpisodeClient(c config) *EpisodeClient {
	return &EpisodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `episode.Hooks(f(g(h())))`.
func (c *EpisodeClient) Use(hooks ...Hook) {
	c.hooks.Episode = append(c.hooks.Episode, hooks...)
}

// Create returns a builder for creating a Episode entity.
func (c *EpisodeClient) Create() *EpisodeCreate {
	mutation := newEpisodeMutation(c.config, OpCreate)
	return &EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Episode entities.
func (c *EpisodeClient) CreateBulk(builders ...*EpisodeCreate) *EpisodeCreateBulk {
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Episode.
func (c *EpisodeClient) Update() *EpisodeUpdate {
	mutation := newEpisodeMutation(c.config, OpUpdate)
	return &EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EpisodeClient) UpdateOne(e *Episode) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisode(e))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EpisodeClient) UpdateOneID(id uuid.UUID) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisodeID(id))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Episode.
func (c *EpisodeClient) Delete() *EpisodeDelete {
	mutation := newEpisodeMutation(c.config, OpDelete)
	return &EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EpisodeClient) DeleteOne(e *Episode) *EpisodeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EpisodeClient) DeleteOneID(id uuid.UUID) *EpisodeDeleteOne {
	builder := c.Delete().Where(episode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EpisodeDeleteOne{builder}
}

// Query returns a query builder for Episode.
func (c *EpisodeClient) Query() *EpisodeQuery {
	return &EpisodeQuery{
		config: c.config,
	}
}

// Get returns a Episode entity by its id.
func (c *EpisodeClient) Get(ctx context.Context, id uuid.UUID) (*Episode, error) {
	return c.Query().Where(episode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EpisodeClient) GetX(ctx context.Context, id uuid.UUID) *Episode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComic queries the comic edge of a Episode.
func (c *EpisodeClient) QueryComic(e *Episode) *ComicQuery {
	query := &ComicQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(comic.Table, comic.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, episode.ComicTable, episode.ComicColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EpisodeClient) Hooks() []Hook {
	return c.hooks.Episode
}

// MagazineClient is a client for the Magazine schema.
type MagazineClient struct {
	config
}

// NewMagazineClient returns a client for the Magazine from the given config.
func NewMagazineClient(c config) *MagazineClient {
	return &MagazineClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `magazine.Hooks(f(g(h())))`.
func (c *MagazineClient) Use(hooks ...Hook) {
	c.hooks.Magazine = append(c.hooks.Magazine, hooks...)
}

// Create returns a builder for creating a Magazine entity.
func (c *MagazineClient) Create() *MagazineCreate {
	mutation := newMagazineMutation(c.config, OpCreate)
	return &MagazineCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Magazine entities.
func (c *MagazineClient) CreateBulk(builders ...*MagazineCreate) *MagazineCreateBulk {
	return &MagazineCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Magazine.
func (c *MagazineClient) Update() *MagazineUpdate {
	mutation := newMagazineMutation(c.config, OpUpdate)
	return &MagazineUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MagazineClient) UpdateOne(m *Magazine) *MagazineUpdateOne {
	mutation := newMagazineMutation(c.config, OpUpdateOne, withMagazine(m))
	return &MagazineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MagazineClient) UpdateOneID(id uuid.UUID) *MagazineUpdateOne {
	mutation := newMagazineMutation(c.config, OpUpdateOne, withMagazineID(id))
	return &MagazineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Magazine.
func (c *MagazineClient) Delete() *MagazineDelete {
	mutation := newMagazineMutation(c.config, OpDelete)
	return &MagazineDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MagazineClient) DeleteOne(m *Magazine) *MagazineDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MagazineClient) DeleteOneID(id uuid.UUID) *MagazineDeleteOne {
	builder := c.Delete().Where(magazine.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MagazineDeleteOne{builder}
}

// Query returns a query builder for Magazine.
func (c *MagazineClient) Query() *MagazineQuery {
	return &MagazineQuery{
		config: c.config,
	}
}

// Get returns a Magazine entity by its id.
func (c *MagazineClient) Get(ctx context.Context, id uuid.UUID) (*Magazine, error) {
	return c.Query().Where(magazine.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MagazineClient) GetX(ctx context.Context, id uuid.UUID) *Magazine {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComics queries the comics edge of a Magazine.
func (c *MagazineClient) QueryComics(m *Magazine) *ComicQuery {
	query := &ComicQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(magazine.Table, magazine.FieldID, id),
			sqlgraph.To(comic.Table, comic.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, magazine.ComicsTable, magazine.ComicsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MagazineClient) Hooks() []Hook {
	return c.hooks.Magazine
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Create returns a builder for creating a Tag entity.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id uuid.UUID) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TagClient) DeleteOneID(id uuid.UUID) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id uuid.UUID) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id uuid.UUID) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComics queries the comics edge of a Tag.
func (c *TagClient) QueryComics(t *Tag) *ComicQuery {
	query := &ComicQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(comic.Table, comic.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.ComicsTable, tag.ComicsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
}
