package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routes {
		var handler http.Handler

		handler = r.HandlerFunc
		handler = Logger(handler, r.Name)

		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(handler)
	}
	return router
}
