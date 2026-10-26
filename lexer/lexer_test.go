package lexer

import (
	"testing"

	"swag/token"
)

func TestNextToken(t *testing.T) {
	input := `var five = 5;
	var two = 2;

	var sum = func(x, y) {
		return x + y;
	};
	var result = sum(five, two);
	!><*-
	10.888 >= 123 == 322 != 12345
	if else elsif
	`

	tests := []struct {
		expectedType 	token.TokenType
		expectedLiteral string
	}{
		{ token.VAR, "var" },
		{ token.ID, "five" },
		{ token.ASSIGN, "=" },
		{ token.INT, "5" },
		{ token.SEMICOLON, ";" },
		{ token.VAR, "var" },
		{ token.ID, "two" },
		{ token.ASSIGN, "=" },
		{ token.INT, "2" },
		{ token.SEMICOLON, ";" },
		{ token.VAR, "var" },
		{ token.ID, "sum" },
		{ token.ASSIGN, "=" },
		{ token.FUNC, "func" },
		{ token.LPAREN, "(" },
		{ token.ID, "x" },
		{ token.COMMA, "," },
		{ token.ID, "y" },
		{ token.RPAREN, ")" },
		{ token.LBRACE, "{" },
		{ token.RETURN, "return" },
		{ token.ID, "x" },
		{ token.PLUS, "+" },
		{ token.ID, "y" },
		{ token.SEMICOLON, ";" },
		{ token.RBRACE, "}" },
		{ token.SEMICOLON, ";" },
		{ token.VAR, "var" },
		{ token.ID, "result" },
		{ token.ASSIGN, "=" },
		{ token.ID, "sum" },
		{ token.LPAREN, "(" },
		{ token.ID, "five" },
		{ token.COMMA, "," },
		{ token.ID, "two" },
		{ token.RPAREN, ")" },
		{ token.SEMICOLON, ";" },
		{ token.BANG, "!" },
		{ token.GTHAN, ">" },
		{ token.LTHAN, "<" },
		{ token.MULTI, "*" },
		{ token.MINUS, "-" },
		{ token.INT, "10.888" },
		{ token.GTHANE, ">=" },
		{ token.INT, "123" },
		{ token.EQUAL, "==" },
		{ token.INT, "322" },
		{ token.NEQUAL, "!=" },
		{ token.INT, "12345" },
		{ token.IF, "if" },
		{ token.ELSE, "else" },
		{ token.ELSIF, "elsif" },
		{ token.EOF, "" },
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		t.Logf(`Current Token: "%s": %s`, tok.Type, tok.Literal)

		if tok.Type != tt.expectedType {
			t.Fatalf(`Unexpected token. Expected: "%s", Got: "%s"`, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(`Unexpected literal. Expected: "%s", Got: "%s"`, tt.expectedLiteral, tok.Literal)
		}
	}
}
