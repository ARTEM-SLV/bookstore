package postgre

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"bookstore/models"
)

type BookRepositoryPg struct {
	pool *pgxpool.Pool
}

func NewBookRepository(pool *pgxpool.Pool) *BookRepositoryPg {
	return &BookRepositoryPg{pool: pool}
}

func (b *BookRepositoryPg) CreateBook(ctx context.Context, book *models.Book) (int, error) {
	conn, err := b.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var query = `INSERT INTO books (title, author_id, year, isbn) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(ctx, query, book.Title, book.AuthorID, book.Year, book.ISBN).Scan(&book.ID)
	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (b *BookRepositoryPg) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book

	conn, err := b.pool.Acquire(ctx)
	if err != nil {
		return books, err
	}
	defer conn.Release()

	query := `SELECT id, title, author_id, year, isbn FROM books`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN)
		if err != nil {
			return books, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (b *BookRepositoryPg) GetBookByID(ctx context.Context, id int) (*models.Book, error) {
	var book models.Book

	conn, err := b.pool.Acquire(ctx)
	if err != nil {
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

func (b *BookRepositoryPg) UpdateBook(ctx context.Context, book *models.Book) error {
	conn, err := b.pool.Acquire(ctx)
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

func (b *BookRepositoryPg) DeleteBook(ctx context.Context, id int) error {
	conn, err := b.pool.Acquire(ctx)
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
