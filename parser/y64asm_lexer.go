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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 24, 205,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 73, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 87, 10, 7, 3, 7, 5, 7, 90, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 5, 9, 107, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5,
	10, 114, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 143,
	10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 5, 19, 160, 10, 19, 3, 19, 6,
	19, 163, 10, 19, 13, 19, 14, 19, 164, 3, 19, 3, 19, 3, 19, 3, 19, 6, 19,
	171, 10, 19, 13, 19, 14, 19, 172, 5, 19, 175, 10, 19, 3, 20, 6, 20, 178,
	10, 20, 13, 20, 14, 20, 179, 3, 20, 7, 20, 183, 10, 20, 12, 20, 14, 20,
	186, 11, 20, 3, 21, 5, 21, 189, 10, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 23, 3, 23, 7, 23, 199, 10, 23, 12, 23, 14, 23, 202, 11, 23,
	3, 23, 3, 23, 2, 2, 24, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 3, 2, 12, 3, 2, 58, 59, 3,
	2, 50, 54, 4, 2, 102, 102, 117, 117, 4, 2, 100, 100, 117, 117, 3, 2, 99,
	102, 3, 2, 50, 59, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 67, 92, 99, 124,
	4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 226, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3,
	2, 2, 2, 2, 45, 3, 2, 2, 2, 3, 47, 3, 2, 2, 2, 5, 49, 3, 2, 2, 2, 7, 51,
	3, 2, 2, 2, 9, 53, 3, 2, 2, 2, 11, 55, 3, 2, 2, 2, 13, 86, 3, 2, 2, 2,
	15, 91, 3, 2, 2, 2, 17, 94, 3, 2, 2, 2, 19, 108, 3, 2, 2, 2, 21, 115, 3,
	2, 2, 2, 23, 120, 3, 2, 2, 2, 25, 125, 3, 2, 2, 2, 27, 130, 3, 2, 2, 2,
	29, 135, 3, 2, 2, 2, 31, 144, 3, 2, 2, 2, 33, 148, 3, 2, 2, 2, 35, 153,
	3, 2, 2, 2, 37, 174, 3, 2, 2, 2, 39, 177, 3, 2, 2, 2, 41, 188, 3, 2, 2,
	2, 43, 192, 3, 2, 2, 2, 45, 196, 3, 2, 2, 2, 47, 48, 7, 42, 2, 2, 48, 4,
	3, 2, 2, 2, 49, 50, 7, 43, 2, 2, 50, 6, 3, 2, 2, 2, 51, 52, 7, 60, 2, 2,
	52, 8, 3, 2, 2, 2, 53, 54, 7, 46, 2, 2, 54, 10, 3, 2, 2, 2, 55, 72, 7,
	48, 2, 2, 56, 57, 7, 115, 2, 2, 57, 58, 7, 119, 2, 2, 58, 59, 7, 99, 2,
	2, 59, 73, 7, 102, 2, 2, 60, 61, 7, 110, 2, 2, 61, 62, 7, 113, 2, 2, 62,
	63, 7, 112, 2, 2, 63, 73, 7, 105, 2, 2, 64, 65, 7, 114, 2, 2, 65, 66, 7,
	113, 2, 2, 66, 73, 7, 117, 2, 2, 67, 68, 7, 99, 2, 2, 68, 69, 7, 110, 2,
	2, 69, 70, 7, 107, 2, 2, 70, 71, 7, 105, 2, 2, 71, 73, 7, 112, 2, 2, 72,
	56, 3, 2, 2, 2, 72, 60, 3, 2, 2, 2, 72, 64, 3, 2, 2, 2, 72, 67, 3, 2, 2,
	2, 73, 12, 3, 2, 2, 2, 74, 75, 7, 99, 2, 2, 75, 76, 7, 102, 2, 2, 76, 87,
	7, 102, 2, 2, 77, 78, 7, 99, 2, 2, 78, 79, 7, 112, 2, 2, 79, 87, 7, 102,
	2, 2, 80, 81, 7, 122, 2, 2, 81, 82, 7, 113, 2, 2, 82, 87, 7, 116, 2, 2,
	83, 84, 7, 117, 2, 2, 84, 85, 7, 119, 2, 2, 85, 87, 7, 100, 2, 2, 86, 74,
	3, 2, 2, 2, 86, 77, 3, 2, 2, 2, 86, 80, 3, 2, 2, 2, 86, 83, 3, 2, 2, 2,
	87, 89, 3, 2, 2, 2, 88, 90, 7, 115, 2, 2, 89, 88, 3, 2, 2, 2, 89, 90, 3,
	2, 2, 2, 90, 14, 3, 2, 2, 2, 91, 92, 7, 107, 2, 2, 92, 93, 5, 13, 7, 2,
	93, 16, 3, 2, 2, 2, 94, 95, 7, 39, 2, 2, 95, 96, 7, 116, 2, 2, 96, 106,
	3, 2, 2, 2, 97, 107, 9, 2, 2, 2, 98, 99, 7, 51, 2, 2, 99, 107, 9, 3, 2,
	2, 100, 101, 9, 4, 2, 2, 101, 107, 7, 107, 2, 2, 102, 103, 9, 5, 2, 2,
	103, 107, 7, 114, 2, 2, 104, 105, 9, 6, 2, 2, 105, 107, 7, 122, 2, 2, 106,
	97, 3, 2, 2, 2, 106, 98, 3, 2, 2, 2, 106, 100, 3, 2, 2, 2, 106, 102, 3,
	2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 18, 3, 2, 2, 2, 108, 109, 7, 111, 2,
	2, 109, 110, 7, 113, 2, 2, 110, 111, 7, 120, 2, 2, 111, 113, 3, 2, 2, 2,
	112, 114, 7, 115, 2, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114,
	20, 3, 2, 2, 2, 115, 116, 7, 107, 2, 2, 116, 117, 7, 116, 2, 2, 117, 118,
	3, 2, 2, 2, 118, 119, 5, 19, 10, 2, 119, 22, 3, 2, 2, 2, 120, 121, 7, 116,
	2, 2, 121, 122, 7, 116, 2, 2, 122, 123, 3, 2, 2, 2, 123, 124, 5, 19, 10,
	2, 124, 24, 3, 2, 2, 2, 125, 126, 7, 111, 2, 2, 126, 127, 7, 116, 2, 2,
	127, 128, 3, 2, 2, 2, 128, 129, 5, 19, 10, 2, 129, 26, 3, 2, 2, 2, 130,
	131, 7, 116, 2, 2, 131, 132, 7, 111, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134,
	5, 19, 10, 2, 134, 28, 3, 2, 2, 2, 135, 142, 7, 108, 2, 2, 136, 137, 7,
	112, 2, 2, 137, 143, 7, 103, 2, 2, 138, 139, 7, 103, 2, 2, 139, 143, 7,
	115, 2, 2, 140, 141, 7, 111, 2, 2, 141, 143, 7, 114, 2, 2, 142, 136, 3,
	2, 2, 2, 142, 138, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 143, 30, 3, 2, 2,
	2, 144, 145, 7, 116, 2, 2, 145, 146, 7, 103, 2, 2, 146, 147, 7, 118, 2,
	2, 147, 32, 3, 2, 2, 2, 148, 149, 7, 101, 2, 2, 149, 150, 7, 99, 2, 2,
	150, 151, 7, 110, 2, 2, 151, 152, 7, 110, 2, 2, 152, 34, 3, 2, 2, 2, 153,
	154, 7, 106, 2, 2, 154, 155, 7, 99, 2, 2, 155, 156, 7, 110, 2, 2, 156,
	157, 7, 118, 2, 2, 157, 36, 3, 2, 2, 2, 158, 160, 7, 38, 2, 2, 159, 158,
	3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 162, 3, 2, 2, 2, 161, 163, 9, 7,
	2, 2, 162, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2,
	164, 165, 3, 2, 2, 2, 165, 175, 3, 2, 2, 2, 166, 167, 7, 50, 2, 2, 167,
	168, 7, 122, 2, 2, 168, 170, 3, 2, 2, 2, 169, 171, 9, 8, 2, 2, 170, 169,
	3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2,
	2, 2, 173, 175, 3, 2, 2, 2, 174, 159, 3, 2, 2, 2, 174, 166, 3, 2, 2, 2,
	175, 38, 3, 2, 2, 2, 176, 178, 9, 9, 2, 2, 177, 176, 3, 2, 2, 2, 178, 179,
	3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 184, 3, 2,
	2, 2, 181, 183, 9, 7, 2, 2, 182, 181, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2,
	184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 40, 3, 2, 2, 2, 186, 184,
	3, 2, 2, 2, 187, 189, 7, 15, 2, 2, 188, 187, 3, 2, 2, 2, 188, 189, 3, 2,
	2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 7, 12, 2, 2, 191, 42, 3, 2, 2, 2,
	192, 193, 9, 10, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195, 8, 22, 2, 2, 195,
	44, 3, 2, 2, 2, 196, 200, 7, 37, 2, 2, 197, 199, 10, 11, 2, 2, 198, 197,
	3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2,
	2, 2, 201, 203, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 204, 8, 23, 2, 2,
	204, 46, 3, 2, 2, 2, 17, 2, 72, 86, 89, 106, 113, 142, 159, 164, 172, 174,
	179, 184, 188, 200, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "':'", "','", "", "", "", "", "", "", "", "", "", "",
	"'ret'", "'call'", "'halt'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "PREPROC", "OP", "IOP", "REG", "MOV", "IRMOV", "RRMOV",
	"MRMOV", "RMMOV", "JMP", "RET", "CALL", "HALT", "NUM", "ID", "NL", "WS",
	"LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "PREPROC", "OP", "IOP", "REG", "MOV", "IRMOV",
	"RRMOV", "MRMOV", "RMMOV", "JMP", "RET", "CALL", "HALT", "NUM", "ID", "NL",
	"WS", "LINE_COMMENT",
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
	Y64asmLexerCALL         = 16
	Y64asmLexerHALT         = 17
	Y64asmLexerNUM          = 18
	Y64asmLexerID           = 19
	Y64asmLexerNL           = 20
	Y64asmLexerWS           = 21
	Y64asmLexerLINE_COMMENT = 22
)
