package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "CreateEvent",
		Method:      "POST",
		Pattern:     "/event",
		HandlerFunc: CreateEvent,
	},
	Route{
		Name:        "GetEvent",
		Method:      "GET",
		Pattern:     "/event/{id}",
		HandlerFunc: GetOneEvent,
	},
	Route{
		Name:        "GetAllEvent",
		Method:      "GET",
		Pattern:     "/event",
		HandlerFunc: GetAllEvent,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		log.Printf("%v method:%v pattern:%v", route.Name, route.Method, route.Pattern)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
