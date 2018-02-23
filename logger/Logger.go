package logger

import (
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

//Logger is an stuct to build the logger
type Logger struct {
	Err error
	Lvl string
	Msg string
}

//writeLog is an internal method of writing received log data
func (s Logger) writeLog(FILE *os.File, msg *[]byte) error {
	fmt.Println("It Works")
	_, err := FILE.Write(*msg)
	defer FILE.Close()
	return err
}

func (s Logger) chkFile(fileName *string) bool {
	os.Chdir(*fileName)
	// fmt.Println(*dir)
	_, statErr := os.Stat(*fileName)
	if statErr == nil {
		fmt.Println("chkFile() true")
		return true

	}

	fmt.Println("chkFile() false")
	return false
}

func (s Logger) getTempdir() string {
	dir, _ := os.Getwd()
	dir = (dir + "/tmp")
	return dir
}

func (s Logger) openFile(fileName *string) (os.File, error) {
	fmt.Println("openFile()")
	fmt.Println(*fileName)
	fileResult, fileErr := os.OpenFile(*fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return *fileResult, fileErr
}
