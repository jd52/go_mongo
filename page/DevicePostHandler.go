package page

import (
	"fmt"
	"go_mongo/database"
	"log"
	"net/http"
)

//DevicePostHandler calls the device.gohtml.  URL is localhost/device.
//This page allows the creation of new items to the MongoDB "Device" Collection.
func DevicePostHandler(res http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	//Received input from the devices.gohtml template and updates addD.
	addD := database.MongoDevice{
		Hostname:   req.FormValue("hostname"),
		IPAddress:  req.FormValue("ipAddress"),
		DeviceType: req.FormValue("deviceType"),
	}
	f, _, fileErr := req.FormFile("bulkFile")
	if fileErr != nil {
		fmt.Println("in devicehandler if")
		fmt.Println(fileErr)
	} else {
		ad := database.ParseCSV(f)
		for _, d := range ad {
			database.StorageCreate(&d)
		}
		fmt.Println("in devicehandler else")
	}
	//Sends the received input and sends it to mongo.AddDevice to add a new
	//entry into the mongoDB database.
	validate, error := database.StorageValidate(&addD)

	if error != nil {
		ErrorHandler(error, res, req)
		return
	}

	if validate == false {
		database.StorageCreate(&addD)
	}
	//deviceList := mongo.ListDevice(&empty, &andOr, res, req)
	err = tpl.ExecuteTemplate(res, "devices.gohtml", validate)
	if err != nil {
		log.Fatalln(err)
	}
}
