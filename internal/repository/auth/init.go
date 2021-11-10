package auth

import (
	"database/sql"
	"deuvox/pkg/db"
	"deuvox/pkg/db/postgres"
)

type Repository struct {
	UserStore     db.UserStore
	ProfileStore  db.ProfileStore
	PasswordStore db.PasswordStore
	SessionStore  db.SessionStore
}

// TODO: ini untuk dependancy injection yang diperlukan db
func New(postgresDB *sql.DB) *Repository {
	return &Repository{
		UserStore:     postgres.NewUserStore(postgresDB),
		ProfileStore:  postgres.NewProfileStore(postgresDB),
		PasswordStore: postgres.NewPasswordStore(postgresDB),
		SessionStore:  postgres.NewSessionStore(postgresDB),
	}
}
