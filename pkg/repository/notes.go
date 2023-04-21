package repository

import (
	"fmt"

	"github.com/gavrylenkoIvan/gonotes"
	"github.com/jmoiron/sqlx"
)

type NotesRepo struct {
	db *sqlx.DB
}

func NewNotesRepo(db *sqlx.DB) *NotesRepo {
	return &NotesRepo{
		db: db,
	}
}

func (r *NotesRepo) GetNotesByUser(id int) ([]gonotes.Note, error) {
	var notesArr []gonotes.Note
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", notes)
	err := r.db.Select(&notesArr, query, id)
	if err != nil {
		return nil, err
	}

	return notesArr, nil
}
