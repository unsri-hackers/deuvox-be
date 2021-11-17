package model

import (
	"database/sql"
	"deuvox/pkg/jwt"
	"time"
)

type LoginRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserAgent string `json:"-"`
	IP        string `json:"-"`
	UserID    string `json:"-"`
}

type LoginResponse struct {
	AccessToken  jwt.Token `json:"access_token"`
	RefreshToken jwt.Token `json:"refresh_token"`
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Fullname        string `json:"fullname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterResponse struct {
	AccessToken  *jwt.Token `json:"access_token"`
	RefreshToken *jwt.Token `json:"refresh_token"`
}

type User struct {
	ID        string       `db:"id"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Verified  bool         `db:"verified"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Session struct {
	JTI       string    `db:"jti"`
	UserID    string    `db:"user_id"`
	Client    string    `db:"client"`
	IP        string    `db:"ip"`
	CreatedAt time.Time `db:"created_at"`
}
