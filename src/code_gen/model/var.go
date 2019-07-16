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
	return &VarList{}
}

func (this *VarList) Append(val *Var) {
	//fmt.Println("this =", this)
	//fmt.Println("this.Vars =", this.Vars)
	//fmt.Println("val =", val)
	this.vars = append(this.vars, val)
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
