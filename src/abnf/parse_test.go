package abnf

import (
	"fmt"
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
		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "%30%31123%3a", true, true, 9, 5},
		{1024, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "ad6789abc", false, true, 0, 0},

		{1, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234abc", false, false, 0, 0},
		{2, "IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234abc", false, false, 5, 0},
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

func BenchmarkParseInCharsetPercentEscapable(b *testing.B) {
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
