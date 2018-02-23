package database

import (
	"github.com/globalsign/mgo"
)

//MongoSession returns a call to a mongoDB
func MongoSession() *mgo.Session {

	session, err := mgo.Dial("10.132.0.5")
	// logger.LogError(&err, "test")
	if err != nil {
		panic(err)
	}

	return session
}
