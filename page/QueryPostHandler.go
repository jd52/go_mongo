package page

import (
	"go_mongo/database"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//QueryPostHandler calls the query.gohtml after recieving a post. URL is localhost/query.
func QueryPostHandler(res http.ResponseWriter, req *http.Request, hrP httprouter.Params) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	//andOr := req.FormValue("anyOr")
	queryDevice := database.MongoDevice{
		Hostname:   req.FormValue("hostname"),
		IPAddress:  req.FormValue("ipAddress"),
		DeviceType: req.FormValue("deviceType"),
	}
	//
	deviceList := database.StorageRead(&queryDevice)
	err = tpl.ExecuteTemplate(res, "query.gohtml", deviceList)
	if err != nil {
		log.Fatalln(err)
	}
}
