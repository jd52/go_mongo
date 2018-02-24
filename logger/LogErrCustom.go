package logger

import (
	"fmt"
	"os"
)

//LogErrCustom allows the operator to post a custom log message in place of the error
//returns error
func LogErrCustom(ERR *error, lvl string, customMsg string) error {
	fmt.Println("LogErrCustom() Processing")
	lg := NewLogger(*ERR, lvl)
	lg.LogFile = "log.txt"
	lg.LogDir = "temp"
	lg.Caller = MyCaller()
	lg.LogCustom = true
	lg.CustomMsg = lg.setCustomMsg(customMsg)
	fmt.Println(lg.CustomMsg)
	fmt.Println(customMsg)
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
