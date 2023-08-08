package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	Bind  string `env:"BIND" envDefault:"127.0.0.1:8080"`
	Debug bool   `env:"DEBUG" envDefault:"false"`
}

func main() {
	log.Info("Starting...")

	if err := run(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	log.Info("Stopped")
}

func run() error {
	log.Info("Parsing environment variables...")
	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	srv := server{
		router: r,
	}
	srv.routes()

	log.Infof("Starting server on %s", cfg.Bind)
	if err := http.ListenAndServe(cfg.Bind, srv.router); err != nil {
		log.Fatal(err)
	}

	return nil
}
