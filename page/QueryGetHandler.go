package page

import (
	"log"
	"net/http"
)

//QueryGetHandler calls the query.gohtml.  URL is localhost/query.
func QueryGetHandler(res http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(res, "query.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
