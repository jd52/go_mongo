package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body><h1>Golang test!\n</body></html>"))
}

import (
        
)

type Person struct {
        Name string
        Phone string
}

func main() {
        session, err := mgo.Dial("10.132.0.5")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	               &Person{"Cla", "+55 53 8402 8510"})
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
	
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":80", r))
}

