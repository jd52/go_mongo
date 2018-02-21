package database

import (
	"fmt"
	"log"

	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//Read returns all hostnames from the Device Collection
func Read(qy *MongoDevice, andOr *string, w http.ResponseWriter, r *http.Request) []MongoDevice {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	result := []MongoDevice{}
	if *andOr == "and" {

		err = deviceCollect.Find(bson.M{"$and": []bson.M{bson.M{"hostname": qy.Hostname}, bson.M{"ipaddress": qy.IPAddress}, bson.M{"devicetype": qy.DeviceType}}}).All(&result)
		fmt.Println(qy)
	} else {
		if qy.Hostname != "" {
			err = deviceCollect.Find(bson.M{"hostname": bson.RegEx{Pattern: qy.Hostname, Options: "i"}}).All(&result)
		}
		if qy.IPAddress != "" {
			err = deviceCollect.Find(bson.M{"ipaddress": bson.RegEx{Pattern: qy.IPAddress, Options: "i"}}).All(&result)
		}
		if qy.DeviceType != "" {
			err = deviceCollect.Find(bson.M{"devicetype": bson.RegEx{Pattern: qy.DeviceType, Options: "i"}}).All(&result)
		}

		fmt.Println(qy)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Printing Results: ", result)
	return result
}
