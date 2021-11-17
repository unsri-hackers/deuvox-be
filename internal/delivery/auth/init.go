package auth

import (
	"context"
	"deuvox/internal/model"
)

//go:generate mockgen -source=./init.go -destination=./_mock/auth_mock.go -package=mock_auth
type authUC interface {
	Login(body model.LoginRequest) (model.LoginResponse, error)
	Register(ctx context.Context, body model.RegisterRequest) (model.RegisterResponse, error)
}

type Delivery struct {
	auth authUC
}

func New(authUC authUC) *Delivery {
	return &Delivery{
		auth: authUC,
	}
}
