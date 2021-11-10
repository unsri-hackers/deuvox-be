package postgres

import (
	"context"
	"database/sql"
	"deuvox/pkg/db"

	"github.com/google/uuid"
)

type ProfileStore struct {
	db *sql.DB
}

func NewProfileStore(db *sql.DB) db.ProfileStore {
	return &ProfileStore{
		db: db,
	}
}

func (s *ProfileStore) AddNewProfile(ctx context.Context, userId, fullname string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	query := `INSERT INTO "profile" (id, user_id, fullname) VALUES ($1, $2, $3)`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id, userId, fullname)
	if err != nil {
		return err
	}

	return nil
}
