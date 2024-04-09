package parser

import (
	"fmt"

	"github.com/dils2k/orpc/ast"
	"github.com/dils2k/orpc/scan"
	"github.com/dils2k/orpc/token"
)

type Parser struct {
	lexer  *scan.Lexer
	schema []ast.Statement
}

func NewParser(lexer *scan.Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) ParseSchema() []ast.Statement {
	for tok := p.lexer.NextToken(); tok.Type != token.EOF; {
		switch tok.Type {
		case token.MESSAGE:
			p.message()
		case token.ILLEGAL:
			fmt.Printf("illegal identifier %s\n", tok.Literal)
		}
	}
	return nil
}

func (p *Parser) message() {
}
