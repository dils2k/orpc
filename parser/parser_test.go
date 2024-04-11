package parser_test

import (
	"testing"

	"github.com/dils2k/orpc/ast"
	"github.com/dils2k/orpc/parser"
	"github.com/dils2k/orpc/scan"
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
