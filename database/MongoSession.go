package database

import (
	"github.com/globalsign/mgo"
)

//MongoSession returns a call to a mongoDB
func MongoSession() *mgo.Session {

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	return session
}
