package parser

import (
	"scheme_interpreter/ast"
	"scheme_interpreter/lexer"
	"testing"
)

func TestQuoteExpression(t *testing.T) {
	input := "(quote 5)"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	tmp := program.Statements[0]

	stmt, ok := tmp.Expression.(*ast.SExpression)
	if !ok {
		t.Fatalf("program.Statements[0].Expression is not ast.ExpressionStatement. got=%T",
			tmp)
	}

	if stmt.Car.Token.Type != "QUOTE" {
		t.Fatalf("exp not QUOTE. got=%T", stmt.Car.Token.Type)
	}

	if stmt.Car.Token.Literal != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			stmt.Car.Token.Literal)
	}

}

func TestQuoteExpression2(t *testing.T) {
	input := "(quote (1 2))"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	tmp := program.Statements[0]

	stmt, ok := tmp.Expression.(*ast.SExpression)
	if !ok {
		t.Fatalf("program.Statements[0].Expression is not ast.ExpressionStatement. got=%T",
			tmp)
	}

	if stmt.Car.Token.Type != "QUOTE" {
		t.Fatalf("exp not QUOTE. got=%T", stmt.Car.Token.Type)
	}

	if stmt.Car.Token.Literal != "(1 2)" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "(1 2)",
			stmt.Car.Token.Literal)
	}

}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt := program.Statements[0]
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
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt := program.Statements[0]
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
