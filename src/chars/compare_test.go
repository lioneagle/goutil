package chars

import (
	"bytes"
	"testing"
	"unsafe"
)

func TestToLowerHex(t *testing.T) {
	src := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	wanted := []byte("0123456789abcdef")

	for i, v := range src {
		u := ToLowerHex(v)
		if u != wanted[i] {
			t.Errorf("TestToLowerHex[%d] failed, ret = %c, wanted = %c\n", i, u, wanted[i])
		}
	}
}

func TestToUpper(t *testing.T) {
	src := []byte(";[]abcdefghigklmnopqrstuvwxyz012-+")
	wanted := []byte(";[]ABCDEFGHIGKLMNOPQRSTUVWXYZ012-+")

	for i, v := range src {
		u := ToUpper(v)
		if u != wanted[i] {
			t.Errorf("TestToUpper[%d] failed, ret = %c, wanted = %c\n", i, u, wanted[i])
		}
	}
}

func TestCompareNoCase(t *testing.T) {

	testdata := []struct {
		s1  string
		s2  string
		ret int
	}{
		{"aBcf", "abck", -1},
		{"aBcf", "abcf", 0},
		{"aBcf", "abcc", 1},
		{"aBcdf", "abcf", 1},
		{"aBcdf", "abcaaf", -1},
		{"089+=abcdefghigklmnopqrstuvwxyz123", "089+=ABCDEFGHIGKLMNOPQRSTUVWXYZ123", 0},
	}

	for i, v := range testdata {
		u := CompareNoCase([]byte(v.s1), []byte(v.s2))
		if u < 0 {
			u = -1
		} else if u > 0 {
			u = 1
		}
		if u != v.ret {
			t.Errorf("TestCompareNoCase[%d] failed, ret = %d, wanted = %d\n", i, u, v.ret)
		}
	}
}

func TestEqualNoCase(t *testing.T) {

	wanted := []struct {
		s1  string
		s2  string
		ret bool
	}{
		{"aBcf", "abck", false},
		{"aBcf", "abcf", true},
		{"aBcf", "abcc", false},
		{"aBcdf", "abcf", false},
		{"aBcdf", "abcaaf", false},
		{"089+=abcdefghigklmnopqrstuvwxyz123", "089+=ABCDEFGHIGKLMNOPQRSTUVWXYZ123", true},
	}

	for i, v := range wanted {
		u := EqualNoCase([]byte(v.s1), []byte(v.s2))
		if u != v.ret {
			t.Errorf("TestEqualNoCase[%d] failed, ret = %v, wanted = %v\n", i, u, v.ret)
		}
	}
}

func BenchmarkEqualNoCase_Equal1(b *testing.B) {
	b.StopTimer()
	var s1 = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var s2 = []byte("abcdefghijklmnopqrstuvwxyz")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase(s1, s2)
	}
}

func BenchmarkEqualNoCase_Equal2(b *testing.B) {
	b.StopTimer()
	s1 := []byte("abcdefghijklmnopqrstuvwxyz")
	s2 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase(s1, s2)
	}
}

func BenchmarkEqualNoCase_Equal3(b *testing.B) {
	b.StopTimer()
	s1 := []byte("abcdefghijklmnopqrstuvwxyz")
	s2 := []byte("abcdefghijklmnopqrstuvwxyz")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase(s1, s2)
	}
}

func BenchmarkEqualNoCase_Equal4(b *testing.B) {
	b.StopTimer()
	s1 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s2 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase(s1, s2)
	}
}

func BenchmarkEqualNoCase_Equal5(b *testing.B) {
	b.StopTimer()
	s1 := []byte("abcdefghijklMnopqrstuvwxyz")
	s2 := []byte("abcdefghijklmnopqrstuvwxyz")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase(s1, s2)
	}
}

func BenchmarkEqualNoCase1_Equal1(b *testing.B) {
	b.StopTimer()
	var s1 = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var s2 = []byte("abcdefghijklmnopqrstuvwxyz")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase1(s1, s2)
	}
}

func BenchmarkEqualNoCase1_Equal2(b *testing.B) {
	b.StopTimer()
	s1 := []byte("abcdefghijklmnopqrstuvwxyz")
	s2 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase1(s1, s2)
	}
}

func BenchmarkEqualNoCase1_Equal3(b *testing.B) {
	b.StopTimer()
	s1 := []byte("abcdefghijklmnopqrstuvwxyz")
	s2 := []byte("abcdefghijklmnopqrstuvwxyz")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase1(s1, s2)
	}
}

func BenchmarkEqualNoCase1_Equal4(b *testing.B) {
	b.StopTimer()
	s1 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s2 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		EqualNoCase1(s1, s2)
	}
}

func BenchmarkBytesEqual(b *testing.B) {
	b.StopTimer()
	s1 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s2 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		bytes.Equal(s1, s2)
	}
}

func BytesEqual2(s1 []byte, s2 []byte) bool {
	len1 := len(s1)
	if len1 != len(s2) {
		return false
	}
	for i := 0; i < len1; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func BytesEqual3(s1 []byte, s2 []byte) bool {
	len1 := len(s1)
	if len1 != len(s2) {
		return false
	}

	p1 := uintptr(unsafe.Pointer(&s1[0]))
	p2 := uintptr(unsafe.Pointer(&s2[0]))
	end := p1 + uintptr(len1)
	end1 := p1 + uintptr((len1>>3)<<3)

	for p1 < end1 {
		if *((*int64)(unsafe.Pointer(p1))) != *((*int64)(unsafe.Pointer(p2))) {
			return false
		}
		p1 += 8
		p2 += 8
	}

	for p1 < end {
		if *((*byte)(unsafe.Pointer(p1))) != *((*byte)(unsafe.Pointer(p2))) {
			return false
		}
		p1++
		p2++
	}
	return true
}

func BenchmarkBytesEqual2(b *testing.B) {
	b.StopTimer()
	s1 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s2 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		BytesEqual2(s1, s2)
	}
}

func BenchmarkBytesEqual3(b *testing.B) {
	b.StopTimer()
	s1 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s2 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		BytesEqual3(s1, s2)
	}
}

func BenchmarkToLower(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ToLower('1')
	}
}

func BenchmarkToLower2(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		toLower('1')
	}
}
