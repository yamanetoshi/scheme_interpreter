package parser

import (
	"scheme_interpreter/ast"
	"scheme_interpreter/lexer"
	"testing"
)

func TestSExpression(t *testing.T) {
	input := "(1 2)"

	l := lexer.New(input)
	p := New(l)
	stmt := p.ParseExpressionStatement()
	checkParserErrors(t, p)

	_, ok := stmt.Expression.(*ast.SExpression)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
}

func TestQuoteExpression(t *testing.T) {
	input := "(quote 5)"

	l := lexer.New(input)
	p := New(l)
	stmt := p.ParseExpressionStatement()
	checkParserErrors(t, p)
	
	exp, ok := stmt.Expression.(*ast.SExpression)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	car, ok := exp.Car.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if car.Token.Type != "QUOTE" {
		t.Fatalf("exp not QUOTE. got=%T", car.Token.Type)
	}

	if car.Token.Literal != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			car.Token.Literal)
	}

}

func TestQuoteExpression2(t *testing.T) {
	input := "(quote (1 2))"

	l := lexer.New(input)
	p := New(l)
	stmt := p.ParseExpressionStatement()
	checkParserErrors(t, p)
	exp, ok := stmt.Expression.(*ast.SExpression)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	car, ok := exp.Car.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if car.Token.Type != "QUOTE" {
		t.Fatalf("exp not QUOTE. got=%T", car.Token.Type)
	}

	if car.Token.Literal != "(1 2)" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "(1 2)",
			car.Token.Literal)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5"

	l := lexer.New(input)
	p := New(l)
	stmt := p.ParseExpressionStatement()
	checkParserErrors(t, p)

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar"

	l := lexer.New(input)
	p := New(l)
	stmt := p.ParseExpressionStatement()
	checkParserErrors(t, p)

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
