package page

import (
	"log"
	"net/http"

	"go_mongo/mongo"
)

//QueryPostHandler calls the query.gohtml after recieving a post. URL is localhost/query.
func QueryPostHandler(res http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	andOr := req.FormValue("anyOr")
	queryDevice := mongo.Device{
		Hostname:   req.FormValue("hostname"),
		IPAddress:  req.FormValue("ipAddress"),
		DeviceType: req.FormValue("deviceType"),
	}
	//
	deviceList := mongo.ListDevice(&queryDevice, &andOr, res, req)
	err = tpl.ExecuteTemplate(res, "query.gohtml", deviceList)
	if err != nil {
		log.Fatalln(err)
	}
}
