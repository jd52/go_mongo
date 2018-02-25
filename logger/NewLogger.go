package logger

//NewLogger generates a new logger strcut
func NewLogger(err error, lvl string) Logger {
	logLevels := []string{"info", "warning", "error", "panic", "fatal", "debug"}
	newLogger := Logger{Err: err, Lvl: lvl}

	//Loops though the []loglevels and matches the received lvl
	//Once a match is made the level is updated on the newLogger strut
	for _, level := range logLevels {
		if lvl == level {
			newLogger.Lvl = level
		}

		return newLogger
	}
	return newLogger
}
