package scan_test

import (
	"testing"

	"github.com/dils2k/orpc/scan"
	"github.com/dils2k/orpc/token"
)

func TestNextToken(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			"lbrace rbrace",
			"{}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			"message declaration",
			"message Person {}",
			[]token.Token{
				{Type: token.MESSAGE, Literal: "message"},
				{Type: token.IDENT, Literal: "Person"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			"invalid char",
			"message % {}",
			[]token.Token{
				{Type: token.MESSAGE, Literal: "message"},
				{Type: token.ILLEGAL, Literal: "%"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			l := scan.NewLexer(tt.input)
			for i := 0; ; i++ {
				tok := l.NextToken()
				exp := tt.expected[i]

				if tok.Type != exp.Type {
					t.Errorf("expected type %s got %s", exp.Type, tok.Type)
				}

				if tok.Type != exp.Type {
					t.Errorf("expected literal %s got %s", exp.Literal, tok.Literal)
				}

				if tok.Type == token.EOF {
					break
				}
			}
		})
	}
}
