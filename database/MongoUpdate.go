package database

import (
	"log"

	"github.com/globalsign/mgo"
)

//Update function is currently a placeholder.
func (md *MongoDevice) Update() {
	var err error
	session := MongoSession()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	device := session.DB("test").C("device")
	err = device.Insert(md)
	if err != nil {
		log.Fatal(err)
	}
}
