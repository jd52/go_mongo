package mongo

import (
	"log"

	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ListDevice returns all hostnames from the Device Collection
func ListDevice(qy *Device, w http.ResponseWriter, r *http.Request) []Device {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	result := []Device{}

	if qy.Hostname == "" {
		err = deviceCollect.Find(bson.M{}).All(&result)
	} else {
		err = deviceCollect.Find(bson.M{"hostname": qy.Hostname}).All(&result)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result

}
