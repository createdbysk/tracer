package adapter

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/tracer/router"
)

// This test passes if it can compile
func TestGorillaMuxRouter(t *testing.T) {
	muxRouter := mux.NewRouter()
	adapter := GorillaMuxRouter{muxRouter}
	// Confirm that the adapter can be passed to a function that expects
	// a router.Router interface.
	testCompile(adapter)
}

func testCompile(router router.Router) {}
