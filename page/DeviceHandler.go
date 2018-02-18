package page

import (
	"fmt"
	"go_mongo/mongo"
	"log"
	"net/http"
)

//DeviceHandler calls the device.gohtml.  URL is localhost/device.
//This page allows the creation of new items to the MongoDB "Device" Collection.
func DeviceHandler(res http.ResponseWriter, req *http.Request) {
	rh := req.Method

	//empty := mongo.Device{}
	//andOr := "or"
	//Passing the empty struct into mongo.ListDevice returns the whole list.
	//deviceList := mongo.ListDevice(&empty, &andOr, res, req)
	if rh == "GET" {
		err := tpl.ExecuteTemplate(res, "devices.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}

	} else {
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		//Received input from the devices.gohtml template and updates addD.
		addD := mongo.Device{
			Hostname:   req.FormValue("hostname"),
			IPAddress:  req.FormValue("ipAddress"),
			DeviceType: req.FormValue("deviceType"),
		}
		f, _, fileErr := req.FormFile("bulkfile")
		if fileErr == nil {
			fmt.Println("in devicehandler if")
			mongo.ParseCSV(f)
		} else {

			fmt.Println("in devicehandler else")
		}
		//Sends the received input and sends it to mongo.AddDevice to add a new
		//entry into the mongoDB database.
		validate := mongo.ValidateAdd(&addD, res, req)

		if validate == false {
			mongo.AddDevice(&addD, res, req)
		}
		//deviceList := mongo.ListDevice(&empty, &andOr, res, req)
		err = tpl.ExecuteTemplate(res, "devices.gohtml", validate)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
