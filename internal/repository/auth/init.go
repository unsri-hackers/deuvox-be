package auth

import (
	"context"
	"database/sql"
	"deuvox/internal/model"
	"deuvox/internal/repository/auth/psql"
)

type datapsql interface {
	GetUserByEmail(email string) (model.User, error)
	InsertNewSession(jti, userId, client, ip string) error
	InsertNewUser(ctx context.Context, id, email, password string) error
	InsertNewProfile(ctx context.Context, id, userId, fullname string) error
	InsertNewPassword(ctx context.Context, id, userId, password string) error
}
type Repository struct {
	datapsql datapsql
}

// TODO: ini untuk dependancy injection yang diperlukan db
func New(postgresDB *sql.DB) *Repository {
	psql, err := psql.New(postgresDB)
	if err != nil {
		panic(err)
	}
	return &Repository{
		datapsql: psql,
	}
}
