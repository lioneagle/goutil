package abnf

import (
	//"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
	"github.com/pkg/errors"
)

var src = []byte("12345678")

func f1() error {
	return New(src, 0, "err of f1")
}

func f2() error {
	err := f1()
	return errors.WithMessage(err, "err of f2")
}

func f3() error {
	err := f2()
	return errors.WithMessage(err, "err of f3")
}

func f4() error {
	return Errorf(src, 0, "%s", "err of f4")
}

func f5() error {
	err := f4()
	return errors.WithMessage(err, "err of f5")
}

func TestError(t *testing.T) {
	wanted := "err of f3: err of f2: err of f1: src[0]: 12345678"
	err := f3()

	test.EXPECT_EQ(t, err.Error(), wanted, "")
}

func TestErrorf(t *testing.T) {
	wanted := "err of f5: err of f4: src[0]: 12345678"
	err := f5()

	test.EXPECT_EQ(t, err.Error(), wanted, "")
}
