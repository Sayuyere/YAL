package lexer

// Token represents a lexical token.
type Token struct {
	Type    string
	Literal string
}

// Lexer represents a lexical scanner.
type Lexer struct {
	input string
}

// New creates a new Lexer instance.
func New(input string) *Lexer {
	return &Lexer{input: input}
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	// TODO: Implement tokenization
	return Token{Type: "EOF", Literal: ""}
}
