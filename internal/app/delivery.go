package app

import (
	"deuvox/internal/delivery/auth"
)

type delivery struct {
	auth *auth.Delivery
}

func (a *App) initDelivery() {
	var delivery delivery
	delivery.auth = auth.New(a.usecase.auth)

	a.delivery = delivery
}
