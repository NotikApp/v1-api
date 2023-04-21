package service

import (
	"github.com/gavrylenkoIvan/gonotes"
	"github.com/gavrylenkoIvan/gonotes/pkg/repository"
)

type Authorization interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	CreateUser(input gonotes.User) (int, error)
}

type Notes interface{}

type Users interface{}

type Service struct {
	Authorization
	Notes
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Users),
	}
}
