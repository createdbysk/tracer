package tracers

import (
	"log"
	"net/http"
	"strings"

	"github.com/tracer/mux"
)

// Register registers handlers for routes.
func Register(r mux.Router) {
	r.POST("/tracers", handleTracers)
}

func handleTracers(responseWriter http.ResponseWriter, request *http.Request) {
	log.Print("handleTracers called with ", request)
	u := request.URL
	u.Host = request.Host
	u.Scheme = "http"
	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}
	u.Path += "0"
	responseWriter.Header().Set("Location", u.String())
}
