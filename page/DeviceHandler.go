package page

import (
	"go_mongo/mongo"
	"log"
	"net/http"
)

//DeviceHandler calls the device.gohtml.  URL is localhost/device.
//This page allows the creation of new items to the MongoDB "Device" Collection.
func DeviceHandler(res http.ResponseWriter, req *http.Request) {
	rh := req.Method
	deviceList := mongo.ListDevice(res, req)

	if rh == "GET" {
		err := tpl.ExecuteTemplate(res, "devices.gohtml", deviceList)
		if err != nil {
			log.Fatalln(err)
		}

	} else if rh == "POST" {

		addD := mongo.Device{
			Hostname:   req.FormValue("hostname"),
			IPAddress:  req.FormValue("ipAddress"),
			DeviceType: req.FormValue("deviceType"),
		}

		mongo.AddDeviceHandler(addD, res, req)
		err := tpl.ExecuteTemplate(res, "devices.gohtml", deviceList)
		if err != nil {
			log.Fatalln(err)

		}
	} else {
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
