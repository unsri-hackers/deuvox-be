package app

import "deuvox/internal/usecase/auth"

type usecase struct {
	auth *auth.Usecase
}

func initUsecase() *usecase {
	r := initRepository()
	return &usecase{
		auth: auth.New(r.auth),
	}
}
