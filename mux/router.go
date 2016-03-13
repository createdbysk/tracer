package mux

import "net/http"

// Router defines the interface to register a handler for a route.
type Router interface {
	// POST specifies a POST path route.
	POST(path string, f func(http.ResponseWriter, *http.Request))
	// Vars retrieves path variables from a request.
	Vars(req *http.Request) map[string]string
	// Provide the http.Handler interface.
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
