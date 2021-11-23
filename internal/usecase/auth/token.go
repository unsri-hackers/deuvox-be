package auth

import (
	"context"
	"deuvox/internal/model"
	"deuvox/pkg/derror"
	"deuvox/pkg/jwt"

	"github.com/rs/zerolog/log"
)

func (u *Usecase) Token(ctx context.Context, token string) (model.RegisterResponse, error) {
	var result model.RegisterResponse

	tokenValue, err := jwt.New().VerifyToken(token)
	if err != nil {
		log.Error().Err(err).Stack().Msg(err.Error())
		return result, derror.New(err.Error(), "AUTH-USC-04")
	}

	userId, _ := tokenValue.Get("userId")
	// jti, _ := tokenValue.Get("id")

	accessToken, _, err := jwt.CreateAccessToken(userId.(string))
	if err != nil {
		return result, derror.New("Failed to create access token.", "AUTH-USC-03")
	}

	refreshToken, _, err := jwt.CreateRefreshToken(userId.(string))
	if err != nil {
		return result, derror.New("Failedto create refresh token.", "AUTH-USC-06")
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken

	return result, nil
}
