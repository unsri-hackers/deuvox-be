package main

import (
	"deuvox/internal/app"
	"deuvox/pkg/db/postgres"
	"os"
	"time"

	"deuvox/pkg/log"
)

func main() {

	log.New()
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
	log.Info().Msg("get connection to postgres")
	postgre, err := postgres.NewPG(postgresCfg)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to postgres")
	}

	app := app.New(serverCfg, postgre)
	app.StartServer()
}
