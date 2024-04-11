package parser

import (
	"fmt"

	"github.com/dils2k/orpc/ast"
	"github.com/dils2k/orpc/scan"
	"github.com/dils2k/orpc/token"
)

type Parser struct {
	lexer                *scan.Lexer
	currToken, peekToken *token.Token
	errors               []string
}

func NewParser(lexer *scan.Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseSchema() []ast.Statement {
	schema := make([]ast.Statement, 0)
	for {
		switch p.currToken.Type {
		case token.MESSAGE:
			schema = append(schema, p.messageStatement())
		}
		p.nextToken()
		if p.currToken.Type == token.EOF {
			break
		}
	}
	return schema
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) messageStatement() *ast.Message {
	stmt := &ast.Message{}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = p.currToken.Literal
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// while next token is identifier we try to register
	// a message field.
	for p.peekToken.Type == token.IDENT {
		p.nextToken()
		field := p.field()
		if field != nil {
			stmt.Fields = append(stmt.Fields, field)
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return stmt
}

func (p *Parser) field() *ast.MessageField {
	field := &ast.MessageField{Name: p.currToken.Literal}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	field.Type = p.currToken
	return field
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
