package repository

import (
	"errors"
	"fmt"
	"strings"

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

func (r *NotesRepo) DeleteNote(id int, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND user_id = $2", notes)
	exec, err := r.db.Exec(query, id, userId)
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

func (r *NotesRepo) UpdateNote(id int, userId int, input gonotes.UpdateNoteStruct) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Important != nil {
		setValues = append(setValues, fmt.Sprintf("important=$%d", argId))
		args = append(args, *input.Important)
		argId++
	}

	if input.Tags != nil {
		setValues = append(setValues, fmt.Sprintf("tags=$%d", argId))
		args = append(args, *input.Tags)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = $%d AND id = $%d", notes, setQuery, argId, argId+1)

	args = append(args, userId, id)

	res, err := r.db.Exec(query, args...)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if rows == 0 {
		return errors.New(zeroRowsAffected)
	}

	return err
}

func (r *NotesRepo) CreateNote(userId int, input gonotes.Note) (gonotes.Note, error) {
	var created gonotes.Note
	fmt.Println(userId, input)
	query := fmt.Sprintf("INSERT INTO %s (title, text, important, tags, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, text, important, tags, user_id", notes)
	row := r.db.QueryRow(query, input.Title, input.Text, input.Important, input.Tags, userId)
	err := row.Scan(&created.ID, &created.Title,
		&created.Text, &created.Important, &created.Tags,
		&created.UserId)

	if err != nil {
		return gonotes.Note{}, err
	}

	return created, nil
}
