package main

import (
	"go_mongo/database"
	"net/http"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func index2(w http.ResponseWriter, req *http.Request, hrP httprouter.Params) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index2.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request, hrP httprouter.Params) {
	u := getUser(w, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/index2", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request, hrP httprouter.Params) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/index2", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		var u database.User

		// get form values
		u.Username = req.FormValue("username")
		bcPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u.Password = bcPass

		// username taken?
		if _, ok := u.Read(); ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		var s database.Session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "gomoje.comsession",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		s.Session = c.Value
		s.Username = u.Username
		err := s.Create

		// store user in dbUsers

		u.Create()

		// redirect
		http.Redirect(w, req, "/index2", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
