package errorCollector

//the Router takes the last entry of the ErrorCollectors error caller and routes based on the condition specefied,
//Default operation will utilzed the default handler
func errRouter(e *ErrCollector) {
	if e.MyErrs[len(e.MyErrs)-1].Caller == "main.main" {
		testHandler(*e)
		return
	}
	//This Routing Entry Needs to be
	if len(e.MyErrs[len(e.MyErrs)-1].Caller) > 0 {
		defaultHandler(*e)
		return
	}
}
