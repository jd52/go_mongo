package page

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

//IndexHandler calls the index.gohmtl. URL is localhost/index
func IndexHandler(res http.ResponseWriter, req *http.Request) {
	//fmt.Println("IndexHandler called")
	rh := req.Method
	if rh == "GET" {

		err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
