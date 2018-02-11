package main

import (
	"go_mongo/httpfunc"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", httpfunc.IndexHandler)
	r.HandleFunc("/index", httpfunc.IndexHandler)
	r.HandleFunc("/device", httpfunc.DeviceHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
