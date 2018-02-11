package httpfunc

import (
	"log"
	"net/http"

	"go_mongo/database"
)

//DeviceHandler calls the device.gohtml.  URL is localhost/device.
//This page allows the creation of new items to the MongoDB "Device" Collection.
func DeviceHandler(res http.ResponseWriter, req *http.Request) {

	rh := req.Method
	if rh == "GET" {

		deviceList := database.ListDevice(res, req)

		err := tpl.ExecuteTemplate(res, "device.gohtml", deviceList)
		if err != nil {
			log.Fatalln(err)
		}

	} else {
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		addD := database.Device{
			Hostname:   req.FormValue("hostname"),
			IPAddress:  req.FormValue("ipAddress"),
			DeviceType: req.FormValue("deviceType"),
		}

		database.AddDeviceHandler(addD, res, req)
	}

}
