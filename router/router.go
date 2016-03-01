package router

import "net/http"

// Router defines the interface to register a handler for a route.
type Router interface {
	// TODO: HandleFunc here expects to return a mux.Route, which strongly
	// 		 couples this implementation to the gorilla mux. Add an abstraction
	// 	     on top of this to get rid of this coupling.
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) interface{}
}
