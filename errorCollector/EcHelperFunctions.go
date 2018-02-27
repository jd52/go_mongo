package errorCollector

import (
	"errors"
	"fmt"
	"os"
	"time"
)

//Function is used within handlers to provide support incase the error itself
//experiences  a error that it cant not handle correctly
//This Function MUST BE PRESENT at the end of every CustomerHandler that is created
func panicErrCollector(e ErrCollector) {
	lastErr := len(e.MyErrs) - 1
	if e.MyErrs[lastErr].err != nil {
		currtime := time.Now().Format("2006-01-02")
		errString := fmt.Sprintf("PANIC - TIME: %s FUNC: %s ERROR:%s\n", currtime, e.MyErrs[len(e.MyErrs)-1].Caller, e.MyErrs[len(e.MyErrs)-1].ErrString)
		tempDir, _ := os.Getwd()
		panicFileName := fmt.Sprintf("%s/dumps/%s_%s", tempDir, currtime, "PANIC_FILE")

		//Validates that the dump folder is present on the disk, if not it will create the directory and file needed
		//for the  panic dump file.
		_, err := os.Stat("dumps")
		if err != nil {
			err = os.Mkdir("dumps", 0777)
			if err != nil {
				panic(errors.New("Default Handler Unable to Create Directory"))
			}

			e.PanicFile, err = os.OpenFile(panicFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				panic(errors.New("Default Handler Unable to Create File"))
			}

			num, err := e.PanicFile.Write([]byte(errString))
			fmt.Println(num)
			if err != nil {
				panic(errors.New("Default Handler Unable to Write File"))
			}
			fmt.Printf("PANIC - TRACE_FILE:\n\t%s\n", e.PanicFile.Name())
			defer e.PanicFile.Close()
			return
		}

		e.PanicFile, err = os.OpenFile(panicFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			panic(errors.New("Default Handler Unable to Create File"))
		}

		num, err := e.PanicFile.Write([]byte(errString))
		fmt.Println(num)
		if err != nil {
			panic(errors.New("Default Handler Unable to Write File"))
		}
		fmt.Printf("PANIC - TRACE_FILE:\n\t%s\n", e.PanicFile.Name())
		defer e.PanicFile.Close()
		return

	}
}
