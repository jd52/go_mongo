package page

import (
	"log"
	"net/http"
)

//UpdatePostHandler calls the deviceinfo.gohtml.  URL is localhost/device.
//This page allows the updating of documents in the MongoDB "Device" Collection.
func UpdatePostHandler(res http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(res, "deviceinfo.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
