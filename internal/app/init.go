package app

import (
	"deuvox/pkg/handler"

	"github.com/go-chi/chi"
)

type App struct {
	R *chi.Mux
}

func New() App {
	r := chi.NewRouter()
	d := initDelivery()
	r.Post("/auth/login", d.auth.Login)
	r.NotFound(handler.NotFound)
	r.MethodNotAllowed(handler.MethodNotAllowed)
	return App{
		R: r,
	}
}
