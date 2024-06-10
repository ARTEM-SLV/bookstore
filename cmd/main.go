package main

import (
	"log"
	"net/http"

	"bookstore/pkg/handlers"
	"bookstore/pkg/repositories"
	"bookstore/pkg/services"
)

func main() {
	rep := repositories.NewRepository()
	srv := services.NewService(rep)
	r := handlers.NewHandler(srv)

	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
