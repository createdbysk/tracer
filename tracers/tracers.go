package tracers

import (
	"fmt"
	"net/http"

	"github.com/tracer/router"
)

// Register registers handlers for routes.
func Register(r router.Router) {
	r.HandleFunc("/tracers/", handleTracers)
}

func handleTracers(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("handleTracers called with ", request)
}
