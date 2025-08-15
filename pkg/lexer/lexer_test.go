package lexer

import "testing"

func TestNextToken(t *testing.T) {
	input := `      // Comment Example

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
    var e = println(c,d);
}`

	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{"COMMENT", "// Comment Example"},
		{"VAR", "var"},
		{"IDENT", "a"},
		{"ASSIGN", "="},
		{"INT", "2"},
		{"SEMICOLON", ";"},
		{"VAR", "var"},
		{"IDENT", "b"},
		{"ASSIGN", "="},
		{"INT", "3"},
		{"SEMICOLON", ";"},
		{"VAR", "var"},
		{"IDENT", "c"},
		{"ASSIGN", "="},
		{"IDENT", "a"},
		{"PLUS", "+"},
		{"IDENT", "b"},
		{"SEMICOLON", ";"},
		{"FN", "fn"},
		{"IDENT", "add"},
		{"LPAREN", "("},
		{"IDENT", "v1"},
		{"COMMA", ","},
		{"IDENT", "v2"},
		{"RPAREN", ")"},
		{"LBRACE", "{"},
		{"RETURN", "return"},
		{"IDENT", "v1"},
		{"PLUS", "+"},
		{"IDENT", "v2"},
		{"RBRACE", "}"},
		{"VAR", "var"},
		{"IDENT", "ma"},
		{"ASSIGN", "="},
		{"STRING", "SAMPLE"},
		{"SEMICOLON", ";"},
		{"VAR", "var"},
		{"IDENT", "mb"},
		{"ASSIGN", "="},
		{"STRING", "2SAMPLE"},
		{"SEMICOLON", ";"},
		{"VAR", "var"},
		{"IDENT", "c1"},
		{"ASSIGN", "="},
		{"CHAR", "a"},
		{"SEMICOLON", ";"},
		{"VAR", "var"},
		{"IDENT", "c2"},
		{"ASSIGN", "="},
		{"CHAR", "b"},
		{"SEMICOLON", ";"},
		{"VAR", "var"},
		{"IDENT", "c3"},
		{"ASSIGN", "="},
		{"IDENT", "c1"},
		{"PLUS", "+"},
		{"IDENT", "c2"},
		{"SEMICOLON", ";"},
		{"FOR", "for"},
		{"VAR", "var"},
		{"IDENT", "d"},
		{"ASSIGN", "="},
		{"INT", "0"},
		{"SEMICOLON", ";"},
		{"IDENT", "d"},
		{"LE", "<="},
		{"INT", "100"},
		{"SEMICOLON", ";"},
		{"IDENT", "d"},
		{"PLUS", "+"},
		{"PLUS", "+"},
		{"LBRACE", "{"},
		{"VAR", "var"},
		{"IDENT", "e"},
		{"ASSIGN", "="},
		{"PRINTLN", "println"},
		{"LPAREN", "("},
		{"IDENT", "c"},
		{"COMMA", ","},
		{"IDENT", "d"},
		{"RPAREN", ")"},
		{"SEMICOLON", ";"},
		{"RBRACE", "}"},
		{"EOF", ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
