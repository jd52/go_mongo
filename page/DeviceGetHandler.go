package page

import (
	"fmt"
	"go_mongo/mongo"
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

	err = req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	//Received input from the devices.gohtml template and updates addD.
	addD := mongo.Device{
		Hostname:   req.FormValue("hostname"),
		IPAddress:  req.FormValue("ipAddress"),
		DeviceType: req.FormValue("deviceType"),
	}
	f, _, fileErr := req.FormFile("bulkFile")
	if fileErr != nil {
		fmt.Println("in devicehandler if")
		fmt.Println(fileErr)
	} else {
		ad := mongo.ParseCSV(f)
		for _, d := range ad {
			mongo.AddDevice(&d, res, req)
		}
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
