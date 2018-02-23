package logger

import (
	"io/ioutil"
	"log"
	"os"
)

/*
the Logger Package provides support to Write the received err to a log file the
way it works is described below

ActiveApp() Calls the LogError(*error, lvl)
LogError(params)
	Accepts the mem ref of the error and the string of lvl - next
	Will call the NewLogger(*error, lvl ) function to return a new logger struct - next
	Checks if the lvl is equal to [info, warn, error, panic, fatal, debug] - next
	Once lvl is matached creates []byte from  lvl + *error.Error()
	writes the []byte to file

*/

//NewLogger generates a new logger strcut
func NewLogger(err error, lvl string) Logger {
	newLogger := Logger{err: err, lvl: lvl}
	return newLogger
}

//MyLogger is an interface identifies the logger capability
type MyLogger interface {
	writeLog()
}

//Logger is an stuct to build the logger
type Logger struct {
	err error
	lvl string
	msg string
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

func (s Logger) openFile(FILE string) (os.File, error) {
	file, err := os.Open(FILE)
	return *file, err
}

//LogError checks for errors and preforms action based on input of b
func LogError(a *error, lvl string) {
	lg := NewLogger(*a, lvl)
	FILE := "/tmp/logs/log.txt"
	if lvl == "INFO" {
		msg := []byte("INFO" + lg.err.Error())
		lg.writeLog(FILE, msg)
		panic(*a)
	}
	if lvl == "WARN" {
		msg := []byte("WARN" + lg.err.Error())
		lg.writeLog(FILE, msg)
		panic(*a)
	}
	if lvl == "DEBUG" {
		msg := []byte("DEBUG" + lg.err.Error())
		lg.writeLog(FILE, msg)
		panic(*a)
	}
	if lvl == "ERROR" {
		msg := []byte("ERROR" + lg.err.Error())
		lg.writeLog(FILE, msg)
		panic(*a)
	}
	if lvl == "FATAL" {
		msg := []byte("FATAL" + lg.err.Error())
		lg.writeLog(FILE, msg)
		panic(*a)
	}
	if lvl == "TRACE" {
		msg := []byte("TRACE" + lg.err.Error())
		lg.writeLog(FILE, msg)
		panic(*a)
	}
	if lvl == "panic" {
		if a != nil {
			panic(a)
		}
	}
	if lvl == "log_fatal" {
		log.Fatal(*a)
	}
}
