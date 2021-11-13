package auth

import (
	"database/sql"
	"deuvox/internal/model"
	"deuvox/internal/repository/auth/psql"
	"deuvox/pkg/db"
	"deuvox/pkg/db/postgres"
)

type datapsql interface {
	GetUserByEmail(email string) (model.User, error)
	InsertNewSession(jti, userID, client, ip string) error
}
type Repository struct {
	UserStore     db.UserStore
	ProfileStore  db.ProfileStore
	PasswordStore db.PasswordStore
	SessionStore  db.SessionStore
	datapsql      datapsql
}

// TODO: ini untuk dependancy injection yang diperlukan db
func New(postgresDB *sql.DB) *Repository {
	psql, err := psql.New(postgresDB)
	if err != nil {
		panic(err)
	}
	return &Repository{
		UserStore:     postgres.NewUserStore(postgresDB),
		ProfileStore:  postgres.NewProfileStore(postgresDB),
		PasswordStore: postgres.NewPasswordStore(postgresDB),
		SessionStore:  postgres.NewSessionStore(postgresDB),
		datapsql:      psql,
	}
}
