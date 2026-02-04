package main

import (
	"log"

	"github.com/MohammedElattar/movie-reservation/internal/app"
	"github.com/MohammedElattar/movie-reservation/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	app, err := app.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		app.Logger.Error("failed to start the app")
	}
}
