package chars

import (
	"bytes"
	"testing"
)

func TestUnescape(t *testing.T) {

	wanted := []struct {
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

	for i, v := range wanted {
		u := Unescape([]byte(v.escaped))
		if bytes.Compare(u, []byte(v.unescaped)) != 0 {
			t.Errorf("TestUnescape[%d] failed, ret = %s, wanted = %s\n", i, string(u), v.unescaped)
		}
	}
}

func TestEscape(t *testing.T) {

	wanted := []struct {
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

	for i, v := range wanted {
		u := Escape(chars, v.isInCharset)
		if !bytes.Equal(Unescape(u), chars) {
			t.Errorf("TestEscape[%d]: %s failed\n", i, v.name)
		}
	}
}

func makeFullCharset() (ret []byte) {
	for i := 0; i < 256; i++ {
		ret = append(ret, byte(i))
	}
	return ret
}
