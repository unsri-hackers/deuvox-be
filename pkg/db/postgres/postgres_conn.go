package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Username string
	Password string
	Database string
}

func NewPG(config PostgresConfig) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Database,
	)

	connectionCfg, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config %w", err)
	}

	connStr := stdlib.RegisterConnConfig(connectionCfg)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
