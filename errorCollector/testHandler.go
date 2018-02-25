package errorCollector

import "fmt"

func testHandler(e ErrCollector) {
	if e.MyErrs[0].err != nil {
		// fmt.Println(e.MyErrs[0].ErrString)
		fmt.Println("Somthing Went Wrong")
	}
}
