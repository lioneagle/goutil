package model

type Function struct {
	name        string
	returnTypes *TypeList
	params      *VarList
	comment     string
	body        Code
}

func NewFunction() *Function {
	return &Function{}
}

func (this *Function) Accept(v CodeVisitor) {
	//v.VisitFunction(this)
}

func (this *Function) AppendParam(val *Var) {
	this.params.Append(val)
}

func (this *Function) AppendReturnType(val *Type) {
	this.returnTypes.Append(val)
}

type FunctionList struct {
	Funcs []*Function
}

func NewFunctionList() *FunctionList {
	return &FunctionList{}
}

func (this *FunctionList) Append(val *Function) {
	this.Funcs = append(this.Funcs, val)
}

func (this *FunctionList) GetMaxTypeNameLen() int {
	max := 0
	for _, v := range this.Funcs {
		len1 := v.returnTypes.GetTypeNameLen()
		if len1 > max {
			max = len1
		}
	}
	return max
}
