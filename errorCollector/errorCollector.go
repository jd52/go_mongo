package errorCollector

import (
	"fmt"
	"go_mongo/logger"
	"time"
)

//ErrCollector allows the operator to create a collector that will be used to point to a err handler
type ErrCollector struct {
	MyErrs  []myErr
	ColMode string
}

type myErr struct {
	TIME      string
	err       error
	ErrString string
	Caller    string
}

//SwallowErr receives the incoming and sends it to the error router
func (ec *ErrCollector) SwallowErr(rerr error) {
	var FRESH myErr
	currtime := time.Now().Format("2006-01-02 15:04:05")
	if rerr == nil {
		return
	}
	FRESH.TIME = currtime
	FRESH.err = rerr
	FRESH.ErrString = rerr.Error()
	FRESH.Caller = logger.MyCaller()

	// e.ErrList = append(e.ErrList, e.Err)
	ec.MyErrs = append(ec.MyErrs, FRESH)
	errRouter(ec)
	return
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
