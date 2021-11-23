package psql

import (
	"database/sql"
)

type statement struct {
	getUserByEmail    *sql.Stmt
	checkEmailExist   *sql.Stmt
	insertNewSession  *sql.Stmt
	insertNewUser     *sql.Stmt
	insertNewProfile  *sql.Stmt
	insertNewPassword *sql.Stmt
}

type psql struct {
	db        *sql.DB
	statement statement
}

func New(postgresDB *sql.DB) (*psql, error) {
	statement, err := initStatement(postgresDB)
	if err != nil {
		return nil, err
	}
	return &psql{
		db:        postgresDB,
		statement: statement,
	}, nil

}

func initStatement(postgresDB *sql.DB) (statement, error) {
	var statement statement
	var err error

	statement.getUserByEmail, err = postgresDB.Prepare(getUserByEmail)
	if err != nil {
		return statement, err
	}

	statement.insertNewSession, err = postgresDB.Prepare(insertNewSession)
	if err != nil {
		return statement, err
	}

	statement.checkEmailExist, err = postgresDB.Prepare(checkEmailExist)
	if err != nil {
		return statement, err
	}

	statement.insertNewUser, err = postgresDB.Prepare(insertNewUser)
	if err != nil {
		return statement, err
	}

	statement.insertNewProfile, err = postgresDB.Prepare(insertNewProfile)
	if err != nil {
		return statement, err
	}

	statement.insertNewPassword, err = postgresDB.Prepare(insertNewPassword)
	if err != nil {
		return statement, err
	}

	return statement, nil
}
