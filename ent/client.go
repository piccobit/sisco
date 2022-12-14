// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"sisco/ent/migrate"

	"sisco/ent/area"
	"sisco/ent/service"
	"sisco/ent/tag"
	"sisco/ent/token"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Area is the client for interacting with the Area builders.
	Area *AreaClient
	// Service is the client for interacting with the Service builders.
	Service *ServiceClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
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
	c.Area = NewAreaClient(c.config)
	c.Service = NewServiceClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.Token = NewTokenClient(c.config)
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
		ctx:     ctx,
		config:  cfg,
		Area:    NewAreaClient(cfg),
		Service: NewServiceClient(cfg),
		Tag:     NewTagClient(cfg),
		Token:   NewTokenClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Area:    NewAreaClient(cfg),
		Service: NewServiceClient(cfg),
		Tag:     NewTagClient(cfg),
		Token:   NewTokenClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Area.
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
	c.Area.Use(hooks...)
	c.Service.Use(hooks...)
	c.Tag.Use(hooks...)
	c.Token.Use(hooks...)
}

// AreaClient is a client for the Area schema.
type AreaClient struct {
	config
}

// NewAreaClient returns a client for the Area from the given config.
func NewAreaClient(c config) *AreaClient {
	return &AreaClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `area.Hooks(f(g(h())))`.
func (c *AreaClient) Use(hooks ...Hook) {
	c.hooks.Area = append(c.hooks.Area, hooks...)
}

// Create returns a builder for creating a Area entity.
func (c *AreaClient) Create() *AreaCreate {
	mutation := newAreaMutation(c.config, OpCreate)
	return &AreaCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Area entities.
func (c *AreaClient) CreateBulk(builders ...*AreaCreate) *AreaCreateBulk {
	return &AreaCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Area.
func (c *AreaClient) Update() *AreaUpdate {
	mutation := newAreaMutation(c.config, OpUpdate)
	return &AreaUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AreaClient) UpdateOne(a *Area) *AreaUpdateOne {
	mutation := newAreaMutation(c.config, OpUpdateOne, withArea(a))
	return &AreaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AreaClient) UpdateOneID(id int) *AreaUpdateOne {
	mutation := newAreaMutation(c.config, OpUpdateOne, withAreaID(id))
	return &AreaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Area.
func (c *AreaClient) Delete() *AreaDelete {
	mutation := newAreaMutation(c.config, OpDelete)
	return &AreaDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AreaClient) DeleteOne(a *Area) *AreaDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AreaClient) DeleteOneID(id int) *AreaDeleteOne {
	builder := c.Delete().Where(area.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AreaDeleteOne{builder}
}

// Query returns a query builder for Area.
func (c *AreaClient) Query() *AreaQuery {
	return &AreaQuery{
		config: c.config,
	}
}

// Get returns a Area entity by its id.
func (c *AreaClient) Get(ctx context.Context, id int) (*Area, error) {
	return c.Query().Where(area.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AreaClient) GetX(ctx context.Context, id int) *Area {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryServices queries the services edge of a Area.
func (c *AreaClient) QueryServices(a *Area) *ServiceQuery {
	query := &ServiceQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(area.Table, area.FieldID, id),
			sqlgraph.To(service.Table, service.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, area.ServicesTable, area.ServicesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AreaClient) Hooks() []Hook {
	return c.hooks.Area
}

// ServiceClient is a client for the Service schema.
type ServiceClient struct {
	config
}

// NewServiceClient returns a client for the Service from the given config.
func NewServiceClient(c config) *ServiceClient {
	return &ServiceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `service.Hooks(f(g(h())))`.
func (c *ServiceClient) Use(hooks ...Hook) {
	c.hooks.Service = append(c.hooks.Service, hooks...)
}

// Create returns a builder for creating a Service entity.
func (c *ServiceClient) Create() *ServiceCreate {
	mutation := newServiceMutation(c.config, OpCreate)
	return &ServiceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Service entities.
func (c *ServiceClient) CreateBulk(builders ...*ServiceCreate) *ServiceCreateBulk {
	return &ServiceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Service.
func (c *ServiceClient) Update() *ServiceUpdate {
	mutation := newServiceMutation(c.config, OpUpdate)
	return &ServiceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServiceClient) UpdateOne(s *Service) *ServiceUpdateOne {
	mutation := newServiceMutation(c.config, OpUpdateOne, withService(s))
	return &ServiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServiceClient) UpdateOneID(id int) *ServiceUpdateOne {
	mutation := newServiceMutation(c.config, OpUpdateOne, withServiceID(id))
	return &ServiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Service.
func (c *ServiceClient) Delete() *ServiceDelete {
	mutation := newServiceMutation(c.config, OpDelete)
	return &ServiceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServiceClient) DeleteOne(s *Service) *ServiceDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServiceClient) DeleteOneID(id int) *ServiceDeleteOne {
	builder := c.Delete().Where(service.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServiceDeleteOne{builder}
}

// Query returns a query builder for Service.
func (c *ServiceClient) Query() *ServiceQuery {
	return &ServiceQuery{
		config: c.config,
	}
}

// Get returns a Service entity by its id.
func (c *ServiceClient) Get(ctx context.Context, id int) (*Service, error) {
	return c.Query().Where(service.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServiceClient) GetX(ctx context.Context, id int) *Service {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTags queries the tags edge of a Service.
func (c *ServiceClient) QueryTags(s *Service) *TagQuery {
	query := &TagQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(service.Table, service.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, service.TagsTable, service.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryArea queries the area edge of a Service.
func (c *ServiceClient) QueryArea(s *Service) *AreaQuery {
	query := &AreaQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(service.Table, service.FieldID, id),
			sqlgraph.To(area.Table, area.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, service.AreaTable, service.AreaColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServiceClient) Hooks() []Hook {
	return c.hooks.Service
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
func (c *TagClient) UpdateOneID(id int) *TagUpdateOne {
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
func (c *TagClient) DeleteOneID(id int) *TagDeleteOne {
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
func (c *TagClient) Get(ctx context.Context, id int) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id int) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryService queries the service edge of a Tag.
func (c *TagClient) QueryService(t *Tag) *ServiceQuery {
	query := &ServiceQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(service.Table, service.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.ServiceTable, tag.ServicePrimaryKey...),
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

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Create returns a builder for creating a Token entity.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id int) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TokenClient) DeleteOneID(id int) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id int) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id int) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	return c.hooks.Token
}
