package model

type Macro struct {
	params *VarList
	code   string
}

func NewMacro() *Macro {
	return &Macro{}
}

func (this *Macro) Append(val *Var) {
	this.params.Append(val)
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
