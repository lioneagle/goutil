package chars

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestUnescape(t *testing.T) {
	testdata := []struct {
		escaped   string
		unescaped string
	}{

		{"a%42c", "aBc"},
		{"a%3B", "a;"},
		{"a%3b%42", "a;B"},
		{"ac%3", "ac%3"},
		{"ac%P3", "ac%P3"},
		{"ac%", "ac%"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			u := Unescape([]byte(v.escaped))
			test.EXPECT_EQ(t, string(u), v.unescaped, "")
		})
	}
}

func TestEscape(t *testing.T) {
	testdata := []struct {
		name        string
		isInCharset func(ch byte) bool
	}{
		{"IsDigit", IsDigit},
		{"IsAlpha", IsAlpha},
		{"IsLower", IsLower},
		{"IsUpper", IsUpper},
		{"IsAlphanum", IsAlphanum},
		{"IsLowerHexAlpha", IsLowerHexAlpha},
		{"IsUpperHexAlpha", IsUpperHexAlpha},
		{"IsLowerHex", IsLowerHex},
		{"IsUpperHex", IsUpperHex},
		{"IsHex", IsHex},
		{"IsCrlfChar", IsCrlfChar},
		{"IsWspChar", IsWspChar},
		{"IsLwsChar", IsLwsChar},
		{"IsAscii", IsAscii},
		{"IsUtf8N1", IsUtf8N1},
		{"IsUtf8N2", IsUtf8N2},
		{"IsUtf8N3", IsUtf8N3},
		{"IsUtf8N4", IsUtf8N4},
		{"IsUtf8N5", IsUtf8N5},
		{"IsUtf8N6", IsUtf8N6},
		{"IsUtf8Cont", IsUtf8Cont},
		{"IsUtf8Char", IsUtf8Char},

		{"IsUriUnreserved", IsUriUnreserved},
		{"IsUriReserved", IsUriReserved},
		{"IsUriUric", IsUriUric},
		{"IsUriUricNoSlash", IsUriUricNoSlash},
		{"IsUriPchar", IsUriPchar},
		{"IsUriScheme", IsUriScheme},
		{"IsUriRegName", IsUriRegName},
	}

	chars := makeFullCharset()

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			u := Escape(chars, v.isInCharset)
			test.EXPECT_TRUE(t, bytes.Equal(Unescape(u), chars), "")
		})
	}
}

func makeFullCharset() (ret []byte) {
	for i := 0; i < 256; i++ {
		ret = append(ret, byte(i))
	}
	return ret
}
