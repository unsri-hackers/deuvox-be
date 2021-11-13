package app

import (
	"deuvox/internal/usecase/auth"
)

type usecase struct {
	auth *auth.Usecase
}

func (a *App) initUsecase() {
	var usecase usecase

	usecase.auth = auth.New(a.repository.auth)
	a.usecase = usecase
}
