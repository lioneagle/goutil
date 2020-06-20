package backends

import (
	"fmt"
	"io"
)

type Indents struct {
	If        int
	Switch    int
	Case      int
	While     int
	For       int
	DoWhile   int
	Block     int
	Struct    int
	FuncParam int
	Enum      int
	Union     int
	Assign    int
	Comment   int
}

func (this *Indents) Init() {
	this.If = 4
	this.Switch = 4
	this.Case = 4
	this.While = 4
	this.For = 4
	this.DoWhile = 4
	this.Block = 4
	this.Struct = 4
	this.FuncParam = 4
	this.Enum = 4
	this.Union = 4
	this.Assign = 1
	this.Comment = 1
}

func (this *Indents) String() string {
	return fmt.Sprintf("indents = %+v", *this)
}

type ConfigCommon interface {
	Indent() *Indents
	SetIndent(val *Indents)

	GenVersion() bool
	SetGenVersion(val bool)

	PrintTimeUsed() bool
	SetPrintTimeUsed(val bool)

	UseTabIndent() bool
	SetUseTabIndent(val bool)

	OutputPath() string
	SetOutputPath(val string)

	DebugFileName() string
	SetDebugFileName(val string)

	DebugFile() io.Writer
	SetDebugFile(val io.Writer)

	VarUseSingleLineComment() bool
	SetVarUseSingleLineComment(val bool)

	ParamsInOneLine() bool
	SetParamsInOneLine(val bool)

	MultiLineCommentDecorate() bool
	SetMultiLineCommentDecorate(val bool)
}

type ConfigCommonImpl struct {
	indent Indents

	genVersion    bool
	printTimeUsed bool
	useTabIndent  bool
	outputPath    string
	debugFileName string
	debugFile     io.Writer

	varUseSingleLineComment  bool
	paramsInOneLine          bool
	multiLineCommentDecorate bool
}

func NewConfigCommon() *Config {
	ret := &Config{}
	return ret
}

func (this *ConfigCommonImpl) Init() {
	this.printTimeUsed = true
	this.outputPath = "./"
	this.indent.Init()
	//this.varUseSingleLineComment = true
	this.paramsInOneLine = true

}

func (this *ConfigCommonImpl) String() string {
	return fmt.Sprintf("config = %+v", *this)
}

func (this *ConfigCommonImpl) Indent() *Indents {
	return &this.indent
}

func (this *ConfigCommonImpl) SetIndent(val *Indents) {
	this.indent = *val
}

func (this *ConfigCommonImpl) GenVersion() bool {
	return this.genVersion
}

func (this *ConfigCommonImpl) SetGenVersion(val bool) {
	this.genVersion = val
}

func (this *ConfigCommonImpl) PrintTimeUsed() bool {
	return this.printTimeUsed
}

func (this *ConfigCommonImpl) SetPrintTimeUsed(val bool) {
	this.printTimeUsed = val
}

func (this *ConfigCommonImpl) UseTabIndent() bool {
	return this.useTabIndent
}

func (this *ConfigCommonImpl) SetUseTabIndent(val bool) {
	this.useTabIndent = val
}

func (this *ConfigCommonImpl) OutputPath() string {
	return this.outputPath
}

func (this *ConfigCommonImpl) SetOutputPath(val string) {
	this.outputPath = val
}

func (this *ConfigCommonImpl) DebugFileName() string {
	return this.debugFileName
}

func (this *ConfigCommonImpl) SetDebugFileName(val string) {
	this.debugFileName = val
}

func (this *ConfigCommonImpl) DebugFile() io.Writer {
	return this.debugFile
}

func (this *ConfigCommonImpl) SetDebugFile(val io.Writer) {
	this.debugFile = val
}

func (this *ConfigCommonImpl) VarUseSingleLineComment() bool {
	return this.varUseSingleLineComment
}

func (this *ConfigCommonImpl) SetVarUseSingleLineComment(val bool) {
	this.varUseSingleLineComment = val
}

func (this *ConfigCommonImpl) ParamsInOneLine() bool {
	return this.paramsInOneLine
}

func (this *ConfigCommonImpl) SetParamsInOneLine(val bool) {
	this.paramsInOneLine = val
}

func (this *ConfigCommonImpl) MultiLineCommentDecorate() bool {
	return this.multiLineCommentDecorate
}

func (this *ConfigCommonImpl) SetMultiLineCommentDecorate(val bool) {
	this.multiLineCommentDecorate = val
}
