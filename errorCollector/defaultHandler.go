package errorCollector

func defaultHandler(e ErrCollector) {
	lastErr := e.MyErrs[len(e.MyErrs)-1]
	switch lastErr.err != nil {
	default:
		panicErrCollector(e)

	}

}
