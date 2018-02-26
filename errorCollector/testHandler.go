package errorCollector

import (
	"fmt"
	"go_mongo/logger"
)

func testHandler(e ErrCollector) {
	if e.MyErrs[0].err != nil {
		// fmt.Println(e.MyErrs[0].ErrString)
		fmt.Println("Test Handler")
		logger.LogError(&e.MyErrs[0].err, "test")
	}
}
