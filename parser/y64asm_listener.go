// Code generated from Y64asm.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Y64asm

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Y64asmListener is a complete listener for a parse tree produced by Y64asmParser.
type Y64asmListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterRefer is called when entering the refer production.
	EnterRefer(c *ReferContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterPreproc is called when entering the preproc production.
	EnterPreproc(c *PreprocContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterInstr0 is called when entering the instr0 production.
	EnterInstr0(c *Instr0Context)

	// EnterInstr1 is called when entering the instr1 production.
	EnterInstr1(c *Instr1Context)

	// EnterInstr2 is called when entering the instr2 production.
	EnterInstr2(c *Instr2Context)

	// EnterInstr is called when entering the instr production.
	EnterInstr(c *InstrContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitRefer is called when exiting the refer production.
	ExitRefer(c *ReferContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitPreproc is called when exiting the preproc production.
	ExitPreproc(c *PreprocContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitInstr0 is called when exiting the instr0 production.
	ExitInstr0(c *Instr0Context)

	// ExitInstr1 is called when exiting the instr1 production.
	ExitInstr1(c *Instr1Context)

	// ExitInstr2 is called when exiting the instr2 production.
	ExitInstr2(c *Instr2Context)

	// ExitInstr is called when exiting the instr production.
	ExitInstr(c *InstrContext)
}
