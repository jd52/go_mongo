package page

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//QueryGetHandler calls the query.gohtml.  URL is localhost/query.
func QueryGetHandler(res http.ResponseWriter, req *http.Request, hrP httprouter.Params) {

	err := tpl.ExecuteTemplate(res, "query.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
