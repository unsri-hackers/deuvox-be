package utils

import (
	"net/http"
	"strings"
)

func GetAuthorizationToken(r *http.Request) string {
	tokenID := r.Header.Get("Authorization")
	tokenID = strings.Replace(tokenID, "BEARER", "", -1)
	tokenID = strings.Replace(tokenID, "Bearer", "", -1)
	tokenID = strings.Replace(tokenID, "bearer", "", -1)
	tokenID = strings.Replace(tokenID, " ", "", -1)
	return tokenID
}
