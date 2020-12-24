// Code generated from Y64asm.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Y64asm

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseY64asmListener is a complete listener for a parse tree produced by Y64asmParser.
type BaseY64asmListener struct{}

var _ Y64asmListener = &BaseY64asmListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseY64asmListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseY64asmListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseY64asmListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseY64asmListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseY64asmListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseY64asmListener) ExitProg(ctx *ProgContext) {}

// EnterLine is called when production line is entered.
func (s *BaseY64asmListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseY64asmListener) ExitLine(ctx *LineContext) {}

// EnterRefer is called when production refer is entered.
func (s *BaseY64asmListener) EnterRefer(ctx *ReferContext) {}

// ExitRefer is called when production refer is exited.
func (s *BaseY64asmListener) ExitRefer(ctx *ReferContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseY64asmListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseY64asmListener) ExitStmt(ctx *StmtContext) {}

// EnterPreproc is called when production preproc is entered.
func (s *BaseY64asmListener) EnterPreproc(ctx *PreprocContext) {}

// ExitPreproc is called when production preproc is exited.
func (s *BaseY64asmListener) ExitPreproc(ctx *PreprocContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseY64asmListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseY64asmListener) ExitLabel(ctx *LabelContext) {}

// EnterInstr0 is called when production instr0 is entered.
func (s *BaseY64asmListener) EnterInstr0(ctx *Instr0Context) {}

// ExitInstr0 is called when production instr0 is exited.
func (s *BaseY64asmListener) ExitInstr0(ctx *Instr0Context) {}

// EnterInstr1 is called when production instr1 is entered.
func (s *BaseY64asmListener) EnterInstr1(ctx *Instr1Context) {}

// ExitInstr1 is called when production instr1 is exited.
func (s *BaseY64asmListener) ExitInstr1(ctx *Instr1Context) {}

// EnterInstr2 is called when production instr2 is entered.
func (s *BaseY64asmListener) EnterInstr2(ctx *Instr2Context) {}

// ExitInstr2 is called when production instr2 is exited.
func (s *BaseY64asmListener) ExitInstr2(ctx *Instr2Context) {}

// EnterInstr is called when production instr is entered.
func (s *BaseY64asmListener) EnterInstr(ctx *InstrContext) {}

// ExitInstr is called when production instr is exited.
func (s *BaseY64asmListener) ExitInstr(ctx *InstrContext) {}
