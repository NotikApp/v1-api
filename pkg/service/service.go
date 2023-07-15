package service

import (
	"github.com/gavrylenkoIvan/gonotes"
	"github.com/gavrylenkoIvan/gonotes/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	CreateUser(input gonotes.SignUpInput, code string) (int, error)
	VerifyUser(userId int, code string) error
	DeleteUser(userId int, code string) error
}

type Notes interface {
	GetNotesByUser(id int) ([]gonotes.Note, error)
	DeleteNote(id int, userId int) error
	UpdateNote(id int, userId int, input gonotes.UpdateNoteStruct) error
	CreateNote(userId int, input gonotes.Note) (gonotes.Note, error)
}

type Users interface{}

type Service struct {
	Authorization
	Notes
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Users),
		Notes:         NewNotesService(repo.Notes),
	}
}
