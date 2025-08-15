package lexer

import "testing"

func TestLexer_NextToken(t *testing.T) {
	lexer := New("print('Hello')")
	tok := lexer.NextToken()
	if tok.Type != "EOF" {
		t.Errorf("expected EOF, got %s", tok.Type)
	}
}
