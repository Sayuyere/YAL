package parser

import (
	"fmt"
	"testing"

	"github.com/sayuyere/yal/pkg/ast"
	"github.com/sayuyere/yal/pkg/lexer"
)

func printAST(node ast.Node, indent string) {
	switch n := node.(type) {
	case *ast.Program:
		fmt.Println(indent + "Program")
		for _, stmt := range n.Statements {
			printAST(stmt, indent+"  ")
		}
	case *ast.VarStatement:
		fmt.Printf("%sVarStatement: %s\n", indent, n.Name.Value)
		if n.Value != nil {
			printAST(n.Value, indent+"  ")
		}
	case *ast.Identifier:
		fmt.Printf("%sIdentifier: %s\n", indent, n.Value)
	default:
		fmt.Printf("%sUnknown node type\n", indent)
	}
}

func TestParseAndPrintMainYAL(t *testing.T) {
	input := `// Comment Example

var a = 2;
var b = 3;
var c = a+b;

fn add(v1,v2) {
    return v1+v2
}
   var ma = "SAMPLE";

var mb = "2SAMPLE";

             var c1 = 'a';

var c2 = 'b';

var c3 = c1+c2;

for var d = 0 ; d<=100 ; d++ { 
    var e = add(c,d);
}`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	printAST(program, "")
}
