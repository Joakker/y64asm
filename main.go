package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Joakker/y64asm/assembler"
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
	case parser.Y64asmLexerIRMOV:
		k, _ := c.GetChild(1).(*parser.ReferContext)
		fmt.Printf("Moving number %s into register %s\n",
			k.GetText(), c.GetStop().GetText(),
		)
	case parser.Y64asmLexerMRMOV:
		k, _ := c.GetChild(1).(*parser.ReferContext)
		fmt.Printf("Moving number at address %s into register %s\n",
			k.GetText(), c.GetStop().GetText(),
		)
	case parser.Y64asmLexerRMMOV:
		k, _ := c.GetChild(3).(*parser.ReferContext)
		fmt.Printf("Moving number from register %s into address %s\n",
			c.GetChild(1), k.GetText(),
		)
	}
}

func (l *AsmListener) EnterInstr1(c *parser.Instr1Context) {
	switch c.GetStart().GetTokenType() {
	case parser.Y64asmLexerCALL:
		fmt.Printf("Calling function: %s\n", c.GetStop().GetText())
	case parser.Y64asmLexerJMP:
		fmt.Printf("Jumping to: %s ", c.GetStop().GetText())
		switch c.GetStart().GetText() {
		case "jmp":
			fmt.Print("unconditionally")
		case "jne":
			fmt.Print("if result wasn't 0")
		case "jeq":
			fmt.Print("if result was 0")
		}
		fmt.Println()
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

func init() {
	version := flag.Bool("V", false, "Show version information")
	flag.Parse()
	if *version {
		fmt.Println("v0.0.0")
		os.Exit(0)
	} else if flag.NArg() == 0 {
		os.Exit(1)
	}
}

func main() {
	is, err := antlr.NewFileStream(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	lexer := parser.NewY64asmLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewY64asmParser(stream)
	a := assembler.New(p)

	fmt.Println(a.Assemble())
}
