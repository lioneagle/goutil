package draw

import (
	"fmt"
	"log"
	"runtime/debug"
)

var (
	logErrors    bool = true
	panicOnError bool = true
)

func LogErrors() bool {
	return logErrors
}

func SetLogErrors(v bool) {
	logErrors = v
}

func PanicOnError() bool {
	return panicOnError
}

func SetPanicOnError(v bool) {
	panicOnError = v
}

type Error struct {
	inner   error
	message string
	stack   []byte
}

func (this *Error) Inner() error {
	return this.inner
}

func (this *Error) Message() string {
	if this.message != "" {
		return this.message
	}

	if this.inner != nil {
		if walkErr, ok := this.inner.(*Error); ok {
			return walkErr.Message()
		} else {
			return this.inner.Error()
		}
	}

	return ""
}

func (this *Error) Error() string {
	return fmt.Sprintf("%s\n\nStack:\n%s", this.Message(), this.stack)
}

func NewError(message string) error {
	return processError(newErr(message))
}

func newErrorNoPanic(message string) error {
	return processErrorNoPanic(newErr(message))
}

func newErr(message string) error {
	return &Error{message: message, stack: debug.Stack()}
}

func processErrorNoPanic(err error) error {
	if logErrors {
		if walkErr, ok := err.(*Error); ok {
			log.Print(walkErr.Error())
		} else {
			log.Printf("%s\n\nStack:\n%s", err, debug.Stack())
		}
	}

	return err
}

func processError(err error) error {
	processErrorNoPanic(err)

	if panicOnError {
		panic(err)
	}

	return err
}
