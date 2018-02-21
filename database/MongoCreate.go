package database

import (
	"fmt"
	"log"

	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"

	"github.com/gorilla/mux"
)

//Device takes three strings, "Hostname", "IPAddress", and "Device".  Used
//to added new devices to the database
type MongoDevice struct {
	Hostname   string `bson:"hostname,omitempty"`
	IPAddress  string `bson:"ipaddress,omitempty"`
	DeviceType string `bson:"devicetype,omitempty"`
}

//Create opens a session to the mongoDB database and adds a type
//Device.
func Create(c *MongoDevice, w http.ResponseWriter, r *http.Request) {
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
	err = device.Insert(c)
	if err != nil {
		log.Fatal(err)
	}
}

//ValidateCreate is used to determind if an entry already exist in the database.
func ValidateCreate(qy *MongoDevice, w http.ResponseWriter, r *http.Request) (bool, error) {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		return true, err
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	resultH := []MongoDevice{}
	resultIP := []MongoDevice{}

	err = deviceCollect.Find(bson.M{"hostname": qy.Hostname}).All(&resultH)

	if err != nil {
		log.Fatal(err)
	}

	err = deviceCollect.Find(bson.M{"ipaddress": qy.IPAddress}).All(&resultIP)

	if err != nil {
		log.Fatal(err)
	}
	if len(resultH) > 0 || len(resultIP) > 0 {
		return true, err
	}

	return false, err

}
