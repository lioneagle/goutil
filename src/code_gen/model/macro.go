package model

type Macro struct {
	name    string
	comment string
	params  *VarList
	body    Code
}

func NewMacro() *Macro {
	ret := &Macro{}
	ret.params = NewVarList()
	return ret
}

func (this *Macro) AppendParam(val *Var) {
	this.params.Append(val)
}

func (this *Macro) GetName() string {
	return this.name
}

func (this *Macro) SetName(name string) {
	this.name = name
}

func (this *Macro) GetComment() string {
	return this.comment
}

func (this *Macro) SetComment(comment string) {
	this.comment = comment
}

func (this *Macro) GetParams() *VarList {
	return this.params
}

func (this *Macro) GetBody() Code {
	return this.body
}

func (this *Macro) SetBody(code Code) {
	this.body = code
}

func (this *Macro) Accept(v CodeVisitor) {
	v.VisitMacro(this)

}

type MacroList struct {
	Macros []*Macro
}

func NewMacroList() *MacroList {
	return &MacroList{}
}

func (this *MacroList) Append(val *Macro) {
	this.Macros = append(this.Macros, val)
}
