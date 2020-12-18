package y64asm

import (
	"bufio"
	"io"
	"unicode"
)

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

// Position represents the current position of the lexer in the input file
type Position struct {
	Row, Col uint
}

// Lexer is the main structure with which lexing is made
type Lexer struct {
	Pos    Position
	reader *bufio.Reader
}

// NewLexer returns a Lexer instance. It receives a reader from which the runes
// to build the Tokens are read.
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		Pos: Position{
			Row: 1,
			Col: 0,
		},
		reader: bufio.NewReader(reader),
	}
}

// Lex returns the next token read from the source.
func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.Pos, EOF, ""
			}
			panic(err)
		}
		l.Pos.Col++
		switch r {
		case '\n':
			l.newLine()
			return l.Pos, NEWL, "\\n"
		case '#':
			for {
				r, _, err := l.reader.ReadRune()
				if err != nil {
					if err == io.EOF {
						return l.Pos, EOF, ""
					}
					panic(err)
				}
				if r == '\n' {
					l.newLine()
					break
				}
			}
		case '.':
			startPos := l.Pos
			lit := l.lexDirec()
			return startPos, DIREC, lit
		case ':':
			return l.Pos, COLON, ":"
		case ',':
			return l.Pos, COMMA, ","
		case '%':
			startPos := l.Pos
			lit := l.lexReg()
			return startPos, REGIST, lit
		case '(':
			return l.Pos, LPAREN, "("
		case ')':
			return l.Pos, RPAREN, ")"
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) || r == '$' {
				startPos := l.Pos
				l.backup()
				lit := l.lexInt()
				return startPos, NUMBER, lit
			} else if unicode.IsLetter(r) {
				startPos := l.Pos
				l.backup()
				lit := l.lexLabel()
				return startPos, IDENT, lit
			} else {
				return l.Pos, ILLEGAL, string(r)
			}
		}
	}
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
	l.Pos.Col--
}

func isHex(r rune) bool {
	return unicode.In(r, unicode.ASCII_Hex_Digit)
}

func (l *Lexer) lexReg() string {
	lit := "%"
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.Pos.Col++
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexDirec() string {
	lit := "."
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.Pos.Col++
		if unicode.IsLetter(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexLabel() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.Pos.Col++
		if unicode.IsLetter(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexInt() string {
	var lit string
	prefix, err := l.reader.Peek(2)
	if err != nil {
		panic(err)
	}
	if prefix[0] == '0' && prefix[1] == 'x' {
		buf := make([]byte, 2)
		n, err := io.ReadFull(l.reader, buf)
		if err != nil || n != 2 {
			panic("Malformed hex literal")
		}
		lit = "0x"
		for {
			r, _, err := l.reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					return lit
				}
			}
			l.Pos.Col++
			if isHex(r) {
				lit = lit + string(r)
			} else {
				l.backup()
				return lit
			}
		}
	} else {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			panic(err)
		}
		l.Pos.Col++
		lit = string(r)
		for {
			r, _, err := l.reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					return lit
				}
			}
			l.Pos.Col++
			if unicode.IsDigit(r) {
				lit = lit + string(r)
			} else {
				l.backup()
				return lit
			}
		}
	}
}

func (l *Lexer) newLine() {
	l.Pos.Row++
	l.Pos.Col = 0
}
