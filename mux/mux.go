// This file provides an adapter to Gorilla mux to eliminate the tight coupling
// with Gorilla mux.
package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

// router implements the mux.Router interface.
type router struct {
	GorillaRouter *mux.Router
}

// NewRouter factors an instance that implements Router
func NewRouter() Router {
	gorillaRouter := mux.NewRouter()
	return router{gorillaRouter}
}

// POST implements the POST function in the Router interface.
func (r router) POST(path string, f func(http.ResponseWriter, *http.Request)) {
	r.GorillaRouter.HandleFunc(path, f).Methods("POST")
}

// ServeHTTP implements the ServeHTTP function in the Router interface
func (r router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.GorillaRouter.ServeHTTP(w, req)
}
