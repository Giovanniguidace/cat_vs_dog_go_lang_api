package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI string
	Method string
	Function func(http.ResponseWriter, *http.Request)
}


// CreateRoutes will be used to create routes using mux project
func CreateRoutes(r *mux.Router) *mux.Router{
	routes := routesVoting

	for _, route := range routes{
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}