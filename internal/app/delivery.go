package app

import "deuvox/internal/delivery/auth"

type delivery struct {
	auth *auth.Delivery
}

func initDelivery() delivery {
	u := initUsecase()
	return delivery{
		auth: auth.New(u.auth),
	}
}
