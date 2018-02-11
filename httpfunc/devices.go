package httpfunc

import (
	"fmt"
	"go_mongo/database"
	"log"
	"net/http"
)

//DeviceHandler calls the device.gohtml.  URL is localhost/device.
//This page allows the creation of new items to the MongoDB "Device" Collection.
func DeviceHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("DeviceHandler called")
	rh := req.Method
	if rh == "GET" {

		//deviceList := database.ListDevice(res, req)

		err := tpl.ExecuteTemplate(res, "device.gohtml", nil)
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
