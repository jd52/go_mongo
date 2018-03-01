package errorCollector

import (
	"bytes"
)

//GenErrBuf allows the operator to generate a new buffer
func GenErrBuf(name []byte) *bytes.Buffer {
	ErrBuff := bytes.NewBuffer(name)
	return ErrBuff
}
