package auth

import (
	"context"
	"deuvox/internal/model"
)

type authRepo interface {
	GetUserByEmail(email string) (model.User, error)
	InsertNewSession(jti, userID, client, ip string) error
	InsertNewUser(ctx context.Context, id, email, password string) error
	InsertNewProfile(ctx context.Context, id, userId, fullname string) error
	InsertNewPassword(ctx context.Context, id, userId, password string) error
}

type Usecase struct {
	auth authRepo
}

func New(authRepo authRepo) *Usecase {
	return &Usecase{
		auth: authRepo,
	}
}
