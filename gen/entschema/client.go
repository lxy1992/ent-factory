// Code generated by ent, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/zaihui/ent-factory/gen/entschema/migrate"

	"github.com/zaihui/ent-factory/gen/entschema/test"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Test is the client for interacting with the Test builders.
	Test *TestClient
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
	c.Test = NewTestClient(c.config)
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
		return nil, errors.New("entschema: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("entschema: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Test:   NewTestClient(cfg),
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
		ctx:    ctx,
		config: cfg,
		Test:   NewTestClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Test.
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
	c.Test.Use(hooks...)
}

// TestClient is a client for the Test schema.
type TestClient struct {
	config
}

// NewTestClient returns a client for the Test from the given config.
func NewTestClient(c config) *TestClient {
	return &TestClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `test.Hooks(f(g(h())))`.
func (c *TestClient) Use(hooks ...Hook) {
	c.hooks.Test = append(c.hooks.Test, hooks...)
}

// Create returns a builder for creating a Test entity.
func (c *TestClient) Create() *TestCreate {
	mutation := newTestMutation(c.config, OpCreate)
	return &TestCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Test entities.
func (c *TestClient) CreateBulk(builders ...*TestCreate) *TestCreateBulk {
	return &TestCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Test.
func (c *TestClient) Update() *TestUpdate {
	mutation := newTestMutation(c.config, OpUpdate)
	return &TestUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TestClient) UpdateOne(t *Test) *TestUpdateOne {
	mutation := newTestMutation(c.config, OpUpdateOne, withTest(t))
	return &TestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TestClient) UpdateOneID(id int) *TestUpdateOne {
	mutation := newTestMutation(c.config, OpUpdateOne, withTestID(id))
	return &TestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Test.
func (c *TestClient) Delete() *TestDelete {
	mutation := newTestMutation(c.config, OpDelete)
	return &TestDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TestClient) DeleteOne(t *Test) *TestDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TestClient) DeleteOneID(id int) *TestDeleteOne {
	builder := c.Delete().Where(test.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TestDeleteOne{builder}
}

// Query returns a query builder for Test.
func (c *TestClient) Query() *TestQuery {
	return &TestQuery{
		config: c.config,
	}
}

// Get returns a Test entity by its id.
func (c *TestClient) Get(ctx context.Context, id int) (*Test, error) {
	return c.Query().Where(test.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TestClient) GetX(ctx context.Context, id int) *Test {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TestClient) Hooks() []Hook {
	return c.hooks.Test
}
