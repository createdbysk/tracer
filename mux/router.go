package mux

import "net/http"

// Router defines the interface to register a handler for a route.
type Router interface {
	POST(path string, f func(http.ResponseWriter, *http.Request))
	// Provide the http.Handler interface.
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
