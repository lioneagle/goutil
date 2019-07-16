package model

type Choice struct {
	condition string
	block     *Block
}

type SingleChoice struct {
	condition  string
	blockTrue  *Block
	blockFalse *Block
}

type MultiChoice struct {
	choices    []*Choice
	lastChoice *Choice
}
