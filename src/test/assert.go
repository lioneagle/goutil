package test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
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
	buf := bytes.NewBuffer(nil)

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
				ok := w.diff(ret1[0].Elem(), val1.Field(j))
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
		t.Errorf("%s:%d\nactual = false, wanted = true\n%s", filepath.Base(file), line, fmt.Sprintf(format, args...))
	}
}

func ASSERT_TRUE(t *testing.T, condition bool, format string, args ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d\nactual = false, wanted = true\n%s", filepath.Base(file), line, fmt.Sprintf(format, args...))
		t.FailNow()
	}
}

func EXPECT_FALSE(t *testing.T, condition bool, format string, args ...interface{}) {
	if condition {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d\nactual = true, wanted = false\n%s", filepath.Base(file), line, fmt.Sprintf(format, args...))
	}
}

func ASSERT_FALSE(t *testing.T, condition bool, format string, args ...interface{}) {
	if condition {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d\nactual = true, wanted = false\n%s", filepath.Base(file), line, fmt.Sprintf(format, args...))
		t.FailNow()
	}
}

func EXPECT_EQ(t *testing.T, actual, wanted interface{}, format string, args ...interface{}) {
	equal, result := DiffEx("", actual, wanted)

	if !equal {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d\n%s%s", filepath.Base(file), line, result, fmt.Sprintf(format, args...))
	}
}

func ASSERT_EQ(t *testing.T, actual, wanted interface{}, format string, args ...interface{}) {
	equal, result := DiffEx("", actual, wanted)

	if !equal {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d\n%s%s", filepath.Base(file), line, result, fmt.Sprintf(format, args...))
		t.FailNow()
	}
}

func EXPECT_NE(t *testing.T, actual, wanted interface{}, format string, args ...interface{}) {
	equal, _ := DiffEx("", actual, wanted)

	if equal {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d\nshould not be equal\n%s", filepath.Base(file), line, fmt.Sprintf(format, args...))
	}
}

func ASSERT_NE(t *testing.T, actual, wanted interface{}, format string, args ...interface{}) {
	equal, _ := DiffEx("", actual, wanted)

	if equal {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d\nshould not be equal\n%s", filepath.Base(file), line, fmt.Sprintf(format, args...))
		t.FailNow()
	}
}
