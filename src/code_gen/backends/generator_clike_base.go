package backends

import (
	"fmt"
	"io"
	"strings"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/goutil/src/chars"

	"github.com/lioneagle/goutil/src/code_gen/model"
)

type CLikeGeneratorBase struct {
	chars.Indent
	w io.Writer

	config CConfig

	maxTypeNameLen int
	maxNameLen     int
	maxValueLen    int
}

func NewCLikeGeneratorBase(w io.Writer, config CConfig) *CLikeGeneratorBase {
	gen := &CLikeGeneratorBase{}
	gen.Init(w, config)
	return gen
}

func (this *CLikeGeneratorBase) Init(w io.Writer, config CConfig) {
	this.w = w
	this.config = config
	this.Indent.Init(0, 4)
}

func (this *CLikeGeneratorBase) genBlockBegin(indent int) {
	if this.config.BraceAtNextLine() {
		this.PrintReturn(this.w)
		this.Fprintln(this.w, "{")

	} else {
		fmt.Fprintln(this.w, " {")
	}

	this.EnterIndent(indent)
}

func (this *CLikeGeneratorBase) genBlockEnd() {
	this.Exit()
	this.Fprintln(this.w, "}")
}

func (this *CLikeGeneratorBase) genSingleLineCommentWithoutIndent(comment string) {
	if this.config.VarUseSingleLineComment() {
		fmt.Fprintf(this.w, "// %s", comment)
	} else {
		fmt.Fprintf(this.w, "/* %s */", comment)
	}
	this.PrintReturn(this.w)
}

func (this *CLikeGeneratorBase) GenMultiLineComment(comment string) {
	lines := strings.Split(comment, "\n")
	if len(lines) <= 0 {
		return
	}

	newLines := make([]string, 0)
	max_len := 0

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
		if len(lines[i]) > 0 {
			newLines = append(newLines, lines[i])
		}

		if len(lines[i]) > max_len {
			max_len = len(lines[i])
		}
	}

	if len(newLines) <= 0 {
		return
	}

	if len(newLines) == 1 {
		this.Fprintfln(this.w, "/* %s */", newLines[0])
	} else if this.config.MultiLineCommentDecorate() {
		this.Fprint(this.w, "/*")
		basic.PrintChars(this.w, '*', max_len+2)
		this.PrintReturn(this.w)

		for i := 0; i < len(newLines); i++ {
			fmt.Fprintf(this.w, " * %s", newLines[i])
			this.PrintReturn(this.w)
		}
		fmt.Fprint(this.w, " ")
		basic.PrintChars(this.w, '*', max_len+3)
		this.Fprintln(this.w, "*/")
	} else {
		this.Fprint(this.w, "/* ")
		fmt.Fprint(this.w, newLines[0])
		this.PrintReturn(this.w)

		for i := 1; i < len(newLines); i++ {
			fmt.Fprintf(this.w, " * %s", newLines[i])
			this.PrintReturn(this.w)
		}
		this.Fprintln(this.w, " */")
	}
}

func (this *CLikeGeneratorBase) VisitBlockBegin(val *model.Block) {
	this.genBlockBegin(this.config.Indent().Block)
}

func (this *CLikeGeneratorBase) VisitBlockEnd(val *model.Block) {
	this.genBlockEnd()
}

func (this *CLikeGeneratorBase) VisitSentence(val *model.Sentence) {
	this.Fprintln(this.w, val.GetCode())
}

func (this *CLikeGeneratorBase) VisitStructBegin(val *model.Struct) {
	this.Fprintf(this.w, "typedef struct tag_%s", val.GetName())
	this.genBlockBegin(this.config.Indent().Struct)
}

func (this *CLikeGeneratorBase) VisitStructEnd(val *model.Struct) {
	this.Exit()
	this.Fprintfln(this.w, "}%s;", val.GetName())
}

func (this *CLikeGeneratorBase) VisitStructRangePublicBegin(val *model.Struct) {
	this.Fprintln(this.w, "/* -------- public begin -------- */")
	this.PrintReturn(this.w)
}

func (this *CLikeGeneratorBase) VisitStructRangePublicEnd(val *model.Struct) {
	this.PrintReturn(this.w)
	this.Fprintln(this.w, "/* -------- public end -------- */")
}

func (this *CLikeGeneratorBase) VisitStructRangeProtectedBegin(val *model.Struct) {
	this.Fprintln(this.w, "/* -------- protected begin -------- */")
	this.PrintReturn(this.w)
}

func (this *CLikeGeneratorBase) VisitStructRangeProtectedEnd(val *model.Struct) {
	this.PrintReturn(this.w)
	this.Fprintln(this.w, "/* -------- protected end -------- */")
}

func (this *CLikeGeneratorBase) VisitStructRangePrivateBegin(val *model.Struct) {
	this.Fprintln(this.w, "/* -------- private begin -------- */")
	this.PrintReturn(this.w)
}

func (this *CLikeGeneratorBase) VisitStructRangePrivateEnd(val *model.Struct) {
	this.PrintReturn(this.w)
	this.Fprintln(this.w, "/* -------- private end -------- */")
}

func (this *CLikeGeneratorBase) VisitStructFieldVarListBegin(val *model.VarList) {
	this.maxTypeNameLen = val.GetMaxTypeNameLen()
	this.maxNameLen = val.GetMaxNameLen()
}

func (this *CLikeGeneratorBase) VisitStructFieldVar(val *model.Var) {
	this.Fprintf(this.w, "%s", val.GetTypeName())
	basic.PrintIndent(this.w, this.maxTypeNameLen-len(val.GetTypeName())+1)
	fmt.Fprintf(this.w, "%s;", val.GetName())

	if len(val.GetComment()) > 0 {
		indent := this.maxNameLen - len(val.GetName()) + this.config.Indent().Comment
		basic.PrintIndent(this.w, indent)
		this.genSingleLineCommentWithoutIndent(val.GetComment())
	} else {
		this.PrintReturn(this.w)
	}
}

func (this *CLikeGeneratorBase) VisitStructFieldVarListEnd(val *model.VarList) {
}

func (this *CLikeGeneratorBase) VisitConstsBegin(val *model.ConstList) {
	this.Fprintf(this.w, "typedef enum tag_%s", val.GetName())
	this.genBlockBegin(this.config.Indent().Enum)

	this.maxTypeNameLen = val.GetMaxTypeNameLen()
	this.maxNameLen = val.GetMaxNameLen()
	this.maxValueLen = val.GetMaxValueLen()
}

func (this *CLikeGeneratorBase) VisitConst(val *model.Var) {
	this.Fprint(this.w, val.GetName())
	if len(val.GetInitValue()) > 0 {
		basic.PrintIndent(this.w, this.maxNameLen-len(val.GetName())+this.config.Indent().Assign)
		fmt.Fprintf(this.w, "= %s,", val.GetInitValue())
	} else {
		fmt.Fprint(this.w, ",")
	}

	if len(val.GetComment()) > 0 {
		indent := this.maxValueLen - len(val.GetInitValue()) + this.config.Indent().Comment
		if len(val.GetInitValue()) <= 0 {
			indent += this.config.Indent().Assign + len("= ")
		}
		basic.PrintIndent(this.w, indent)
		this.genSingleLineCommentWithoutIndent(val.GetComment())
	} else {
		this.PrintReturn(this.w)
	}
}

func (this *CLikeGeneratorBase) VisitConstsEnd(val *model.ConstList) {
	this.Exit()
	this.Fprintfln(this.w, "}%s;", val.GetName())
}
