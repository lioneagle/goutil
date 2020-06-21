package abnf

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestParseUint64(t *testing.T) {
	testdata := []struct {
		src    string
		pos    uint
		digit  uint64
		newPos Pos
		ok     bool
	}{
		{"1234567890.abc", 0, 1234567890, 10, true},
		{"10.40.1.1", 0, 10, 2, true},
		{"aa10.40.1.1", 2, 10, 4, true},
		{"18446744073709551615", 0, 18446744073709551615, 20, true},
		{"18446744073709551615.", 0, 18446744073709551615, 20, true},
		{"1844674407370955161", 0, 1844674407370955161, 19, true},
		{"1844674407370955161.", 0, 1844674407370955161, 19, true},

		{"", 0, 0, 0, false},
		{"10", 3, 0, 3, false},
		{"18446744073709551616", 0, 0, 19, false},
		{"18446744073709551626", 0, 0, 19, false},
		{"aa18446744073709551616", 2, 0, 21, false},
		{"184467440737095516155", 0, 0, 20, false},
		{"aa184467440737095516155", 2, 0, 22, false},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			digit1, newPos, ok := ParseUint64([]byte(v.src), v.pos)

			if v.ok {
				test.EXPECT_TRUE(t, ok, "")
				test.EXPECT_EQ(t, uint64(digit1), v.digit, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
			} else {
				test.EXPECT_FALSE(t, ok, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
			}
		})
	}
}

func TestParseUint32(t *testing.T) {
	testdata := []struct {
		src    string
		digit  uint64
		newPos Pos
		ok     bool
	}{
		{"123456789.abc", 123456789, 9, true},
		{"4294967295", 4294967295, 10, true},
		{"4294967295.", 4294967295, 10, true},

		{"", 0, 0, false},
		{"4294967296", 0, 10, false},
		{"42949672956", 0, 11, false},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			digit1, newPos, ok := ParseUint32([]byte(v.src), 0)

			if v.ok {
				test.EXPECT_TRUE(t, ok, "")
				test.EXPECT_EQ(t, uint64(digit1), v.digit, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
			} else {
				test.EXPECT_FALSE(t, ok, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
			}
		})
	}
}

func TestParseUint16(t *testing.T) {
	testdata := []struct {
		src    string
		digit  uint64
		newPos Pos
		ok     bool
	}{
		{"1234.abc", 1234, 4, true},
		{"65535", 65535, 5, true},
		{"65535.", 65535, 5, true},

		{"", 0, 0, false},
		{"65536", 0, 5, false},
		{"655351", 0, 6, false},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			digit1, newPos, ok := ParseUint16([]byte(v.src), 0)

			if v.ok {
				test.EXPECT_TRUE(t, ok, "")
				test.EXPECT_EQ(t, uint64(digit1), v.digit, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
			} else {
				test.EXPECT_FALSE(t, ok, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
			}
		})
	}
}

func TestParseUint8(t *testing.T) {
	testdata := []struct {
		src    string
		digit  uint64
		newPos Pos
		ok     bool
	}{
		{"12.abc", 12, 2, true},
		{"255", 255, 3, true},
		{"255.", 255, 3, true},

		{"", 0, 0, false},
		{"256", 0, 3, false},
		{"2551", 0, 4, false},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			digit1, newPos, ok := ParseUint8([]byte(v.src), 0)

			if v.ok {
				test.EXPECT_TRUE(t, ok, "")
				test.EXPECT_EQ(t, uint64(digit1), v.digit, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
			} else {
				test.EXPECT_FALSE(t, ok, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
			}
		})
	}
}

func TestEncodeUint(t *testing.T) {
	testdata := []struct {
		digit  uint64
		wanted string
	}{
		{0, "0"},
		{4, "4"},
		{13, "13"},
		{123, "123"},
		{65536, "65536"},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			buf := NewByteBuffer(nil)
			EncodeUint(buf, v.digit)
			test.EXPECT_EQ(t, buf.String(), v.wanted, "")
		})
	}
}

func TestEncodeUintWithWidth(t *testing.T) {
	testdata := []struct {
		digit  uint64
		width  int
		wanted string
	}{
		{0, 0, "0"},
		{2, 0, "2"},
		{2, 3, "  2"},
		{25, 1, "25"},
		{25, 3, " 25"},
		{123, 7, "    123"},
		{65536, 5, "65536"},
		{65536, 10, "     65536"},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			buf := NewByteBuffer(nil)
			EncodeUintWithWidth(buf, v.digit, v.width)
			test.EXPECT_EQ(t, buf.String(), v.wanted, "")
		})
	}
}

func BenchmarkParseUint64_1(b *testing.B) {
	b.StopTimer()
	src := []byte("12345")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ParseUint64(src, 0)
	}
}

func BenchmarkParseUint64_Strconv_ParseUint(b *testing.B) {
	b.StopTimer()
	src := "12345"
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		strconv.ParseUint(src, 10, 64)
	}
}

func BenchmarkParseUint64_Strconv_Atoi(b *testing.B) {
	b.StopTimer()
	src := "12345"
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		strconv.Atoi(src)
	}
}

func BenchmarkParseUint32_1(b *testing.B) {
	b.StopTimer()
	src := []byte("12345")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ParseUint32(src, 0)
	}
}

func BenchmarkParseUint16_1(b *testing.B) {
	b.StopTimer()
	src := []byte("12345")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ParseUint16(src, 0)
	}
}

func BenchmarkParseUint8_1(b *testing.B) {
	b.StopTimer()
	src := []byte("123")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ParseUint8(src, 0)
	}
}

func BenchmarkEncodeUint_1(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		EncodeUint(buf, 1234567)
	}
}

func BenchmarkEncodeUint_2(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		EncodeUint(buf, 123)
	}
}

func BenchmarkEncodeUint_3(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		EncodeUint(buf, 20)
	}
}

func BenchmarkEncodeUint_1_Strconv_1(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.WriteString(strconv.FormatUint(1234567, 10))
	}
}

func BenchmarkEncodeUint_2_Strconv(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.WriteString(strconv.FormatUint(123, 10))
	}
}

func BenchmarkEncodeUint_3_Strconv(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.WriteString(strconv.FormatUint(20, 10))
	}
}

func BenchmarkEncodeUintWithWidth_1(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		EncodeUintWithWidth(buf, 1234567, 10)
	}
}

func BenchmarkEncodeUintWithWidth_2(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		EncodeUintWithWidth(buf, 123, 10)
	}
}

func BenchmarkEncodeUintWithWidth_3(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		EncodeUintWithWidth(buf, 20, 10)
	}
}

func BenchmarkEncodeUintWithWidth_1_fmt_Sprintf(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString(fmt.Sprintf("%10d", 1234567))
	}
}

func BenchmarkEncodeUintWithWidth_2_fmt_Sprintf(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString(fmt.Sprintf("%10d", 123))
	}
}

func BenchmarkEncodeUintWithWidth_3_fmt_Sprintf(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024))
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString(fmt.Sprintf("%10d", 20))
	}
}
