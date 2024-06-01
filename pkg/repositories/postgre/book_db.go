package postgre

import (
	"bookstore/models"
)

type bookRepository struct{}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (b bookRepository) CreateBook(book *models.Book) (int, error) {
	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0, err
	}
	defer conn.Release()

	var query = `INSERT INTO books (title, author_id, year, isbn) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(ctx, query, book.Title, book.AuthorID, book.Year, book.ISBN).Scan(&book.ID)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0, err
	}

	return book.ID, nil
}

func (b bookRepository) GetAllBooks() ([]*models.BookAuthor, error) {
	var books []*models.BookAuthor

	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		return books, err
	}
	defer conn.Release()

	query := `SELECT b.id, b.title, a.id, CONCAT(a.first_name, ' ', a.last_name) as author, b.year, b.isbn FROM public.books as b INNER JOIN public.authors as a on b.author_id = a.id`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.BookAuthor
		err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Author, &book.Year, &book.ISBN)
		if err != nil {
			return books, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (b bookRepository) GetBookByID(id int) (*models.Book, error) {
	var book models.Book

	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return &book, err
	}
	defer conn.Release()

	query := `SELECT id, title, author_id, year, isbn FROM books WHERE id=$1`
	err = conn.QueryRow(ctx, query, id).Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN)
	if err != nil {
		return &book, err
	}

	return &book, nil
}

func (b bookRepository) UpdateBook(book *models.Book) error {
	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	query := "UPDATE books SET title=$1, author_id=$2, year=$3, isbn=$4 WHERE id=$5"
	_, err = conn.Exec(ctx, query, book.Title, book.AuthorID, book.Year, book.ISBN, book.ID)
	if err != nil {
		return err
	}

	return nil
}

func (b bookRepository) DeleteBook(id int) error {
	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	query := "DELETE FROM books WHERE id=$1"
	_, err = conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
