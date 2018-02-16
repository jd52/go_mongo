package mongo

import (
	"fmt"
	"log"

	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type dInt interface {
	makeMap() map[string]string
}

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
	mongoMap := qy.makeMongoString()
	if qy.Hostname+qy.DeviceType+qy.IPAddress == "" {
		err = deviceCollect.Find(bson.M{}).All(&result)
	} else {

		err = deviceCollect.Find(mongoMap["hostname"]).All(&result)
		fmt.Println(qy)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result

}

func (d *Device) makeMongoString() map[string]string {
	newMap := map[string]string{"hostname": d.Hostname}
	newMap["ipaddress"] = d.IPAddress
	newMap["devicetype"] = d.DeviceType

	return newMap
}
