package router

import (
	"devbook-api/src/router/routes"

	"github.com/gorilla/mux"
)

// Create retorna um router com as rotas configuradas
func Create() *mux.Router {
	return routes.Config(mux.NewRouter())
}
