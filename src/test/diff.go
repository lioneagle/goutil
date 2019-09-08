package test

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"runtime"

	"github.com/lioneagle/goutil/src/mathex"
)

type diffWriter struct {
	writer io.Writer
	label  string
}

func DiffEx(label string, actual, wanted interface{}) (bool, string) {
	buf := bytes.NewBuffer(nil)
	w := &diffWriter{writer: buf, label: label}
	ok := w.diff(reflect.ValueOf(actual), reflect.ValueOf(wanted))
	return ok, buf.String()
}

func Diff(actual, wanted interface{}) (bool, string) {
	return DiffEx(funcName(2), actual, wanted)
}

func (this *diffWriter) diff(actual, wanted reflect.Value) bool {
	if !actual.IsValid() && !wanted.IsValid() {
		return true
	}

	if !actual.IsValid() && wanted.IsValid() {
		this.print(nil, wanted)
		return false
	}

	if actual.IsValid() && !wanted.IsValid() {
		this.print(actual, nil)
		return false
	}

	typeOfActual := actual.Type()
	typeOfWanted := wanted.Type()

	if typeOfActual != typeOfWanted {
		this.printf("actual type: %v, wanted type: %v\n", typeOfActual, typeOfWanted)
		return false
	}

	switch kind := typeOfActual.Kind(); kind {
	case reflect.Bool:
		if a, b := actual.Bool(), wanted.Bool(); a != b {
			this.print(a, b)
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if a, b := actual.Int(), wanted.Int(); a != b {
			this.print(a, b)
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if a, b := actual.Uint(), wanted.Uint(); a != b {
			this.print(a, b)
			return false
		}
	case reflect.Float32:
		if a, b := actual.Float(), wanted.Float(); mathex.CompareFloat64Ex(a, b, mathex.MinAccuracyFloat32) != 0 {
			this.print(a, b)
			return false
		}

	case reflect.Float64:
		if a, b := actual.Float(), wanted.Float(); mathex.CompareFloat64Ex(a, b, mathex.MinAccuracyFloat64) != 0 {
			this.print(a, b)
			return false
		}
	case reflect.Complex64:
		if a, b := actual.Complex(), wanted.Complex(); !mathex.EqualComplex128Ex(a, b, mathex.MinAccuracyFloat32) {
			this.print(a, b)
			return false
		}

	case reflect.Complex128:
		if a, b := actual.Complex(), wanted.Complex(); !mathex.EqualComplex128Ex(a, b, mathex.MinAccuracyFloat64) {
			this.print(a, b)
			return false
		}
	case reflect.String:
		if a, b := actual.String(), wanted.String(); a != b {
			this.print(a, b)
			return false
		}
	case reflect.Array:
		//fmt.Println("enter Array")
		ret := true
		n := actual.Len()
		for i := 0; i < n; i++ {
			if !this.reLabel(fmt.Sprintf("[%d]", i)).diff(actual.Index(i), wanted.Index(i)) {
				ret = false
			}
		}
		return ret
	case reflect.Chan:
		//fmt.Println("enter Chan")
		if a, b := actual.Pointer(), wanted.Pointer(); a != b {
			this.print(a, b)
			return false
		}
		return true
	case reflect.Interface:
		//fmt.Println("enter Interface")
		var err error
		if actual.Type().Implements(reflect.TypeOf(&err).Elem()) && wanted.Type().Implements(reflect.TypeOf(&err).Elem()) {
			if actual != wanted {
				var s1, s2 reflect.Value
				if actual.IsNil() {
					s1 = reflect.ValueOf("nil")
				} else {
					m1 := actual.MethodByName("Error")
					if m1.CanInterface() {
						s1 = m1.Call(nil)[0]
					} else {
						s1 = actual.Elem()
					}
				}

				if wanted.IsNil() {
					s2 = reflect.ValueOf("nil")
				} else {
					m2 := wanted.MethodByName("Error")
					if m2.CanInterface() {
						s2 = m2.Call(nil)[0]
					} else {
						s2 = wanted.Elem()
					}
				}

				//this.printf("actual = %v, wanted = %v\n", s1, s2)
				this.print(s1, s2)

				return false
			}
			return true
		}

		return this.diff(actual.Elem(), wanted.Elem())
	case reflect.Map:
		//fmt.Println("enter Map")
	case reflect.Ptr:
		//fmt.Println("enter Ptr")
		switch {
		case actual.IsNil() && !wanted.IsNil():
			this.print(nil, wanted)
			return false
		case !actual.IsNil() && wanted.IsNil():
			this.print(actual, nil)
			return false
		case !actual.IsNil() && !wanted.IsNil():
			return this.diff(actual.Elem(), wanted.Elem())
		default:
			return true
		}
	case reflect.Slice:
		//fmt.Println("enter Slice")
		len1 := actual.Len()
		len2 := wanted.Len()
		if len1 != len2 {
			if (!actual.IsNil() && actual.Index(0).Kind() == reflect.Uint8) ||
				(!wanted.IsNil() && wanted.Index(0).Kind() == reflect.Uint8) {
				this.printf("actual: %v(%q), wanted: %v(%q)\n", actual, actual, wanted, wanted)
			} else {
				this.printf("actual: len = %v, wanted: len = %v\n", len1, len2)
			}
			return false
		}

		ret := true
		for i := 0; i < len1; i++ {
			//fmt.Printf("actual.Index(%d) = %#v, wanted.Index(%d) = %#v\n", i, actual.Index(i), i, wanted.Index(i))
			if !this.reLabel(fmt.Sprintf("[%d]", i)).diff(actual.Index(i), wanted.Index(i)) {
				ret = false
			}
		}
		return ret
	case reflect.Struct:
		//fmt.Println("enter Struct")
		ret := true
		for i := 0; i < actual.NumField(); i++ {
			if !this.reLabel(typeOfActual.Field(i).Name).diff(actual.Field(i), wanted.Field(i)) {
				ret = false
				if typeOfActual.Field(i).Tag == "break" {
					this.writer.Write([]byte("break\n"))
					return false
				}
			}
		}
		return ret

	case reflect.Func:
		if actual.Addr().Elem().Pointer() == wanted.Addr().Elem().Pointer() {
			return true
		}
		this.print(actual.Addr().Elem(), wanted.Addr().Elem())
		return false

	default:
		//fmt.Println("enter Default, typeOfActual.Kind() =", typeOfActual.Kind())
		if actual != wanted {
			this.print(actual, wanted)
			return false
		}
	}

	return true
}

func (this *diffWriter) reLabel(name string) *diffWriter {
	w := *this
	if this.label != "" && name[0] != '[' {
		w.label += "."
	}
	w.label += name
	return &w
}

type MyError struct {
	x1 int
	x2 int
}

func (this *MyError) Error() string {
	return "123"
}

func (this *diffWriter) print(actual, expected interface{}) {
	this.printf("actual = %v, wanted = %v\n", actual, expected)
}

func (this *diffWriter) printf(format string, args ...interface{}) {
	label := this.label
	if label != "" {
		label += ": "
	}
	fmt.Fprintf(this.writer, label+format, args...)
}

func funcName(depth int) string {
	pc, _, _, ok := runtime.Caller(depth)
	if !ok {
		return ""
	}
	return runtime.FuncForPC(pc).Name()
}

func FuncName() string {
	return funcName(1)
}
