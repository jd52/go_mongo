package logger

import (
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

//

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
