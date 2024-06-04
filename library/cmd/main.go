package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	server "library/internal"
)

type application struct {
	logger *slog.Logger
	grpc   *server.Server
}

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	server := server.NewServer("localhost:50051", "localhost:50052")
	app := &application{
		logger: logger,
		grpc:   server,
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("Starting server on :8080")
	err := srv.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)
}
