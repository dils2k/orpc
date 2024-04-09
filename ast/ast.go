package ast

type Statement interface {
	statement()
}

type Message struct {
	Name   string
	Fields []MessageField
}

func (Message) statement() {}

type MessageField struct {
	Name string
	Type string
}
