package chars

import (
	"fmt"
	"io"
)

func PrintIndent(w io.Writer, indent int) {
	if indent <= 0 {
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%%%ds", indent), "")
}

type Indent struct {
	UseTab        bool
	DefaultIndent int
	Stack         []int
	returns       Return
}

func NewIndent(initIndent, defaultIndent int) *Indent {
	ret := &Indent{}
	ret.Init(initIndent, defaultIndent)
	return ret
}

func (this *Indent) Init(initIndent, defaultIndent int) {
	this.DefaultIndent = defaultIndent
	this.Stack = nil
	this.Stack = append(this.Stack, initIndent)
	this.returns.Init()
}

func (this *Indent) CopyStack(rhs *Indent) {
	copy(this.Stack, rhs.Stack)
}

func (this *Indent) Enter() {
	this.EnterIndent(this.DefaultIndent)
}

func (this *Indent) SetReturnString(ret string) {
	this.returns.SetReturnString(ret)
}

func (this *Indent) EnterIndent(indent int) {
	if len(this.Stack) == 0 {
		if !this.UseTab {
			this.Stack = append(this.Stack, indent)
		} else {
			this.Stack = append(this.Stack, 1)
		}
	} else {
		if !this.UseTab {
			this.Stack = append(this.Stack, this.Stack[len(this.Stack)-1]+indent)
		} else {
			this.Stack = append(this.Stack, this.Stack[len(this.Stack)-1]+1)
		}
	}
}

func (this *Indent) Exit() {
	if len(this.Stack) == 0 {
		return
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *Indent) printIndent(w io.Writer) {
	if len(this.Stack) <= 0 {
		return
	}

	num := this.Stack[len(this.Stack)-1]
	if this.UseTab {
		for i := 0; i < num; i++ {
			fmt.Fprint(w, "\t")
		}
	} else {
		for i := 0; i < num; i++ {
			fmt.Fprint(w, " ")
		}
	}
}

func (this *Indent) Fprint(w io.Writer, args ...interface{}) {
	this.printIndent(w)
	fmt.Fprint(w, args...)
}

func (this *Indent) Fprintln(w io.Writer, args ...interface{}) {
	if len(args) > 0 {
		this.printIndent(w)
		fmt.Fprint(w, args...)
	}

	this.PrintReturn(w)
}

func (this *Indent) Fprintf(w io.Writer, format string, args ...interface{}) {
	this.printIndent(w)
	fmt.Fprintf(w, format, args...)
}

func (this *Indent) Fprintfln(w io.Writer, format string, args ...interface{}) {
	this.printIndent(w)
	fmt.Fprintf(w, format, args...)
	//fmt.Fprint(w)
	this.PrintReturn(w)
}

func (this *Indent) PrintReturn(w io.Writer) {
	fmt.Fprint(w, this.returns.returnString)
}
