package model

import (
	_ "fmt"
)

type Var struct {
	typeName  string
	name      string
	comment   string
	initValue string
}

func NewVar() *Var {
	return &Var{}
}

func (this *Var) GetTypeName() string {
	return this.typeName
}

func (this *Var) GetName() string {
	return this.name
}

func (this *Var) GetInitValue() string {
	return this.initValue
}

func (this *Var) GetComment() string {
	return this.comment
}

func (this *Var) SetTypeName(val string) {
	this.typeName = val
}

func (this *Var) SetName(val string) {
	this.name = val
}

func (this *Var) SetInitValue(val string) {
	this.initValue = val
}

func (this *Var) SetComment(val string) {
	this.comment = val
}

type StructFieldDefine Var

/*func (this *StructFieldDefine) Accept(visitor CodeVisitor) {
	visitor.VisitStructFieldDefine(this)
}*/

type VarList struct {
	vars []*Var
}

func NewVarList() *VarList {
	return &VarList{
		vars: make([]*Var, 0),
	}
}

func (this *VarList) AcceptAsFuncReturns(visitor CodeVisitor) {
	if len(this.vars) <= 0 {
		visitor.VisitFuncNoReturn()
		return
	}

	visitor.VisitFuncReturnFirst(this.vars[0])
	for _, v := range this.vars {
		visitor.VisitFuncReturnNonFirst(v)
	}
}

func (this *VarList) AcceptAsStructField(visitor CodeVisitor) {
	if len(this.vars) <= 0 {
		return
	}

	visitor.VisitStructFieldVarListBegin(this)
	for _, v := range this.vars {
		visitor.VisitStructFieldVar(v)
	}
	visitor.VisitStructFieldVarListEnd(this)
}

func (this *VarList) AcceptAsFuncParmList(visitor CodeVisitor) {
	if len(this.vars) <= 0 {
		return
	}
	visitor.VisitFuncParamVarFirst(this.vars[0])

	if len(this.vars) == 1 {
		return
	}

	visitor.VisitFuncParamVarNonFirstBegin()
	for i := 1; i < len(this.vars); i++ {
		visitor.VisitFuncParamVarNonFirst(this.vars[i])
	}
	visitor.VisitFuncParamVarNonFirstEnd()
}

func (this *VarList) AcceptAsMacroParmList(visitor CodeVisitor) {
	if len(this.vars) <= 0 {
		return
	}
	visitor.VisitMacroParamVarFirst(this.vars[0])

	if len(this.vars) == 1 {
		return
	}

	visitor.VisitMacroParamVarNonFirstBegin()
	for i := 1; i < len(this.vars); i++ {
		visitor.VisitMacroParamVarNonFirst(this.vars[i])
	}
	visitor.VisitMacroParamVarNonFirstEnd()
}
func (this *VarList) Append(val ...*Var) *VarList {
	//fmt.Println("this =", this)
	//fmt.Println("this.Vars =", this.Vars)
	//fmt.Println("val =", val)
	this.vars = append(this.vars, val...)
	return this
}

func (this *VarList) GetMaxNameLen() int {
	max := 0
	for _, v := range this.vars {
		if len(v.name) > max {
			max = len(v.name)
		}
	}
	return max
}

func (this *VarList) GetMaxTypeNameLen() int {
	max := 0
	for _, v := range this.vars {
		if len(v.typeName) > max {
			max = len(v.typeName)
		}
	}
	return max
}

func (this *VarList) GetMaxValueLen() int {
	max := 0
	for _, v := range this.vars {
		if len(v.initValue) > max {
			max = len(v.initValue)
		}
	}
	return max
}
