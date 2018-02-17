package mongo

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
	Hostname   string `bson:"hostname,omitempty"`
	IPAddress  string `bson:"ipaddress,omitempty"`
	DeviceType string `bson:"devicetype,omitempty"`
}

//AddDevice opens a session to the mongoDB database and adds a type
//Device.
func AddDevice(a *Device, w http.ResponseWriter, r *http.Request) {
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

//ValidateAdd is used to determind if an entry already exist in the database.
func ValidateAdd(a *Device, w http.ResponseWriter, r *http.Request) bool {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	result := []Device{}

	err = deviceCollect.Find(bson.M{"hostname": a.Hostname, "ipaddress": a.IPAddress, "devicetype": a.DeviceType}).All(&result)

	if err != nil {
		log.Fatal(err)
	}
	if len(result) > 0 {
		return false
	}

	return true

}
