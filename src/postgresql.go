package main

import (
	"database/sql"
	"fmt"
	"os"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewPostgreSQL() (*PostgreSQL, error) {
	postgresqlURL := os.Getenv("POSTGRESQL_URL")
	if postgresqlURL == "" {
		return nil, fmt.Errorf("POSTGRESQL_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", postgresqlURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreSQL{db: db}, nil
}
