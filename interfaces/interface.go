package interfaces

import (
	"bytes"
	"io"
	"strings"
)

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}
