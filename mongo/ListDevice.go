package mongo

import (
	"fmt"
	"log"

	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	} else {
		var hn, ipa, dt bson.M
		if qy.Hostname != "" {
			hn = bson.M{"$regex": "'" + qy.Hostname + "'"}
			fmt.Println("set hostname var")
			err = deviceCollect.Find(bson.M{"hostname": hn}).All(&result)
			fmt.Println(bson.M{"hostname": hn})
		}
		if qy.IPAddress != "" {
			ipa = bson.M{"$regex": "'" + qy.IPAddress + "'"}
			fmt.Println("set ipaddress var")
			err = deviceCollect.Find(bson.M{"ipaddress": ipa}).All(&result)
			fmt.Println(bson.M{"ipaddress": ipa})
		}
		if qy.DeviceType != "" {
			dt = bson.M{"$regex": "'" + qy.DeviceType + "'"}
			fmt.Println("set devicetype var")
			err = deviceCollect.Find(bson.M{"devicetype": dt}).All(&result)
			fmt.Println(bson.M{"devicetype": dt})
		}

		fmt.Println(result)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result

}
