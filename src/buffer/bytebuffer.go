package buffer

import (
	"fmt"
	"io"
	"strconv"

	"github.com/lioneagle/goutil/src/chars"
)

type Writer interface {
	Write(val []byte) (int, error)
	WriteString(val string) (int, error)
	WriteByte(val byte) error
}

type ByteBuffer struct {
	data []byte
}

func NewByteBuffer(buf []byte) *ByteBuffer {
	ret := &ByteBuffer{data: buf}
	ret.Reset()
	return ret
}

func (this *ByteBuffer) Alloc(size int) {
	this.data = make([]byte, size)
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

func (this *ByteBuffer) WriteByteN(val byte, num int) error {
	for i := 0; i < num; i++ {
		this.data = append(this.data, val)
	}
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

func (this *ByteBuffer) WritePercentEscape(val []byte, charset *[256]uint32, mask uint32) (int, error) {
	charset1 := charset
	for _, v := range val {
		if (charset1[v] & mask) != 0 {
			this.WriteByte(v)
		} else {
			this.WriteByte('%')
			this.WriteByte(chars.ToUpperHex(v >> 4))
			this.WriteByte(chars.ToUpperHex(v))
		}
	}
	return len(val), nil
}

func (this *ByteBuffer) Print(args ...interface{}) {
	fmt.Fprint(this, args...)
}

func (this *ByteBuffer) Printf(format string, args ...interface{}) {
	fmt.Fprintf(this, format, args...)
}

func (this *ByteBuffer) Println(args ...interface{}) {
	fmt.Fprintln(this, args...)
}

func (this *ByteBuffer) Printfln(format string, args ...interface{}) {
	fmt.Fprintf(this, format, args...)
	fmt.Fprintln(this)
}

func (this *ByteBuffer) Reset() {
	this.data = this.data[:0]
}

func (this *ByteBuffer) PrintAsHex(w Writer, begin, end int) {
	PrintAsHex(w, this.data, begin, end)
}

func PrintAsHex(w io.Writer, data []byte, begin, end int) {
	size := len(data)
	if size == 0 {
		return
	}

	if begin < 0 {
		begin = 0
	}

	if end > size {
		end = size
	}

	if begin >= end {
		return
	}

	size = end - begin

	lines := size / 16
	last := size % 16

	for i := 0; i < lines; i++ {
		printHexOneline(w, data, begin, begin+16)
		begin += 16
	}

	if last > 0 {
		printHexOneline(w, data, begin, begin+last)
	}
}

func printHexOneline(w io.Writer, data []byte, begin, end int) {
	num := end - begin
	fmt.Fprintf(w, "%08xh: ", begin)

	if num < 8 {
		for i := begin; i < end; i++ {
			fmt.Fprintf(w, "%02X ", data[i])
		}
		for i := 0; i < 8-num; i++ {
			fmt.Fprintf(w, "   ")
		}
	} else {
		for i := begin; i < begin+8; i++ {
			fmt.Fprintf(w, "%02X ", data[i])
		}
	}

	fmt.Fprintf(w, " ")

	if num < 8 {
		for i := 0; i < 8; i++ {
			fmt.Fprintf(w, "   ")
		}
	} else {
		for i := begin + 8; i < end; i++ {
			fmt.Fprintf(w, "%02X ", data[i])
		}
		for i := 0; i < 16-num; i++ {
			fmt.Fprintf(w, "   ")
		}
	}

	fmt.Fprintf(w, "; ")

	for i := begin; i < end; i++ {
		if strconv.IsPrint(rune(data[i])) {
			if data[i] < 128 {
				fmt.Fprintf(w, "%c", data[i])
			} else {
				fmt.Fprintf(w, "?")
			}
		} else {
			fmt.Fprintf(w, ".")
		}
		if i == (begin + 7) {
			fmt.Fprintf(w, " ")
		}
	}
	fmt.Fprintf(w, "\n")
}
