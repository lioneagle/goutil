package model

type Block struct {
	children []Code
}

func NewBlock() *Block {
	return &Block{}
}

func (this *Block) Accept(visitor CodeVisitor) {
	visitor.VisitBlockBegin(this)
	for _, v := range this.children {
		v.Accept(visitor)
	}
	visitor.VisitBlockEnd(this)
}

func (this *Block) AppendCode(child Code) {
	this.children = append(this.children, child)
}
