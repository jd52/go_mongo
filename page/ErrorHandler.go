package page

import (
	"log"
	"net/http"
)

//ErrorHandler calls the error.gohmtl. URL is localhost/index
func ErrorHandler(e error, res http.ResponseWriter, req *http.Request) {
	//fmt.Println("IndexHandler called")

	err := tpl.ExecuteTemplate(res, "error.gohtml", e)
	if err != nil {
		log.Fatalln(err)
	}
}
