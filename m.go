package main

import (
	"log/slog"
	"net/http"
	"rocketseat/api"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	db := make(map[string]string)
	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         "localhost:8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
