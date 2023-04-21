package gonotes

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	Important bool     `json:"important"`
	Tags      []string `json:"tags"`
}
