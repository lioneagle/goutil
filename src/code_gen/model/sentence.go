package model

type Sentence struct {
	code string
}

func NewSentence(code string) *Sentence {
	return &Sentence{code: code}
}

func (this *Sentence) Accept(visitor CodeVisitor) {
	visitor.VisitSentence(this)
}

func (this *Sentence) GetCode() string {
	return this.code
}

type SentenceList struct {
	codes []*Sentence
}

func NewSentenceList() *SentenceList {
	return &SentenceList{}
}

func (this *SentenceList) Append(val *Sentence) {
	this.codes = append(this.codes, val)
}
