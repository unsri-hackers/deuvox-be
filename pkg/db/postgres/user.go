package postgres

import (
	"context"
	"database/sql"
	"deuvox/pkg/db"

	"github.com/google/uuid"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) db.UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) CheckEmailExist(ctx context.Context, email string) (bool, error) {
	query := `SELECT "id", "email" FROM "user" WHERE "deleted_at" IS NULL AND "email" = $1`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return false, err
	}

	var user db.User
	err = stmt.QueryRowContext(ctx, email).Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (us *UserStore) AddNewUser(ctx context.Context, email, password string) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return id.String(), err
	}

	query := `INSERT INTO "user" (id, email, password) VALUES ($1, $2, $3)`

	stmt, err := us.db.Prepare(query)
	if err != nil {
		return id.String(), err
	}

	_, err = stmt.ExecContext(ctx, id, email, password)
	if err != nil {
		return id.String(), err
	}

	return id.String(), nil
}
