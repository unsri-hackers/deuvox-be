package auth

import (
	"deuvox/internal/model"
	"deuvox/pkg/derror"
)

func (u *Usecase) Login(body model.LoginRequest) (model.LoginResponse, error) {
	var result model.LoginResponse
	if !u.auth.IsAuthValid(body) {
		return result, derror.New("Invalid username or password.", "AUTH-USC-01")
	}
	// TODO: JWT things
	return result, nil
}

func (u *Usecase) Register(body model.RegisterRequest) (model.RegisterResponse, error) {
	var result model.RegisterResponse
	if !u.auth.AddNewUser(body) {
		return result, derror.New("Invalid username or password.", "AUTH-USC-01")
	}
	// TODO: JWT things
	return result, nil
}
