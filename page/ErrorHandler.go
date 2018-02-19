package page

import (
	"log"
	"net/http"
)

//ErrorHandler calls the error.gohmtl. URL is localhost/index
func ErrorHandler(res http.ResponseWriter, req *http.Request) {
	//fmt.Println("IndexHandler called")
	rh := req.Method
	if rh == "GET" {

		err := tpl.ExecuteTemplate(res, "error.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
