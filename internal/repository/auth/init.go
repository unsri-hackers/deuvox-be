package auth

import "database/sql"

type Repository struct {
	PostgresDB *sql.DB
}

// TODO: ini untuk dependancy injection yang diperlukan db
func New(postgresDB *sql.DB) *Repository {
	return &Repository{
		PostgresDB: postgresDB,
	}
}
