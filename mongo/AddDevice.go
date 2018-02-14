package mongo

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

//Device takes three strings, "Hostname", "IPAddress", and "Device".  Used
//to added new devices to the database
type Device struct {
	Hostname   string `json:"hostname"`
	IPAddress  string `json:"ipaddress"`
	DeviceType string `json:"devicetype"`
}

//AddDevice opens a session to the mongoDB database and adds a type
//Device.
func AddDevice(a Device, w http.ResponseWriter, r *http.Request) {
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
	err = device.Insert(a)
	if err != nil {
		log.Fatal(err)
	}
}
