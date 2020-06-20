package model

type ConstList struct {
	name   string
	consts *VarList
}

func NewConstList(name string) *ConstList {
	return &ConstList{name: name, consts: NewVarList()}
}

func (this *ConstList) GetName() string {
	return this.name
}

func (this *ConstList) SetName(val string) {
	this.name = val
}

func (this *ConstList) Accept(visitor CodeVisitor) {
	num := len(this.consts.vars)
	if num == 0 {
		return
	}

	if num == 1 {
		visitor.VisitConst(this.consts.vars[0])
		return
	}

	visitor.VisitConstsBegin(this)
	for _, v := range this.consts.vars {
		visitor.VisitConst(v)
	}
	visitor.VisitConstsEnd(this)
}

func (this *ConstList) AppendConst(Const ...*Var) *ConstList {
	this.consts.Append(Const...)
	return this
}

func (this *ConstList) GetMaxNameLen() int {
	return this.consts.GetMaxNameLen()
}

func (this *ConstList) GetMaxTypeNameLen() int {
	return this.consts.GetMaxTypeNameLen()
}

func (this *ConstList) GetMaxValueLen() int {
	return this.consts.GetMaxValueLen()
}

type ConstGroup struct {
	name   string
	consts []*ConstList
}

func NewConstGroup() *ConstGroup {
	return &ConstGroup{}
}

func (this *ConstGroup) GetName() string {
	return this.name
}

func (this *ConstGroup) SetName(val string) {
	this.name = val
}

func (this *ConstGroup) AppendConstList(val ...*ConstList) *ConstGroup {
	this.consts = append(this.consts, val...)
	return this
}
