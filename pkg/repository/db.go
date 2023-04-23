package repository

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	users            = "users"
	notes            = "notes"
	zeroRowsAffected = "this note does not exist"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSL      string
}

func InitDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
