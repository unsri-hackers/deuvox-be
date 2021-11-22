package psql

import (
	"context"
	"database/sql"
	"deuvox/internal/model"
	"deuvox/pkg/derror"
)

func (p *psql) GetUserByEmail(email string) (model.User, error) {
	var result model.User
	row := p.statement.getUserByEmail.QueryRow(email)
	if err := row.Scan(&result.ID, &result.Email, &result.Password, &result.Verified, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt); err != nil {
		if err != sql.ErrNoRows {
			return result, derror.New(err.Error(), "AUTH-PSQL-001")
		}
	}
	return result, nil
}

func (p *psql) InsertNewSession(jti, userID, client, ip string) error {
	result, err := p.statement.insertNewSession.Exec(jti, userID, client, ip)
	if err != nil {
		return derror.New(err.Error(), "AUTH-PSQL-001")
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return derror.New(err.Error(), "AUTH-PSQL-001")
	}
	if rowAffected == 0 {
		return derror.New("Failed to insert new session", "AUTH-PSQL-001")
	}
	return nil
}

func (p *psql) InsertNewUser(ctx context.Context, id, email, password string) error {
	result, err := p.statement.insertNewUser.ExecContext(ctx, id, email, password)
	if err != nil {
		return derror.New(err.Error(), "AUTH-PSQL-001")
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return derror.New(err.Error(), "AUTH-PSQL-001")
	}
	if rowAffected == 0 {
		return derror.New("Failed to insert new user", "AUTH-PSQL-001")
	}
	return nil
}

func (p *psql) InsertNewProfile(ctx context.Context, id, userId, fullname string) error {
	result, err := p.statement.insertNewProfile.ExecContext(ctx, id, userId, fullname)
	if err != nil {
		return derror.New(err.Error(), "AUTH-PSQL-001")
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return derror.New(err.Error(), "AUTH-PSQL-001")
	}
	if rowAffected == 0 {
		return derror.New("Failed to insert new profile", "AUTH-PSQL-001")
	}
	return nil
}

func (p *psql) InsertNewPassword(ctx context.Context, id, userId, password string) error {
	result, err := p.statement.insertNewPassword.ExecContext(ctx, id, userId, password)
	if err != nil {
		return derror.New(err.Error(), "AUTH-PSQL-001")
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return derror.New(err.Error(), "AUTH-PSQL-001")
	}
	if rowAffected == 0 {
		return derror.New("Failed to insert new password", "AUTH-PSQL-001")
	}
	return nil
}
