package mongo

import (
	"log"

	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ListDevice returns all hostnames from the Device Collection
func ListDevice(qy Device, w http.ResponseWriter, r *http.Request) []Device {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	result := []Device{}

	qy.marshallJSON
	if qy.Hostname != "" {
		err = deviceCollect.Find(bson.M{}).All(&result)
	} else {
		err = deviceCollect.Find(bson.D{qy.marshallJSON}).All(&result)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result

}

func (d Device) marshallJSON() {

	d, _ := bson.Marshal(d)

}
