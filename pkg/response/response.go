package response

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message   string      `json:"message"`
	ErrorCode string      `json:"errorCode,omitempty"`
	Data      interface{} `json:"data"`
}

func Write(w http.ResponseWriter, status int, message string, data interface{}, errCode string) {
	r := response{
		Message:   message,
		ErrorCode: errCode,
	}
	r.Data = data
	jsonResponse, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
