package parser

import (
	"testing"

	"github.com/yourusername/yal/pkg/lexer"
)

func TestParser_ParseProgram(t *testing.T) {
	l := lexer.New("print('Hello')")
	p := New(l)
	program := p.ParseProgram()
	if program != nil {
		t.Logf("Parsed program: %v", program)
	}
}
