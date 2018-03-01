package errorCollector

import (
	"bufio"
	"bytes"
	"fmt"
	"go_mongo/logger"
	"io"
	"os"
	"time"
)

//ErrCollector allows the operator to create a collector that will be used to point to a err handler
type ErrCollector struct {
	MyErrs    []myErr
	ColMode   string
	PanicFile *os.File
	Buffer    *bytes.Buffer
	Writer    *bufio.Writer
}

//ConsumeError receives the incoming and sends it to the error router
func (ec *ErrCollector) ConsumeError(rerr error) {

	var MyErr myErr
	currtime := time.Now().Format("2006-01-02 15:04:05")

	if rerr == nil {
		return
	}
	MyErr.TIME = currtime
	MyErr.err = rerr
	MyErr.ErrString = rerr.Error()
	MyErr.Caller = logger.MyCaller()

	ec.MyErrs = append(ec.MyErrs, MyErr)

	errRouter(ec)
	return
}

//GetErrString allows the operatore to return a string version of the error message
func (ec *ErrCollector) GetErrString() string {
	errString := ec.MyErrs[len(ec.MyErrs)-1].err.Error()
	return errString
}

//GetErrs allows the operator to return a list of errors listed in the cache
func (ec *ErrCollector) GetErrs() string {
	var NS string
	fmt.Println(len(ec.MyErrs))
	for _, errs := range ec.MyErrs {
		fmt.Println(errs)
	}
	return NS
}

//SetErrBuf allows the operator to set the current buffer for the collected errors
func (ec *ErrCollector) SetErrBuf() {
	var MyWr io.Writer
	MyWrE := bufio.NewWriter(MyWr)
	ec.Writer = MyWrE

}
