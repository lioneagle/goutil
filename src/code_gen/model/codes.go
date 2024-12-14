package model

type Codes struct {
	comment string
	codes   []Code
}

func NewCodes() *Codes {
	return &Codes{
		codes: make([]Code, 0),
	}
}

func (this *Codes) SetComment(comment string) {
	this.comment = comment
}

func (this *Codes) GetComment() string {
	return this.comment
}

func (this *Codes) Accept(visitor CodeVisitor) {
	visitor.VisitCodesBegin(this)
	for _, v := range this.codes {
		v.Accept(visitor)
	}
	visitor.VisitCodesEnd(this)
}

func (this *Codes) Append(codes ...Code) *Codes {
	this.codes = append(this.codes, codes...)
	return this
}
