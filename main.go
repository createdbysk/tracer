package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tracer/tracers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	tracers.Register(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
