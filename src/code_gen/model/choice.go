package model

import (
	_ "fmt"
)

const (
	CHOICE_TYPE_NORAML       = 0
	CHOICE_TYPE_MACRO_IF     = 1
	CHOICE_TYPE_MACRO_IFDEF  = 2
	CHOICE_TYPE_MACRO_IFNDEF = 3
)

type Choice struct {
	choiceType int
	comment   string
	condition string
	code      Code
}

func NewChoice() *Choice {
	return &Choice{}
}

func (this *Choice) SetChoiceType(choiceType int) {
	this.choiceType = choiceType
}

func (this *Choice) GetChoiceType() int {
	return this.choiceType
}

func (this *Choice) AcceptAsFirst(visitor CodeVisitor) {
	visitor.VisitChoiceFirstBegin(this)
	this.code.Accept(visitor)
	visitor.VisitChoiceFirstEnd(this)
}

func (this *Choice) AcceptAsFirstMacro(visitor CodeVisitor) {
	visitor.VisitMacroChoiceFirstBegin(this)
	this.code.Accept(visitor)
	visitor.VisitMacroChoiceFirstEnd(this)
}

func (this *Choice) AcceptAsNonFirst(visitor CodeVisitor) {
	visitor.VisitChoiceNonFirstBegin(this)
	this.code.Accept(visitor)
	visitor.VisitChoiceNonFirstEnd(this)
}

func (this *Choice) AcceptAsNonFirstMacro(visitor CodeVisitor) {
	visitor.VisitMacroChoiceNonFirstBegin(this)
	this.code.Accept(visitor)
	visitor.VisitMacroChoiceNonFirstEnd(this)
}

func (this *Choice) AcceptAsChoiceGropuItem(visitor CodeVisitor) {
	visitor.VisitChoiceGroupItemBegin(this)
	this.code.Accept(visitor)
	visitor.VisitChoiceGroupItemEnd(this)
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

const (
	MULTI_CHOICE_TYPE_NORAML = 0
	MULTI_CHOCIE_TYPE_MACRO  = 1
)

type MultiChoice struct {
	choiceType int
	comment  string
	choices  []*Choice
	lastCode Code
	endComment string
}

func NewMultiChoice() *MultiChoice {
	return &MultiChoice{
		choices: make([]*Choice, 0),
	}
}

func (this *MultiChoice) SetChoiceType(choiceType int) {
	this.choiceType = choiceType
}

func (this *MultiChoice) GetChoiceType() int {
	return this.choiceType
}

func (this *MultiChoice) SetComment(comment string) {
	this.comment = comment
}

func (this *MultiChoice) GetComment() string {
	return this.comment
}

func (this *MultiChoice) SetEndComment(comment string) {
	this.endComment = comment
}

func (this *MultiChoice) GetEndComment() string {
	return this.endComment
}

func (this *MultiChoice) AppendChoice(choice ...*Choice) *MultiChoice {
	this.choices = append(this.choices, choice...)
	return this
}

func (this *MultiChoice) SetLastCode(code Code) {
	this.lastCode = code
}

func (this *MultiChoice) ChoiceLen() int {
	return len(this.choices)
}

func (this *MultiChoice) Accept(visitor CodeVisitor) {
	if len(this.choices) <= 0 {
		return
	}

	if this.choiceType == MULTI_CHOICE_TYPE_NORAML {
		this.AcceptNormal(visitor)
	} else {
		this.AcceptAsMacro(visitor)
	}
}

func (this *MultiChoice) AcceptNormal(visitor CodeVisitor) {
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

func (this *MultiChoice) AcceptAsMacro(visitor CodeVisitor) {
	if len(this.choices) <= 0 {
		return
	}

	visitor.VisitMacroMultiChoiceBegin(this)

	this.choices[0].AcceptAsFirstMacro(visitor)

	for i := 1; i < len(this.choices); i++ {
		this.choices[i].AcceptAsNonFirstMacro(visitor)
	}

	visitor.VisitMacroMultiChoiceLastCode(this.lastCode)

	visitor.VisitMacroMultiChoiceEnd(this)

}

type ChoiceGroup struct {
	comment     string
	condition   string
	hasDefault  bool
	choices     []*Choice
	defaultCode Code
}

func NewChoiceGroup() *ChoiceGroup {
	return &ChoiceGroup{
		choices: make([]*Choice, 0),
	}
}

func (this *ChoiceGroup) SetComment(comment string) {
	this.comment = comment
}

func (this *ChoiceGroup) GetComment() string {
	return this.comment
}

func (this *ChoiceGroup) SetCondition(condition string) {
	this.condition = condition
}

func (this *ChoiceGroup) GetCondition() string {
	return this.condition
}

func (this *ChoiceGroup) AppendChoice(choice ...*Choice) *ChoiceGroup {
	this.choices = append(this.choices, choice...)
	return this
}

func (this *ChoiceGroup) SetDefaultCode(code Code) {
	this.defaultCode = code
	this.hasDefault = true
}

func (this *ChoiceGroup) Accept(visitor CodeVisitor) {
	if len(this.choices) <= 0 {
		return
	}

	visitor.VisitChoiceGroupBegin(this)

	for i := 0; i < len(this.choices); i++ {
		this.choices[i].AcceptAsChoiceGropuItem(visitor)
	}

	if this.hasDefault {
		visitor.VisitChoiceGroupDefaultBegin(this.defaultCode)
		this.defaultCode.Accept(visitor)
		visitor.VisitChoiceGroupDefaultEnd(this.defaultCode)
	}

	visitor.VisitChoiceGroupEnd(this)

}
