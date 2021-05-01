package rest

import (
	usecases2 "github.com/carrot-systems/cs-core/src/core/usecases"
)

type RoutesHandler struct {
	Usecases usecases2.Usecases
}

func NewRouter(ucHandler usecases2.Usecases) RoutesHandler {
	return RoutesHandler{Usecases: ucHandler}
}
