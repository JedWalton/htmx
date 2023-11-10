package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq" // This is the PostgreSQL driver for database/sql
)

var (
	once        sync.Once
	instance    *sql.DB
	instanceErr error
)

type PostgreSQL struct {
	db *sql.DB
}

func NewPostgreSQL() (*sql.DB, error) {
	once.Do(func() {
		postgresqlURL := os.Getenv("POSTGRESQL_URL")
		if postgresqlURL == "" {
			instanceErr = fmt.Errorf("POSTGRESQL_URL environment variable is not set")
			return
		}

		var db *sql.DB
		db, instanceErr = sql.Open("postgres", postgresqlURL)
		if instanceErr != nil {
			return
		}

		if instanceErr = db.Ping(); instanceErr != nil {
			return
		}

		instance = db
	})

	return instance, instanceErr
}
