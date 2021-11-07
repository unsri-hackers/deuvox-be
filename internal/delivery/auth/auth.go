package auth

import (
	"deuvox/internal/model"
	"deuvox/pkg/derror"
	"deuvox/pkg/response"
	"encoding/json"
	"net/http"
	"net/mail"
)

func (d *Delivery) Login(w http.ResponseWriter, r *http.Request) {
	var body model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Write(w, http.StatusBadRequest, "Invalid body request.", nil, "AUTH-DLV-01")
		return
	}

	if body.Email == "" {
		response.Write(w, http.StatusBadRequest, "Email cannot be empty.", nil, "AUTH-DLV-02")
		return
	}
	_, err := mail.ParseAddress(body.Email)
	if err != nil {
		response.Write(w, http.StatusBadRequest, "Invalid email format.", nil, "AUTH-DLV-03")
		return
	}

	if body.Password == "" {
		response.Write(w, http.StatusBadRequest, "Password cannot be empty.", nil, "AUTH-DLV-04")
		return
	}
	if len(body.Password) < 8 {
		response.Write(w, http.StatusBadRequest, "Password is too short.", nil, "AUTH-DLV-05")
		return
	}

	res, err := d.auth.Login(body)
	if err != nil {
		e, ok := err.(*derror.DError)
		if !ok {
			response.Write(w, http.StatusInternalServerError, "Our server encounter a problem.", nil, "BAD-ERROR")
			return
		}
		response.Write(w, http.StatusBadRequest, e.Error(), nil, e.ErrorCode)
		return
	}
	response.Write(w, http.StatusOK, "Login successful", res, "")
	return
}
