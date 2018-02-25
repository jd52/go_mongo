package errorCollector

import "fmt"

func defaultHandler(e ErrCollector) {
	if e.MyErrs[0].err != nil {
		// fmt.Println(e.MyErrs[0].ErrString)
		fmt.Println("Default Handler")
	}
}
