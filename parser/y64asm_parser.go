// Code generated from Y64asm.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Y64asm

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 98, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 7, 2, 24, 10, 2,
	12, 2, 14, 2, 27, 11, 2, 3, 3, 3, 3, 5, 3, 31, 10, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 38, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 44, 10, 5, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 58, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 91, 10, 10, 3, 11, 3, 11, 3, 11, 5, 11, 96, 10, 11, 3,
	11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 3, 4, 2, 17, 18, 20,
	20, 2, 103, 2, 25, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 37, 3, 2, 2, 2, 8,
	43, 3, 2, 2, 2, 10, 45, 3, 2, 2, 2, 12, 48, 3, 2, 2, 2, 14, 51, 3, 2, 2,
	2, 16, 57, 3, 2, 2, 2, 18, 90, 3, 2, 2, 2, 20, 95, 3, 2, 2, 2, 22, 24,
	5, 4, 3, 2, 23, 22, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2,
	25, 26, 3, 2, 2, 2, 26, 3, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 30, 5, 8,
	5, 2, 29, 31, 7, 23, 2, 2, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31,
	5, 3, 2, 2, 2, 32, 38, 7, 21, 2, 2, 33, 34, 7, 3, 2, 2, 34, 35, 7, 10,
	2, 2, 35, 38, 7, 4, 2, 2, 36, 38, 7, 22, 2, 2, 37, 32, 3, 2, 2, 2, 37,
	33, 3, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 7, 3, 2, 2, 2, 39, 44, 5, 12, 7,
	2, 40, 44, 5, 20, 11, 2, 41, 44, 5, 10, 6, 2, 42, 44, 7, 23, 2, 2, 43,
	39, 3, 2, 2, 2, 43, 40, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 42, 3, 2, 2,
	2, 44, 9, 3, 2, 2, 2, 45, 46, 7, 7, 2, 2, 46, 47, 7, 21, 2, 2, 47, 11,
	3, 2, 2, 2, 48, 49, 7, 22, 2, 2, 49, 50, 7, 5, 2, 2, 50, 13, 3, 2, 2, 2,
	51, 52, 9, 2, 2, 2, 52, 15, 3, 2, 2, 2, 53, 54, 7, 16, 2, 2, 54, 58, 5,
	6, 4, 2, 55, 56, 7, 19, 2, 2, 56, 58, 5, 6, 4, 2, 57, 53, 3, 2, 2, 2, 57,
	55, 3, 2, 2, 2, 58, 17, 3, 2, 2, 2, 59, 60, 7, 8, 2, 2, 60, 61, 7, 10,
	2, 2, 61, 62, 7, 6, 2, 2, 62, 91, 7, 10, 2, 2, 63, 64, 7, 9, 2, 2, 64,
	65, 5, 6, 4, 2, 65, 66, 7, 6, 2, 2, 66, 67, 7, 10, 2, 2, 67, 91, 3, 2,
	2, 2, 68, 69, 7, 12, 2, 2, 69, 70, 5, 6, 4, 2, 70, 71, 7, 6, 2, 2, 71,
	72, 7, 10, 2, 2, 72, 91, 3, 2, 2, 2, 73, 74, 7, 13, 2, 2, 74, 75, 7, 10,
	2, 2, 75, 76, 7, 6, 2, 2, 76, 91, 7, 10, 2, 2, 77, 78, 7, 14, 2, 2, 78,
	79, 5, 6, 4, 2, 79, 80, 7, 6, 2, 2, 80, 81, 7, 10, 2, 2, 81, 91, 3, 2,
	2, 2, 82, 83, 7, 15, 2, 2, 83, 84, 7, 10, 2, 2, 84, 85, 7, 6, 2, 2, 85,
	91, 5, 6, 4, 2, 86, 87, 7, 13, 2, 2, 87, 88, 7, 10, 2, 2, 88, 89, 7, 6,
	2, 2, 89, 91, 7, 10, 2, 2, 90, 59, 3, 2, 2, 2, 90, 63, 3, 2, 2, 2, 90,
	68, 3, 2, 2, 2, 90, 73, 3, 2, 2, 2, 90, 77, 3, 2, 2, 2, 90, 82, 3, 2, 2,
	2, 90, 86, 3, 2, 2, 2, 91, 19, 3, 2, 2, 2, 92, 96, 5, 14, 8, 2, 93, 96,
	5, 16, 9, 2, 94, 96, 5, 18, 10, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2,
	2, 95, 94, 3, 2, 2, 2, 96, 21, 3, 2, 2, 2, 9, 25, 30, 37, 43, 57, 90, 95,
}
var literalNames = []string{
	"", "'('", "')'", "':'", "','", "", "", "", "", "", "", "", "", "", "",
	"'ret'", "'nop'", "'call'", "'halt'",
}
var symbolicNames = []string{
	"", "", "", "", "", "PREPROC", "OP", "IOP", "REG", "MOV", "IRMOV", "RRMOV",
	"MRMOV", "RMMOV", "JMP", "RET", "NOP", "CALL", "HALT", "NUM", "ID", "NL",
	"WS", "LINE_COMMENT",
}

var ruleNames = []string{
	"prog", "line", "refer", "stmt", "preproc", "label", "instr0", "instr1",
	"instr2", "instr",
}

type Y64asmParser struct {
	*antlr.BaseParser
}

// NewY64asmParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Y64asmParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewY64asmParser(input antlr.TokenStream) *Y64asmParser {
	this := new(Y64asmParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Y64asm.g4"

	return this
}

// Y64asmParser tokens.
const (
	Y64asmParserEOF          = antlr.TokenEOF
	Y64asmParserT__0         = 1
	Y64asmParserT__1         = 2
	Y64asmParserT__2         = 3
	Y64asmParserT__3         = 4
	Y64asmParserPREPROC      = 5
	Y64asmParserOP           = 6
	Y64asmParserIOP          = 7
	Y64asmParserREG          = 8
	Y64asmParserMOV          = 9
	Y64asmParserIRMOV        = 10
	Y64asmParserRRMOV        = 11
	Y64asmParserMRMOV        = 12
	Y64asmParserRMMOV        = 13
	Y64asmParserJMP          = 14
	Y64asmParserRET          = 15
	Y64asmParserNOP          = 16
	Y64asmParserCALL         = 17
	Y64asmParserHALT         = 18
	Y64asmParserNUM          = 19
	Y64asmParserID           = 20
	Y64asmParserNL           = 21
	Y64asmParserWS           = 22
	Y64asmParserLINE_COMMENT = 23
)

// Y64asmParser rules.
const (
	Y64asmParserRULE_prog    = 0
	Y64asmParserRULE_line    = 1
	Y64asmParserRULE_refer   = 2
	Y64asmParserRULE_stmt    = 3
	Y64asmParserRULE_preproc = 4
	Y64asmParserRULE_label   = 5
	Y64asmParserRULE_instr0  = 6
	Y64asmParserRULE_instr1  = 7
	Y64asmParserRULE_instr2  = 8
	Y64asmParserRULE_instr   = 9
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ProgContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *Y64asmParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Y64asmParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Y64asmParserPREPROC)|(1<<Y64asmParserOP)|(1<<Y64asmParserIOP)|(1<<Y64asmParserIRMOV)|(1<<Y64asmParserRRMOV)|(1<<Y64asmParserMRMOV)|(1<<Y64asmParserRMMOV)|(1<<Y64asmParserJMP)|(1<<Y64asmParserRET)|(1<<Y64asmParserNOP)|(1<<Y64asmParserCALL)|(1<<Y64asmParserHALT)|(1<<Y64asmParserID)|(1<<Y64asmParserNL))) != 0 {
		{
			p.SetState(20)
			p.Line()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Stmt() IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *LineContext) NL() antlr.TerminalNode {
	return s.GetToken(Y64asmParserNL, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *Y64asmParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Y64asmParserRULE_line)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Stmt()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(27)
			p.Match(Y64asmParserNL)
		}

	}

	return localctx
}

// IReferContext is an interface to support dynamic dispatch.
type IReferContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferContext differentiates from other interfaces.
	IsReferContext()
}

type ReferContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferContext() *ReferContext {
	var p = new(ReferContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_refer
	return p
}

func (*ReferContext) IsReferContext() {}

func NewReferContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferContext {
	var p = new(ReferContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_refer

	return p
}

func (s *ReferContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferContext) NUM() antlr.TerminalNode {
	return s.GetToken(Y64asmParserNUM, 0)
}

func (s *ReferContext) REG() antlr.TerminalNode {
	return s.GetToken(Y64asmParserREG, 0)
}

func (s *ReferContext) ID() antlr.TerminalNode {
	return s.GetToken(Y64asmParserID, 0)
}

func (s *ReferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterRefer(s)
	}
}

func (s *ReferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitRefer(s)
	}
}

func (p *Y64asmParser) Refer() (localctx IReferContext) {
	localctx = NewReferContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Y64asmParserRULE_refer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(35)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Y64asmParserNUM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Match(Y64asmParserNUM)
		}

	case Y64asmParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(Y64asmParserT__0)
		}
		{
			p.SetState(32)
			p.Match(Y64asmParserREG)
		}
		{
			p.SetState(33)
			p.Match(Y64asmParserT__1)
		}

	case Y64asmParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(Y64asmParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *StmtContext) Instr() IInstrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstrContext)
}

func (s *StmtContext) Preproc() IPreprocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreprocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreprocContext)
}

func (s *StmtContext) NL() antlr.TerminalNode {
	return s.GetToken(Y64asmParserNL, 0)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *Y64asmParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Y64asmParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Y64asmParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Label()
		}

	case Y64asmParserOP, Y64asmParserIOP, Y64asmParserIRMOV, Y64asmParserRRMOV, Y64asmParserMRMOV, Y64asmParserRMMOV, Y64asmParserJMP, Y64asmParserRET, Y64asmParserNOP, Y64asmParserCALL, Y64asmParserHALT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Instr()
		}

	case Y64asmParserPREPROC:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Preproc()
		}

	case Y64asmParserNL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.Match(Y64asmParserNL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPreprocContext is an interface to support dynamic dispatch.
type IPreprocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPreprocContext differentiates from other interfaces.
	IsPreprocContext()
}

type PreprocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreprocContext() *PreprocContext {
	var p = new(PreprocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_preproc
	return p
}

func (*PreprocContext) IsPreprocContext() {}

func NewPreprocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreprocContext {
	var p = new(PreprocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_preproc

	return p
}

func (s *PreprocContext) GetParser() antlr.Parser { return s.parser }

func (s *PreprocContext) PREPROC() antlr.TerminalNode {
	return s.GetToken(Y64asmParserPREPROC, 0)
}

func (s *PreprocContext) NUM() antlr.TerminalNode {
	return s.GetToken(Y64asmParserNUM, 0)
}

func (s *PreprocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreprocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterPreproc(s)
	}
}

func (s *PreprocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitPreproc(s)
	}
}

func (p *Y64asmParser) Preproc() (localctx IPreprocContext) {
	localctx = NewPreprocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Y64asmParserRULE_preproc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(Y64asmParserPREPROC)
	}
	{
		p.SetState(44)
		p.Match(Y64asmParserNUM)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) ID() antlr.TerminalNode {
	return s.GetToken(Y64asmParserID, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *Y64asmParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Y64asmParserRULE_label)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(Y64asmParserID)
	}
	{
		p.SetState(47)
		p.Match(Y64asmParserT__2)
	}

	return localctx
}

// IInstr0Context is an interface to support dynamic dispatch.
type IInstr0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstr0Context differentiates from other interfaces.
	IsInstr0Context()
}

type Instr0Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr0Context() *Instr0Context {
	var p = new(Instr0Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_instr0
	return p
}

func (*Instr0Context) IsInstr0Context() {}

func NewInstr0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr0Context {
	var p = new(Instr0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_instr0

	return p
}

func (s *Instr0Context) GetParser() antlr.Parser { return s.parser }

func (s *Instr0Context) RET() antlr.TerminalNode {
	return s.GetToken(Y64asmParserRET, 0)
}

func (s *Instr0Context) HALT() antlr.TerminalNode {
	return s.GetToken(Y64asmParserHALT, 0)
}

func (s *Instr0Context) NOP() antlr.TerminalNode {
	return s.GetToken(Y64asmParserNOP, 0)
}

func (s *Instr0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterInstr0(s)
	}
}

func (s *Instr0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitInstr0(s)
	}
}

func (p *Y64asmParser) Instr0() (localctx IInstr0Context) {
	localctx = NewInstr0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Y64asmParserRULE_instr0)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Y64asmParserRET)|(1<<Y64asmParserNOP)|(1<<Y64asmParserHALT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInstr1Context is an interface to support dynamic dispatch.
type IInstr1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstr1Context differentiates from other interfaces.
	IsInstr1Context()
}

type Instr1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr1Context() *Instr1Context {
	var p = new(Instr1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_instr1
	return p
}

func (*Instr1Context) IsInstr1Context() {}

func NewInstr1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr1Context {
	var p = new(Instr1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_instr1

	return p
}

func (s *Instr1Context) GetParser() antlr.Parser { return s.parser }

func (s *Instr1Context) JMP() antlr.TerminalNode {
	return s.GetToken(Y64asmParserJMP, 0)
}

func (s *Instr1Context) Refer() IReferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferContext)
}

func (s *Instr1Context) CALL() antlr.TerminalNode {
	return s.GetToken(Y64asmParserCALL, 0)
}

func (s *Instr1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterInstr1(s)
	}
}

func (s *Instr1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitInstr1(s)
	}
}

func (p *Y64asmParser) Instr1() (localctx IInstr1Context) {
	localctx = NewInstr1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Y64asmParserRULE_instr1)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Y64asmParserJMP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Match(Y64asmParserJMP)
		}
		{
			p.SetState(52)
			p.Refer()
		}

	case Y64asmParserCALL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(Y64asmParserCALL)
		}
		{
			p.SetState(54)
			p.Refer()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstr2Context is an interface to support dynamic dispatch.
type IInstr2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstr2Context differentiates from other interfaces.
	IsInstr2Context()
}

type Instr2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr2Context() *Instr2Context {
	var p = new(Instr2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_instr2
	return p
}

func (*Instr2Context) IsInstr2Context() {}

func NewInstr2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr2Context {
	var p = new(Instr2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_instr2

	return p
}

func (s *Instr2Context) GetParser() antlr.Parser { return s.parser }

func (s *Instr2Context) OP() antlr.TerminalNode {
	return s.GetToken(Y64asmParserOP, 0)
}

func (s *Instr2Context) AllREG() []antlr.TerminalNode {
	return s.GetTokens(Y64asmParserREG)
}

func (s *Instr2Context) REG(i int) antlr.TerminalNode {
	return s.GetToken(Y64asmParserREG, i)
}

func (s *Instr2Context) IOP() antlr.TerminalNode {
	return s.GetToken(Y64asmParserIOP, 0)
}

func (s *Instr2Context) Refer() IReferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferContext)
}

func (s *Instr2Context) IRMOV() antlr.TerminalNode {
	return s.GetToken(Y64asmParserIRMOV, 0)
}

func (s *Instr2Context) RRMOV() antlr.TerminalNode {
	return s.GetToken(Y64asmParserRRMOV, 0)
}

func (s *Instr2Context) MRMOV() antlr.TerminalNode {
	return s.GetToken(Y64asmParserMRMOV, 0)
}

func (s *Instr2Context) RMMOV() antlr.TerminalNode {
	return s.GetToken(Y64asmParserRMMOV, 0)
}

func (s *Instr2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterInstr2(s)
	}
}

func (s *Instr2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitInstr2(s)
	}
}

func (p *Y64asmParser) Instr2() (localctx IInstr2Context) {
	localctx = NewInstr2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Y64asmParserRULE_instr2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(Y64asmParserOP)
		}
		{
			p.SetState(58)
			p.Match(Y64asmParserREG)
		}
		{
			p.SetState(59)
			p.Match(Y64asmParserT__3)
		}
		{
			p.SetState(60)
			p.Match(Y64asmParserREG)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(Y64asmParserIOP)
		}
		{
			p.SetState(62)
			p.Refer()
		}
		{
			p.SetState(63)
			p.Match(Y64asmParserT__3)
		}
		{
			p.SetState(64)
			p.Match(Y64asmParserREG)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Match(Y64asmParserIRMOV)
		}
		{
			p.SetState(67)
			p.Refer()
		}
		{
			p.SetState(68)
			p.Match(Y64asmParserT__3)
		}
		{
			p.SetState(69)
			p.Match(Y64asmParserREG)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.Match(Y64asmParserRRMOV)
		}
		{
			p.SetState(72)
			p.Match(Y64asmParserREG)
		}
		{
			p.SetState(73)
			p.Match(Y64asmParserT__3)
		}
		{
			p.SetState(74)
			p.Match(Y64asmParserREG)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.Match(Y64asmParserMRMOV)
		}
		{
			p.SetState(76)
			p.Refer()
		}
		{
			p.SetState(77)
			p.Match(Y64asmParserT__3)
		}
		{
			p.SetState(78)
			p.Match(Y64asmParserREG)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(80)
			p.Match(Y64asmParserRMMOV)
		}
		{
			p.SetState(81)
			p.Match(Y64asmParserREG)
		}
		{
			p.SetState(82)
			p.Match(Y64asmParserT__3)
		}
		{
			p.SetState(83)
			p.Refer()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(84)
			p.Match(Y64asmParserRRMOV)
		}
		{
			p.SetState(85)
			p.Match(Y64asmParserREG)
		}
		{
			p.SetState(86)
			p.Match(Y64asmParserT__3)
		}
		{
			p.SetState(87)
			p.Match(Y64asmParserREG)
		}

	}

	return localctx
}

// IInstrContext is an interface to support dynamic dispatch.
type IInstrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstrContext differentiates from other interfaces.
	IsInstrContext()
}

type InstrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstrContext() *InstrContext {
	var p = new(InstrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Y64asmParserRULE_instr
	return p
}

func (*InstrContext) IsInstrContext() {}

func NewInstrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstrContext {
	var p = new(InstrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Y64asmParserRULE_instr

	return p
}

func (s *InstrContext) GetParser() antlr.Parser { return s.parser }

func (s *InstrContext) Instr0() IInstr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr0Context)
}

func (s *InstrContext) Instr1() IInstr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr1Context)
}

func (s *InstrContext) Instr2() IInstr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr2Context)
}

func (s *InstrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.EnterInstr(s)
	}
}

func (s *InstrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Y64asmListener); ok {
		listenerT.ExitInstr(s)
	}
}

func (p *Y64asmParser) Instr() (localctx IInstrContext) {
	localctx = NewInstrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Y64asmParserRULE_instr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Y64asmParserRET, Y64asmParserNOP, Y64asmParserHALT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Instr0()
		}

	case Y64asmParserJMP, Y64asmParserCALL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Instr1()
		}

	case Y64asmParserOP, Y64asmParserIOP, Y64asmParserIRMOV, Y64asmParserRRMOV, Y64asmParserMRMOV, Y64asmParserRMMOV:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Instr2()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
