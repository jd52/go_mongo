
package mongo

import "log"

type logger struct {
	msg string
	
}

//CheckError checks for errors and preforms action based on input of b
func CheckError(a *error, b string) {
	if b == "panic" {
		if a != nil {
			panic(a)
		}
	} else if b == "log_fatal" {
		log.Fatal(a)
	}

