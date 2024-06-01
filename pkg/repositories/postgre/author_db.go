package postgre

import (
	"bookstore/models"
	"log"
)

type authorRepository struct{}

func NewAuthorRepository() *authorRepository {
	return &authorRepository{}
}

func (a authorRepository) CreateAuthor(author *models.AuthorWithoutTime) (int, error) {
	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	query := `INSERT INTO authors (first_name, last_name, biography, birth_date) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(ctx, query, author.FirstName, author.LastName, author.Biography, author.BirthDate).Scan(&author.ID)
	if err != nil {
		return 0, err
	}

	return author.ID, nil
}

func (a authorRepository) GetAllAuthors() ([]*models.Author, error) {
	var authors []*models.Author

	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		return authors, err
	}
	defer conn.Release()

	query := `SELECT id, first_name, last_name, biography, birth_date FROM authors`
	rows, err := conn.Query(ctx, query)
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

func (a authorRepository) GetAuthorByID(id int) (*models.Author, error) {
	var author models.Author

	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		return &author, err
	}
	defer conn.Release()

	query := "SELECT id, first_name, last_name, biography, birth_date FROM authors WHERE id=$1"
	err = conn.QueryRow(ctx, query, id).Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate)
	if err != nil {
		return &author, err
	}

	return &author, nil
}

func (a authorRepository) UpdateAuthor(author *models.AuthorWithoutTime) error {
	conn, err := GetDB().Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	query := "UPDATE authors SET first_name=$1, last_name=$2, biography=$3, birth_date=$4 WHERE id=$5"
	_, err = conn.Exec(ctx, query, author.FirstName, author.LastName, author.Biography, author.BirthDate, author.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a authorRepository) DeleteAuthor(id int) error {
	tx, err := GetDB().Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := "DELETE FROM books WHERE author_id=$1"
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	query = "DELETE FROM authors WHERE id=$1"
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
