package errorCollector

import (
	"bytes"
	"fmt"
)

type ErrBuf struct {
	Err myErr
}

func GenErrBuf(name []byte) *bytes.Buffer {
	ErrBuff := bytes.NewBuffer(name)
	return ErrBuff
}

func CatchError(e error) *bytes.Buffer {
	if e != nil {
		panic(e)
	}
	ErrByte := []byte(fmt.Sprintf("%s \n", e.Error()))
	ErrBuff := bytes.NewBuffer(ErrByte)
	return ErrBuff
}
