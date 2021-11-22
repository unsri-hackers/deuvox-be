package auth

import (
	"context"
	"deuvox/internal/model"
)

func (r *Repository) GetUserByEmail(email string) (model.User, error) {
	return r.datapsql.GetUserByEmail(email)
}

func (r *Repository) InsertNewSession(jti, userID, client, ip string) error {
	return r.datapsql.InsertNewSession(jti, userID, client, ip)
}

func (r *Repository) InsertNewUser(ctx context.Context, id, email, password string) error {
	return r.datapsql.InsertNewUser(ctx, id, email, password)
}

func (r *Repository) InsertNewProfile(ctx context.Context, id, userId, fullname string) error {
	return r.datapsql.InsertNewProfile(ctx, id, userId, fullname)
}

func (r *Repository) InsertNewPassword(ctx context.Context, id, userId, password string) error {
	return r.datapsql.InsertNewPassword(ctx, id, userId, password)
}
