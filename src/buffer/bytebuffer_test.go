package buffer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestBytesBufferWrite(t *testing.T) {
	testdata := []struct {
		data string
	}{
		{"abc"},
		{"124aadf"},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			buf := NewByteBuffer(nil)
			len1, err := buf.Write([]byte(v.data))

			test.ASSERT_EQ(t, err, nil, "")
			test.EXPECT_EQ(t, len1, len(v.data), "")
			test.EXPECT_EQ(t, len1, buf.Len(), "")
			test.EXPECT_EQ(t, buf.String(), v.data, "")
			test.EXPECT_EQ(t, bytes.Equal(buf.Bytes(), []byte(v.data)), true, "")
		})
	}
}

func TestBytesBufferWritef(t *testing.T) {
	buf := NewByteBuffer(nil)
	buf.Writef("abc_%d", 101)
	wanted := "abc_101"

	str := buf.String()

	test.EXPECT_EQ(t, str, wanted, "")
}

func TestWriterWriteln(t *testing.T) {
	buf := NewByteBuffer(nil)
	buf.WriteByte(';')
	buf.WriteString("cde")
	buf.Writeln("abc_%d", 101)
	wanted := ";cdeabc_101" + fmt.Sprintln()

	str := buf.String()

	test.EXPECT_EQ(t, str, wanted, "")
}

func TestWriterPrintAsHex1(t *testing.T) {
	buf := NewByteBuffer(nil)
	buf.WriteString("123")

	buf2 := NewByteBuffer(nil)
	buf.PrintAsHex(buf2, 0, buf.Len())

	str := buf2.String()
	wanted := "00000000h: 31 32 33                                         ; 123\n"

	test.EXPECT_EQ(t, str, wanted, "")
}

func TestWriterPrintAsHex2(t *testing.T) {
	buf := NewByteBuffer(nil)
	buf.WriteString("12345678")
	buf.WriteByte(0)
	buf.WriteByte(0xA2)

	buf2 := NewByteBuffer(nil)
	buf.PrintAsHex(buf2, -1, buf.Len()+1)

	str := buf2.String()
	wanted := "00000000h: 31 32 33 34 35 36 37 38  00 A2                   ; 12345678 .?\n"

	test.EXPECT_EQ(t, str, wanted, "")
}

func TestWriterPrintAsHex3(t *testing.T) {
	buf := NewByteBuffer(nil)
	buf.WriteString("12345678abcdef-=")
	buf.WriteByte(0)
	buf.WriteByte(0xA2)

	buf2 := NewByteBuffer(nil)
	buf.PrintAsHex(buf2, -1, buf.Len()+1)

	str := buf2.String()
	wanted := "00000000h: 31 32 33 34 35 36 37 38  61 62 63 64 65 66 2D 3D ; 12345678 abcdef-=\n" +
		"00000010h: 00 A2                                            ; .?\n"

	test.EXPECT_EQ(t, str, wanted, "")
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

func BenchmarkByteBufferWriteString(b *testing.B) {
	b.StopTimer()
	buf := NewByteBuffer(make([]byte, 1024*100))
	src := "foobarbaz"
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString(src)
	}
}
