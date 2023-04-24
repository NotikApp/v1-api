package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	users            = "users"
	notes            = "notes"
	zeroRowsAffected = "this note does not exist"
)

func InitDB(dblink string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dblink)
	if err != nil {
		return nil, err
	}

	return db, nil
}
