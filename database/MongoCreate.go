package database

import (
	"go_mongo/logger"

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
	var err error
	//vars := mux.Vars(r)
	//fmt.Println(vars["hostname"])

	session := MongoSession()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	device := session.DB("test").C("device")
	err = device.Insert(md)
	if err != nil {
		logger.LogError(&err, "test")
	}

}
