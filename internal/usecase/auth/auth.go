package auth

import (
	"deuvox/internal/model"
	"errors"
)

func (u *Usecase) Login(body model.LoginRequest) (model.LoginResponse, error) {
	var result model.LoginResponse
	if !u.auth.IsAuthValid(body) {
		return result, errors.New("Not Valid")
	}
	// TODO: JWT things
	return result, nil
}
