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

	if qy.Hostname == "" {
		err = deviceCollect.Find(bson.M{}).All(&result)
	} else {
		_, ipA, _ := qy.makeMongoString()
		err = deviceCollect.Find(ipA).All(&result)
		fmt.Println(qy)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result

}

func (d *Device) makeMongoString() (map[string]string, map[string]string, map[string]string) {
	hostN := map[string]string{"hostname": d.Hostname}

	ipA := map[string]string{"ipaddress": d.IPAddress}
	devT := map[string]string{"devicetype": d.DeviceType}

	return hostN, ipA, devT
}
