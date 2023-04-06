// Code generated from MyDSL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MyDSLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var mydsllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mydsllexerLexerInit() {
	staticData := &mydsllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'import'", "'service'", "'{'", "'}'", "'method'", "'response'",
		"'payload'", "'description'", "'field'", "'required'", "'optional'",
		"'match'", "'type'", "'='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "EQUAL_SIGN",
		"IDENTIFIER", "COMMENT", "MULTILINE_COMMENT", "STRING_LITERAL", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "EQUAL_SIGN", "IDENTIFIER", "COMMENT",
		"MULTILINE_COMMENT", "STRING_LITERAL", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 192, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 134,
		8, 14, 10, 14, 12, 14, 137, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 143,
		8, 15, 10, 15, 12, 15, 146, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 5, 16, 154, 8, 16, 10, 16, 12, 16, 157, 9, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 168, 8, 17, 10, 17, 12,
		17, 171, 9, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 178, 8, 17, 10,
		17, 12, 17, 181, 9, 17, 1, 17, 3, 17, 184, 8, 17, 1, 18, 4, 18, 187, 8,
		18, 11, 18, 12, 18, 188, 1, 18, 1, 18, 1, 155, 0, 19, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 1, 0, 6, 2, 0, 65, 90,
		97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 10, 10, 13, 13, 2, 0, 34,
		34, 92, 92, 2, 0, 92, 92, 96, 96, 3, 0, 9, 10, 13, 13, 32, 32, 200, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1,
		0, 0, 0, 3, 46, 1, 0, 0, 0, 5, 54, 1, 0, 0, 0, 7, 56, 1, 0, 0, 0, 9, 58,
		1, 0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 74, 1, 0, 0, 0, 15, 82, 1, 0, 0, 0,
		17, 94, 1, 0, 0, 0, 19, 100, 1, 0, 0, 0, 21, 109, 1, 0, 0, 0, 23, 118,
		1, 0, 0, 0, 25, 124, 1, 0, 0, 0, 27, 129, 1, 0, 0, 0, 29, 131, 1, 0, 0,
		0, 31, 138, 1, 0, 0, 0, 33, 149, 1, 0, 0, 0, 35, 183, 1, 0, 0, 0, 37, 186,
		1, 0, 0, 0, 39, 40, 5, 105, 0, 0, 40, 41, 5, 109, 0, 0, 41, 42, 5, 112,
		0, 0, 42, 43, 5, 111, 0, 0, 43, 44, 5, 114, 0, 0, 44, 45, 5, 116, 0, 0,
		45, 2, 1, 0, 0, 0, 46, 47, 5, 115, 0, 0, 47, 48, 5, 101, 0, 0, 48, 49,
		5, 114, 0, 0, 49, 50, 5, 118, 0, 0, 50, 51, 5, 105, 0, 0, 51, 52, 5, 99,
		0, 0, 52, 53, 5, 101, 0, 0, 53, 4, 1, 0, 0, 0, 54, 55, 5, 123, 0, 0, 55,
		6, 1, 0, 0, 0, 56, 57, 5, 125, 0, 0, 57, 8, 1, 0, 0, 0, 58, 59, 5, 109,
		0, 0, 59, 60, 5, 101, 0, 0, 60, 61, 5, 116, 0, 0, 61, 62, 5, 104, 0, 0,
		62, 63, 5, 111, 0, 0, 63, 64, 5, 100, 0, 0, 64, 10, 1, 0, 0, 0, 65, 66,
		5, 114, 0, 0, 66, 67, 5, 101, 0, 0, 67, 68, 5, 115, 0, 0, 68, 69, 5, 112,
		0, 0, 69, 70, 5, 111, 0, 0, 70, 71, 5, 110, 0, 0, 71, 72, 5, 115, 0, 0,
		72, 73, 5, 101, 0, 0, 73, 12, 1, 0, 0, 0, 74, 75, 5, 112, 0, 0, 75, 76,
		5, 97, 0, 0, 76, 77, 5, 121, 0, 0, 77, 78, 5, 108, 0, 0, 78, 79, 5, 111,
		0, 0, 79, 80, 5, 97, 0, 0, 80, 81, 5, 100, 0, 0, 81, 14, 1, 0, 0, 0, 82,
		83, 5, 100, 0, 0, 83, 84, 5, 101, 0, 0, 84, 85, 5, 115, 0, 0, 85, 86, 5,
		99, 0, 0, 86, 87, 5, 114, 0, 0, 87, 88, 5, 105, 0, 0, 88, 89, 5, 112, 0,
		0, 89, 90, 5, 116, 0, 0, 90, 91, 5, 105, 0, 0, 91, 92, 5, 111, 0, 0, 92,
		93, 5, 110, 0, 0, 93, 16, 1, 0, 0, 0, 94, 95, 5, 102, 0, 0, 95, 96, 5,
		105, 0, 0, 96, 97, 5, 101, 0, 0, 97, 98, 5, 108, 0, 0, 98, 99, 5, 100,
		0, 0, 99, 18, 1, 0, 0, 0, 100, 101, 5, 114, 0, 0, 101, 102, 5, 101, 0,
		0, 102, 103, 5, 113, 0, 0, 103, 104, 5, 117, 0, 0, 104, 105, 5, 105, 0,
		0, 105, 106, 5, 114, 0, 0, 106, 107, 5, 101, 0, 0, 107, 108, 5, 100, 0,
		0, 108, 20, 1, 0, 0, 0, 109, 110, 5, 111, 0, 0, 110, 111, 5, 112, 0, 0,
		111, 112, 5, 116, 0, 0, 112, 113, 5, 105, 0, 0, 113, 114, 5, 111, 0, 0,
		114, 115, 5, 110, 0, 0, 115, 116, 5, 97, 0, 0, 116, 117, 5, 108, 0, 0,
		117, 22, 1, 0, 0, 0, 118, 119, 5, 109, 0, 0, 119, 120, 5, 97, 0, 0, 120,
		121, 5, 116, 0, 0, 121, 122, 5, 99, 0, 0, 122, 123, 5, 104, 0, 0, 123,
		24, 1, 0, 0, 0, 124, 125, 5, 116, 0, 0, 125, 126, 5, 121, 0, 0, 126, 127,
		5, 112, 0, 0, 127, 128, 5, 101, 0, 0, 128, 26, 1, 0, 0, 0, 129, 130, 5,
		61, 0, 0, 130, 28, 1, 0, 0, 0, 131, 135, 7, 0, 0, 0, 132, 134, 7, 1, 0,
		0, 133, 132, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135,
		136, 1, 0, 0, 0, 136, 30, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139, 5,
		47, 0, 0, 139, 140, 5, 47, 0, 0, 140, 144, 1, 0, 0, 0, 141, 143, 8, 2,
		0, 0, 142, 141, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0,
		144, 145, 1, 0, 0, 0, 145, 147, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147,
		148, 6, 15, 0, 0, 148, 32, 1, 0, 0, 0, 149, 150, 5, 47, 0, 0, 150, 151,
		5, 42, 0, 0, 151, 155, 1, 0, 0, 0, 152, 154, 9, 0, 0, 0, 153, 152, 1, 0,
		0, 0, 154, 157, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0,
		156, 158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 159, 5, 42, 0, 0, 159,
		160, 5, 47, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 6, 16, 0, 0, 162, 34,
		1, 0, 0, 0, 163, 169, 5, 34, 0, 0, 164, 165, 5, 92, 0, 0, 165, 168, 9,
		0, 0, 0, 166, 168, 8, 3, 0, 0, 167, 164, 1, 0, 0, 0, 167, 166, 1, 0, 0,
		0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170,
		172, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 184, 5, 34, 0, 0, 173, 179,
		5, 96, 0, 0, 174, 175, 5, 92, 0, 0, 175, 178, 9, 0, 0, 0, 176, 178, 8,
		4, 0, 0, 177, 174, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0,
		0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 182, 184, 5, 96, 0, 0, 183, 163, 1, 0, 0, 0, 183, 173,
		1, 0, 0, 0, 184, 36, 1, 0, 0, 0, 185, 187, 7, 5, 0, 0, 186, 185, 1, 0,
		0, 0, 187, 188, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0,
		189, 190, 1, 0, 0, 0, 190, 191, 6, 18, 0, 0, 191, 38, 1, 0, 0, 0, 10, 0,
		135, 144, 155, 167, 169, 177, 179, 183, 188, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MyDSLLexerInit initializes any static state used to implement MyDSLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMyDSLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyDSLLexerInit() {
	staticData := &mydsllexerLexerStaticData
	staticData.once.Do(mydsllexerLexerInit)
}

// NewMyDSLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMyDSLLexer(input antlr.CharStream) *MyDSLLexer {
	MyDSLLexerInit()
	l := new(MyDSLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &mydsllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "MyDSL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MyDSLLexer tokens.
const (
	MyDSLLexerT__0              = 1
	MyDSLLexerT__1              = 2
	MyDSLLexerT__2              = 3
	MyDSLLexerT__3              = 4
	MyDSLLexerT__4              = 5
	MyDSLLexerT__5              = 6
	MyDSLLexerT__6              = 7
	MyDSLLexerT__7              = 8
	MyDSLLexerT__8              = 9
	MyDSLLexerT__9              = 10
	MyDSLLexerT__10             = 11
	MyDSLLexerT__11             = 12
	MyDSLLexerT__12             = 13
	MyDSLLexerEQUAL_SIGN        = 14
	MyDSLLexerIDENTIFIER        = 15
	MyDSLLexerCOMMENT           = 16
	MyDSLLexerMULTILINE_COMMENT = 17
	MyDSLLexerSTRING_LITERAL    = 18
	MyDSLLexerWS                = 19
)
