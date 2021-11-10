package auth

import (
	"deuvox/internal/model"
)

type authUC interface {
	Login(body model.LoginRequest) (model.LoginResponse, error)
	Register(body model.RegisterRequest) (model.RegisterResponse, error)
}

type Delivery struct {
	auth authUC
}

func New(authUC authUC) *Delivery {
	return &Delivery{
		auth: authUC,
	}
}
