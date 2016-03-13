package metadata

import (
	"net/http"

	"github.com/tracer/mux"
)

// Register registers the API supported by this package.
func Register(r mux.Router) {
	r.POST("/tracers/{id}/metadata", handleMetadata)
}

func handleMetadata(responseWriter http.ResponseWriter, request *http.Request) {

}
