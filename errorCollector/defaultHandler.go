package errorCollector

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func defaultHandler(e ErrCollector) {
	if e.MyErrs[len(e.MyErrs)-1].err != nil {
		currtime := time.Now().Format("2006-01-02")
		errString := fmt.Sprintf("PANIC - TIME: %s FUNC: %s ERROR:%s", currtime, e.MyErrs[len(e.MyErrs)-1].Caller, e.MyErrs[len(e.MyErrs)-1].ErrString)
		tempDir, _ := os.Getwd()
		panicFile := fmt.Sprintf("%s/dumps/%s_%s", tempDir, currtime, "PANIC_FILE")
		_, err := os.Stat("dumps")
		if err != nil {
			err = os.Mkdir("dumps", 0777)
			if err != nil {
				panic(errors.New("Default Handler Unable to Create Directory"))
			}
			fmt.Println("Stat Failed")
		}
		err = ioutil.WriteFile(panicFile, []byte(errString), 0666)
		if err != nil {
			fmt.Println(err)
			fmt.Println(panicFile)
			panic(errors.New("Default Handler Unable to Write File"))
		}
		fmt.Printf("PANIC - TRACE_FILE:\n\t%s\n", panicFile)

	}
}
