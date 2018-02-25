package errorCollector

func routeError(e *ErrCollector) {
	if e.MyErrs[0].Caller == "main.main" {
		testHandler(*e)
	}
}
