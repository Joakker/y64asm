// Code generated from Y64asm.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 211,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 75, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 89, 10, 7, 3, 7, 5,
	7, 92, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 109, 10, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 5, 10, 116, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 5, 15, 145, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 20, 5, 20, 166, 10, 20, 3, 20, 6, 20, 169, 10, 20, 13, 20, 14, 20,
	170, 3, 20, 3, 20, 3, 20, 3, 20, 6, 20, 177, 10, 20, 13, 20, 14, 20, 178,
	5, 20, 181, 10, 20, 3, 21, 6, 21, 184, 10, 21, 13, 21, 14, 21, 185, 3,
	21, 7, 21, 189, 10, 21, 12, 21, 14, 21, 192, 11, 21, 3, 22, 5, 22, 195,
	10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 7, 24,
	205, 10, 24, 12, 24, 14, 24, 208, 11, 24, 3, 24, 3, 24, 2, 2, 25, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 3, 2, 12, 3, 2, 58, 59, 3, 2, 50, 54, 4, 2, 102,
	102, 117, 117, 4, 2, 100, 100, 117, 117, 3, 2, 99, 102, 3, 2, 50, 59, 5,
	2, 50, 59, 67, 72, 99, 104, 4, 2, 67, 92, 99, 124, 4, 2, 11, 11, 34, 34,
	4, 2, 12, 12, 15, 15, 2, 232, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3,
	2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2, 5, 51, 3, 2, 2, 2, 7, 53,
	3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 57, 3, 2, 2, 2, 13, 88, 3, 2, 2, 2,
	15, 93, 3, 2, 2, 2, 17, 96, 3, 2, 2, 2, 19, 110, 3, 2, 2, 2, 21, 117, 3,
	2, 2, 2, 23, 122, 3, 2, 2, 2, 25, 127, 3, 2, 2, 2, 27, 132, 3, 2, 2, 2,
	29, 137, 3, 2, 2, 2, 31, 146, 3, 2, 2, 2, 33, 150, 3, 2, 2, 2, 35, 154,
	3, 2, 2, 2, 37, 159, 3, 2, 2, 2, 39, 180, 3, 2, 2, 2, 41, 183, 3, 2, 2,
	2, 43, 194, 3, 2, 2, 2, 45, 198, 3, 2, 2, 2, 47, 202, 3, 2, 2, 2, 49, 50,
	7, 42, 2, 2, 50, 4, 3, 2, 2, 2, 51, 52, 7, 43, 2, 2, 52, 6, 3, 2, 2, 2,
	53, 54, 7, 60, 2, 2, 54, 8, 3, 2, 2, 2, 55, 56, 7, 46, 2, 2, 56, 10, 3,
	2, 2, 2, 57, 74, 7, 48, 2, 2, 58, 59, 7, 115, 2, 2, 59, 60, 7, 119, 2,
	2, 60, 61, 7, 99, 2, 2, 61, 75, 7, 102, 2, 2, 62, 63, 7, 110, 2, 2, 63,
	64, 7, 113, 2, 2, 64, 65, 7, 112, 2, 2, 65, 75, 7, 105, 2, 2, 66, 67, 7,
	114, 2, 2, 67, 68, 7, 113, 2, 2, 68, 75, 7, 117, 2, 2, 69, 70, 7, 99, 2,
	2, 70, 71, 7, 110, 2, 2, 71, 72, 7, 107, 2, 2, 72, 73, 7, 105, 2, 2, 73,
	75, 7, 112, 2, 2, 74, 58, 3, 2, 2, 2, 74, 62, 3, 2, 2, 2, 74, 66, 3, 2,
	2, 2, 74, 69, 3, 2, 2, 2, 75, 12, 3, 2, 2, 2, 76, 77, 7, 99, 2, 2, 77,
	78, 7, 102, 2, 2, 78, 89, 7, 102, 2, 2, 79, 80, 7, 99, 2, 2, 80, 81, 7,
	112, 2, 2, 81, 89, 7, 102, 2, 2, 82, 83, 7, 122, 2, 2, 83, 84, 7, 113,
	2, 2, 84, 89, 7, 116, 2, 2, 85, 86, 7, 117, 2, 2, 86, 87, 7, 119, 2, 2,
	87, 89, 7, 100, 2, 2, 88, 76, 3, 2, 2, 2, 88, 79, 3, 2, 2, 2, 88, 82, 3,
	2, 2, 2, 88, 85, 3, 2, 2, 2, 89, 91, 3, 2, 2, 2, 90, 92, 7, 115, 2, 2,
	91, 90, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 14, 3, 2, 2, 2, 93, 94, 7,
	107, 2, 2, 94, 95, 5, 13, 7, 2, 95, 16, 3, 2, 2, 2, 96, 97, 7, 39, 2, 2,
	97, 98, 7, 116, 2, 2, 98, 108, 3, 2, 2, 2, 99, 109, 9, 2, 2, 2, 100, 101,
	7, 51, 2, 2, 101, 109, 9, 3, 2, 2, 102, 103, 9, 4, 2, 2, 103, 109, 7, 107,
	2, 2, 104, 105, 9, 5, 2, 2, 105, 109, 7, 114, 2, 2, 106, 107, 9, 6, 2,
	2, 107, 109, 7, 122, 2, 2, 108, 99, 3, 2, 2, 2, 108, 100, 3, 2, 2, 2, 108,
	102, 3, 2, 2, 2, 108, 104, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 18, 3,
	2, 2, 2, 110, 111, 7, 111, 2, 2, 111, 112, 7, 113, 2, 2, 112, 113, 7, 120,
	2, 2, 113, 115, 3, 2, 2, 2, 114, 116, 7, 115, 2, 2, 115, 114, 3, 2, 2,
	2, 115, 116, 3, 2, 2, 2, 116, 20, 3, 2, 2, 2, 117, 118, 7, 107, 2, 2, 118,
	119, 7, 116, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 5, 19, 10, 2, 121, 22,
	3, 2, 2, 2, 122, 123, 7, 116, 2, 2, 123, 124, 7, 116, 2, 2, 124, 125, 3,
	2, 2, 2, 125, 126, 5, 19, 10, 2, 126, 24, 3, 2, 2, 2, 127, 128, 7, 111,
	2, 2, 128, 129, 7, 116, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 5, 19, 10,
	2, 131, 26, 3, 2, 2, 2, 132, 133, 7, 116, 2, 2, 133, 134, 7, 111, 2, 2,
	134, 135, 3, 2, 2, 2, 135, 136, 5, 19, 10, 2, 136, 28, 3, 2, 2, 2, 137,
	144, 7, 108, 2, 2, 138, 139, 7, 112, 2, 2, 139, 145, 7, 103, 2, 2, 140,
	141, 7, 103, 2, 2, 141, 145, 7, 115, 2, 2, 142, 143, 7, 111, 2, 2, 143,
	145, 7, 114, 2, 2, 144, 138, 3, 2, 2, 2, 144, 140, 3, 2, 2, 2, 144, 142,
	3, 2, 2, 2, 145, 30, 3, 2, 2, 2, 146, 147, 7, 116, 2, 2, 147, 148, 7, 103,
	2, 2, 148, 149, 7, 118, 2, 2, 149, 32, 3, 2, 2, 2, 150, 151, 7, 112, 2,
	2, 151, 152, 7, 113, 2, 2, 152, 153, 7, 114, 2, 2, 153, 34, 3, 2, 2, 2,
	154, 155, 7, 101, 2, 2, 155, 156, 7, 99, 2, 2, 156, 157, 7, 110, 2, 2,
	157, 158, 7, 110, 2, 2, 158, 36, 3, 2, 2, 2, 159, 160, 7, 106, 2, 2, 160,
	161, 7, 99, 2, 2, 161, 162, 7, 110, 2, 2, 162, 163, 7, 118, 2, 2, 163,
	38, 3, 2, 2, 2, 164, 166, 7, 38, 2, 2, 165, 164, 3, 2, 2, 2, 165, 166,
	3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167, 169, 9, 7, 2, 2, 168, 167, 3, 2,
	2, 2, 169, 170, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2,
	171, 181, 3, 2, 2, 2, 172, 173, 7, 50, 2, 2, 173, 174, 7, 122, 2, 2, 174,
	176, 3, 2, 2, 2, 175, 177, 9, 8, 2, 2, 176, 175, 3, 2, 2, 2, 177, 178,
	3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2,
	2, 2, 180, 165, 3, 2, 2, 2, 180, 172, 3, 2, 2, 2, 181, 40, 3, 2, 2, 2,
	182, 184, 9, 9, 2, 2, 183, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185,
	183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 190, 3, 2, 2, 2, 187, 189,
	9, 7, 2, 2, 188, 187, 3, 2, 2, 2, 189, 192, 3, 2, 2, 2, 190, 188, 3, 2,
	2, 2, 190, 191, 3, 2, 2, 2, 191, 42, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2,
	193, 195, 7, 15, 2, 2, 194, 193, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	196, 3, 2, 2, 2, 196, 197, 7, 12, 2, 2, 197, 44, 3, 2, 2, 2, 198, 199,
	9, 10, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 8, 23, 2, 2, 201, 46, 3, 2,
	2, 2, 202, 206, 7, 37, 2, 2, 203, 205, 10, 11, 2, 2, 204, 203, 3, 2, 2,
	2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207,
	209, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 210, 8, 24, 2, 2, 210, 48,
	3, 2, 2, 2, 17, 2, 74, 88, 91, 108, 115, 144, 165, 170, 178, 180, 185,
	190, 194, 206, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "':'", "','", "", "", "", "", "", "", "", "", "", "",
	"'ret'", "'nop'", "'call'", "'halt'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "PREPROC", "OP", "IOP", "REG", "MOV", "IRMOV", "RRMOV",
	"MRMOV", "RMMOV", "JMP", "RET", "NOP", "CALL", "HALT", "NUM", "ID", "NL",
	"WS", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "PREPROC", "OP", "IOP", "REG", "MOV", "IRMOV",
	"RRMOV", "MRMOV", "RMMOV", "JMP", "RET", "NOP", "CALL", "HALT", "NUM",
	"ID", "NL", "WS", "LINE_COMMENT",
}

type Y64asmLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewY64asmLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *Y64asmLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewY64asmLexer(input antlr.CharStream) *Y64asmLexer {
	l := new(Y64asmLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Y64asm.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Y64asmLexer tokens.
const (
	Y64asmLexerT__0         = 1
	Y64asmLexerT__1         = 2
	Y64asmLexerT__2         = 3
	Y64asmLexerT__3         = 4
	Y64asmLexerPREPROC      = 5
	Y64asmLexerOP           = 6
	Y64asmLexerIOP          = 7
	Y64asmLexerREG          = 8
	Y64asmLexerMOV          = 9
	Y64asmLexerIRMOV        = 10
	Y64asmLexerRRMOV        = 11
	Y64asmLexerMRMOV        = 12
	Y64asmLexerRMMOV        = 13
	Y64asmLexerJMP          = 14
	Y64asmLexerRET          = 15
	Y64asmLexerNOP          = 16
	Y64asmLexerCALL         = 17
	Y64asmLexerHALT         = 18
	Y64asmLexerNUM          = 19
	Y64asmLexerID           = 20
	Y64asmLexerNL           = 21
	Y64asmLexerWS           = 22
	Y64asmLexerLINE_COMMENT = 23
)
