package page

import (
	"log"
	"net/http"
)

//DeviceGetHandler calls the device.gohtml.  URL is localhost/device.
//This page allows the creation of new items to the MongoDB "Device" Collection.
func DeviceGetHandler(res http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(res, "devices.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
