package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/nikitsingh/forky/backend/internal/config"

	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", config.Envs.DB_URL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
