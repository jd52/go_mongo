package main

import (
	"go_mongo/page"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	//fmt.Println("Router created")
	// Routes consist of a path and a handler function.
	router.GET("/", page.IndexHandler)
	router.GET("/index", page.IndexHandler)
	router.GET("/devices", page.DeviceGetHandler)
	router.POST("/devices", page.DevicePostHandler)
	router.GET("/query", page.QueryGetHandler)
	router.POST("/query", page.QueryPostHandler)
	router.POST("/queryresponse", page.QueryResponseHandler)
	//r.HandleFunc("/error", page.ErrorHandler).Methods("POST")

	//PathPrefix allows local files to be served
	router.ServeFiles("/public/*filepath", http.Dir("./public"))

	// Bind to a port and pass our router in
	//log.Fatal(http.ListenAndServeTLS(":8080", "/etc/letsencrypt/live/gomoje.com/fullchain.pem", "/etc/letsencrypt/live/gomoje.com/privkey.pem", r))
	log.Fatal(http.ListenAndServe(":8080", router))

}
