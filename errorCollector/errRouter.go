package errorCollector

//the Router takes the last entry of the ErrorCollectors error caller and routes based on the condition specefied,
//Default operation will utilzed the default handler
func errRouter(e *ErrCollector) {
	lastErr := e.MyErrs[len(e.MyErrs)-1]

	if lastErr.Caller == "main.eee" {
		testHandler(*e)
		return
	}
	//This Routing Entry Needs to be
	if len(lastErr.Caller) > 0 {
		defaultHandler(*e)
		return
	}
}
