package evaluator

import (
	"testing"

	"github.com/yourusername/yal/pkg/ast"
)

func TestEval(t *testing.T) {
	var node ast.Node // TODO: create a sample AST node
	result := Eval(node)
	if result != nil {
		t.Logf("Eval result: %v", result)
	}
}
