package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jd52/go_mongo/httpfunc"
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
