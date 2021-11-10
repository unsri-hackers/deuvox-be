package auth

import (
	"deuvox/internal/model"
)

type authRepo interface {
	IsAuthValid(req model.LoginRequest) bool
	AddNewUser(req model.RegisterRequest) bool
}

type Usecase struct {
	auth authRepo
}

func New(authRepo authRepo) *Usecase {
	return &Usecase{
		auth: authRepo,
	}
}
