package main

import (
	"bookstore/pkg/handlers"
	"bookstore/pkg/repositories"
	"bookstore/pkg/services"
)

func main() {
	rep := repositories.NewRepository()
	srv := services.NewService(rep)
	handlers.InitHandler(srv)
}
