package postgre

import (
	"bookstore/models"
	"context"
	"log"
)

type authorPgRepository struct {
	ctx context.Context
}

func NewAuthorRepository() *authorPgRepository {
	a := authorPgRepository{ctx: context.Background()}

	return &a
}

func (a authorPgRepository) CreateAuthor(author *models.AuthorTimeS) (int, error) {
	conn, err := GetDB().Acquire(a.ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	query := `INSERT INTO authors (first_name, last_name, biography, birth_date) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(a.ctx, query, author.FirstName, author.LastName, author.Biography, author.BirthDate).Scan(&author.ID)
	if err != nil {
		return 0, err
	}

	return author.ID, nil
}

func (a authorPgRepository) GetAllAuthors() ([]*models.Author, error) {
	var authors []*models.Author

	conn, err := GetDB().Acquire(a.ctx)
	if err != nil {
		return authors, err
	}
	defer conn.Release()

	query := `SELECT id, first_name, last_name, biography, birth_date FROM authors`
	rows, err := conn.Query(a.ctx, query)
	if err != nil {
		log.Println("ошибка выполнения запроса")
		return authors, err
	}
	defer rows.Close()

	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate)
		if err != nil {
			return authors, err
		}
		authors = append(authors, &author)
	}

	return authors, nil
}

func (a authorPgRepository) GetAuthorByID(id int) (*models.Author, error) {
	var author models.Author

	conn, err := GetDB().Acquire(a.ctx)
	if err != nil {
		return &author, err
	}
	defer conn.Release()

	query := "SELECT id, first_name, last_name, biography, birth_date FROM authors WHERE id=$1"
	err = conn.QueryRow(a.ctx, query, id).Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate)
	if err != nil {
		return &author, err
	}

	return &author, nil
}

func (a authorPgRepository) UpdateAuthor(author *models.AuthorTimeS) error {
	conn, err := GetDB().Acquire(a.ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	query := "UPDATE authors SET first_name=$1, last_name=$2, biography=$3, birth_date=$4 WHERE id=$5"
	_, err = conn.Exec(a.ctx, query, author.FirstName, author.LastName, author.Biography, author.BirthDate, author.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a authorPgRepository) DeleteAuthor(id int) error {
	tx, err := GetDB().Begin(a.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(a.ctx)

	query := "DELETE FROM books WHERE author_id=$1"
	_, err = tx.Exec(a.ctx, query, id)
	if err != nil {
		return err
	}

	query = "DELETE FROM authors WHERE id=$1"
	_, err = tx.Exec(a.ctx, query, id)
	if err != nil {
		return err
	}

	if err := tx.Commit(a.ctx); err != nil {
		return err
	}

	return nil
}
