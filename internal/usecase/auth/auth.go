package auth

import (
	"context"
	"deuvox/internal/model"
	"deuvox/pkg/derror"
	"deuvox/pkg/jwt"

	"github.com/rs/zerolog/log"
)

func (u *Usecase) Login(body model.LoginRequest) (model.LoginResponse, error) {
	var result model.LoginResponse
	if !u.auth.IsAuthValid(body) {
		return result, derror.New("Invalid username or password.", "AUTH-USC-01")
	}
	// TODO: JWT things
	return result, nil
}

func (u *Usecase) Register(ctx context.Context, body model.RegisterRequest) (model.RegisterResponse, error) {
	var result model.RegisterResponse
	jti, userId, err := u.auth.AddNewUser(ctx, body)
	if err != nil {
		return result, derror.New("Error adding new user.", "AUTH-USC-02")
	}

	accessToken, err := jwt.New().CreateToken(jti, userId, jwt.AccessTokenExpiration)
	if err != nil {
		log.Error().Err(err).Stack().Msg("Error create jwt access token")
		return result, derror.New("Error create jwt access token.", "AUTH-USC-03")
	}

	refreshToken, err := jwt.New().CreateToken(jti, userId, jwt.RefreshTokenExpiration)
	if err != nil {
		return result, derror.New("Error create jwt access token.", "AUTH-USC-03")
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken

	return result, nil
}
