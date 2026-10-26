package lexer

import (
	"slices"
	"swag/token"
)

type Lexer struct {
	input	string
	currPos int
	nextPos int
	char	byte
}

// Creates new instance of Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Reads next character of an input into l.char for further processing
// Maintains current and next char indices for easier reference
func (l *Lexer) readChar() {
	if (l.nextPos >= len(l.input)) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPos]
	}
	l.currPos = l.nextPos
	l.nextPos += 1
}

// "Eats" all consequent whitespaces, new lines, tabs etc
func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.char) {
		l.readChar()
	}
}

// Returns an instance of a Token parsed from given input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	var seq []byte

	l.skipWhitespace()

	for isValidSequence(seq, l.char) {
		seq = append(seq, l.char)
		l.readChar()
	}

	tok = newToken(string(seq))
	return tok
}

// Creates an instance of a token, used as a helper for NextToken()
func newToken(literal string) token.Token {
	tt := token.LookupID(literal)

	if tt != token.USER {
		tok := token.Token{Type: tt, Literal: literal}
		return tok
	}

	if isDigit(literal[0]) {
		tok := token.Token{Type: token.INT, Literal: literal}
		return tok
	}
	if isLetter(literal[0]) {
		tok := token.Token{Type: token.ID, Literal: literal}
		return tok
	}
	if literal[0] == 0 {
		tok := token.Token{Type: token.EOF, Literal: ""}
		return tok
	}

	tok := token.Token{Type: token.ILLEGAL, Literal: literal}
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// Ensures that current sequence of characters being processed is a valid sequence for 1 token
func isValidSequence(seq []byte, char byte) bool {
	if len(seq) == 0 {
		return true
	}
	if isWhitespace(char) {
		return false
	}
	if isLetter(seq[0]) {
		return isLetter(char)
	}
	if isDigit(seq[0]) {
		return isDigit(char) || (char == '.' && !slices.Contains(seq, '.'))
	}

	separators := []byte{',', ';', '(', ')', '{', '}', }
	if slices.Contains(separators, char) {
		return false
	}

	operators := []byte{'!', '>', '<', '=', '+', }
	if char == '=' && len(seq) == 1 && slices.Contains(operators, seq[0]) {
		return true
	}

	return false
}
