package main

import (
	"log/slog"
	"net/http"
	"time"

	"rocketseat/internal/api"
	"rocketseat/internal/store"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	store := store.NewStore(rdb)
	handler := api.NewHandler(store)

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
