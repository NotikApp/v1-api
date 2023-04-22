package repository

import (
	"errors"
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

func (r *UserRepo) CreateUser(user gonotes.User, code string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, code) values ($1, $2, $3, $4) RETURNING id", users)
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password, code)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetUser(email, password string) (gonotes.User, error) {
	var user gonotes.User
	query := fmt.Sprintf("SELECT u.password, u.email, u.username, u.verified FROM %s u WHERE u.password = $1 AND u.email = $2", users)
	err := r.db.Get(&user, query, password, email)

	return user, err
}

func (r *UserRepo) VerifyUser(userId int, code string) error {
	query := fmt.Sprintf("UPDATE %s SET verified = true, code='' WHERE code = $1 AND id = $2", users)
	exec, err := r.db.Exec(query, code, userId)
	if err != nil {
		return err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New(zeroRowsAffected)
	}

	return nil
}
