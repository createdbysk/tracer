package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tracer/router/adapter"
	"github.com/tracer/tracers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// Use the GorillaMuxRouter to adapt the mux.Router to provide the
	// router.Router interface that tracers.Register expects.
	adapted := adapter.GorillaMuxRouter{router}
	tracers.Register(adapted)

	log.Fatal(http.ListenAndServe(":8080", router))
}
