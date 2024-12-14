package diff

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"runtime"
)

const (
	MinAccuracyFloat32 = 1.1920928955078125e-7                 // 1 / 2**23
	MinAccuracyFloat64 = 2.2204460492503130808472633361816e-16 // 1 / 2**52
)

func CompareFloat64Ex(x, y, epsilon float64) int {
	d := x - y

	if d < -epsilon {
		return -1
	} else if d > epsilon {
		return 1
	}

	return 0
}

func EqualComplex128Ex(x, y complex128, epsilon float64) bool {
	return (CompareFloat64Ex(real(x), real(y), epsilon) == 0) &&
		(CompareFloat64Ex(imag(x), imag(y), epsilon) == 0)
}

type DiffWriter struct {
	writer io.Writer
	label  string
}

func DiffEx(label string, actual, wanted interface{}) (bool, string) {
	buf := bytes.NewBuffer(nil)
	w := &DiffWriter{writer: buf, label: label}
	ok := w.Diff(reflect.ValueOf(actual), reflect.ValueOf(wanted))
	return ok, buf.String()
}

func Diff(actual, wanted interface{}) (bool, string) {
	return DiffEx(funcName(2), actual, wanted)
}

func NewDiffWriter(writer io.Writer, label string) *DiffWriter {
	return &DiffWriter{
		writer: writer,
		label:  label,
	}
}

func (this *DiffWriter) Diff(actual, wanted reflect.Value) bool {
	if !actual.IsValid() && !wanted.IsValid() {
		return true
	}

	if !actual.IsValid() && wanted.IsValid() {
		this.Print(nil, wanted)
		return false
	}

	if actual.IsValid() && !wanted.IsValid() {
		this.Print(actual, nil)
		return false
	}

	typeOfActual := actual.Type()
	typeOfWanted := wanted.Type()

	if typeOfActual != typeOfWanted {
		this.Printf("actual type: %v, wanted type: %v\n", typeOfActual, typeOfWanted)
		return false
	}

	switch kind := typeOfActual.Kind(); kind {
	case reflect.Bool:
		if a, b := actual.Bool(), wanted.Bool(); a != b {
			this.Print(a, b)
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if a, b := actual.Int(), wanted.Int(); a != b {
			this.Print(a, b)
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if a, b := actual.Uint(), wanted.Uint(); a != b {
			this.Print(a, b)
			return false
		}
	case reflect.Float32:
		if a, b := actual.Float(), wanted.Float(); CompareFloat64Ex(a, b, MinAccuracyFloat32) != 0 {
			this.Print(a, b)
			return false
		}

	case reflect.Float64:
		if a, b := actual.Float(), wanted.Float(); CompareFloat64Ex(a, b, MinAccuracyFloat64) != 0 {
			this.Print(a, b)
			return false
		}
	case reflect.Complex64:
		if a, b := actual.Complex(), wanted.Complex(); !EqualComplex128Ex(a, b, MinAccuracyFloat32) {
			this.Print(a, b)
			return false
		}

	case reflect.Complex128:
		if a, b := actual.Complex(), wanted.Complex(); !EqualComplex128Ex(a, b, MinAccuracyFloat64) {
			this.Print(a, b)
			return false
		}
	case reflect.String:
		if a, b := actual.String(), wanted.String(); a != b {
			this.Print(a, b)
			return false
		}
	case reflect.Array:
		//fmt.Println("enter Array")
		ret := true
		n := actual.Len()
		for i := 0; i < n; i++ {
			if !this.reLabel(fmt.Sprintf("[%d]", i)).Diff(actual.Index(i), wanted.Index(i)) {
				ret = false
			}
		}
		return ret
	case reflect.Chan:
		//fmt.Println("enter Chan")
		if a, b := actual.Pointer(), wanted.Pointer(); a != b {
			this.Print(a, b)
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

				//this.Printf("actual = %v, wanted = %v\n", s1, s2)
				this.Print(s1, s2)

				return false
			}
			return true
		}

		return this.Diff(actual.Elem(), wanted.Elem())
	case reflect.Map:
		//fmt.Println("enter Map")
	case reflect.Ptr:
		//fmt.Println("enter Ptr")
		switch {
		case actual.IsNil() && !wanted.IsNil():
			this.Print(nil, wanted)
			return false
		case !actual.IsNil() && wanted.IsNil():
			this.Print(actual, nil)
			return false
		case !actual.IsNil() && !wanted.IsNil():
			return this.Diff(actual.Elem(), wanted.Elem())
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
				this.Printf("actual: %v(%q), wanted: %v(%q)\n", actual, actual, wanted, wanted)
			} else {
				this.Printf("actual: len = %v, wanted: len = %v\n", len1, len2)
			}
			return false
		}

		ret := true
		for i := 0; i < len1; i++ {
			//fmt.Printf("actual.Index(%d) = %#v, wanted.Index(%d) = %#v\n", i, actual.Index(i), i, wanted.Index(i))
			if !this.reLabel(fmt.Sprintf("[%d]", i)).Diff(actual.Index(i), wanted.Index(i)) {
				ret = false
			}
		}
		return ret
	case reflect.Struct:
		//fmt.Println("enter Struct")
		ret := true
		for i := 0; i < actual.NumField(); i++ {
			if !this.reLabel(typeOfActual.Field(i).Name).Diff(actual.Field(i), wanted.Field(i)) {
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
		this.Print(actual.Addr().Elem(), wanted.Addr().Elem())
		return false

	default:
		//fmt.Println("enter Default, typeOfActual.Kind() =", typeOfActual.Kind())
		if actual != wanted {
			this.Print(actual, wanted)
			return false
		}
	}

	return true
}

func (this *DiffWriter) reLabel(name string) *DiffWriter {
	w := *this
	if this.label != "" && name[0] != '[' {
		w.label += "."
	}
	w.label += name
	return &w
}

/*type MyError struct {
	x1 int
	x2 int
}

func (this *MyError) Error() string {
	return "123"
}*/

func (this *DiffWriter) Print(actual, expected interface{}) {
	this.Printf("actual = %v, wanted = %v\n", actual, expected)
}

func (this *DiffWriter) Printf(format string, args ...interface{}) {
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
