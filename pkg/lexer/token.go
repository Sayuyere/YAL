package lexer

// TokenType is a string alias for token types.
type TokenType string

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	COMMENT TokenType = "COMMENT"

	// Identifiers + literals
	IDENT  TokenType = "IDENT"  // main, a, b, c
	INT    TokenType = "INT"    // 12345
	STRING TokenType = "STRING" // "SAMPLE"
	CHAR   TokenType = "CHAR"   // 'a'

	// Operators
	ASSIGN   TokenType = "ASSIGN"   // =
	PLUS     TokenType = "PLUS"     // +
	MINUS    TokenType = "MINUS"    // -
	ASTERISK TokenType = "ASTERISK" // *
	SLASH    TokenType = "SLASH"    // /
	LT       TokenType = "LT"       // <
	GT       TokenType = "GT"       // >
	EQ       TokenType = "EQ"       // ==
	NOT_EQ   TokenType = "NOT_EQ"   // !=
	LE       TokenType = "LE"       // <=
	GE       TokenType = "GE"       // >=

	// Delimiters
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE    TokenType = "RBRACE"

	// Keywords
	VAR     TokenType = "VAR"
	FOR     TokenType = "FOR"
	PRINTLN TokenType = "PRINTLN"
	FN      TokenType = "FN"     // fn
	RETURN  TokenType = "RETURN" // return
)

// Keywords lookup
var keywords = map[string]TokenType{
	"var":     VAR,
	"for":     FOR,
	"println": PRINTLN,
	"fn":      FN,
	"return":  RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Token represents a lexical token with location information.
type Token struct {
	Type    string // token type
	Literal string // token value
	Line    int    // line number in source (1-based)
	Column  int    // column number in source (1-based)
}
