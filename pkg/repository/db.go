package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	users = "users"
	notes = "notes"
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSL)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
