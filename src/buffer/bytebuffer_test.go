package buffer

import (
	"testing"
)

func TestBytesBufferWrite(t *testing.T) {
	funcName := "TestBytesBufferWrite"

	testdata := []struct {
		data string
	}{
		{"abc"},
		{"124aadf"},
	}

	buf := NewByteBuffer(nil)

	for i, v := range testdata {
		len1, err := buf.Write([]byte(v.data))
		if err != nil {
			t.Errorf("%s[%d] failed, err = %v, wanted = nil\n", funcName, i, err)
		}

		if len1 != len(v.data) {
			t.Errorf("%s[%d] failed, len = %v, wanted = %v\n", funcName, i, len1, len(v.data))
		}
	}
}
