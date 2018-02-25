package page

import (
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

	// Creates a database.MongoDevice using input from the form input
	// on the queryResponse.gohtml page.
	queryDevice := database.MongoDevice{
		Hostname:   req.FormValue("hostname"),
		IPAddress:  req.FormValue("ipAddress"),
		DeviceType: req.FormValue("deviceType"),
		ID:         req.FormValue("_id"),
	}

	// Checks to see if the delete checkbox has been selected.  If it has been
	// selected, then queryDevice.ID will not be empty.
	if queryDevice.ID != "" {
		multi := req.Form["_id"]
		for _, id := range multi {
			deleteM := database.MongoDevice{ID: id}
			database.StorageDelete(&deleteM)
		}
	}

	// If the delete checkbox has not been selected, a simple query to the mongoDB
	// will be made using the newly created database.MongoDevice.
	deviceList := database.StorageRead(&queryDevice)
	err = tpl.ExecuteTemplate(res, "query.gohtml", deviceList)
	if err != nil {
		log.Fatalln(err)
	}

}
