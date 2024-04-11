package scan_test

import (
	"testing"

	"github.com/dils2k/orpc/scan"
	"github.com/dils2k/orpc/token"
	"github.com/google/go-cmp/cmp"
)

func TestNextToken(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected []*token.Token
	}{
		{
			"lbrace rbrace",
			"{}",
			[]*token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			"message declaration",
			"message Person {}",
			[]*token.Token{
				{Type: token.MESSAGE, Literal: "message"},
				{Type: token.IDENT, Literal: "Person"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			"message declaration with number",
			"message Person123 {}",
			[]*token.Token{
				{Type: token.MESSAGE, Literal: "message"},
				{Type: token.IDENT, Literal: "Person123"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			"message declaration with fields",
			`message Person {
				a int32
				b int64
				c string
			}`,
			[]*token.Token{
				{Type: token.MESSAGE, Literal: "message"},
				{Type: token.IDENT, Literal: "Person"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.IDENT, Literal: "a"},
				{Type: token.IDENT, Literal: "int32"},
				{Type: token.IDENT, Literal: "b"},
				{Type: token.IDENT, Literal: "int64"},
				{Type: token.IDENT, Literal: "c"},
				{Type: token.IDENT, Literal: "string"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			"invalid char",
			"message % {}",
			[]*token.Token{
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

			toks := make([]*token.Token, 0)
			for {
				tok := l.NextToken()
				toks = append(toks, tok)
				if tok.Type == token.EOF {
					break
				}
			}

			if diff := cmp.Diff(tt.expected, toks); diff != "" {
				t.Fatalf("token mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
