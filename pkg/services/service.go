package services

import (
	"bookstore/pkg/repositories/postgre"
)

var bookRep = postgre.NewBookRepository()
var authorRep = postgre.NewAuthorRepository()
var authorAndBookRep = postgre.NewAuthorAndBookRepository()
