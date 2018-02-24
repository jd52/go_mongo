package database

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//Read returns user from the users Collection
func (us *User) Read() bool {

	var err error
	session := MongoSession()
	defer session.Close()

	deviceCollect := session.DB("test").C("users")

	session.SetMode(mgo.Monotonic, true)

	if us.Username != "" {
		err = deviceCollect.Find(bson.M{"username": bson.RegEx{Pattern: us.Username, Options: "i"}}).One(&us)
		if err != nil {
			return true
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	return false
}
