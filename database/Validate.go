package database

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//Validate is used to determind if an entry already exist in the database.
func (md *MongoDevice) Validate() (bool, error) {

	var err error

	session := MongoSession()
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	resultH := []MongoDevice{}
	resultIP := []MongoDevice{}

	err = deviceCollect.Find(bson.M{"hostname": md.Hostname}).All(&resultH)

	if err != nil {
		log.Fatal(err)
	}

	err = deviceCollect.Find(bson.M{"ipaddress": md.IPAddress}).All(&resultIP)

	if err != nil {
		log.Fatal(err)
	}
	if len(resultH) > 0 || len(resultIP) > 0 {
		return true, err
	}

	return false, err

}
