package app

import (
	"database/sql"
	"deuvox/internal/usecase/auth"
)

type usecase struct {
	auth *auth.Usecase
}

func initUsecase(postgresDB *sql.DB) *usecase {
	r := initRepository(postgresDB)
	return &usecase{
		auth: auth.New(r.auth),
	}
}
