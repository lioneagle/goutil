package model

type Function struct {
	name    string
	comment string
	returns *VarList
	params  *VarList
	body    *Block
}

func NewFunction() *Function {
	ret := &Function{}
	ret.returns = NewVarList()
	ret.params = NewVarList()
	ret.body = NewBlock()
	return ret
}

func (this *Function) GetName() string {
	return this.name
}

func (this *Function) SetName(name string) {
	this.name = name
}

func (this *Function) GetComment() string {
	return this.comment
}

func (this *Function) SetComment(comment string) {
	this.comment = comment
}

func (this *Function) GetParams() *VarList {
	return this.params
}

func (this *Function) GetReturnList() *VarList {
	return this.returns
}

func (this *Function) AppendCode(code ...Code) *Function {
	this.body.AppendCode(code...)
	return this
}

func (this *Function) GetBody() *Block {
	return this.body
}

func (this *Function) SetBody(body *Block) {
	this.body = body
}

func (this *Function) AcceptAsDeclare(v CodeVisitor) {
	v.VisitFuncDeclare(this)
}

func (this *Function) AcceptAsDefine(v CodeVisitor) {
	v.VisitFuncDefine(this)
}

func (this *Function) AppendParam(val ...*Var) *Function {
	this.params.Append(val...)
	return this
}

func (this *Function) AppendReturnType(val ...*Var) *Function {
	this.returns.Append(val...)
	return this
}

type FunctionList struct {
	Funcs []*Function
}

func NewFunctionList() *FunctionList {
	return &FunctionList{}
}

func (this *FunctionList) Append(val ...*Function) *FunctionList {
	this.Funcs = append(this.Funcs, val...)
	return this
}
