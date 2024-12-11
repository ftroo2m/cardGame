// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"cardGame/ent/migrate"

	"cardGame/ent/card"
	"cardGame/ent/monster"
	"cardGame/ent/user"
	"cardGame/ent/userconfig"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Card is the client for interacting with the Card builders.
	Card *CardClient
	// Monster is the client for interacting with the Monster builders.
	Monster *MonsterClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserConfig is the client for interacting with the UserConfig builders.
	UserConfig *UserConfigClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Card = NewCardClient(c.config)
	c.Monster = NewMonsterClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserConfig = NewUserConfigClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
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
		Card:       NewCardClient(cfg),
		Monster:    NewMonsterClient(cfg),
		User:       NewUserClient(cfg),
		UserConfig: NewUserConfigClient(cfg),
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
		Card:       NewCardClient(cfg),
		Monster:    NewMonsterClient(cfg),
		User:       NewUserClient(cfg),
		UserConfig: NewUserConfigClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Card.
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
	c.Card.Use(hooks...)
	c.Monster.Use(hooks...)
	c.User.Use(hooks...)
	c.UserConfig.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Card.Intercept(interceptors...)
	c.Monster.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.UserConfig.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CardMutation:
		return c.Card.mutate(ctx, m)
	case *MonsterMutation:
		return c.Monster.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *UserConfigMutation:
		return c.UserConfig.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CardClient is a client for the Card schema.
type CardClient struct {
	config
}

// NewCardClient returns a client for the Card from the given config.
func NewCardClient(c config) *CardClient {
	return &CardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `card.Hooks(f(g(h())))`.
func (c *CardClient) Use(hooks ...Hook) {
	c.hooks.Card = append(c.hooks.Card, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `card.Intercept(f(g(h())))`.
func (c *CardClient) Intercept(interceptors ...Interceptor) {
	c.inters.Card = append(c.inters.Card, interceptors...)
}

// Create returns a builder for creating a Card entity.
func (c *CardClient) Create() *CardCreate {
	mutation := newCardMutation(c.config, OpCreate)
	return &CardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Card entities.
func (c *CardClient) CreateBulk(builders ...*CardCreate) *CardCreateBulk {
	return &CardCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CardClient) MapCreateBulk(slice any, setFunc func(*CardCreate, int)) *CardCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CardCreateBulk{err: fmt.Errorf("calling to CardClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CardCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Card.
func (c *CardClient) Update() *CardUpdate {
	mutation := newCardMutation(c.config, OpUpdate)
	return &CardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CardClient) UpdateOne(ca *Card) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCard(ca))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CardClient) UpdateOneID(id int) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCardID(id))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Card.
func (c *CardClient) Delete() *CardDelete {
	mutation := newCardMutation(c.config, OpDelete)
	return &CardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CardClient) DeleteOne(ca *Card) *CardDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CardClient) DeleteOneID(id int) *CardDeleteOne {
	builder := c.Delete().Where(card.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CardDeleteOne{builder}
}

// Query returns a query builder for Card.
func (c *CardClient) Query() *CardQuery {
	return &CardQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCard},
		inters: c.Interceptors(),
	}
}

// Get returns a Card entity by its id.
func (c *CardClient) Get(ctx context.Context, id int) (*Card, error) {
	return c.Query().Where(card.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardClient) GetX(ctx context.Context, id int) *Card {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CardClient) Hooks() []Hook {
	return c.hooks.Card
}

// Interceptors returns the client interceptors.
func (c *CardClient) Interceptors() []Interceptor {
	return c.inters.Card
}

func (c *CardClient) mutate(ctx context.Context, m *CardMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CardCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CardUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CardDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Card mutation op: %q", m.Op())
	}
}

// MonsterClient is a client for the Monster schema.
type MonsterClient struct {
	config
}

// NewMonsterClient returns a client for the Monster from the given config.
func NewMonsterClient(c config) *MonsterClient {
	return &MonsterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `monster.Hooks(f(g(h())))`.
func (c *MonsterClient) Use(hooks ...Hook) {
	c.hooks.Monster = append(c.hooks.Monster, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `monster.Intercept(f(g(h())))`.
func (c *MonsterClient) Intercept(interceptors ...Interceptor) {
	c.inters.Monster = append(c.inters.Monster, interceptors...)
}

// Create returns a builder for creating a Monster entity.
func (c *MonsterClient) Create() *MonsterCreate {
	mutation := newMonsterMutation(c.config, OpCreate)
	return &MonsterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Monster entities.
func (c *MonsterClient) CreateBulk(builders ...*MonsterCreate) *MonsterCreateBulk {
	return &MonsterCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MonsterClient) MapCreateBulk(slice any, setFunc func(*MonsterCreate, int)) *MonsterCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MonsterCreateBulk{err: fmt.Errorf("calling to MonsterClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MonsterCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MonsterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Monster.
func (c *MonsterClient) Update() *MonsterUpdate {
	mutation := newMonsterMutation(c.config, OpUpdate)
	return &MonsterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MonsterClient) UpdateOne(m *Monster) *MonsterUpdateOne {
	mutation := newMonsterMutation(c.config, OpUpdateOne, withMonster(m))
	return &MonsterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MonsterClient) UpdateOneID(id int) *MonsterUpdateOne {
	mutation := newMonsterMutation(c.config, OpUpdateOne, withMonsterID(id))
	return &MonsterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Monster.
func (c *MonsterClient) Delete() *MonsterDelete {
	mutation := newMonsterMutation(c.config, OpDelete)
	return &MonsterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MonsterClient) DeleteOne(m *Monster) *MonsterDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MonsterClient) DeleteOneID(id int) *MonsterDeleteOne {
	builder := c.Delete().Where(monster.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MonsterDeleteOne{builder}
}

// Query returns a query builder for Monster.
func (c *MonsterClient) Query() *MonsterQuery {
	return &MonsterQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMonster},
		inters: c.Interceptors(),
	}
}

// Get returns a Monster entity by its id.
func (c *MonsterClient) Get(ctx context.Context, id int) (*Monster, error) {
	return c.Query().Where(monster.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MonsterClient) GetX(ctx context.Context, id int) *Monster {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MonsterClient) Hooks() []Hook {
	return c.hooks.Monster
}

// Interceptors returns the client interceptors.
func (c *MonsterClient) Interceptors() []Interceptor {
	return c.inters.Monster
}

func (c *MonsterClient) mutate(ctx context.Context, m *MonsterMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MonsterCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MonsterUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MonsterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MonsterDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Monster mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// UserConfigClient is a client for the UserConfig schema.
type UserConfigClient struct {
	config
}

// NewUserConfigClient returns a client for the UserConfig from the given config.
func NewUserConfigClient(c config) *UserConfigClient {
	return &UserConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userconfig.Hooks(f(g(h())))`.
func (c *UserConfigClient) Use(hooks ...Hook) {
	c.hooks.UserConfig = append(c.hooks.UserConfig, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `userconfig.Intercept(f(g(h())))`.
func (c *UserConfigClient) Intercept(interceptors ...Interceptor) {
	c.inters.UserConfig = append(c.inters.UserConfig, interceptors...)
}

// Create returns a builder for creating a UserConfig entity.
func (c *UserConfigClient) Create() *UserConfigCreate {
	mutation := newUserConfigMutation(c.config, OpCreate)
	return &UserConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserConfig entities.
func (c *UserConfigClient) CreateBulk(builders ...*UserConfigCreate) *UserConfigCreateBulk {
	return &UserConfigCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserConfigClient) MapCreateBulk(slice any, setFunc func(*UserConfigCreate, int)) *UserConfigCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserConfigCreateBulk{err: fmt.Errorf("calling to UserConfigClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserConfigCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserConfig.
func (c *UserConfigClient) Update() *UserConfigUpdate {
	mutation := newUserConfigMutation(c.config, OpUpdate)
	return &UserConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserConfigClient) UpdateOne(uc *UserConfig) *UserConfigUpdateOne {
	mutation := newUserConfigMutation(c.config, OpUpdateOne, withUserConfig(uc))
	return &UserConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserConfigClient) UpdateOneID(id int) *UserConfigUpdateOne {
	mutation := newUserConfigMutation(c.config, OpUpdateOne, withUserConfigID(id))
	return &UserConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserConfig.
func (c *UserConfigClient) Delete() *UserConfigDelete {
	mutation := newUserConfigMutation(c.config, OpDelete)
	return &UserConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserConfigClient) DeleteOne(uc *UserConfig) *UserConfigDeleteOne {
	return c.DeleteOneID(uc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserConfigClient) DeleteOneID(id int) *UserConfigDeleteOne {
	builder := c.Delete().Where(userconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserConfigDeleteOne{builder}
}

// Query returns a query builder for UserConfig.
func (c *UserConfigClient) Query() *UserConfigQuery {
	return &UserConfigQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUserConfig},
		inters: c.Interceptors(),
	}
}

// Get returns a UserConfig entity by its id.
func (c *UserConfigClient) Get(ctx context.Context, id int) (*UserConfig, error) {
	return c.Query().Where(userconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserConfigClient) GetX(ctx context.Context, id int) *UserConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserConfigClient) Hooks() []Hook {
	return c.hooks.UserConfig
}

// Interceptors returns the client interceptors.
func (c *UserConfigClient) Interceptors() []Interceptor {
	return c.inters.UserConfig
}

func (c *UserConfigClient) mutate(ctx context.Context, m *UserConfigMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserConfigCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserConfigDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown UserConfig mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Card, Monster, User, UserConfig []ent.Hook
	}
	inters struct {
		Card, Monster, User, UserConfig []ent.Interceptor
	}
)
