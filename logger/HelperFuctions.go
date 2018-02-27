package logger

import (
	"io/ioutil"
	"os"
	"time"
)

//getTempdir is an internal method that gets the cwd and
func getTempdir() (string, error) {
	dir, wdErr := os.Getwd()
	if wdErr != nil {
		return "", wdErr
	}

	newDir := (dir + "/temp")

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
	switch lvl {
	case "test":
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(&msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(&msg)
		if err != nil {
			return err
		}
		return err

	case "info":
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(&msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(&msg)
		if err != nil {
			return err
		}
		return err
	case "warning":
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(&msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(&msg)
		if err != nil {
			return err
		}
		return err

	case "debug":
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(&msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(&msg)
		if err != nil {
			return err
		}
		return err
	case "error":
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(&msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(&msg)
		if err != nil {
			return err
		}
		return err

	case "panic":
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(&msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(&msg)
		if err != nil {
			return err
		}
		return err

	case "fatal":
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(&msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(&msg)
		if err != nil {
			return err
		}
		return err

	default:
		timeNow := time.Now()
		Currtime := timeNow.Format("2006-01-02 15:04:05")

		if lg.LogCustom == false {
			msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.Err.Error() + "\n")
			err := lg.WriteLog(&msg)
			if err != nil {
				return err
			}
		}
		msg := []byte("Level:" + lvl + " " + Currtime + " FUNC:" + lg.Caller + " MSG:" + lg.CustomMsg + "\n")
		err := lg.WriteLog(&msg)
		if err != nil {
			return err
		}
		return err

	}
}
