package database

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//Read returns all hostnames from the Device Collection
func (us *User) Read() User {

	var err error
	session := MongoSession()
	defer session.Close()

	deviceCollect := session.DB("test").C("users")

	session.SetMode(mgo.Monotonic, true)

	result := User{}

	if us.Username != "" {
		err = deviceCollect.Find(bson.M{"username": bson.RegEx{Pattern: us.Username, Options: "i"}}).All(&result)
	}

	if err != nil {
		log.Fatal(err)
	}
	return result
}
