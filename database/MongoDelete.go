package database

import (
	"log"

	"github.com/globalsign/mgo"
)

//Delete function is currently a placeholder.
func (md *MongoDevice) Delete() {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//err = device.Insert(md)
	if err != nil {
		log.Fatal(err)
	}
}
