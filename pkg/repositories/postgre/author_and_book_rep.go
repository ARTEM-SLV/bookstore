package postgre

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"bookstore/internal/logger"
	"bookstore/models"
)

type AuthorAndBookPgRep struct {
	pool *pgxpool.Pool
}

func NewAuthorAndBookRepository(pool *pgxpool.Pool) *AuthorAndBookPgRep {
	return &AuthorAndBookPgRep{pool: pool}
}

func (ab AuthorAndBookPgRep) UpdateBookAndAuthor(ctx context.Context, book *models.Book, author *models.Author) error {
	tx, err := ab.pool.Begin(ctx)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "UPDATE books SET title=$1, author_id=$2, year=$3, isbn=$4 WHERE id=$5",
		book.Title, book.AuthorID, book.Year, book.ISBN, book.ID)
	if err != nil {
		logger.Log.Info(err.Error())
		return err
	}

	_, err = tx.Exec(ctx, "UPDATE authors SET first_name=$1, last_name=$2, biography=$3, birth_date=$4 WHERE id=$5",
		author.FirstName, author.LastName, author.Biography, author.BirthDate, author.ID)
	if err != nil {
		logger.Log.Info(err.Error())
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (ab AuthorAndBookPgRep) GetAuthorAndBooks(ctx context.Context, id int) (*models.Author, []*models.Book, error) {
	var author models.Author
	var books []*models.Book

	conn, err := ab.pool.Acquire(ctx)
	if err != nil {
		logger.Log.Error(err.Error())
		return &author, books, err
	}
	defer conn.Release()

	query := "SELECT id, first_name, last_name, biography, birth_date FROM authors WHERE id=$1"
	err = conn.QueryRow(ctx, query, id).Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate)
	if err != nil {
		logger.Log.Info(fmt.Sprintf("%s (id: %d)", err.Error(), id))
		return &author, books, err
	}

	query = `SELECT id, title, author_id, year, isbn FROM books WHERE author_id=$1`
	rows, err := conn.Query(ctx, query, id)
	if err != nil {
		logger.Log.Info(fmt.Sprintf("%s (author_id: %d)", err.Error(), id))
		return &author, books, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN)
		if err != nil {
			logger.Log.Error(err.Error())
			return &author, books, err
		}
		books = append(books, &book)
	}

	return &author, books, nil
}
