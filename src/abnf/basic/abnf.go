package abnf

import (
	"github.com/lioneagle/goutil/src/buffer"
)

type Pos = uint
type ByteBuffer = buffer.ByteBuffer

func NewByteBuffer(buf []byte) *ByteBuffer {
	return buffer.NewByteBuffer(buf)
}
