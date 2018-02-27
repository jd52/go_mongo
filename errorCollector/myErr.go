package errorCollector

type myErr struct {
	TIME      string
	err       error
	ErrString string
	Caller    string
}

//Allows  this type to indetify as an io.Writer interface type
func (ec myErr) Write(p []byte) (n int, err error) {
	return len(p), err
}
