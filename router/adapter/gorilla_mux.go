package adapter

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GorillaMuxRouter is an adapter for the gorilla mux Router.
// Usage: GorillaMuxRouter{router}, where router *mux.Router
type GorillaMuxRouter struct {
	// The instance of mux.Router to adapt to the Router interface.
	Router *mux.Router
}

// HandleFunc implements the HandleFunc method of the router interface
func (router GorillaMuxRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) interface{} {
	// TODO: Handle nil router.Router
	return router.Router.HandleFunc(path, f)
}
