package model

type Choice struct {
	comment   string
	condition string
	code      Code
}

func NewChoice() *Choice {
	return &Choice{}
}

func (this *Choice) SetCode(code Code) {
	this.code = code
}

func (this *Choice) SetComment(comment string) {
	this.comment = comment
}

func (this *Choice) GetComment() string {
	return this.comment
}

func (this *Choice) SetCondiition(condition string) {
	this.condition = condition
}

func (this *Choice) GetCondiition() string {
	return this.condition
}

type SingleChoice struct {
	comment   string
	condition string
	codeTrue  Code
	codeFalse Code
}

func NewSingleChoice() *SingleChoice {
	return &SingleChoice{}
}

func (this *SingleChoice) Accept(visitor CodeVisitor) {
	visitor.VisitSingleChoiceBegin(this)
	if this.codeTrue != nil {
		visitor.VisitSingleChoiceTrueBegin(this)
		this.codeTrue.Accept(visitor)
		visitor.VisitSingleChoiceTrueEnd(this)
	}

	if this.codeFalse != nil {
		visitor.VisitSingleChoiceFalseBegin(this)
		this.codeFalse.Accept(visitor)
		visitor.VisitSingleChoiceFalseEnd(this)
	}

	visitor.VisitSingleChoiceEnd(this)
}

func (this *SingleChoice) SetComment(comment string) {
	this.comment = comment
}

func (this *SingleChoice) GetComment() string {
	return this.comment
}

func (this *SingleChoice) SetCondition(condition string) {
	this.condition = condition
}

func (this *SingleChoice) GetCondition() string {
	return this.condition
}

func (this *SingleChoice) SetCodeTrue(code Code) {
	this.codeTrue = code
}

func (this *SingleChoice) SetCodeFalse(code Code) {
	this.codeFalse = code
}

type MultiChoice struct {
	comment string
	choices []*Choice
}

func NewMultiChoice() *MultiChoice {
	return &MultiChoice{}
}

func (this *MultiChoice) SetComment(comment string) {
	this.comment = comment
}

func (this *MultiChoice) GetComment() string {
	return this.comment
}

func (this *MultiChoice) AppendChoice(choice *Choice) {
	this.choices = append(this.choices, choice)
}
