package metadata

import (
	"io/ioutil"
	"net/http"

	"github.com/tracer/mux"
)

type router struct {
	r mux.Router
	s Storage
}

// Register registers the API supported by this package.
func Register(r mux.Router, s Storage) {
	storedRouter := router{r, s}
	r.POST("/tracers/{id}/metadata", storedRouter.handleMetadata)
}

func (r *router) handleMetadata(responseWriter http.ResponseWriter, request *http.Request) {
	vars := r.r.Vars(request)
	data, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	r.s.Store(vars["id"], string(data))
}
