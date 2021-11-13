package app

import (
	"database/sql"
	"deuvox/internal/repository/auth"
)

type repository struct {
	auth *auth.Repository
}

func (a *App) initRepository(db *sql.DB) {
	var repository repository
	repository.auth = auth.New(db)
	a.repository = repository
}
