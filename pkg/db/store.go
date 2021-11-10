package db

import "context"

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Verified  bool   `json:"verified" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type Profile struct {
	ID        string `json:"id"`
	UserId    string `json:"user_id"`
	Fullname  string `json:"fullname" validate:"required"`
	Picture   string `json:"picture"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Password struct {
	ID        string `json:"id"`
	UserId    string `json:"user_id"`
	Password  string `json:"password" validate:"required"`
	CreatedAt string `json:"created_at"`
}

type Session struct {
	JTI       string `json:"JTI"`
	UserId    string `json:"user_id"`
	Client    string `json:"client"`
	IP        string `json:"ip"`
	CreatedAt string `json:"created_at"`
}

type UserStore interface {
	CheckEmailExist(ctx context.Context, email string) (bool, error)
	AddNewUser(ctx context.Context, email, password string) (string, error)
}

type ProfileStore interface {
	AddNewProfile(ctx context.Context, userId, fullname string) error
}

type PasswordStore interface {
	AddNewPassword(ctx context.Context, userId, password string) error
}

type SessionStore interface {
	AddNewSession(ctx context.Context, userId string) (string, error)
}
