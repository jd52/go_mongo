package database

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Device takes three strings, "Hostname", "IPAddress", and "Device".  Used
//to added new devices to the database
type Device struct {
	Hostname   string
	IPAddress  string
	DeviceType string
}

//MgoDeviceHandler opens a session to the mongoDB database and adds a type
//Device.
func AddDeviceHandler(a Device, w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddDeviceHandler called")
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

//ListDevice returns all hostnames from the Device Collection
func ListDevice(w http.ResponseWriter, r *http.Request) string {

	fmt.Println("ListDevice called")
	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	result := Device{}
	err = deviceCollect.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Hostname

}
