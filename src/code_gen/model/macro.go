package model

type Macro interface {
	Code
	IsMacro() bool
}

type MacroDefine struct {
	name      string
	comment   string
	hasParams bool
	params    *VarList
	body      Code
}

func NewMacroDefine() *MacroDefine {
	ret := &MacroDefine{}
	ret.params = NewVarList()
	return ret
}

func (this *MacroDefine) AppendParam(val ...*Var) {
	if len(val) > 0 {
		this.params.Append(val...)
		this.hasParams = true
	}
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

func (this *MacroDefine) SetHasParams() {
	this.hasParams = true
}

func (this *MacroDefine) SetNoParams() {
	this.hasParams = false
}

func (this *MacroDefine) HasParams() bool {
	return this.hasParams
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

type MacroUndefine struct {
	comment string
	value   string
}

func NewMacroUndefine() *MacroUndefine {
	return &MacroUndefine{}
}

func (this *MacroUndefine) GetComment() string {
	return this.comment
}

func (this *MacroUndefine) SetComment(comment string) {
	this.comment = comment
}

func (this *MacroUndefine) GetValue() string {
	return this.value
}

func (this *MacroUndefine) SetValue(value string) {
	this.value = value
}

func (this *MacroUndefine) IsMacro() bool {
	return true
}

func (this *MacroUndefine) Accept(v CodeVisitor) {
	v.VisitMacroUndefine(this)
}

type MacroDefineList struct {
	macros []*MacroDefine
}

func NewMacroDefineList() *MacroDefineList {
	return &MacroDefineList{
		macros: make([]*MacroDefine, 0),
	}
}

func (this *MacroDefineList) Append(val ...*MacroDefine) {
	this.macros = append(this.macros, val...)
}

func (this *MacroDefineList) IsMacro() bool {
	return true
}

type MacroList struct {
	macros []*Macro
}

func NewMacroList() *MacroList {
	return &MacroList{
		macros: make([]*Macro, 0),
	}
}

func (this *MacroList) Append(val ...*Macro) *MacroList {
	this.macros = append(this.macros, val...)
	return this
}

func (this *MacroList) IsMacro() bool {
	return true
}
