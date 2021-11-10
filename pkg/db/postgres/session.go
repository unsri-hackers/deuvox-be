package postgres

import (
	"context"
	"database/sql"
	"deuvox/pkg/db"

	"github.com/google/uuid"
)

type SessionStore struct {
	db *sql.DB
}

func NewSessionStore(db *sql.DB) db.SessionStore {
	return &SessionStore{
		db: db,
	}
}

func (s *SessionStore) AddNewSession(ctx context.Context, userId string) (string, error) {
	jti, err := uuid.NewRandom()
	if err != nil {
		return jti.String(), err
	}

	query := `INSERT INTO "session" (jti, user_id, client, ip) VALUES ($1, $2, $3, $4)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return jti.String(), err
	}

	_, err = stmt.ExecContext(ctx, jti, userId, "", "")
	if err != nil {
		return jti.String(), err
	}

	return jti.String(), nil
}
