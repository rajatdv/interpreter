package lexer

import (
	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// The purpose of readChar is to give us the next character and advance our position in the input
// x`string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII for NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken returns the next token from the input string.
func (l *Lexer) NextToken() token.Token {
	// Skip any whitespace characters.
	l.skipWhitespace()
	// Initialize an empty token.
	tok := token.Token{}

	// Switch statement to handle different characters.
	switch l.ch {
	case ';', '(', ')', ',', '+', '{', '}', '-', '*', '/':
		// For single character tokens, create a new token with the corresponding type.
		tok = newToken(token.TokenType(l.ch), l.ch)
	case '<', '>':
		// Handle comparison operators (e.g., <, >, <=, >=).
		tok = handleComparison(l)
	case '!':
		// Handle the not equals operator (!=).
		tok = makeTwoCharToken(l, token.BANG, token.NOT_EQ)
	case '=':
		// Handle the equals operator (=, ==).
		tok = makeTwoCharToken(l, token.ASSIGN, token.EQ)
	case 0:
		// End of file, set token type to EOF.
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// Handle identifiers and digits.
		tok = handleIdentifierOrDigit(l)
	}

	// Move to the next character in the input string.
	l.readChar()
	return tok
}

// handleComparison handles comparison operators like <, >, <=, >=.
func handleComparison(l *Lexer) token.Token {
	var tok token.Token
	// Check if the next character is '=' to determine the token type.
	if l.peekChar() == '=' {
		ch := l.ch
		l.readChar()
		// Create a token with the corresponding type and literal.
		tok = token.Token{Type: token.TokenType(ch), Literal: string(ch) + string(l.ch)}
	} else {
		// For single character comparison operators, create a new token.
		tok = newToken(token.TokenType(l.ch), l.ch)
	}
	return tok
}

// makeTwoCharToken handles comparison operators like != and ==.
func makeTwoCharToken(l *Lexer, singleType token.TokenType, doubleType token.TokenType) token.Token {
	var tok token.Token
	// Check if the next character is '=' to determine the token type.
	if l.peekChar() == '=' {
		ch := l.ch
		l.readChar()
		// Create a token with the double type (e.g., EQ or NOT_EQ) and corresponding literal.
		tok = token.Token{Type: doubleType, Literal: string(ch) + string(l.ch)}
	} else {
		// If the next character is not '=', create a token with the single type (e.g., BANG or ASSIGN).
		tok = newToken(singleType, l.ch)
	}
	return tok
}

// handleIdentifierOrDigit handles identifiers and digits.
func handleIdentifierOrDigit(l *Lexer) token.Token {
	var tok token.Token
	if isLetter(l.ch) {
		// Read the identifier and lookup its type.
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
	} else if isDigit(l.ch) {
		// Read the digit and set token type to INT.
		tok.Literal = l.readDigit()
		tok.Type = token.INT
	} else {
		// If the character is not a letter or digit, create an ILLEGAL token.
		tok = newToken(token.ILLEGAL, l.ch)
	}
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readDigit() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
