package auth_test

import (
	"bytes"
	"deuvox/internal/delivery/auth"
	mock "deuvox/internal/delivery/auth/_mock"
	"deuvox/internal/model"
	"deuvox/pkg/derror"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initRequest(r *http.Request, m *mock.MockauthUC) *httptest.ResponseRecorder {
	handler := auth.New(m)
	router := chi.NewRouter()
	w := httptest.NewRecorder()

	router.Post("/login", handler.Login)
	router.ServeHTTP(w, r)
	return w
}

func TestDelivery_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockauthUC(ctrl)

	testCase := []struct {
		name       string
		body       model.LoginRequest
		ucCalled   bool
		ucError    error
		wantStatus int
	}{
		{"success", model.LoginRequest{Email: "test@test.com", Password: "testyttttttttt"}, true, nil, http.StatusOK},
		{"empty email", model.LoginRequest{}, false, nil, http.StatusBadRequest},
		{"invalid email", model.LoginRequest{Email: "LUL"}, false, nil, http.StatusBadRequest},
		{"empty password", model.LoginRequest{Email: "test@test.com", Password: ""}, false, nil, http.StatusBadRequest},
		{"short password", model.LoginRequest{Email: "test@test.com", Password: "LUL"}, false, nil, http.StatusBadRequest},
		{"uc derror", model.LoginRequest{Email: "test@test.com", Password: "testyttttttttt"}, true, derror.New("", ""), http.StatusBadRequest},
		{"uc not derror", model.LoginRequest{Email: "test@test.com", Password: "testyttttttttt"}, true, errors.New(""), http.StatusInternalServerError},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tc.body)
			if err != nil {
				t.Error(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
			if tc.ucCalled {
				m.EXPECT().Login(gomock.Any()).Return(model.LoginResponse{}, tc.ucError)
			}
			res := initRequest(req, m).Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}
