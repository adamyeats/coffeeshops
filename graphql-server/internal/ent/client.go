// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/ent/migrate"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/ent/coffeeshop"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Coffeeshop is the client for interacting with the Coffeeshop builders.
	Coffeeshop *CoffeeshopClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Coffeeshop = NewCoffeeshopClient(c.config)
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
		ctx:        ctx,
		config:     cfg,
		Coffeeshop: NewCoffeeshopClient(cfg),
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
		ctx:        ctx,
		config:     cfg,
		Coffeeshop: NewCoffeeshopClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Coffeeshop.
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
	c.Coffeeshop.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Coffeeshop.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CoffeeshopMutation:
		return c.Coffeeshop.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CoffeeshopClient is a client for the Coffeeshop schema.
type CoffeeshopClient struct {
	config
}

// NewCoffeeshopClient returns a client for the Coffeeshop from the given config.
func NewCoffeeshopClient(c config) *CoffeeshopClient {
	return &CoffeeshopClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coffeeshop.Hooks(f(g(h())))`.
func (c *CoffeeshopClient) Use(hooks ...Hook) {
	c.hooks.Coffeeshop = append(c.hooks.Coffeeshop, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `coffeeshop.Intercept(f(g(h())))`.
func (c *CoffeeshopClient) Intercept(interceptors ...Interceptor) {
	c.inters.Coffeeshop = append(c.inters.Coffeeshop, interceptors...)
}

// Create returns a builder for creating a Coffeeshop entity.
func (c *CoffeeshopClient) Create() *CoffeeshopCreate {
	mutation := newCoffeeshopMutation(c.config, OpCreate)
	return &CoffeeshopCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Coffeeshop entities.
func (c *CoffeeshopClient) CreateBulk(builders ...*CoffeeshopCreate) *CoffeeshopCreateBulk {
	return &CoffeeshopCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Coffeeshop.
func (c *CoffeeshopClient) Update() *CoffeeshopUpdate {
	mutation := newCoffeeshopMutation(c.config, OpUpdate)
	return &CoffeeshopUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CoffeeshopClient) UpdateOne(co *Coffeeshop) *CoffeeshopUpdateOne {
	mutation := newCoffeeshopMutation(c.config, OpUpdateOne, withCoffeeshop(co))
	return &CoffeeshopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CoffeeshopClient) UpdateOneID(id int) *CoffeeshopUpdateOne {
	mutation := newCoffeeshopMutation(c.config, OpUpdateOne, withCoffeeshopID(id))
	return &CoffeeshopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Coffeeshop.
func (c *CoffeeshopClient) Delete() *CoffeeshopDelete {
	mutation := newCoffeeshopMutation(c.config, OpDelete)
	return &CoffeeshopDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CoffeeshopClient) DeleteOne(co *Coffeeshop) *CoffeeshopDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CoffeeshopClient) DeleteOneID(id int) *CoffeeshopDeleteOne {
	builder := c.Delete().Where(coffeeshop.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CoffeeshopDeleteOne{builder}
}

// Query returns a query builder for Coffeeshop.
func (c *CoffeeshopClient) Query() *CoffeeshopQuery {
	return &CoffeeshopQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCoffeeshop},
		inters: c.Interceptors(),
	}
}

// Get returns a Coffeeshop entity by its id.
func (c *CoffeeshopClient) Get(ctx context.Context, id int) (*Coffeeshop, error) {
	return c.Query().Where(coffeeshop.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CoffeeshopClient) GetX(ctx context.Context, id int) *Coffeeshop {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CoffeeshopClient) Hooks() []Hook {
	return c.hooks.Coffeeshop
}

// Interceptors returns the client interceptors.
func (c *CoffeeshopClient) Interceptors() []Interceptor {
	return c.inters.Coffeeshop
}

func (c *CoffeeshopClient) mutate(ctx context.Context, m *CoffeeshopMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CoffeeshopCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CoffeeshopUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CoffeeshopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CoffeeshopDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Coffeeshop mutation op: %q", m.Op())
	}
}
