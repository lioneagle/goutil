package backends

import (
	"fmt"
	"io"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/goutil/src/chars"

	"github.com/lioneagle/goutil/src/code_gen/backends"
	"github.com/lioneagle/goutil/src/code_gen/model"
)

type GolangGenerator struct {
	chars.Indent
	w io.Writer

	config backends.CConfig

	maxTypeNameLen int
	maxNameLen     int
	maxValueLen    int
}

func NewGolangGenerator(w io.Writer, config backends.CConfig) *GolangGenerator {
	gen := &GolangGenerator{}
	gen.Init(w, config)
	return gen
}

func (this *GolangGenerator) Init(w io.Writer, config backends.CConfig) {
	this.w = w
	this.config = config
	this.Indent.Init(0, 4)
}

func (this *GolangGenerator) genBlockBegin(indent int) {
	if this.config.BraceAtNextLine() {
		this.PrintReturn(this.w)
		this.Fprintln(this.w, "{")

	} else {
		fmt.Fprintln(this.w, " {")
	}

	this.EnterIndent(indent)
}

func (this *GolangGenerator) genBlockEnd() {
	this.Exit()
	this.Fprintln(this.w, "}")
}

func (this *GolangGenerator) genSingleLineCommentWithoutIndent(comment string) {
	if this.config.VarUseSingleLineComment() {
		fmt.Fprintf(this.w, "// %s", comment)
	} else {
		fmt.Fprintf(this.w, "/* %s */", comment)
	}
	this.PrintReturn(this.w)
}

func (this *GolangGenerator) VisitBlockBegin(val *model.Block) {
	this.genBlockBegin(this.config.Indent().Block)
}

func (this *GolangGenerator) VisitBlockEnd(val *model.Block) {
	this.genBlockEnd()
}

func (this *GolangGenerator) VisitSentence(val *model.Sentence) {
	this.Fprintln(this.w, val.GetCode())
}

func (this *GolangGenerator) VisitStructDefineBegin(val *model.Struct) {

}

func (this *GolangGenerator) VisitStructField(val *model.Struct) {

}

func (this *GolangGenerator) VisitStructDefineEnd(val *model.Struct) {

}

func (this *GolangGenerator) VisitConstsBegin(val *model.ConstList) {
	this.Fprintf(this.w, "typedef enum tag_%s", val.GetName())
	this.genBlockBegin(this.config.Indent().Enum)

	this.maxTypeNameLen = val.GetMaxTypeNameLen()
	this.maxNameLen = val.GetMaxNameLen()
	this.maxValueLen = val.GetMaxValueLen()
}

func (this *GolangGenerator) VisitConst(val *model.Var) {
	this.Fprint(this.w, val.GetName())
	if len(val.GetInitValue()) > 0 {
		basic.PrintIndent(this.w, this.maxNameLen-len(val.GetName())+this.config.Indent().Assign)
		fmt.Fprintf(this.w, "= %s,", val.GetInitValue())
	} else {
		fmt.Fprint(this.w, ",")
	}

	if len(val.GetComment()) > 0 {
		basic.PrintIndent(this.w, this.maxValueLen-len(val.GetInitValue())+this.config.Indent().Comment)
		this.genSingleLineCommentWithoutIndent(val.Comment)
	} else {
		this.PrintReturn(this.w)
	}
}

func (this *GolangGenerator) VisitConstsEnd(val *model.ConstList) {
	this.Exit()
	this.Fprintfln(this.w, "}%s;", val.GetName())
}
