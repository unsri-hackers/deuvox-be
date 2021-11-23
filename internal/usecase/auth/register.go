package auth

import (
	"context"
	"deuvox/internal/model"
	"deuvox/pkg/crypto"
	"deuvox/pkg/derror"
	"deuvox/pkg/jwt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (u *Usecase) Register(ctx context.Context, body model.RegisterRequest) (model.RegisterResponse, error) {
	var result model.RegisterResponse

	user, err := u.auth.GetUserByEmail(body.Email)
	if err != nil {
		return result, err
	}

	hashPassword, err := crypto.HashPassword(body.Password)
	if err != nil {
		return result, err
	}

	userId, err := uuid.NewRandom()
	if err != nil {
		return result, derror.New("Failed to create random uuid: %v", "AUTH-USC-01")
	}

	if err := u.auth.InsertNewUser(ctx, userId.String(), body.Email, hashPassword); err != nil {
		return result, derror.New("Error add user", "AUTH-USC-01")
	}

	profileId, err := uuid.NewRandom()
	if err != nil {
		return result, derror.New("Failed to create random uuid: %v", "AUTH-USC-01")
	}

	if err := u.auth.InsertNewProfile(ctx, profileId.String(), userId.String(), body.Fullname); err != nil {
		return result, derror.New("Error add profile", "AUTH-USC-01")
	}

	passwordId, err := uuid.NewRandom()
	if err != nil {
		return result, derror.New("Failed to create random uuid: %v", "AUTH-USC-01")
	}

	if err := u.auth.InsertNewPassword(ctx, passwordId.String(), userId.String(), hashPassword); err != nil {
		log.Error().Err(err).Stack().Msg("Error add password")
		return result, derror.New("Error add password", "AUTH-USC-01")
	}

	accessToken, jti, err := jwt.CreateAccessToken(user.ID)
	if err != nil {
		log.Error().Err(err).Stack().Msg("Failed to create access token.")
		return result, derror.New("Failed to create access token.", "AUTH-USC-03")
	}

	if err := u.auth.InsertNewSession(jti, userId.String(), body.UserAgent, body.IP); err != nil {
		return result, derror.New("Error add session", "AUTH-USC-01")
	}

	refreshToken, _, err := jwt.CreateRefreshToken(user.ID)
	if err != nil {
		return result, derror.New("Failed to create refresh token.", "AUTH-USC-06")
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken

	return result, nil
}
