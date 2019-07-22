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

func (this *CLikeGeneratorBase) genBlockBegin(indent int, spaceBeforBrace int) {
	if this.config.BraceAtNextLine() {
		this.PrintReturn(this.w)
		this.Fprintln(this.w, "{")

	} else {
		chars.PrintIndent(this.w, spaceBeforBrace)
		fmt.Fprintln(this.w, "{")
	}

	this.EnterIndent(indent)
}

func (this *CLikeGeneratorBase) genBlockEnd() {
	this.Exit()
	this.Fprintln(this.w, "}")
}

func (this *CLikeGeneratorBase) genSingleLineCommentWithoutIndent(comment string) {
	if len(comment) == 0 {
		return
	}

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

func (this *CLikeGeneratorBase) VisitComment(val *model.Comment) {
	if val.IsSingleLine() {
		this.genSingleLineCommentWithoutIndent(val.GetComment())
	} else {
		this.GenMultiLineComment(val.GetComment())
	}
}

func (this *CLikeGeneratorBase) VisitBlockBegin(val *model.Block) {
	this.genBlockBegin(this.config.Indent().Block, 1)
}

func (this *CLikeGeneratorBase) VisitBlockEnd(val *model.Block) {
	this.genBlockEnd()
}

func (this *CLikeGeneratorBase) VisitSentence(val *model.Sentence) {
	this.Fprintln(this.w, val.GetCode())
}

func (this *CLikeGeneratorBase) VisitStructBegin(val *model.Struct) {
	this.Fprintf(this.w, "typedef struct tag%s", val.GetName())
	this.genBlockBegin(this.config.Indent().Struct, 1)
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
	this.Fprintf(this.w, "typedef enum tag%s", val.GetName())
	this.genBlockBegin(this.config.Indent().Enum, 1)

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

func (this *CLikeGeneratorBase) VisitSingleChoiceBegin(val *model.SingleChoice) {
	this.GenMultiLineComment(val.GetComment())

	this.Fprint(this.w, "if (")
	fmt.Fprintf(this.w, val.GetCondition())
	fmt.Fprint(this.w, ")")
	this.genBlockBegin(this.config.Indent().If, 1)
}

func (this *CLikeGeneratorBase) VisitSingleChoiceEnd(val *model.SingleChoice) {
	this.PrintReturn(this.w)
}

func (this *CLikeGeneratorBase) VisitSingleChoiceTrueBegin(val *model.SingleChoice) {
	// do nothing now
}

func (this *CLikeGeneratorBase) VisitSingleChoiceTrueEnd(val *model.SingleChoice) {
	this.Exit()
	this.Fprint(this.w, "}")
}

func (this *CLikeGeneratorBase) VisitSingleChoiceFalseBegin(val *model.SingleChoice) {
	if !this.config.BraceAtNextLine() {
		fmt.Fprintf(this.w, " else {")
		this.PrintReturn(this.w)
	} else {
		this.PrintReturn(this.w)
		this.Fprintln(this.w, "else")
		this.Fprintln(this.w, "{")
	}
	this.EnterIndent(this.config.Indent().If)
}

func (this *CLikeGeneratorBase) VisitSingleChoiceFalseEnd(val *model.SingleChoice) {
	this.Exit()
	this.Fprint(this.w, "}")
}

func (this *CLikeGeneratorBase) VisitRepeatAsForBegin(val *model.Repeat) {
	this.GenMultiLineComment(val.GetComment())
	this.Fprintf(this.w, "for (%s)", val.GetCondition())
	this.genBlockBegin(this.config.Indent().For, 1)
}

func (this *CLikeGeneratorBase) VisitRepeatAsForEnd(val *model.Repeat) {
	this.genBlockEnd()
}

func (this *CLikeGeneratorBase) VisitRepeatAsWhileBegin(val *model.Repeat) {
	this.GenMultiLineComment(val.GetComment())
	this.Fprintf(this.w, "while (%s)", val.GetCondition())
	this.genBlockBegin(this.config.Indent().While, 1)
}

func (this *CLikeGeneratorBase) VisitRepeatAsWhileEnd(val *model.Repeat) {
	this.genBlockEnd()
}

func (this *CLikeGeneratorBase) VisitRepeatAsDoWhileBegin(val *model.Repeat) {
	this.GenMultiLineComment(val.GetComment())
	this.Fprint(this.w, "do")
	this.genBlockBegin(this.config.Indent().DoWhile, 0)
}

func (this *CLikeGeneratorBase) VisitRepeatAsDoWhileEnd(val *model.Repeat) {
	this.Exit()
	this.Fprintfln(this.w, "}while(%s);", val.GetCondition())
}

func (this *CLikeGeneratorBase) VisitParamVarFirst(val *model.Var) {
	if this.config.ParamsInOneLine() {
		fmt.Fprintf(this.w, "%s %s", val.GetTypeName(), val.GetName())
	} else {
		this.Fprintf(this.w, "%s %s", val.GetTypeName(), val.GetName())
	}
}

func (this *CLikeGeneratorBase) VisitParamVarNonFirstBegin() {
	this.EnterIndent(this.config.Indent().FuncParam)
}

func (this *CLikeGeneratorBase) VisitParamVarNonFirst(val *model.Var) {
	if this.config.ParamsInOneLine() {
		fmt.Fprintf(this.w, ", %s %s", val.GetTypeName(), val.GetName())
	} else {
		fmt.Fprint(this.w, ",")
		this.PrintReturn(this.w)
		this.Fprintf(this.w, "%s %s", val.GetTypeName(), val.GetName())
	}
}

func (this *CLikeGeneratorBase) VisitParamVarNonFirstEnd() {
	this.Exit()
}

func (this *CLikeGeneratorBase) VisitFuncDeclare(val *model.Function) {
	this.visitFuncDeclareBase(val)
	fmt.Fprint(this.w, ";")
	this.PrintReturn(this.w)

}

func (this *CLikeGeneratorBase) VisitFuncDefine(val *model.Function) {
	this.visitFuncDeclareBase(val)
	val.GetBody().Accept(this)
}

func (this *CLikeGeneratorBase) VisitFuncNoReturn() {
	this.Fprint(this.w, "void ")
}

func (this *CLikeGeneratorBase) visitFuncDeclareBase(val *model.Function) {
	this.GenMultiLineComment(val.GetComment())
	val.GetReturnList().AcceptAsFuncReturns(this)
	fmt.Fprintf(this.w, "%s(", val.GetName())
	val.GetParams().AcceptAsParmList(this)
	fmt.Fprint(this.w, ")")

}

func (this *CLikeGeneratorBase) VisitFuncReturnFirst(val *model.Var) {
	fmt.Fprintf(this.w, "%s ", val.GetTypeName())
}

func (this *CLikeGeneratorBase) VisitFuncReturnNonFirst(val *model.Var) {

}
