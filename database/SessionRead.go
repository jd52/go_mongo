package database

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//Read returns all hostnames from the Device Collection
func (se *Session) Read() Session {

	var err error
	session := MongoSession()
	defer session.Close()

	deviceCollect := session.DB("test").C("authsessions")

	session.SetMode(mgo.Monotonic, true)

	result := Session{}

	if se.Username != "" {
		err = deviceCollect.Find(bson.M{"username": bson.RegEx{Pattern: se.Username, Options: "i"}}).All(&result)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}
