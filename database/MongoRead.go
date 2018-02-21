package database

import (
	"log"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

//Read returns all hostnames from the Device Collection
func (md *MongoDevice) Read() []Device {

	var err error
	// session, err := mgo.Dial("10.132.0.5")
	// if err != nil {
	// 	panic(err)
	// }
	session := MongoSession()
	defer session.Close()

	deviceCollect := session.DB("test").C("device")

	session.SetMode(mgo.Monotonic, true)

	result := []Device{}
	// if *andOr == "and" {

	// 	err = deviceCollect.Find(bson.M{"$and": []bson.M{bson.M{"hostname": md.Hostname}, bson.M{"ipaddress": md.IPAddress}, bson.M{"devicetype": md.DeviceType}}}).All(&result)
	// } else {

	if md.Hostname != "" {
		err = deviceCollect.Find(bson.M{"hostname": bson.RegEx{Pattern: md.Hostname, Options: "i"}}).All(&result)
	}
	if md.IPAddress != "" {
		err = deviceCollect.Find(bson.M{"ipaddress": bson.RegEx{Pattern: md.IPAddress, Options: "i"}}).All(&result)
	}
	if md.DeviceType != "" {
		err = deviceCollect.Find(bson.M{"devicetype": bson.RegEx{Pattern: md.DeviceType, Options: "i"}}).All(&result)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}
