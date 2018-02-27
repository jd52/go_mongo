package errorCollector

import (
	"bufio"
	"bytes"
	"fmt"
	"go_mongo/logger"
	"io"
	"time"
)

//ErrCollector allows the operator to create a collector that will be used to point to a err handler
type ErrCollector struct {
	MyErrs  []myErr
	ColMode string
	Buffer  *bytes.Buffer
	Writer  *bufio.Writer
}

//SwallowErr receives the incoming and sends it to the error router
func (ec *ErrCollector) SwallowErr(rerr error) {

	var MyErr myErr
	var COLLECT []byte
	currtime := time.Now().Format("2006-01-02 15:04:05")

	ec.Buffer = GenErrBuf(COLLECT)
	// fmt.Println(ec.Buffer.Cap())
	if rerr == nil {
		return
	}
	MyErr.TIME = currtime
	MyErr.err = rerr
	MyErr.ErrString = rerr.Error()
	MyErr.Caller = logger.MyCaller()

	// e.ErrList = append(e.ErrList, e.Err)
	ec.MyErrs = append(ec.MyErrs, MyErr)
	ec.Buffer.WriteTo(MyErr)

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

func (ec *ErrCollector) SetErrBuf() {
	var MyWr io.Writer
	MyWrE := bufio.NewWriter(MyWr)
	ec.Writer = MyWrE

}
