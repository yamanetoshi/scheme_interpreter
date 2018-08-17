package lexer

import (
	"testing"

	"scheme_interpreter/token"
)

func TestNextToken(t *testing.T) {
        input := `5
x
(quote (1 . ()))
(quote xxx)
(cons 1 ())
(define x 5)
(define x (lambda (x) (+ 1 x)))
(define add4 (let ((x 4)) (lambda (y) (+ x y))))
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
	        {token.INT, "5"},
		{token.IDENT, "x"},
		{token.LPAREN, "("},
		{token.QUOTE, "quote"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.DOT, "."},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.QUOTE, "quote"},
		{token.IDENT, "xxx"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.CONS, "cons"},
		{token.INT, "1"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.DEFINE, "define"},
		{token.IDENT, "x"},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.DEFINE, "define"},
		{token.IDENT, "x"},
		{token.LPAREN, "("},
		{token.LAMBDA, "lambda"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.IDENT, "+"},
		{token.INT, "1"},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.DEFINE, "define"},
		{token.IDENT, "add4"},
		{token.LPAREN, "("},
		{token.LET, "let"},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.INT, "4"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.LAMBDA, "lambda"},
		{token.LPAREN, "("},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.IDENT, "+"},
		{token.IDENT, "x"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
