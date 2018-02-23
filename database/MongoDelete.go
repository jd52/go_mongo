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

	session := MongoSession()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//The id variable changes the received ID of type string and converts it to a
	//bson hex.
	id := bson.ObjectIdHex(md.ID)

	err = session.DB("test").C("device").RemoveId(id)
	if err != nil {
		fmt.Println(err)
	}
}
