package abnf

import (
	"fmt"
	"strconv"
	"testing"

	//"github.com/lioneagle/goutil/src/buffer"
	"github.com/lioneagle/goutil/src/chars"
	"github.com/lioneagle/goutil/src/mem"
	"github.com/lioneagle/goutil/src/test"
)

func TestParseInCharset(t *testing.T) {
	testdata := []struct {
		name    string
		charset *[256]uint32
		mask    uint32
		src     string
		newPos  Pos
	}{

		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234abc", 5},
		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "56789=bc", 5},
		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "ad6789abc", 0},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			newPos := ParseInCharset([]byte(v.src), 0, v.charset, v.mask)
			test.EXPECT_EQ(t, newPos, v.newPos, "")
		})
	}
}

func TestParseInCharsetAndAlloc(t *testing.T) {
	testdata := []struct {
		name             string
		charset          *[256]uint32
		mask             uint32
		src              string
		outputIsNotEmpty bool
		newPos           Pos
		outputLen        int
	}{

		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234abc", true, 5, 5},
		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "56789=bc", true, 5, 5},
		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "ad6789abc", false, 0, 5},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			allocator := mem.NewArenaAllocator(1024, 1)
			addr, newPos := ParseInCharsetAndAlloc(allocator, []byte(v.src), 0, v.charset, v.mask)

			if v.outputIsNotEmpty {
				test.EXPECT_NE(t, addr, mem.MEM_PTR_NIL, "")
				test.EXPECT_EQ(t, allocator.Strlen(addr), v.outputLen, "")
			} else {
				test.EXPECT_EQ(t, addr, mem.MEM_PTR_NIL, "")
			}

			test.EXPECT_EQ(t, newPos, v.newPos, "")
		})
	}
}

func TestParseInCharsetPercentEscapable(t *testing.T) {
	testdata := []struct {
		memCapacity      uint32
		name             string
		charset          *[256]uint32
		mask             uint32
		src              string
		outputIsNotEmpty bool
		isOk             bool
		newPos           Pos
		outputLen        int
	}{

		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234abc", true, true, 5, 5},
		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "56789=bc", true, true, 5, 5},
		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "%301234abc", true, true, 7, 5},
		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "%30%311234abc", true, true, 10, 6},
		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "%311234%30", true, true, 10, 6},
		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "%30%31123%3a", true, true, 12, 6},
		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "ad6789abc", false, true, 0, 0},

		{1, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234abc", false, false, 0, 0},
		{2, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234abc", false, false, 5, 0},
		{2, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234", false, false, 5, 0},
		{4, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "012%3134abc", false, false, 3, 0},
		{4, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "012%3G34abc", false, false, 3, 0},
		{4, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "012%g334abc", false, false, 3, 0},
		{4, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "012%3", false, false, 3, 0},
		{4, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "012%", false, false, 3, 0},
		{5, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "012%31ab", false, false, 3, 0},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			allocator := mem.NewArenaAllocator(v.memCapacity, 1)
			addr, newPos, err := ParseInCharsetPercentEscapable(allocator, []byte(v.src), 0, v.charset, v.mask)

			if !v.isOk {
				test.EXPECT_NE(t, err, nil, "")
				test.EXPECT_EQ(t, addr, mem.MEM_PTR_NIL, "")
				test.EXPECT_EQ(t, newPos, v.newPos, "")
				return
			}

			test.EXPECT_EQ(t, err, nil, "%s", err)

			if v.outputIsNotEmpty {
				test.EXPECT_NE(t, addr, mem.MEM_PTR_NIL, "")
				test.EXPECT_EQ(t, allocator.Strlen(addr), v.outputLen, "")
			} else {
				test.EXPECT_EQ(t, addr, mem.MEM_PTR_NIL, "")
			}

			test.EXPECT_EQ(t, newPos, v.newPos, "")
		})
	}
}

func TestParseUint64(t *testing.T) {
	testdata := []struct {
		src    string
		digit  uint64
		newPos Pos
		ok     bool
	}{
		{"1234567890.abc", 1234567890, 10, true},
		{"10.40.1.1", 10, 2, true},
		{"18446744073709551615", 18446744073709551615, 20, true},
		{"18446744073709551615.", 18446744073709551615, 20, true},
		{"1844674407370955161", 1844674407370955161, 19, true},
		{"1844674407370955161.", 1844674407370955161, 19, true},

		{"", 0, 0, false},
		{"18446744073709551616", 0, 19, false},
		{"18446744073709551626", 0, 19, false},
		{"184467440737095516155", 0, 20, false},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			digit1, newPos, ok := ParseUint64([]byte(v.src), 0)

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

func BenchmarkParseCharset(b *testing.B) {
	b.StopTimer()
	allocator := mem.NewArenaAllocator(1024, 1)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		newPos := ParseInCharset(data, 0, &chars.Charsets0, chars.MASK_DIGIT)
		if newPos != Pos(len(data)) {
			return
		}
	}
}

func BenchmarkParseCharsetAndAlloc(b *testing.B) {
	b.StopTimer()
	allocator := mem.NewArenaAllocator(1024, 1)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		addr, newPos := ParseInCharsetAndAlloc(allocator, data, 0, &chars.Charsets0, chars.MASK_DIGIT)
		if addr == mem.MEM_PTR_NIL || newPos != Pos(len(data)) {
			return
		}
	}
}

func BenchmarkParseInCharsetPercentEscapable1(b *testing.B) {
	b.StopTimer()
	allocator := mem.NewArenaAllocator(1024, 1)
	data := []byte("001234567890%330123456789")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		_, _, err := ParseInCharsetPercentEscapable(allocator, data, 0, &chars.Charsets0, chars.MASK_DIGIT)
		if err != nil {
			return
		}
	}
}

func BenchmarkParseInCharsetPercentEscapable2(b *testing.B) {
	b.StopTimer()
	allocator := mem.NewArenaAllocator(1024, 1)
	data := []byte("%30%30%31%32%33%34%35%36%37%38%39%30%33%301%32%33%34%35%36%37%38%39")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		_, _, err := ParseInCharsetPercentEscapable(allocator, data, 0, &chars.Charsets0, chars.MASK_DIGIT)
		if err != nil {
			return
		}
	}
}

func BenchmarkParseInCharsetPercentEscapableErr(b *testing.B) {
	b.StopTimer()
	allocator := mem.NewArenaAllocator(1024, 1)
	data := []byte("001234567890%330123456789%")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		_, _, err := ParseInCharsetPercentEscapable(allocator, data, 0, &chars.Charsets0, chars.MASK_DIGIT)
		if err == nil {
			return
		}
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
