package buffer

import (
	"github.com/lioneagle/goutil/src/chars"
)

type ByteBuffer struct {
	data []byte
}

func (this *ByteBuffer) Len() int {
	return len(this.data)
}

func (this *ByteBuffer) Bytes() []byte {
	return this.data
}

func (this *ByteBuffer) String() string {
	return chars.ByteSliceToString(this.data)
}

func (this *ByteBuffer) WriteByte(val byte) error {
	this.data = append(this.data, val)
	return nil
}

func (this *ByteBuffer) Write(val []byte) (int, error) {
	this.data = append(this.data, val...)
	return len(val), nil
}

func (this *ByteBuffer) WriteString(val string) (int, error) {
	this.data = append(this.data, val...)
	return len(val), nil
}

func (this *ByteBuffer) Reset() {
	this.data = this.data[:0]
}

func NewByteBuffer(buf []byte) *ByteBuffer {
	ret := &ByteBuffer{data: buf}
	ret.Reset()
	return ret
}
