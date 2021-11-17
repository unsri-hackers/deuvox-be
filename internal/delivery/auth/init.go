package auth

import (
	"context"
	"deuvox/internal/model"
)

type authUC interface {
	Login(body model.LoginRequest) (model.LoginResponse, error)
	Register(ctx context.Context, body model.RegisterRequest) (model.RegisterResponse, error)
	Token(ctx context.Context, token string) (model.RegisterResponse, error)
}

type Delivery struct {
	auth authUC
}

func New(authUC authUC) *Delivery {
	return &Delivery{
		auth: authUC,
	}
}
