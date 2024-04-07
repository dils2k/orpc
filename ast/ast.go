package ast

type Statement interface {
	statement()
}

type MessageStatement struct {
	Name   string
	Fields []MessageField
}

func (MessageStatement) statement() {}

type MessageField struct {
	Name string
	Type string
}
