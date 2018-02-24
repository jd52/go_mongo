package logger

import (
	"os"
)

/*
The Logger Package provides support to Write the received err or a custom message
to a log file. Two main functions can be used directly with the pacakge.
#####################
LogError(err error, level string) - Used to log captured error message and returns an error if encountered
LogErrorCustom(err error, level string, CustomMsg string) - used to log custom messages returns an error if encountered
#####################
To used the packaged as an abstracted instances

#Import logger package
#Declare Logger type
ex: var lg Logger

#Set LogFile
ex: lg.SetLogFile("test1333.txt") returns error based on if the file was able to be opened

#Write to Log
ex: lg.WriteLog(s []byte) returns error based on if file is able to be written


*/

//Logger is an stuct to build the logger
type Logger struct {
	Err       error
	Lvl       string
	Msg       string
	LogFile   string
	LogDir    string
	CwDir     string
	TempDir   string
	FileFQN   string
	CustomMsg string
	LogCustom bool
	Caller    string
	File      os.File
}

//LogLevel is a struct to refrence Logging levels
type LogLevel struct {
	Info    string `json:"info"`
	Warning string `json:"warning"`
	Err     string `json:"error"`
	Panic   string `json:"panic"`
	Fatal   string `json:"fatal"`
	Debug   string `json:"debug"`
}

//WriteLog is an internal method of writing received log data
func (s Logger) WriteLog(msg *[]byte) error {
	_, err := s.File.Write(*msg)
	if err != nil {
		return err
	}
	defer s.File.Close()
	return err
}

//chkFile is an internal method that checks the status of a given file,
//returns bool if a match is encountered
func (s Logger) chkFile(fileName *string) error {
	_, statErr := os.Stat("temp/" + *fileName)
	if statErr != nil {
		dirErr := os.Chdir(s.FileFQN)
		if dirErr != nil {
			return dirErr
		}

		file, err := os.Create(s.LogFile)
		if err != nil {
			return err
		}
		defer file.Close()
		return err

	}
	return statErr
}

func (s *Logger) openFile(fileName *string) (os.File, error) {
	fileResult, fileErr := os.OpenFile(*fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	return *fileResult, fileErr
}

//SetLogFile is an method that allows the operator to change the default log name
func (s *Logger) SetLogFile(fileName string) error {
	s.LogFile = fileName
	file, err := os.OpenFile(s.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	s.File = *file
	return err
}

//SetLogDir is an method that allows the operator to change the default log directory
func (s *Logger) SetLogDir(dir string) {
	s.LogDir = dir
}

//SetFile is an method that allows the operator to set the os.File struct that will be used
func (s *Logger) SetFile(file os.File) {
	s.File = file
}

//GetCwDir is an method that allows the operator to set the current working director
func (s *Logger) GetCwDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// s.CwDir = dir
	return dir, err
}

//GetDir is an method that allows the operator to change the cwd and set the CwDir attribute
//to the new directory
func (s *Logger) GetDir(dir *string) error {
	err := os.Chdir(*dir)
	if err != nil {
		return err
	}
	s.CwDir, err = os.Getwd()
	return err
}

//setLogdir is a internal method that sets the TempDir attribute
func (s *Logger) setLogdir() string {
	TempDir := (s.CwDir + "/" + s.LogDir)
	return TempDir
}

func (s *Logger) setFilefqn() string {
	FileFQN := (s.CwDir + "/" + s.LogDir + "/" + s.LogFile)
	return FileFQN
}

func (s *Logger) setCustomMsg(msg string) string {
	s.CustomMsg = msg
	return s.CustomMsg
}
