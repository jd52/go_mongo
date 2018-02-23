package logger

import (
	"errors"
	"fmt"
	"os"
)

/*
the Logger Package provides support to Write the received err to a log file the
way it works is described below

ActiveApp() Calls the LogError(*error, lvl)
LogError(params)
	Accepts the &error and the string lvl - next
	Will call the NewLogger(*error, lvl ) function to return a new logger struct - next
	Checks if the lvl is equal to [info, warn, error, panic, fatal, debug] - next
	Once lvl is matached creates []byte from  lvl + *error.Error()
	writes the []byte to file

*/

//LogLevel is a struct to refrence Logging levels
type LogLevel struct {
	Info    string `json:"info"`
	Warning string `json:"warning"`
	Err     string `json:"error"`
	Panic   string `json:"panic"`
	Fatal   string `json:"fatal"`
	Debug   string `json:"debug"`
}

//chckLvl is an internal func supporting LogError() used to evaluate the results of &lvl
func checkLvl(lg *Logger, err *error, lvl string, FILE *os.File) (bool, error) {
	logLevels := []string{"info", "warning", "error", "panic", "fatal", "debug"}

	for _, levels := range logLevels {
		fmt.Println(levels)
	}
	// = []string{"info", "warning", "error", "panic", "fatal", "debug"}
	if lvl == "test" {
		msg := []byte("TEST: " + lg.Err.Error() + "\n")
		err := lg.writeLog(FILE, &msg)
		return false, err
	}
	//  if *lvl == "panic" {
	// 	if *err != nil {
	// 		panic(*err)
	// 	}
	// } else if *lvl == "log_fatal" {
	// 	log.Fatal(*err)
	// } else if *lvl == "" {
	// 	return false, nil
	// }
	return false, nil
}

//LogError checks for errors and preforms action based on received logging level
//returns bool based on success of logger and standard error
func LogError(err *error, lvl string) (bool, error) {
	lg := NewLogger(*err, lvl)
	dir := lg.getTempdir()
	file := (dir + "/log.txt")
	DefaultError := errors.New("LogError() unable to process log")

	var results bool
	results = lg.chkFile(&file)

	if results == true {
		fmt.Println("LogError() working")
		file, _ := lg.openFile(&file)
		chkBool, chkErr := checkLvl(&lg, &lg.Err, lvl, &file)
		if chkBool == true {
			return chkBool, chkErr
		}

		return false, DefaultError
	}

	return false, *err
}
