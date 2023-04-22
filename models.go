package gonotes

import "errors"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Verified bool   `json:"verified"`
}

type Note struct {
	ID        int    `json:"id"`
	Title     string `json:"title" binding:"required"`
	Text      string `json:"text" binding:"required"`
	Important bool   `json:"important" binding:"required"`
	Tags      string `json:"tags" binding:"required"`
	UserId    int    `json:"user_id" db:"user_id"`
}

type UpdateNoteStruct struct {
	Title     *string `json:"title"`
	Text      *string `json:"text"`
	Important *bool   `json:"important"`
	Tags      *string `json:"tags"`
}

func (u UpdateNoteStruct) Validate() error {
	if u.Title == nil && u.Important == nil && u.Text == nil && u.Tags == nil {
		return errors.New("empty update struct")
	}

	return nil
}
