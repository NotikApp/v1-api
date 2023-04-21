package repository

import (
	"github.com/gavrylenkoIvan/gonotes"
	"github.com/jmoiron/sqlx"
)

type Notes interface {
	GetNotesByUser(id int) ([]gonotes.Note, error)
}

type Users interface {
	GetUser(email string, hash string) (gonotes.User, error)
	CreateUser(input gonotes.User) (int, error)
}

type Repository struct {
	Notes
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUserRepo(db),
	}
}
