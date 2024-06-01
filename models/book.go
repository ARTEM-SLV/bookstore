package models

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	Year     int    `json:"year"`
	ISBN     string `json:"isbn"`
}

type BookAuthor struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	ISBN     string `json:"isbn"`
}

type BookWithAuthor struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
	ISBN  string `json:"isbn"`
}
