package parser

import (
	"github.com/sayuyere/yal/pkg/ast"
	"github.com/sayuyere/yal/pkg/lexer"
)

// Parser represents a parser for YAL source code.
type Parser struct {
	l       *lexer.Lexer
	curTok  lexer.Token
	nextTok lexer.Token
}

// New creates a new Parser instance.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.nextTok
	p.nextTok = p.l.NextToken()
}

// ParseProgram parses the input and returns the root AST node.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}
	for p.curTok.Type != string(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case string(lexer.VAR):
		return p.parseVarStatement()
	case string(lexer.FN):
		return p.parseFunctionStatement()
	case string(lexer.FOR):
		return p.parseForStatement()
	default:
		return nil // Only var, fn, for statements for now
	}
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curTok}
	p.nextToken()
	stmt.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	p.nextToken()
	if p.curTok.Type != string(lexer.ASSIGN) {
		return nil
	}
	p.nextToken()
	stmt.Value = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal} // placeholder for expression
	return stmt
}

func (p *Parser) parseFunctionStatement() ast.Statement {
	fs := &ast.FunctionStatement{Token: p.curTok}
	p.nextToken() // fn
	fs.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	p.nextToken() // name
	if p.curTok.Type != string(lexer.LPAREN) {
		return nil
	}
	fs.Parameters = p.parseFunctionParameters()
	if p.curTok.Type != string(lexer.LBRACE) {
		return nil
	}
	fs.Body = p.parseBlockStatement()
	return fs
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	params := []*ast.Identifier{}
	p.nextToken() // skip '('
	if p.curTok.Type == string(lexer.RPAREN) {
		p.nextToken()
		return params
	}
	param := &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	params = append(params, param)
	p.nextToken()
	for p.curTok.Type == string(lexer.COMMA) {
		p.nextToken()
		param := &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
		params = append(params, param)
		p.nextToken()
	}
	if p.curTok.Type == string(lexer.RPAREN) {
		p.nextToken()
	}
	return params
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curTok, Statements: []ast.Statement{}}
	p.nextToken() // skip '{'
	for p.curTok.Type != string(lexer.RBRACE) && p.curTok.Type != string(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseForStatement() ast.Statement {
	fs := &ast.ForStatement{Token: p.curTok}
	p.nextToken() // for
	fs.Init = p.parseVarStatement()
	if p.curTok.Type != string(lexer.SEMICOLON) {
		return nil
	}
	p.nextToken()                                                            // ;
	fs.Condition = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal} // placeholder
	p.nextToken()
	if p.curTok.Type != string(lexer.SEMICOLON) {
		return nil
	}
	p.nextToken()                   // ;
	fs.Post = p.parseVarStatement() // placeholder, should be expression/statement
	if p.curTok.Type != string(lexer.LBRACE) {
		return nil
	}
	fs.Body = p.parseBlockStatement()
	return fs
}
