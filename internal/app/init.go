package app

import (
	"deuvox/pkg/db/postgres"
	"deuvox/pkg/handler"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type App struct {
	R *chi.Mux
}

func New() App {
	r := chi.NewRouter()
	d := initDelivery()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/auth/login", d.auth.Login)
	r.NotFound(handler.NotFound)
	r.MethodNotAllowed(handler.MethodNotAllowed)

	// example conn
	postgresCfg := postgres.PostgresConfig{
		Host:     os.Getenv("POSTGRES_ADDR"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	log.Info().Msg("get connection to postgres")
	_, err := postgres.NewPG(postgresCfg)
	if err != nil {
		log.Error().Err(err).Stack().Msg("failed to connect to postgres")
	}

	return App{
		R: r,
	}
}
