package database

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//MongoDevice takes three strings, "Hostname", "IPAddress", and "Device".  Used
//to added new devices to the database
type MongoDevice struct {
	Hostname   string `bson:"hostname,omitempty"`
	IPAddress  string `bson:"ipaddress,omitempty"`
	DeviceType string `bson:"devicetype,omitempty"`
}

//Create opens a session to the mongoDB database and adds a type
//Device.
func (md *MongoDevice) Create() {
	//vars := mux.Vars(r)
	//fmt.Println(vars["hostname"])
	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	device := session.DB("test").C("device")
	err = device.Insert(md)
	if err != nil {
		log.Fatal(err)
	}
}

//ValidateCreate is used to determind if an entry already exist in the database.
func (md *MongoDevice) ValidateCreate() (bool, error) {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		return true, err
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	resultH := []MongoDevice{}
	resultIP := []MongoDevice{}

	err = deviceCollect.Find(bson.M{"hostname": md.Hostname}).All(&resultH)

	if err != nil {
		log.Fatal(err)
	}

	err = deviceCollect.Find(bson.M{"ipaddress": md.IPAddress}).All(&resultIP)

	if err != nil {
		log.Fatal(err)
	}
	if len(resultH) > 0 || len(resultIP) > 0 {
		return true, err
	}

	return false, err

}
