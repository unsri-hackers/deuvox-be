package auth

import "deuvox/internal/model"

func (r *Repository) IsAuthValid(req model.LoginRequest) bool {
	// TODO: ini ngambil data dari DB, terus ngecek email dan passwordnya bener atau nggak
	return true
}

func (r *Repository) AddNewUser(req model.RegisterRequest) bool {
	// TODO: ini ngambil data dari DB, terus ngecek email dan passwordnya bener atau nggak
	return true
}
