package model

type Macro interface {
	Code
	IsMacro() bool
}

type MacroDefine struct {
	name    string
	comment string
	params  *VarList
	body    Code
}

func NewMacroDefine() *MacroDefine {
	ret := &MacroDefine{}
	ret.params = NewVarList()
	return ret
}

func (this *MacroDefine) AppendParam(val *Var) {
	this.params.Append(val)
}

func (this *MacroDefine) IsMacro() bool {
	return true
}

func (this *MacroDefine) GetName() string {
	return this.name
}

func (this *MacroDefine) SetName(name string) {
	this.name = name
}

func (this *MacroDefine) GetComment() string {
	return this.comment
}

func (this *MacroDefine) SetComment(comment string) {
	this.comment = comment
}

func (this *MacroDefine) GetParams() *VarList {
	return this.params
}

func (this *MacroDefine) GetBody() Code {
	return this.body
}

func (this *MacroDefine) SetBody(code Code) {
	this.body = code
}

func (this *MacroDefine) Accept(v CodeVisitor) {
	v.VisitMacroDefine(this)

}

type MacroDefineList struct {
	macros []*MacroDefine
}

func NewMacroDefineList() *MacroDefineList {
	return &MacroDefineList{}
}

func (this *MacroDefineList) Append(val *MacroDefine) {
	this.macros = append(this.macros, val)
}

func (this *MacroDefineList) IsMacro() bool {
	return true
}

type MacroList struct {
	macros []*Macro
}

func NewMacroList() *MacroList {
	return &MacroList{}
}

func (this *MacroList) Append(val *Macro) {
	this.macros = append(this.macros, val)
}

func (this *MacroList) IsMacro() bool {
	return true
}
