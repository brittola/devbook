package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da API
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Config adiciona todas as rotas dentro do router
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
