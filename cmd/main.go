package main

import (
	"bookstore/pkg/handlers"
	"bookstore/pkg/repositories"
)

func main() {
	repositories.InitRepositoryPG()
	handlers.InitHandlerMux()
}
