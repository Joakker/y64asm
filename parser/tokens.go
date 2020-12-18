package parser

// Token is the numeric representation of a token, as determined by it's type
type Token int

const (
	// EOF is emmited when the end of file is reached
	EOF = iota
	// ILLEGAL is emmited when a token isn't recognized
	ILLEGAL
	// IDENT represents any identifier, including instructions and labels
	IDENT
	// DIREC represents a directive, of the form '.name'
	DIREC
	// COLON is the literal ':'
	COLON
	// NUMBER is any numeric constant
	NUMBER
	// COMMA is the literal ','
	COMMA
	// NEWL is the literal '\n'
	NEWL

	// REGIST is a register of the form '%rXX'
	REGIST
	// LPAREN is the literal '('
	LPAREN
	// RPAREN is the literal ')'
	RPAREN
)

var tokenString = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	DIREC:   "DIREC",
	COLON:   "COLON",
	NUMBER:  "NUMBER",
	COMMA:   "COMMA",
	NEWL:    "NEWL",
	REGIST:  "REGIST",
	LPAREN:  "LPAREN",
	RPAREN:  "RPAREN",
}

// String is the string representation of a Token
func (t Token) String() string {
	return tokenString[t]
}
