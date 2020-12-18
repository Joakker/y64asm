package parser

import (
	"bufio"
	"io"
	"unicode"
)

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
