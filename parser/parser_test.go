package parser_test

import (
	"testing"

	"github.com/dils2k/orpc/ast"
	"github.com/dils2k/orpc/parser"
	"github.com/dils2k/orpc/scan"
	"github.com/dils2k/orpc/token"
	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected []ast.Statement
	}{
		{
			"valid message",
			`message Request123 {}`,
			[]ast.Statement{
				&ast.Message{Name: "Request123"},
			},
		},
		{
			"valid message with fields",
			`
			message Request {
				a int32
				b int64
				c string
			}
			`,
			[]ast.Statement{
				&ast.Message{
					Name: "Request",
					Fields: []*ast.MessageField{
						{Name: "a", Type: &token.Token{Type: token.IDENT, Literal: "int32"}},
						{Name: "b", Type: &token.Token{Type: token.IDENT, Literal: "int64"}},
						{Name: "c", Type: &token.Token{Type: token.IDENT, Literal: "string"}},
					},
				},
			},
		},
		{
			"two valid messages",
			`
			message Request {}
			message Request2 {}
			`,
			[]ast.Statement{
				&ast.Message{Name: "Request"},
				&ast.Message{Name: "Request2"},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lexer := scan.NewLexer(tt.input)
			parser := parser.NewParser(lexer)
			if errs := parser.Errors(); errs != nil {
				t.Fatalf("parsing finished with errors::\n%s", errs)
			}
			schema := parser.ParseSchema()
			if diff := cmp.Diff(tt.expected, schema); diff != "" {
				t.Errorf("ParserSchema() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
