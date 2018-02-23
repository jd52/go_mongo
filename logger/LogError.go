package logger

import (
	"io/ioutil"
	"log"
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

//NewLogger generates a new logger strcut
func NewLogger(err error, lvl string) Logger {
	newLogger := Logger{Err: err, Lvl: lvl}
	return newLogger
}

//MyLogger is an interface identifies the logger capability
type MyLogger interface {
	writeLog()
}

//Logger is an stuct to build the logger
type Logger struct {
	Err error
	Lvl string
	Msg string
}

//writeLog is a method of writing received log data
func (s Logger) writeLog(FILE string, msg []byte) error {
	// errString := []byte(myerr.Error())
	err := ioutil.WriteFile(FILE, msg, 0644)
	if err != nil {
		return err
	}
	return nil
}

//chckLvl is used to evaluate the results of  &lvl
func checkLvl(lg *Logger, err *error, lvl *string, FILE *string) (bool, error) {
	if *lvl == "INFO" {
		msg := []byte("INFO" + lg.Err.Error())
		lg.writeLog(*FILE, msg)
		panic(*err)
	}
	if *lvl == "WARN" {
		msg := []byte("WARN" + lg.Err.Error())
		lg.writeLog(*FILE, msg)
		panic(*err)
	}
	if *lvl == "DEBUG" {
		msg := []byte("DEBUG" + lg.Err.Error())
		lg.writeLog(*FILE, msg)
		panic(*err)
	}
	if *lvl == "ERROR" {
		msg := []byte("ERROR" + lg.Err.Error())
		lg.writeLog(*FILE, msg)
		panic(*err)
	}
	if *lvl == "FATAL" {
		msg := []byte("FATAL" + lg.Err.Error())
		lg.writeLog(*FILE, msg)
		panic(*err)
	}
	if *lvl == "TRACE" {
		msg := []byte("TRACE" + lg.Err.Error())
		lg.writeLog(*FILE, msg)
		panic(*err)
	}
	if *lvl == "panic" {
		if *err != nil {
			panic(*err)
		}
	}
	if *lvl == "log_fatal" {
		log.Fatal(*err)
	}
	if *lvl == "" {
		return false, true
	}
	return false, false
}

//LogError checks for errors and preforms action based on received logging level
//returns bool based on success of logger
func LogError(err *error, lvl string) bool {
	lg := NewLogger(*err, lvl)
	file := "/tmp/logs/log.txt"
	results, nerr := checkLvl(&lg, &lg.Err, &lvl, &file)
	return results
}
