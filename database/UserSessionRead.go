package database

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//Read returns all hostnames from the Device Collection
func (se *Session) Read() (Session,bool) {

	var err error
	session := MongoSession()
	defer session.Close()

	deviceCollect := session.DB("test").C("authsessions")

	session.SetMode(mgo.Monotonic, true)

	result := Session{}

	if se.Session != "" {
		err = deviceCollect.Find(bson.M{"session": bson.RegEx{Pattern: se.Session, Options: "i"}}).All(&result)
		if err != nil {
			return result,false
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	return result,true
}
