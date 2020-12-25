package main

import (
	"fmt"

	"github.com/Joakker/y64asm/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type AsmListener struct {
	*parser.BaseY64asmListener
	currAddr uint
	labels   map[string]uint
}

func NewAsmListener() *AsmListener {
	return &AsmListener{
		currAddr: 0,
		labels:   map[string]uint{},
	}
}

func (l *AsmListener) EnterPreproc(c *parser.PreprocContext) {
	fmt.Printf("Preprocessing directive: %s, operand %s\n",
		c.GetStart().GetText(), c.GetStop().GetText())
}

func (l *AsmListener) EnterLabel(c *parser.LabelContext) {
	l.labels[c.ID().GetText()] = l.currAddr
	fmt.Println("Label: ", c.ID().GetText())
}

func (l *AsmListener) EnterInstr2(c *parser.Instr2Context) {
	switch c.GetStart().GetTokenType() {
	case parser.Y64asmLexerOP:
		op := c.GetStart().GetText()
		fmt.Printf("%s between registers: %s\n",
			op, c.GetTokens(parser.Y64asmLexerREG))
	case parser.Y64asmLexerIOP:
		op := c.GetStart().GetText()[1:]
		k, _ := c.GetChild(1).(*parser.ReferContext)
		fmt.Printf("%s number %s into register %s\n",
			op, k.GetText(), c.GetStop().GetText(),
		)
	}
}

func (l *AsmListener) EnterInstr1(c *parser.Instr1Context) {
	switch c.GetStart().GetTokenType() {
	case parser.Y64asmLexerCALL:
		fmt.Printf("Calling function: %s\n", c.GetStop().GetText())
	case parser.Y64asmLexerJMP:
		fmt.Printf("Jumping to: %s\n", c.GetStop().GetText())
	}
}

func (l *AsmListener) EnterInstr0(c *parser.Instr0Context) {
	switch c.GetStart().GetTokenType() {
	case parser.Y64asmLexerRET:
		fmt.Println("Returning")
	case parser.Y64asmLexerHALT:
		fmt.Println("HALT")
	}
}

func (l *AsmListener) ExitInstr0(c *parser.Instr0Context) { l.currAddr++ }

func (l *AsmListener) ExitInstr1(c *parser.Instr1Context) { l.currAddr += 2 }

func (l *AsmListener) ExitInstr2(c *parser.Instr2Context) { l.currAddr += 3 }

func main() {
	is := antlr.NewInputStream(`
.pos 0
	jmp		label1
label1:
	irmovq	$10, %r10
	irmovq	0x10, %r9
	add		%r9, %r10
	call	label2
	halt
label2:
	iadd	0x5, %r10
	ret
`)
	lexer := parser.NewY64asmLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewY64asmParser(stream)

	l := NewAsmListener()

	antlr.ParseTreeWalkerDefault.Walk(l, p.Prog())
	fmt.Println(l.labels)
}
