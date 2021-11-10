package main

import (
	"deuvox/internal/app"
	"deuvox/pkg/db/postgres"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	serverCfg := app.AppConfig{
		Host:            os.Getenv("SERVER_HOST"),
		Port:            os.Getenv("SERVER_PORT"),
		ReadTimeout:     500 * time.Millisecond,
		WriteTimeout:    500 * time.Millisecond,
		ShutdownTimeout: 10 * time.Second,
	}

	postgresCfg := postgres.PostgresConfig{
		Host:     os.Getenv("POSTGRES_ADDR"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	app := app.New(serverCfg, postgresCfg)
	log.Info().Msg("Starting api server in localhost:8080")
	app.StartServer()
}
