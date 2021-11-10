package app

import (
	"database/sql"
	"deuvox/internal/delivery/auth"
)

type delivery struct {
	auth *auth.Delivery
}

func initDelivery(db *sql.DB) delivery {
	u := initUsecase(db)
	return delivery{
		auth: auth.New(u.auth),
	}
}
