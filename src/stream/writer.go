package stream

import (
	"fmt"

	"github.com/lioneagle/goutil/src/buffer"
	"github.com/lioneagle/goutil/src/chars"
)

type Writer struct {
	buf buffer.ByteBuffer
}

func (this *Writer) Write(format string, args ...interface{}) {
	fmt.Fprintf(&this.buf, format, args...)
}

func (this *Writer) WriteString(val string) {
	this.buf.WriteString(val)
}

func (this *Writer) WriteByte(val byte) {
	this.buf.WriteByte(val)
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
