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

func (this *ByteBuffer) PrintAsHex(w Writer, begin, end int) {
	size := len(this.data)
	if size == 0 {
		return
	}

	if begin < 0 {
		begin = 0
	}

	if end >= size {
		end = size
	}

	if begin >= end {
		return
	}

	size = end - begin

	lines := size / 16
	last := size % 16

	for i := 0; i < lines; i++ {
		this.printHexOneline(w, begin, begin+16)
		begin += 16
	}

	if last > 0 {
		this.printHexOneline(w, begin, begin+last)
	}
}

func (this *ByteBuffer) printHexOneline(w Writer, begin, end int) {
	num := end - begin
	fmt.Fprintf(w, "%08xh: ", begin)

	if num < 8 {
		for i := begin; i < end; i++ {
			fmt.Fprintf(w, "%02X ", this.data[i])
		}
		for i := 0; i < 8-num; i++ {
			w.WriteString("   ")
		}
	} else {
		for i := begin; i < begin+8; i++ {
			fmt.Fprintf(w, "%02X ", this.data[i])
		}
	}

	w.WriteByte(' ')

	if num < 8 {
		for i := 0; i < 8; i++ {
			w.WriteString("   ")
		}
	} else {
		for i := begin + 8; i < end; i++ {
			fmt.Fprintf(w, "%02X ", this.data[i])
		}

		for i := 0; i < 16-num; i++ {
			w.WriteString("   ")
		}
	}

	w.WriteByte(';')
	w.WriteByte(' ')
	for i := begin; i < end; i++ {
		if strconv.IsPrint(rune(this.data[i])) {
			if this.data[i] < 128 {
				w.WriteByte(this.data[i])
			} else {
				w.WriteByte('?')
			}
		} else {
			w.WriteByte('.')
		}
	}
	w.WriteByte('\n')
}

func PrintAsHex(w io.Writer, data []byte, begin, end int) {
	size := len(data)
	if size == 0 {
		return
	}

	if begin < 0 {
		begin = 0
	}

	if end >= size {
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
	}
	fmt.Fprintf(w, "\n")
}
