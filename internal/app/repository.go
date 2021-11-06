package app

import "deuvox/internal/repository/auth"

type repository struct {
	auth *auth.Repository
}

func initRepository() *repository {
	return &repository{
		auth: auth.New(),
	}
}
