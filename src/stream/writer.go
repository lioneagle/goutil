package stream

import (
	"bytes"
	"fmt"

	"chars"
)

type Writer struct {
	buf bytes.Buffer
}

func (this *Writer) Write(format string, args ...interface{}) {
	//this.buf.WriteString(fmt.Sprintf(format, args...))
	fmt.Fprintf(&this.buf, format, args...)
}

func (this *Writer) Writeln(format string, args ...interface{}) {
	//this.buf.WriteString(fmt.Sprintf(format, args...))
	//this.buf.WriteString(fmt.Sprintln())
	fmt.Fprintf(&this.buf, format, args...)
	fmt.Fprintln(&this.buf)
}

func (this *Writer) String() string {
	//return this.buf.String()
	return chars.ByteSliceToString(this.Bytes())
}

func (this *Writer) Bytes() []byte {
	return this.buf.Bytes()
}
