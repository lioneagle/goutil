package model

type Code interface {
	Accept(visitor CodeVisitor)
}
