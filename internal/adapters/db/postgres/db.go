package postgres

import (
	"DDD-HEX/config"
	utils "DDD-HEX/internal/application/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Repositories struct {
	Db *sql.DB
	mu sync.Mutex
}

// NewRepositories create new my sql instance for other repositories
func NewRepositories(appConfig config.AppConfig, config config.PostgresConfig) (*Repositories, error) {
	dsn := utils.GeneratePostgresConnectionString(config)
	db, err := sql.Open(appConfig.DbType, dsn)
	if err != nil {
		fmt.Println("db connection failed", err)
	}
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(2)

	if pingErr := db.Ping(); pingErr != nil {
		log.Println("Err postgres ping", pingErr)
	} else {
		log.Println("Success postgres connection is ok")
	}

	return &Repositories{
		Db: db,
	}, nil
}

// Ping pings the database connection
func (mr *Repositories) Ping() error {
	if err := mr.Db.Ping(); err != nil {
		return err
	}

	return nil
}

// DB returns the database connection
func (mr *Repositories) DB() *sql.DB {
	return mr.Db
}

// Close closes the  database connection
func (mr *Repositories) Close() {
	_ = mr.Db.Close()
}
