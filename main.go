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

func (l *AsmListener) ExitInstr(c *parser.InstrContext) {
	i := uint(0)
	for _, e := range c.GetChildren() {
		if k, ok := e.(*parser.ReferContext); ok {
			fmt.Println(k.GetText())
			i++
		} else if k, ok := e.(*antlr.TerminalNodeImpl); ok {
			if k.GetText() == "," {
				continue
			}
			fmt.Println(k.GetText())
			i++
		}
	}
	fmt.Println(i)
	l.currAddress += i
}

func main() {
	is := antlr.NewInputStream("init:\nirmovq 0x10, %r10\nirmovq stack, %rsp\nstack:")
	lexer := parser.NewY64asmLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewY64asmParser(stream)

	l := NewAsmListener()

	antlr.ParseTreeWalkerDefault.Walk(l, p.Prog())
	fmt.Println(l.currAddress)
	fmt.Println(l.labels)
}
