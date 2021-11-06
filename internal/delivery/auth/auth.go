package auth

import (
	"deuvox/internal/model"
	"deuvox/pkg/response"
	"encoding/json"
	"net/http"
)

func (d *Delivery) Login(w http.ResponseWriter, r *http.Request) {
	var body model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Write(w, http.StatusBadRequest, "Invalid body request!", nil, "ERR01")
		return
	}

	if body.Email == "" {
		response.Write(w, http.StatusBadRequest, "Email cannot be empty!", nil, "ERR01")
		return
	}
	if body.Password == "" {
		response.Write(w, http.StatusBadRequest, "Password cannot be empty!", nil, "ERR01")
		return
	}

	res, err := d.auth.Login(body)
	if err != nil {
		response.Write(w, http.StatusBadRequest, err.Error(), nil, "ERR01")
		return
	}
	response.Write(w, http.StatusOK, "Login successful", res, "")
	return
}
