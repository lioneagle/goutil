package detailerr

import (
	//"fmt"
	"path/filepath"
	"runtime"

	"github.com/lioneagle/goutil/src/buffer"
)

type Error struct {
	description string
	fileName    string
	pc          uintptr
	line        int
}

func (this *Error) Error() string {
	return this.String()
}

func (this *Error) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.Write(buf)
	return buf.String()
}

func (this *Error) Write(buf *buffer.ByteBuffer) {
	description := this.description

	if len(description) == 0 {
		description = "unknown error"
	}
	buf.Printf("%s:%d: %s", this.fileName, this.line, this.description)
}

type Errors struct {
	src    []byte
	errors []*Error
}

func (this *Errors) Len() int {
	return len(this.errors)
}

func (this *Errors) Add(description string, call_depth int) {
	fileName, pc, line := GetCallerInfoN(call_depth)
	this.errors = append(this.errors, &Error{description: description, fileName: fileName, pc: pc, line: line})
}

func (this *Errors) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.Write(buf)
	return buf.String()
}

func (this *Errors) Write(buf *buffer.ByteBuffer) {
	len1 := len(this.errors)
	for i := len1 - 1; i >= 0; i-- {
		buf.Printf("[%d]: ", len1-i)
		this.errors[i].Write(buf)
		buf.Printfln("")
	}
}

func GetCallerInfoN(n int) (fileName string, pc uintptr, line int) {
	pc, fileName, line, ok := runtime.Caller(n)
	if ok {
		return filepath.Base(fileName), pc, line
	}
	return "unknown-file", 0, -1
}
