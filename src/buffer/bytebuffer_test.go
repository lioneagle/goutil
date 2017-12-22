package buffer

import (
	"bytes"
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

func BenchmarkBytesBufferWrite(b *testing.B) {
	b.StopTimer()
	var buf bytes.Buffer

	src := []byte("foobarbaz")

	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(src)
	}
}

func BenchmarkByteBufferWrite(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024*100))
	src := []byte("foobarbaz")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(src)
	}
}
