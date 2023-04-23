package repository

import (
	"github.com/gavrylenkoIvan/gonotes"
	"github.com/jmoiron/sqlx"
)

type Notes interface {
	GetNotesByUser(id int) ([]gonotes.Note, error)
	DeleteNote(id int, userId int) error
	UpdateNote(id int, userId int, input gonotes.UpdateNoteStruct) error
	CreateNote(userId int, input gonotes.Note) (gonotes.Note, error)
}

type Users interface {
	GetUser(email string, hash string) (gonotes.User, error)
	CreateUser(input gonotes.SignUpInput, code string) (int, error)
	VerifyUser(userId int, code string) error
	DeleteUser(userId int, code string) error
}

type Repository struct {
	Notes
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUserRepo(db),
		Notes: NewNotesRepo(db),
	}
}
