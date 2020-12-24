package main

import (
	"fmt"

	"github.com/Joakker/y64asm/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type AsmListener struct {
	*parser.BaseY64asmListener
	currAddress uint
	labels      map[string]uint
}

func NewAsmListener() *AsmListener {
	return &AsmListener{
		currAddress: 0,
		labels:      map[string]uint{},
	}
}

func (l *AsmListener) EnterLabel(c *parser.LabelContext) {
	l.labels[c.ID().GetText()] = l.currAddress
}

func (l *AsmListener) ExitInstr0(c *parser.Instr0Context) {
	l.currAddress++
}

func (l *AsmListener) ExitInstr1(c *parser.Instr1Context) {
	l.currAddress += 2
}

func (l *AsmListener) ExitInstr2(c *parser.Instr2Context) {
	l.currAddress += 3
}

func main() {
	is := antlr.NewInputStream(`
label1:
	irmovq $10, %r10
	irmovq 0x10, %r9
laber2:
`)
	lexer := parser.NewY64asmLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewY64asmParser(stream)

	l := NewAsmListener()

	antlr.ParseTreeWalkerDefault.Walk(l, p.Prog())
	fmt.Println(l.labels)
}
