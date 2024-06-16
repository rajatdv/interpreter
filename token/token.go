package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	//Operators
	PLUS     = "+"
	ASSIGN   = "="
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"

	EQ     = "=="
	NOT_EQ = "!="

	LT = "<"
	GT = ">"

	//DELIMETER
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
	TRUE     = "true"
	FALSE    = "false"

	STRING = "STRING"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
