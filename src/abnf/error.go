package abnf

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

const (
	ABNF_ERROR_MAX_SRC_OUTPUT_LEN = 40
)

type Error struct {
	msg string
	src []byte
	pos Pos
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	buf := NewByteBuffer(nil)
	e.Write(buf)
	return buf.String()
}

func (e *Error) Write(w io.Writer) {
	if e.pos < Pos(len(e.src)) {
		num := Pos(ABNF_ERROR_MAX_SRC_OUTPUT_LEN)
		if (e.pos + num) > Pos(len(e.src)) {
			num = Pos(len(e.src)) - e.pos
		}
		fmt.Fprintf(w, "%s: src[%d]: %s", e.msg, e.pos, string(e.src[e.pos:e.pos+num]))
		return
	}
}

func NewError(src []byte, pos Pos, message string) error {
	err := &Error{src: src, pos: pos, msg: message}
	//return errors.WithStack(err)
	return errors.WithStack(err)
	//return err
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(src []byte, pos Pos, format string, args ...interface{}) error {
	err := &Error{src: src, pos: pos}
	err.msg = fmt.Sprintf(format, args...)
	return errors.WithStack(err)
}
