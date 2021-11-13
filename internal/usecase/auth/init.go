package auth

import (
	"context"
	"deuvox/internal/model"
)

type authRepo interface {
	GetUserByEmail(email string) (model.User, error)
	AddNewUser(ctx context.Context, req model.RegisterRequest) (string, string, error)
	InsertNewSession(jti, userID, client, ip string) error
}

type Usecase struct {
	auth authRepo
}

func New(authRepo authRepo) *Usecase {
	return &Usecase{
		auth: authRepo,
	}
}
