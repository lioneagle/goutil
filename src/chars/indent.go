package chars

import (
	"fmt"
	"io"
)

type Indent struct {
	indent int
	delta  int
}

func NewIndent(initIndent, delta int) *Indent {
	return &Indent{indent: initIndent, delta: delta}
}

func (this *Indent) Enter() {
	this.indent += this.delta
}

func (this *Indent) Exit() {
	this.indent -= this.delta
}

func (this *Indent) Print(w io.Writer) {
	for i := 0; i < this.indent; i++ {
		fmt.Fprint(w, " ")
	}
}
