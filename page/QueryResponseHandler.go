package page

import (
	"go_mongo/database"
	"log"
	"net/http"
)

//QueryResponseHandler is used for query.
func QueryResponseHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	andOr := req.FormValue("anyOr")
	queryDevice := database.MongoDevice{
		Hostname:   req.FormValue("hostname"),
		IPAddress:  req.FormValue("ipAddress"),
		DeviceType: req.FormValue("deviceType"),
	}
	//
	deviceList := database.Read(&queryDevice, &andOr, res, req)
	err = tpl.ExecuteTemplate(res, "queryResponse.gohtml", deviceList)
	if err != nil {
		log.Fatalln(err)
	}
}
