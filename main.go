package main

import (
	"log"
	"net/http"
	"fmt"

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

	port := 8080
	fmt.Printf(":%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
