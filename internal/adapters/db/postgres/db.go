package postgres

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/utils"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
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
		logrus.Info("db connection failed", err)
	}
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(2)

	if pingErr := db.Ping(); pingErr != nil {
		logrus.Info("Err postgres ping", pingErr)
	} else {
		logrus.Info("Success postgres connection is ok")
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
