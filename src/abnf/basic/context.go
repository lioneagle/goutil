package abnf

import (
	"github.com/lioneagle/goutil/src/mem"
)

type Context struct {
	allocator *mem.ArenaAllocator
	parseSrc  []byte
	parsePos  Pos
	srcLen    Pos
}

func NewContext(allocator *mem.ArenaAllocator, src []byte) *Context {
	return &Context{allocator: allocator, parseSrc: src, srcLen: Pos(len(src))}
}

func (this *Context) SetParseSrc(src []byte) {
	this.parseSrc = src
	this.srcLen = Pos(len(src))
}

func (this *Context) SetParsePos(pos Pos) {
	this.parsePos = pos
}

func (this *Context) SetAllocator(allocator *mem.ArenaAllocator) {
	this.allocator = allocator
}
