package main

import (
	"fmt"
	"go_mongo/database"
	"net/http"

	"github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) database.User {
	// get cookie
	c, err := req.Cookie("gomoje.comsession")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "gomoje.comsession",
			Value: sID.String(),
		}
		fmt.Println(c)

	}

	http.SetCookie(w, c)

	// if the user exists already, get user
	var u database.User
	var s database.Session

	s.SessionID = c.Value

	if s.Read() {
		u.Username = s.Username
		u.Read()
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("gomoje.comsession")
	if err != nil {
		return false
	}
	
	var s database.Session
	
	s.SessionID = c.Value
	ok := s.Read()
	
	return ok
}
