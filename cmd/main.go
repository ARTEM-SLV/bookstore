package main

import (
	"bookstore/pkg/repositories/postgre"
	"bookstore/pkg/service"
)

func main() {
	//handlers.NewHandler()
	postgre.InitRepository()
	//repositories.NewRepository()
	service.StartServer()
}
