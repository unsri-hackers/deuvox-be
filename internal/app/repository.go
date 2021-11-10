package app

import (
	"database/sql"
	"deuvox/internal/repository/auth"
)

type repository struct {
	auth *auth.Repository
}

func initRepository(postgresDB *sql.DB) *repository {
	return &repository{
		auth: auth.New(postgresDB),
	}
}
