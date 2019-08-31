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
		newPos  int
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
		name    string
		charset *[256]uint32
		mask    uint32
		src     string
		isOk    bool
		newPos  int
	}{

		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "01234abc", true, 5},
		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "56789=bc", true, 5},
		{"IsDigit", &chars.Charsets0, chars.MASK_DIGIT, "ad6789abc", false, 0},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			allocator := mem.NewArenaAllocator(1024, 1)
			addr, newPos := ParseInCharsetAndAlloc(allocator, []byte(v.src), 0, v.charset, v.mask)

			if v.isOk {
				test.EXPECT_NE(t, addr, mem.MEM_PTR_NIL, "")
				test.EXPECT_EQ(t, allocator.Strlen(addr), v.newPos, "")
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
		if newPos != len(data) {
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
		if addr == mem.MEM_PTR_NIL || newPos != len(data) {
			return
		}
	}
}
