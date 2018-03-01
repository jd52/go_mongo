package database

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Delete function receives a pointer of MongoDevice and deletes the ObjectID
//from the MongoDB database.
func (md *MongoDevice) Delete() {
	var err error

	// Creates a session to the mongoDB using the preconfigurations found in
	// database.MongoSession.
	session := MongoSession()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// The received md.ID needs to be converted to Hex before it can be used by
	// the func (c *Collection) RemoveAll(selector interface{}).  This conversion
	// is done by database.FormatMongoID.
	id := bson.ObjectIdHex(FormatMongoID(md.ID))

	// Uses the mongoDB session to remove documents solely by the mongoDB "_id".
	err = session.DB("test").C("device").RemoveId(id)

	if err != nil {
		fmt.Println(err)
	}
	return
}

//MongoDelete function receives a slice of string, converts the string to
//hex,and deletes the ObjectID from the MongoDB database.
func MongoDelete(id []string) {
	var err error

	// Creates a session to the mongoDB using the preconfigurations found in
	// database.MongoSession.
	session := MongoSession()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// The received md.ID needs to be converted to Hex before it can be used by
	// the func (c *Collection) RemoveAll(selector interface{}).  This conversion
	// is done by database.FormatMongoID.
	for _, i := range id {
		hex := bson.ObjectIdHex(FormatMongoID(i))

		// Uses the mongoDB session to remove documents solely by the mongoDB "_id".
		err = session.DB("test").C("device").RemoveId(hex)
	}
	if err != nil {
		fmt.Println(err)
	}
	return
}
