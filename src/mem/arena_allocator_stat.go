package mem

import (
	"fmt"
	"io"

	"github.com/lioneagle/goutil/src/buffer"
)

type StatNumber = uint64

type ArenaAllocatorStat struct {
	allocNum    StatNumber
	allocNumOk  StatNumber
	freeAllNum  StatNumber
	freePartNum StatNumber
}

func (this *ArenaAllocatorStat) Init() {
	this.allocNum = 0
	this.allocNumOk = 0
	this.freeAllNum = 0
	this.freePartNum = 0

}

func (this *ArenaAllocatorStat) ClearAllocNum() {
	this.allocNum = 0
}

func (this *ArenaAllocatorStat) AllocNum() StatNumber {
	return this.allocNum
}

func (this *ArenaAllocatorStat) AllocNumOk() StatNumber {
	return this.allocNumOk
}

func (this *ArenaAllocatorStat) FreeAllNum() StatNumber {
	return this.freeAllNum
}

func (this *ArenaAllocatorStat) FreePartNum() StatNumber {
	return this.freePartNum
}

func (this *ArenaAllocatorStat) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.Print(buf)
	return buf.String()
}

func (this *ArenaAllocatorStat) Print(w io.Writer) {
	stat := []struct {
		name string
		num  StatNumber
	}{
		{"alloc num", this.allocNum},
		{"alloc num ok", this.allocNumOk},
		{"free all num", this.freeAllNum},
		{"free part num", this.freePartNum},
	}

	hasNonZero := false

	for _, v := range stat {
		if v.num > 0 {
			hasNonZero = true
			io.WriteString(w, v.name)
			io.WriteString(w, ": ")
			fmt.Fprintf(w, "%d\n", v.num)
		}
	}

	if !hasNonZero {
		io.WriteString(w, "all stats are zero")
	}
}
