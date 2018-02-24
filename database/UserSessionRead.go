package database

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//Read returns session from the sessions Collection
func (se *Session) Read() bool{

	var err error
	session := MongoSession()
	defer session.Close()

	deviceCollect := session.DB("test").C("authsessions")

	session.SetMode(mgo.Monotonic, true)

	if se.SessionID != "" {
		err = deviceCollect.Find(bson.M{"sessionid": bson.RegEx{Pattern: se.SessionID, Options: "i"}}).One(&se)
		if err != nil {
			return false
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	return true
}
