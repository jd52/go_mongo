package main

import (
	"go_mongo/httpfunc"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	//fmt.Println("Router created")
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", httpfunc.IndexHandler)
	r.HandleFunc("/index", httpfunc.IndexHandler)
	r.HandleFunc("/devices", httpfunc.DeviceHandler)
	r.HandleFunc("/query", httpfunc.DeviceHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
