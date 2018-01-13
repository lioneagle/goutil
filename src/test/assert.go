package test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/lioneagle/goutil/src/buffer"
)

func TestGroup(t *testing.T, testdata interface{}, fn func(interface{}) interface{}) (bool, string) {
	if testdata == nil {
		//t.Errorf("testdata is nil")
		return false, "testdata is nil"
	}

	//_, file, line, _ := runtime.Caller(1)

	typeOfTestData := reflect.TypeOf(testdata)
	if typeOfTestData.Kind() != reflect.Slice {
		//t.Errorf("testdata is not slice")
		return false, "testdata is not slice"
	}
	valueOfTestData := reflect.ValueOf(testdata)
	n := valueOfTestData.Len()

	//label := fmt.Sprintf("%s:%d  %s", filepath.Base(file), line, filepath.Base(funcName(2)))

	fn1 := reflect.ValueOf(fn)

	ret := true
	buf := buffer.NewByteBuffer(nil)

	for i := 0; i < n; i++ {
		val1 := valueOfTestData.Index(i)
		type1 := val1.Type()
		//label1 := fmt.Sprintf("%s[%d]", label, i)
		label1 := fmt.Sprintf("[%d].", i)

		if type1.Kind() != reflect.Struct {
			//t.Errorf("testdata is not slice")
			return false, "testdata is not slice"
		}

		args := []reflect.Value{val1}
		ret1 := fn1.Call(args)
		typeOfRet1 := ret1[0].Elem().Type()

		for j := 0; j < val1.NumField(); j++ {
			typeOfFiled := type1.Field(j).Type

			if typeOfFiled == typeOfRet1 {

				w := &diffWriter{writer: buf, label: label1}
				ok := w.diff(val1.Field(j), ret1[0].Elem())
				if !ok {
					ret = false
					//return false, buf.String()
					//t.Errorf("\n" + buf.String())
				}
			}
		}
	}

	return ret, buf.String()
}

func EXPECT_TRUE(t *testing.T, condition bool, format string, args ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d", filepath.Base(file), line)
		//t.FailNow()
	}
}
