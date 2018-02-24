package main

import (
	"fmt"
	"go_mongo/page"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

func Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}

var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

func main() {

	user := "gordon"
	pass := "secret!"

	router := httprouter.New()
	//fmt.Println("Router created")
	// Routes consist of a path and a handler function.
	router.GET("/", page.IndexHandler)
	router.GET("/index", page.IndexHandler)
	router.GET("/index2", index2)
	router.GET("/protected", BasicAuth(Protected, user, pass))
	router.GET("/devices", page.DeviceGetHandler)
	router.POST("/devices", page.DevicePostHandler)
	router.GET("/query", page.QueryGetHandler)
	router.POST("/query", page.QueryPostHandler)
	router.POST("/queryresponse", page.QueryResponseHandler)
	router.GET("/bar", bar)
	router.GET("/signup", signup)
	router.POST("/signup", signup)
	//r.HandleFunc("/error", page.ErrorHandler).Methods("POST")

	//Serviles allows all files in the public folder to be served
	router.ServeFiles("/public/*filepath", http.Dir("./public"))

	// Bind to a port and pass our router in
	//log.Fatal(http.ListenAndServeTLS(":8080", "/etc/letsencrypt/live/gomoje.com/fullchain.pem", "/etc/letsencrypt/live/gomoje.com/privkey.pem", r))
	log.Fatal(http.ListenAndServe(":8080", router))

}
