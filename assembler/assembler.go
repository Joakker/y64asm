package assembler

import (
	"fmt"

	"github.com/Joakker/y64asm/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Assembler is the main structure defined by this package. It's the one that
// does the heavy lifting.
type Assembler struct {
	*parser.BaseY64asmListener
	parser *parser.Y64asmParser
	cur    uint64
	labels map[string]uint64
	code   []uint16
}

// New returns a new instance of Assembler
func New(parser *parser.Y64asmParser) *Assembler {
	return &Assembler{
		parser: parser,
		cur:    0,
		labels: map[string]uint64{},
	}
}

// Assemble is the function that will actually kick off the assembly process
func (a *Assembler) Assemble() (code []byte) {
	antlr.ParseTreeWalkerDefault.Walk(a, a.parser.Prog())
	for _, b := range a.code {
		code = append(code, byte(b))
	}
	return
}

// ExitInstr0 updates the current address after an instruction with no args
func (a *Assembler) ExitInstr0(ctx *parser.Instr0Context) { a.cur += 1 }

// ExitInstr1 updates the current address after an instruction with 1 arg
func (a *Assembler) ExitInstr1(ctx *parser.Instr1Context) { a.cur += 2 }

// ExitInstr2 updates the current address after an instruction with 2 args
func (a *Assembler) ExitInstr2(ctx *parser.Instr2Context) { a.cur += 3 }

// EnterInstr0 enters an instruction with 0 arguments (ret, halt, nop)
func (a *Assembler) EnterInstr0(ctx *parser.Instr0Context) {
	if debug {
		fmt.Println(ctx.GetStart())
	}
	switch t := ctx.GetStart().GetTokenType(); t {
	case parser.Y64asmLexerRET:
		a.emitInstr(Ret)
	case parser.Y64asmLexerHALT:
		a.emitInstr(Halt)
	case parser.Y64asmLexerNOP:
		a.emitInstr(Nop)
	}
}

func (a *Assembler) EnterInstr1(ctx *parser.Instr1Context) {
	if debug {
		fmt.Println(ctx.GetStart())
	}
	switch t := ctx.GetStart().GetTokenType(); t {
	case parser.Y64asmLexerJMP:
		s := ctx.GetStart().GetText()
		a.emitInstr(Jmp(s[1:]))
	case parser.Y64asmLexerCALL:
		a.emitInstr(Call)
	}
}

func (a *Assembler) emitInstr(i uint16) {
	a.code = append(a.code, i)
}
