package model

const (
	REPEAT_TYPE_FOR = iota
	REPEAT_TYPE_WHILE
	REPEAT_TYPE_DO_WHILE
)

type Repeat struct {
	acceptType int
	comment    string
	condition  string
	code       Code
}

func NewRepeat() *Repeat {
	return &Repeat{}
}

func (this *Repeat) SetAcceptType(acceptType int) {
	this.acceptType = acceptType
}

func (this *Repeat) GetComment() string {
	return this.comment
}

func (this *Repeat) SetComment(comment string) {
	this.comment = comment
}

func (this *Repeat) SetCondition(condition string) {
	this.condition = condition
}

func (this *Repeat) GetCondition() string {
	return this.condition
}

func (this *Repeat) SetCode(code Code) {
	this.code = code
}

func (this *Repeat) Accept(visitor CodeVisitor) {
	accepts := []func(CodeVisitor){
		this.acceptAsFor, this.acceptAsWhile, this.acceptAsDoWhile,
	}
	accepts[this.acceptType](visitor)
}

func (this *Repeat) acceptAsFor(visitor CodeVisitor) {
	visitor.VisitRepeatAsForBegin(this)
	this.code.Accept(visitor)
	visitor.VisitRepeatAsForEnd(this)
}

func (this *Repeat) acceptAsWhile(visitor CodeVisitor) {
	visitor.VisitRepeatAsWhileBegin(this)
	this.code.Accept(visitor)
	visitor.VisitRepeatAsWhileEnd(this)
}

func (this *Repeat) acceptAsDoWhile(visitor CodeVisitor) {
	visitor.VisitRepeatAsDoWhileBegin(this)
	this.code.Accept(visitor)
	visitor.VisitRepeatAsDoWhileEnd(this)
}
