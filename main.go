package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tracer/mux"
	"github.com/tracer/tracers"
)

func main() {
	router := mux.NewRouter()
	tracers.Register(router)

	port := 8080
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
