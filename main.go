package main

import (
	"go_mongo/page"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	//fmt.Println("Router created")
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", page.IndexHandler)
	r.HandleFunc("/index", page.IndexHandler)
	r.HandleFunc("/devices", page.DeviceHandler)
	r.HandleFunc("/query", page.QueryHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("css"))))

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
