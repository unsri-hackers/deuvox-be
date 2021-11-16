package auth

import (
	"deuvox/pkg/derror"
	"deuvox/pkg/response"
	"deuvox/pkg/utils"
	"net/http"
)

func (d *Delivery) Token(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := utils.GetAuthorizationToken(r)

	if token == "" {
		response.Write(w, http.StatusBadRequest, "Authorization cannot be empty.", nil, "AUTH-TKN-01")
		return
	}

	res, err := d.auth.Token(ctx, token)
	if err != nil {
		e, ok := err.(*derror.DError)
		if !ok {
			response.Write(w, http.StatusInternalServerError, "Our server encounter a problem.", nil, "BAD-ERROR")
			return
		}
		response.Write(w, http.StatusBadRequest, e.Error(), nil, e.ErrorCode)
		return
	}
	response.Write(w, http.StatusOK, "Register successful", res, "")
}
