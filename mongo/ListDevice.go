package mongo

import (
	"fmt"
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
	mongoMap := qy.makeMongoString()
	if qy.Hostname+qy.DeviceType+qy.IPAddress == "" {
		err = deviceCollect.Find(bson.M{}).All(&result)
	} else {
		bsO, _ := bson.Marshal(&mongoMap)
		err = deviceCollect.Find(bsO).All(&result)
		fmt.Println(qy)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result

}

func (d *Device) makeMongoString() mongoDevice {
	newDev := mongoDevice{
		Hostname:   map[string]string{"hostname": d.Hostname},
		IPAddress:  map[string]string{"ipaddress": d.IPAddress},
		DeviceType: map[string]string{"devicetype": d.DeviceType},
	}

	return newDev
}

type mongoDevice struct {
	Hostname   map[string]string `bson:"hostname"`
	IPAddress  map[string]string `bson:"ipaddress"`
	DeviceType map[string]string `bson:"devicetype"`
}

type mdInt interface {
	makeMongoString() mongoDevice
}
