package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Println(vars["hostname"])
	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	device := session.DB("test").C("device")
	err = device.Insert(&Device{"R1", "192.168.1.1", "router"},
		&Device{"SW1", "192.168.1.2", "switch"})
	if err != nil {
		log.Fatal(err)
	}

	result := Device{}
	err = device.Find(bson.M{"hostname": vars["hostname"]}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("IP Address:", result.IPAdress)
	webString := "<html><body><h1>Golang test!\n" + result.IPAdress + " " + result.DeviceType + "\n</body></html>"
	w.Write([]byte(webString))
}

type Person struct {
	Name  string
	Phone string
}

type Device struct {
	Hostname   string
	IPAdress   string
	DeviceType string
}

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/device/{hostname}", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":80", r))
}
