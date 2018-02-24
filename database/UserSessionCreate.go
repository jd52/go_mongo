package database

import (
	"log"

	"github.com/globalsign/mgo"
)

//Session struct stores an authenticated username to its matching session uuid
type Session struct {
	Username string `bson:"username"`
	SessionID  string `bson:"sessionid"`
	SessionExist bool `bson:"sessionexist"`
}

//Create opens a session to the mongoDB database and adds a type
//Device.
func (se *Session) Create() {
	var err error

	session := MongoSession()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	device := session.DB("test").C("authsessions")
	err = device.Insert(se)
	if err != nil {
		log.Fatal(err)
	}
}
