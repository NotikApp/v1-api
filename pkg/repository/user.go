package repository

import (
	"fmt"

	"github.com/gavrylenkoIvan/gonotes"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user gonotes.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password) values ($1, $2, $3) RETURNING id", users)
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetUser(email, password string) (gonotes.User, error) {
	var user gonotes.User
	query := fmt.Sprintf("SELECT * FROM %s u WHERE u.password = $1 AND u.email = $2", users)
	err := r.db.Get(&user, query, password, email)

	return user, err
}
