package mongo

import (
	"fmt"
	"log"

	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//ListDevice returns all hostnames from the Device Collection
func ListDevice(qy *Device, andOr *string, w http.ResponseWriter, r *http.Request) []Device {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	result := []Device{}

	if *andOr == "and" {

		err = deviceCollect.Find(bson.M{"$and": []bson.M{bson.M{"hostname": qy.Hostname}, bson.M{"ipaddress": qy.IPAddress}, bson.M{"devicetype": qy.DeviceType}}}).All(&result)
		fmt.Println(qy)
	} else {

		err = deviceCollect.Find(bson.M{"$or": []bson.M{bson.M{"hostname": "//" + qy.Hostname + "//"}, bson.M{"ipaddress": qy.IPAddress}, bson.M{"devicetype": qy.DeviceType}}}).All(&result)
		fmt.Println(qy)
		fmt.Println(bson.M{"hostname": "/" + qy.Hostname + "/"})
		fmt.Println([]bson.M{bson.M{"hostname": "/" + qy.Hostname + "/"}})
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}
