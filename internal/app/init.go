package app

import (
	"github.com/go-chi/chi"
)

type App struct {
	R *chi.Mux
}

func New() App {
	r := chi.NewRouter()
	d := initDelivery()
	r.Post("/auth/login", d.auth.Login)
	return App{
		R: r,
	}
}
