package parser

import (
	"testing"

	"github.com/sayuyere/yal/pkg/ast"
	"github.com/sayuyere/yal/pkg/lexer"
)

func TestParser_ParseProgram(t *testing.T) {
	l := lexer.New("print('Hello')")
	p := New(l)
	program := p.ParseProgram()
	if program != nil {
		t.Logf("Parsed program: %v", program)
	}
}

func TestParseVarStatements(t *testing.T) {
	input := `var a = 2;
var b = 3;`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"b"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		varStmt, ok := stmt.(*ast.VarStatement)
		if !ok {
			t.Fatalf("stmt not *ast.VarStatement. got=%T", stmt)
		}
		if varStmt.Name.Value != tt.expectedIdentifier {
			t.Errorf("varStmt.Name.Value not '%s'. got=%s", tt.expectedIdentifier, varStmt.Name.Value)
		}
	}
}
