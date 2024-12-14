package model

type Block struct {
	codes *Codes
}

func NewBlock() *Block {
	return &Block{
		codes: NewCodes(),
	}
}

func (this *Block) Accept(visitor CodeVisitor) {
	visitor.VisitBlockBegin(this)
	this.codes.Accept(visitor)
	visitor.VisitBlockEnd(this)
}

func (this *Block) GetCodes() *Codes {
	return this.codes
}

func (this *Block) SetCodes(codes *Codes) {
	this.codes = codes
}

func (this *Block) AppendCode(code Code) {
	this.codes.Append(code)
}
