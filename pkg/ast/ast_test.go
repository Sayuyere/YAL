package ast

import (
	"testing"

	"github.com/sayuyere/yal/pkg/lexer"
)

func TestIdentifierNode(t *testing.T) {
	tok := lexer.Token{Type: "IDENT", Literal: "foobar", Line: 1, Column: 5}
	ident := &Identifier{Token: tok, Value: "foobar"}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() wrong. got=%q", ident.TokenLiteral())
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value wrong. got=%q", ident.Value)
	}
}

func TestVarStatementNode(t *testing.T) {
	tok := lexer.Token{Type: "VAR", Literal: "var", Line: 1, Column: 1}
	name := &Identifier{Token: lexer.Token{Type: "IDENT", Literal: "x", Line: 1, Column: 5}, Value: "x"}
	stmt := &VarStatement{Token: tok, Name: name, Value: name}

	if stmt.TokenLiteral() != "var" {
		t.Errorf("stmt.TokenLiteral() wrong. got=%q", stmt.TokenLiteral())
	}
	if stmt.Name.Value != "x" {
		t.Errorf("stmt.Name.Value wrong. got=%q", stmt.Name.Value)
	}
}

// PrintAST has been moved to pkg/ast/print.go for reuse in tests and integration tests.

// Remove TestParseAndPrintMainYAL from ast_test.go to avoid import cycle. Use the integration test in parser package instead.
