package psql

import (
	"deuvox/internal/model"
	"deuvox/pkg/derror"
)

func (p *psql) GetUserByEmail(email string) (model.User, error) {
	var result model.User
	row := p.statement.getUserByEmail.QueryRow(email)
	if err := row.Scan(&result.ID, &result.Email, &result.Password, &result.Verified, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt); err != nil {
		if err.Error() != "sql: no rows in result set" {
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
