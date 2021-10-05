package interfaces

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sonalys/ddd/application"
)

type Router struct {
	*mux.Router
	*Configuration
}

type Configuration struct {
	Address string `json:"address"`
}

type Dependency struct {
	*Configuration
	application.Cart
}

func NewRouter(d Dependency) Router {
	router := Router{mux.NewRouter(), d.Configuration}

	cartRouter := router.PathPrefix("/cart").Subrouter()
	newCartHandler(d.Cart, cartRouter)

	return router
}

func (r *Router) Start() {
	http.ListenAndServe(r.Configuration.Address, r)
}
