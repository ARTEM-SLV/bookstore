package main

import (
	"bookstore/internal/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"bookstore/pkg/handlers"
	"bookstore/pkg/repositories"
	"bookstore/pkg/services"
)

func main() {
	logger.InitLogger()
	defer logger.Log.CloseLog()

	// Обработка сигнала прерывания
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		logger.Log.CloseLog()

		os.Exit(0)
	}()

	rep := repositories.NewRepository()
	srv := services.NewService(rep)
	r := handlers.NewHandler(srv)

	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
