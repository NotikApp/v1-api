package service

import (
	"github.com/gavrylenkoIvan/gonotes"
	"github.com/gavrylenkoIvan/gonotes/pkg/repository"
)

type NotesService struct {
	repo repository.Notes
}

func NewNotesService(repo repository.Notes) *NotesService {
	return &NotesService{
		repo: repo,
	}
}

func (s *NotesService) GetNotesByUser(id int) ([]gonotes.Note, error) {
	return s.repo.GetNotesByUser(id)
}

func (s *NotesService) DeleteNote(id int, userId int) error {
	return s.repo.DeleteNote(id, userId)
}

func (s *NotesService) UpdateNote(id int, userId int, input gonotes.UpdateNoteStruct) error {
	return s.repo.UpdateNote(id, userId, input)
}

func (s *NotesService) CreateNote(userId int, input gonotes.Note) (int, error) {
	return s.repo.CreateNote(userId, input)
}
