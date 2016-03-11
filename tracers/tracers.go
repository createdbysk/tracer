package tracers

import (
	"log"
	"net/http"

	"github.com/tracer/mux"
)

// Register registers handlers for routes.
func Register(r mux.Router) {
	r.POST("/tracers", handleTracers)
}

func handleTracers(responseWriter http.ResponseWriter, request *http.Request) {
	log.Print("handleTracers called with ", request)
	//if request.Method == "POST" {
	response := ([]byte)("This is a test.")
	responseWriter.Write(response)
	//} else {
	//	http.Error(responseWriter, "Not Allowed", http.StatusMethodNotAllowed)
	//}
}
