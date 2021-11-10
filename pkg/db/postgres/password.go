package postgres

import (
	"context"
	"database/sql"
	"deuvox/pkg/db"

	"github.com/google/uuid"
)

type PasswordStore struct {
	db *sql.DB
}

func NewPasswordStore(db *sql.DB) db.PasswordStore {
	return &PasswordStore{
		db: db,
	}
}

func (s *PasswordStore) AddNewPassword(ctx context.Context, userId, password string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	query := `INSERT INTO "password" (id, user_id, password) VALUES ($1, $2, $3)`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id, userId, password)
	if err != nil {
		return err
	}

	return nil
}
