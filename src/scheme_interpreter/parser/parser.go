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

func (p *Parser) ParseExpressionStatement() ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	for !p.curTokenIs(token.EOF) {
		switch p.curToken.Type {
		case token.INT:
			stmt.Expression = p.parseIntegerLiteral()
		case token.IDENT:
			stmt.Expression = p.parseIdentifier()
		case token.QUOTE:
			stmt.Expression = &ast.ExpressionStatement{Token: p.curToken}
		default:
			p.nextToken() // thru LPAREN
			stmt.Expression = p.parseExpression()
		}

		p.nextToken()
	}

	return *stmt
}

func (p *Parser) parseExpressionInner() ast.Expression {
	switch p.curToken.Type {
	case token.INT:
		return p.parseIntegerLiteral()
	case token.IDENT:
		return p.parseIdentifier()
	case token.QUOTE:
		return &ast.ExpressionStatement{Token: p.curToken}
	default:
		p.nextToken() // thru LPAREN
		return p.parseExpression()
	}
}

func (p *Parser) parseExpression() ast.Expression {
	// null list ()
	// single (1)
	// double (1 2)
	// other (1 2 3 ...)
	
	se := &ast.SExpression{}

	if p.curToken.Type == token.RPAREN {
		// null list
		return se
	}
	
	se.Car = p.parseExpressionInner()
	
	p.nextToken()
	if !p.curTokenIs(token.RPAREN) {
		se.Cdr = p.parseExpression()
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
