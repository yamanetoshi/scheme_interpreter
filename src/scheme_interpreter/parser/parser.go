package parser

import (
	"fmt"
	"scheme_interpreter/ast"
	"scheme_interpreter/lexer"
	"scheme_interpreter/token"
	"strconv"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.ExpressionStatement{}

	for !p.curTokenIs(token.EOF) {
		//		stmt := p.parseStatement()
		stmt := p.parseExpressionStatement()
		if &stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}

	return program
}

func (p *Parser) parseExpressionStatement() ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	switch p.curToken.Type {
	case token.INT:
		stmt.Expression = p.parseIntegerLiteral()
		p.nextToken()
	case token.IDENT:
		stmt.Expression = p.parseIdentifier()
		p.nextToken()
	case token.QUOTE:
		return *stmt
	default:
		stmt.Expression = p.parseExpression()
	}

	return *stmt
}

func (p *Parser) parseExpression() ast.Expression {
	se := &ast.SExpression{}
	
	p.nextToken()
	se.Car = p.parseExpressionStatement()
	
	p.nextToken()
	if !p.curTokenIs(token.RPAREN) {
		se.Cdr = p.parseExpressionStatement()
	}
	p.nextToken()
	
	return se
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}
