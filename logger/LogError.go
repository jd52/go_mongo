package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
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

//getTempdir is an internal method that gets the cwd and
func getTempdir() (string, error) {
	fmt.Println("getTempdir() processing")
	dir, wdErr := os.Getwd()
	if wdErr != nil {
		return "", wdErr
	}

	newDir := (dir + "/temp")
	fmt.Println("New Direcotry")
	fmt.Println(newDir)
	fmt.Println("Created Directory")
	os.Mkdir("temp", 0777)
	return newDir, wdErr

}

//CreateLogFile is a helper function used to create new log files if once is not present
func createLogFile(filename *string) error {
	newWords := []byte("New File\n")
	err := ioutil.WriteFile(*filename, newWords, 0666)
	return err

}

//chckLvl is an internal func supporting LogError() used to evaluate the results of &lvl
func checkLvl(lg *Logger, ERR *error, lvl string, FILE *os.File) error {
	var err error
	if lvl == "test" {
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		// fmt.Println(callerFunc)
		fmt.Println(lg.CustomMsg)
		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(FILE, &msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(FILE, &msg)
		if err != nil {
			return err
		}
		return err
	}

	return err
}

//LogError checks for errors and preforms action based on received logging level
//returns error
func LogError(ERR *error, lvl string) error {
	lg := NewLogger(*ERR, lvl)
	lg.Caller = MyCaller()
	lg.LogFile = "log.txt"
	lg.LogCustom = false
	lg.LogDir = "temp"

	var err error

	//Retrevies the Current Working Directory
	lg.CwDir, err = lg.GetCwDir()
	if err != nil {
		return err
	}

	//Sets the log directory based on the lg.LogFile Decleration
	lg.TempDir = lg.setLogdir()

	//Validates that the log directory exsists, if not creates the temp dir
	_, err = os.Stat(lg.TempDir)
	if err != nil {
		tempDir, tempErr := getTempdir()
		if tempErr != nil {
			return tempErr
		}
		lg.TempDir = tempDir
	}

	//Sets the Fully Qualified File Name(FQFN) of the Log Dir and Log File
	lg.FileFQN = lg.setFilefqn()

	//Validates the the FQFN exsists, f not it creates the File based on the lg.LogFile decleration
	_, err = os.Stat(lg.FileFQN)
	if err != nil {
		err = createLogFile(&lg.FileFQN)
		if err != nil {
			return err
		}

		return err
	}

	//Opens the log file for Writing
	file, err := lg.openFile(&lg.FileFQN)
	if err != nil {
		return err
	}
	lg.File = file

	//Validates received Level
	err = checkLvl(&lg, &lg.Err, lvl, &lg.File)
	if err != nil {
		return err
	}
	return err

}
