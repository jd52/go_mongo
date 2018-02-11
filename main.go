package main

import (
	"log"
	"net/http"

	"go_mongo/httpfunc"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", httpfunc.IndexHandler)
	r.HandleFunc("/index", httpfunc.IndexHandler)
	r.HandleFunc("/device", httpfunc.DeviceHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
