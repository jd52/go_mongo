package errorCollector

func defaultHandler(e ErrCollector) {
	switch e.MyErrs[len(e.MyErrs)-1] {
	default:
		panicErrCollector(e)

	}

}
