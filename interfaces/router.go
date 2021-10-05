package interfaces

import (
	"github.com/gorilla/mux"
	"github.com/sonalys/ddd/application"
)

type Router struct {
	*mux.Router
}

type Dependency struct {
	application.Cart
}

func NewRouter(d Dependency) Router {
	router := Router{mux.NewRouter()}

	cartRouter := router.PathPrefix("/cart").Subrouter()
	newCartHandler(d.Cart, cartRouter)

	return router
}
