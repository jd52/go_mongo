package logger

import "io"

//NewSafeWriter creates a new writer reader
func NewSafeWriter(w io.Writer) *SafeWriter {
	return &SafeWriter{w: w}
}

//SafeWriter contrusts an Error writer pipe to rewrite an received err value
type SafeWriter struct {
	w   io.Writer
	err error
}

//Err method returns the err message to STD Libs Error() method call on a function
func (sw *SafeWriter) Err() error {
	return sw.err
}

func (sw *SafeWriter) Write(p []byte) (n int, err error) {
	if sw.err != nil {
		return 0, err
	}
	n, sw.err = sw.w.Write(p)
	return n, sw.err
}
