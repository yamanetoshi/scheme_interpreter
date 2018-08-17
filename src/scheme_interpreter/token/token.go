package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 1343456
	SYMBOL = "SYMBOL" //

	// Operators
//	BANG     = "!"
	QUOTE    = "QUOTE"
	DOT      = "."

	LT = "<"
	GT = ">"

	// Delimiters
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	DEFINE   = "DEFINE"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	LAMBDA   = "LAMBDA"
	CONS     = "CONS"
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"define": DEFINE,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"lambda": LAMBDA,
	"quote":  QUOTE,
	"cons":   CONS,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
