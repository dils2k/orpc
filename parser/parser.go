package parser

import (
	"fmt"

	"github.com/dils2k/orpc/ast"
	"github.com/dils2k/orpc/scan"
	"github.com/dils2k/orpc/token"
)

type Parser struct {
	lexer                *scan.Lexer
	schema               []ast.Statement
	currToken, peekToken *token.Token
	errors               []string
}

func NewParser(lexer *scan.Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	return p
}

func (p *Parser) ParseSchema() []ast.Statement {
	for {
		switch p.currToken.Type {
		case token.MESSAGE:
			p.messageStatement()
		}

		p.nextToken()
		if p.currToken.Type == token.EOF {
			break
		}
	}

	return nil
}

func (p *Parser) nextToken() {
	p.currToken = p.lexer.NextToken()
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) messageStatement() *ast.Message {
	stmt := &ast.Message{}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = p.currToken.Literal
	p.expectPeek(token.LBRACE)
	p.expectPeek(token.RBRACE)
	return stmt
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekTokenIs(typ token.Type) bool {
	return p.peekToken.Type == typ
}
