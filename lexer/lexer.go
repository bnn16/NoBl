package lexer

import (
	"nobl/token"
)

// in order to support full unicode and utf-8
// this needs to be use rune instead of byte
// but this is a future improvement
type Lexer struct {
	input        string
	position     int  // current position, points to char
	readPosition int  // current reading position, after current char
	char         byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// we have reached the end of the input
	// return ASCII NUL ( 0 )
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LEFT_PAREN, l.char)
	case ')':
		tok = newToken(token.RIGHT_PAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '{':
		tok = newToken(token.LEFT_BRACE, l.char)
	case '}':
		tok = newToken(token.RIGHT_BRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}
