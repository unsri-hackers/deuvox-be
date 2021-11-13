package auth

import (
	"context"
	"deuvox/internal/model"
	"deuvox/pkg/derror"
	"deuvox/pkg/jwt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) Login(body model.LoginRequest) (model.LoginResponse, error) {
	var result model.LoginResponse

	user, err := u.auth.GetUserByEmail(body.Email)
	if err != nil {
		return result, err
	}
	if user.Email == "" {
		return result, derror.New("Email not found.", "AUTH-USC-01")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return result, derror.New("Wrong password.", "AUTH-USC-04")
	}
	jti, err := uuid.NewRandom()
	if err != nil {
		return result, derror.New("Failed to create jti", "AUTH-USC-05")
	}
	if err := u.auth.InsertNewSession(jti.String(), user.ID, body.UserAgent, body.IP); err != nil {
		return result, err
	}
	accessToken, err := jwt.New().CreateToken(jti.String(), user.ID, jwt.AccessTokenExpiration)
	if err != nil {
		return result, derror.New("Failed to create access token.", "AUTH-USC-03")
	}
	refreshToken, err := jwt.New().CreateToken(jti.String(), user.ID, jwt.RefreshTokenExpiration)
	if err != nil {
		return result, derror.New("Failedto create refresh token.", "AUTH-USC-06")
	}
	result.AccessToken = accessToken
	result.RefreshToken = refreshToken
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
