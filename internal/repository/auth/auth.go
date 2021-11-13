package auth

import (
	"context"
	"deuvox/internal/model"
	"deuvox/pkg/crypto"
	"fmt"
)

func (r *Repository) GetUserByEmail(email string) (model.User, error) {
	return r.datapsql.GetUserByEmail(email)
}
func (r *Repository) InsertNewSession(jti, userID, client, ip string) error {
	return r.datapsql.InsertNewSession(jti, userID, client, ip)
}

func (r *Repository) AddNewUser(ctx context.Context, req model.RegisterRequest) (string, string, error) {
	exist, err := r.UserStore.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return "", "", err
	}

	if exist {
		return "", "", fmt.Errorf("Email already exists")
	}

	hashPassword, err := crypto.HashPassword(req.Password)
	if err != nil {
		return "", "", err
	}

	userId, err := r.UserStore.AddNewUser(ctx, req.Email, hashPassword)
	if err != nil {
		return "", "", err
	}

	if err := r.ProfileStore.AddNewProfile(ctx, userId, req.Fullname); err != nil {
		return "", "", err
	}

	if err := r.PasswordStore.AddNewPassword(ctx, userId, hashPassword); err != nil {
		return "", "", err
	}

	jti, err := r.SessionStore.AddNewSession(ctx, userId)
	if err != nil {
		return "", "", err
	}

	return jti, userId, nil
}
