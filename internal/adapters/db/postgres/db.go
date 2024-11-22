// internal/application/db/postgres/postgres.go
package postgres

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/ports/clients"
	"context"
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"time"
)

type PostgresDB struct {
	Db     *sql.DB
	Config config.DatabaseConfig
}

func NewPostgresDB(config config.DatabaseConfig) clients.Database {
	return &PostgresDB{
		Config: config,
	}
}

func (p *PostgresDB) Connect() error {
	dsn := utils.GeneratePostgresConnectionString(p.Config)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(p.Config.MaxOpenConns) // Max number of open connections
	db.SetMaxIdleConns(p.Config.MaxIdleConns) // Max number of idle connections
	db.SetConnMaxLifetime(30 * time.Minute)   // Max lifetime of a connection
	p.Db = db
	return p.Ping()
}

func (p *PostgresDB) Close() error {
	return p.Db.Close()
}

func (p *PostgresDB) Ping() error {
	return p.Db.Ping()
}

// Query executes a query with context
func (p *PostgresDB) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.Db.QueryContext(ctx, query, args...)
}

// QueryRow executes a query and returns a single row with context
func (p *PostgresDB) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.Db.QueryRowContext(ctx, query, args...)
}

// Exec executes a query that doesn’t return rows with context
func (p *PostgresDB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.Db.ExecContext(ctx, query, args...)
}

// Prepare prepares a query for later execution with context
func (p *PostgresDB) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	return p.Db.PrepareContext(ctx, query)
}
