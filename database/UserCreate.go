package database

import (
	"log"

	"github.com/globalsign/mgo"
)

//User takes three strings, "Hostname", "IPAddress", and "Device".  Used
//to added new devices to the database
type User struct {
	Username string `bson:"username"`
	Password []byte `bson:"password"`
}

//Create opens a session to the mongoDB database and adds a type
//Device.
func (us *User) Create() {
	var err error

	session := MongoSession()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	device := session.DB("test").C("users")
	err = device.Insert(us)
	if err != nil {
		log.Fatal(err)
	}
}
