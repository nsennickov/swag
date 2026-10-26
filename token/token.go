package token;

type TokenType string

type Token struct {
	Type 	TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"

	ID		= "ID"
	INT		= "INT"

	// operators
	ASSIGN	= "ASSIGN"
	PLUS	= "PLUS"
	MINUS	= "MINUS"
	MULTI	= "MULTI"
	DIVIDE	= "DIVIDE"
	BANG	= "BANG"
	GTHAN	= "GTHAN"
	LTHAN	= "LTHAN"
	GTHANE	= "GTHANE"
	LTHANE	= "LTHANE"
	EQUAL	= "EQUAL"
	NEQUAL	= "NEQUAL"

	COMMA		= "COMMA"
	SEMICOLON	= "SEMICOLON"

	LPAREN	= "LPAREN" // (
	RPAREN	= "RPAREN" // )
	LBRACE	= "LBRACE" // {
	RBRACE	= "RBRACE" // }

	// keywords
	VAR		= "VAR"
	FUNC	= "FUNC"
	IF		= "IF"
	ELSE	= "ELSE"
	ELSIF	= "ELSIF"
	RETURN	= "RETURN"

	// user
	USER	= "USER"
)

var keywords = map[string]TokenType{
	"var": VAR,
	"func": FUNC,
	"if": IF,
	"else": ELSE,
	"elsif": ELSIF,
	"return": RETURN,
}

var tokensMap = map[string]TokenType{
	// keywords
	"var": VAR,
	"func": FUNC,
	"if": IF,
	"else": ELSE,
	"elsif": ELSIF,
	"return": RETURN,

	// operators
	"=": ASSIGN,
	"+": PLUS,
	"-": MINUS,
	"*": MULTI,
	"/": DIVIDE,
	"!": BANG,
	">": GTHAN,
	"<": LTHAN,
	">=": GTHANE,
	"<=": LTHANE,
	"==": EQUAL,
	"!=": NEQUAL,

	// separators
	",": COMMA,
	";": SEMICOLON,
	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
}

func LookupID(literal string) TokenType {
	if tokType, ok := tokensMap[literal]; ok {
		return tokType
	}

	return USER
}
