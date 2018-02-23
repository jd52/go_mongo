package page

import (
	"fmt"
	"go_mongo/database"
	"log"
	"net/http"
)

//QueryPostHandler calls the query.gohtml after recieving a post. URL is localhost/query.
func QueryPostHandler(res http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	//onlyID converts the received value from the webpage and removes the
	//ObjectIdHex("ID") wrapper because it is received as a string. Once it
	//is sent to Delete() it is converted back to a ObjectIdHex.
	onlyID := req.FormValue("_id")[13 : len(req.FormValue("_id"))-2]

	queryDevice := database.MongoDevice{
		Hostname:   req.FormValue("hostname"),
		IPAddress:  req.FormValue("ipAddress"),
		DeviceType: req.FormValue("deviceType"),
		ID:         onlyID,
	}

	if queryDevice.ID != "" {
		fmt.Println(queryDevice.ID)
		database.StorageDelete(&queryDevice)
	}

	deviceList := database.StorageRead(&queryDevice)
	err = tpl.ExecuteTemplate(res, "query.gohtml", deviceList)
	if err != nil {
		log.Fatalln(err)
	}

}
