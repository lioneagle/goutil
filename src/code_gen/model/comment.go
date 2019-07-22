package model

type Comment struct {
	comment      string
	isSingleLine bool
}

func NewComment() *Comment {
	return &Comment{}
}

func (this *Comment) SetComment(comment string) {
	this.comment = comment
}

func (this *Comment) GetComment() string {
	return this.comment
}

func (this *Comment) IsSingleLine() bool {
	return this.isSingleLine
}

func (this *Comment) SetSingleLine() {
	this.isSingleLine = true
}

func (this *Comment) SetMultiLine() {
	this.isSingleLine = false
}

func (this *Comment) Accept(visitor CodeVisitor) {
	visitor.VisitComment(this)
}
