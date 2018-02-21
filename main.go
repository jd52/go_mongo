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
	r.HandleFunc("/devices", page.DeviceGetHandler).Methods("GET")
	r.HandleFunc("/devices", page.DevicePostHandler).Methods("POST")
	r.HandleFunc("/query", page.QueryGetHandler).Methods("GET")
	r.HandleFunc("/query", page.QueryPostHandler).Methods("POST")
	r.HandleFunc("/queryresponse", page.QueryResponseHandler).Methods("POST")
	//r.HandleFunc("/error", page.ErrorHandler).Methods("POST")

	//PathPrefix allows local files to be served
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// Bind to a port and pass our router in
	//log.Fatal(http.ListenAndServeTLS(":8080", "/etc/letsencrypt/live/gomoje.com/fullchain.pem", "/etc/letsencrypt/live/gomoje.com/privkey.pem", r))
	log.Fatal(http.ListenAndServe(":8080", r))
}
