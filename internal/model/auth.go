package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Fullname        string `json:"fullname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
