// Package mux provides an adapter to Gorilla mux to eliminate the tight coupling
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

// Vars implements the Vars() function in the Router interface.
func (r router) Vars(request *http.Request) map[string]string {
	return mux.Vars(request)
}

// ServeHTTP implements the ServeHTTP function in the Router interface
func (r router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.GorillaRouter.ServeHTTP(w, req)
}
