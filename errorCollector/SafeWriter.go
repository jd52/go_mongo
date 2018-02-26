package errorCollector

import "io"

//NewSafeWriter a writer that streams returned error messages
func NewSafeWriter(w io.Writer) *SafeWriter {
	return &SafeWriter{w: w}
}

//SafeWriter a writer that streams returned error messages
type SafeWriter struct {
	w   io.Writer
	err error
}

//Err returns a en error message
func (sw *SafeWriter) Err() error { return sw.err }

func (sw *SafeWriter) Write(p []byte) (n int, err error) {
	// var myCaller = logger.MyCaller()

	// myCaller := logger.MyCaller()
	if sw.err != nil {
		return 0, err
	}
	n, sw.err = sw.w.Write(p)
	return n, sw.err
}
