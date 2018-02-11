package httpfunc

import (
	"log"
	"net/http"
)

//IndexHandler calls the index.gohmtl. URL is localhost/index
func IndexHandler(res http.ResponseWriter, req *http.Request) {

	rh := req.Method
	if rh == "GET" {

		err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
