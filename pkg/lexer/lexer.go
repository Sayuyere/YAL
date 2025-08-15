package lexer

import (
	"strings"
	"unicode"
)

// Lexer represents a lexical scanner.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number (1-based)
	column       int  // current column number (1-based)
}

// New creates a new Lexer instance.
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespace()
	startLine := l.line
	startColumn := l.column

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: string(EQ), Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: string(ASSIGN), Literal: string(l.ch)}
		}
	case '+':
		tok = Token{Type: string(PLUS), Literal: string(l.ch)}
	case '-':
		tok = Token{Type: string(MINUS), Literal: string(l.ch)}
	case '*':
		tok = Token{Type: string(ASTERISK), Literal: string(l.ch)}
	case '/':
		if l.peekChar() == '/' {
			comment := l.readComment()
			tok = Token{Type: string(COMMENT), Literal: comment}
			return tok
		} else {
			tok = Token{Type: string(SLASH), Literal: string(l.ch)}
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: string(LE), Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: string(LT), Literal: string(l.ch)}
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: string(GE), Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: string(GT), Literal: string(l.ch)}
		}
	case ';':
		tok = Token{Type: string(SEMICOLON), Literal: string(l.ch)}
	case ',':
		tok = Token{Type: string(COMMA), Literal: string(l.ch)}
	case '(':
		tok = Token{Type: string(LPAREN), Literal: string(l.ch)}
	case ')':
		tok = Token{Type: string(RPAREN), Literal: string(l.ch)}
	case '{':
		tok = Token{Type: string(LBRACE), Literal: string(l.ch)}
	case '}':
		tok = Token{Type: string(RBRACE), Literal: string(l.ch)}
	case 0:
		tok.Literal = ""
		tok.Type = string(EOF)
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokType := LookupIdent(literal)
			tok = Token{Type: string(tokType), Literal: literal, Line: startLine, Column: startColumn}
			return tok
		} else if isDigit(l.ch) {
			tok = Token{Type: string(INT), Literal: l.readNumber(), Line: startLine, Column: startColumn}
			return tok
		} else if l.ch == '"' {
			tok = Token{Type: string(STRING), Literal: l.readString(), Line: startLine, Column: startColumn}
			return tok
		} else if l.ch == '\'' {
			tok = Token{Type: string(CHAR), Literal: l.readCharLiteral(), Line: startLine, Column: startColumn}
			return tok
		} else {
			tok = Token{Type: string(ILLEGAL), Literal: string(l.ch), Line: startLine, Column: startColumn}
		}
	}
	l.readChar()
	tok.Line = startLine
	tok.Column = startColumn
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readComment() string {
	position := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return strings.TrimSpace(l.input[position:l.position])
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readString() string {
	l.readChar() // skip opening quote
	position := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	str := l.input[position:l.position]
	l.readChar() // skip closing quote
	return str
}

func (l *Lexer) readCharLiteral() string {
	l.readChar() // skip opening quote
	ch := l.ch
	l.readChar() // move to closing quote
	l.readChar() // skip closing quote
	return string(ch)
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
