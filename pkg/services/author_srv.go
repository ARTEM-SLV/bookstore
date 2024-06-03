package services

import (
	"bookstore/models"
)

func CreateAuthor(author *models.AuthorTimeS) (int, error) {
	return authorRep.CreateAuthor(author)
}

func GetAllAuthors() ([]*models.Author, error) {
	return authorRep.GetAllAuthors()
}

func GetAuthorByID(id int) (*models.Author, error) {
	return authorRep.GetAuthorByID(id)
}

func UpdateAuthor(author *models.AuthorTimeS) error {
	return authorRep.UpdateAuthor(author)
}

func DeleteAuthor(id int) error {
	return authorRep.DeleteAuthor(id)
}
