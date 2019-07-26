package model

import (
	_ "fmt"
)

type Choice struct {
	comment   string
	condition string
	code      Code
}

func NewChoice() *Choice {
	return &Choice{}
}

func (this *Choice) AcceptAsFirst(visitor CodeVisitor) {
	visitor.VisitChoiceFirstBegin(this)
	this.code.Accept(visitor)
	visitor.VisitChoiceFirstEnd(this)
}

func (this *Choice) AcceptAsNonFirst(visitor CodeVisitor) {
	visitor.VisitChoiceNonFirstBegin(this)
	this.code.Accept(visitor)
	visitor.VisitChoiceNonFirstEnd(this)
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

func (this *Choice) SetCondition(condition string) {
	this.condition = condition
}

func (this *Choice) GetCondition() string {
	return this.condition
}

type MultiChoice struct {
	comment  string
	choices  []*Choice
	lastCode Code
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

func (this *MultiChoice) SetLastCode(code Code) {
	this.lastCode = code
}

func (this *MultiChoice) Accept(visitor CodeVisitor) {
	if len(this.choices) <= 0 {
		return
	}

	visitor.VisitMultiChoiceBegin(this)

	this.choices[0].AcceptAsFirst(visitor)

	for i := 1; i < len(this.choices); i++ {
		this.choices[i].AcceptAsNonFirst(visitor)
	}

	visitor.VisitMultiChoiceLastCode(this.lastCode)

	visitor.VisitMultiChoiceEnd(this)

}

type ChoiceGroup struct {
	comment  string
	choices  []*Choice
	lastCode Code
}

func NewChoiceGroup() *ChoiceGroup {
	ret := &ChoiceGroup{}
	ret.choices = make([]*Choice, 0)
	return ret
}

func (this *ChoiceGroup) SetComment(comment string) {
	this.comment = comment
}

func (this *ChoiceGroup) GetComment() string {
	return this.comment
}

func (this *ChoiceGroup) AppendChoice(choice *Choice) {
	this.choices = append(this.choices, choice)
}

func (this *ChoiceGroup) SetLastCode(code Code) {
	this.lastCode = code
}
