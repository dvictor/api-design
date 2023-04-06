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
		"'payload'", "'description'", "'field'", "'required'", "'true'", "'false'",
		"'match'", "'type'", "'='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "EQUAL_SIGN",
		"IDENTIFIER", "COMMENT", "MULTILINE_COMMENT", "STRING_LITERAL", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "EQUAL_SIGN", "IDENTIFIER",
		"COMMENT", "MULTILINE_COMMENT", "STRING_LITERAL", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 196, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 5, 15, 138, 8, 15, 10, 15, 12, 15, 141, 9, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 5, 16, 147, 8, 16, 10, 16, 12, 16, 150, 9, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 158, 8, 17, 10, 17, 12, 17, 161,
		9, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5,
		18, 172, 8, 18, 10, 18, 12, 18, 175, 9, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 5, 18, 182, 8, 18, 10, 18, 12, 18, 185, 9, 18, 1, 18, 3, 18, 188,
		8, 18, 1, 19, 4, 19, 191, 8, 19, 11, 19, 12, 19, 192, 1, 19, 1, 19, 1,
		159, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 1, 0, 6, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122,
		2, 0, 10, 10, 13, 13, 2, 0, 34, 34, 92, 92, 2, 0, 92, 92, 96, 96, 3, 0,
		9, 10, 13, 13, 32, 32, 204, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 48, 1, 0,
		0, 0, 5, 56, 1, 0, 0, 0, 7, 58, 1, 0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 67,
		1, 0, 0, 0, 13, 76, 1, 0, 0, 0, 15, 84, 1, 0, 0, 0, 17, 96, 1, 0, 0, 0,
		19, 102, 1, 0, 0, 0, 21, 111, 1, 0, 0, 0, 23, 116, 1, 0, 0, 0, 25, 122,
		1, 0, 0, 0, 27, 128, 1, 0, 0, 0, 29, 133, 1, 0, 0, 0, 31, 135, 1, 0, 0,
		0, 33, 142, 1, 0, 0, 0, 35, 153, 1, 0, 0, 0, 37, 187, 1, 0, 0, 0, 39, 190,
		1, 0, 0, 0, 41, 42, 5, 105, 0, 0, 42, 43, 5, 109, 0, 0, 43, 44, 5, 112,
		0, 0, 44, 45, 5, 111, 0, 0, 45, 46, 5, 114, 0, 0, 46, 47, 5, 116, 0, 0,
		47, 2, 1, 0, 0, 0, 48, 49, 5, 115, 0, 0, 49, 50, 5, 101, 0, 0, 50, 51,
		5, 114, 0, 0, 51, 52, 5, 118, 0, 0, 52, 53, 5, 105, 0, 0, 53, 54, 5, 99,
		0, 0, 54, 55, 5, 101, 0, 0, 55, 4, 1, 0, 0, 0, 56, 57, 5, 123, 0, 0, 57,
		6, 1, 0, 0, 0, 58, 59, 5, 125, 0, 0, 59, 8, 1, 0, 0, 0, 60, 61, 5, 109,
		0, 0, 61, 62, 5, 101, 0, 0, 62, 63, 5, 116, 0, 0, 63, 64, 5, 104, 0, 0,
		64, 65, 5, 111, 0, 0, 65, 66, 5, 100, 0, 0, 66, 10, 1, 0, 0, 0, 67, 68,
		5, 114, 0, 0, 68, 69, 5, 101, 0, 0, 69, 70, 5, 115, 0, 0, 70, 71, 5, 112,
		0, 0, 71, 72, 5, 111, 0, 0, 72, 73, 5, 110, 0, 0, 73, 74, 5, 115, 0, 0,
		74, 75, 5, 101, 0, 0, 75, 12, 1, 0, 0, 0, 76, 77, 5, 112, 0, 0, 77, 78,
		5, 97, 0, 0, 78, 79, 5, 121, 0, 0, 79, 80, 5, 108, 0, 0, 80, 81, 5, 111,
		0, 0, 81, 82, 5, 97, 0, 0, 82, 83, 5, 100, 0, 0, 83, 14, 1, 0, 0, 0, 84,
		85, 5, 100, 0, 0, 85, 86, 5, 101, 0, 0, 86, 87, 5, 115, 0, 0, 87, 88, 5,
		99, 0, 0, 88, 89, 5, 114, 0, 0, 89, 90, 5, 105, 0, 0, 90, 91, 5, 112, 0,
		0, 91, 92, 5, 116, 0, 0, 92, 93, 5, 105, 0, 0, 93, 94, 5, 111, 0, 0, 94,
		95, 5, 110, 0, 0, 95, 16, 1, 0, 0, 0, 96, 97, 5, 102, 0, 0, 97, 98, 5,
		105, 0, 0, 98, 99, 5, 101, 0, 0, 99, 100, 5, 108, 0, 0, 100, 101, 5, 100,
		0, 0, 101, 18, 1, 0, 0, 0, 102, 103, 5, 114, 0, 0, 103, 104, 5, 101, 0,
		0, 104, 105, 5, 113, 0, 0, 105, 106, 5, 117, 0, 0, 106, 107, 5, 105, 0,
		0, 107, 108, 5, 114, 0, 0, 108, 109, 5, 101, 0, 0, 109, 110, 5, 100, 0,
		0, 110, 20, 1, 0, 0, 0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 114, 0, 0,
		113, 114, 5, 117, 0, 0, 114, 115, 5, 101, 0, 0, 115, 22, 1, 0, 0, 0, 116,
		117, 5, 102, 0, 0, 117, 118, 5, 97, 0, 0, 118, 119, 5, 108, 0, 0, 119,
		120, 5, 115, 0, 0, 120, 121, 5, 101, 0, 0, 121, 24, 1, 0, 0, 0, 122, 123,
		5, 109, 0, 0, 123, 124, 5, 97, 0, 0, 124, 125, 5, 116, 0, 0, 125, 126,
		5, 99, 0, 0, 126, 127, 5, 104, 0, 0, 127, 26, 1, 0, 0, 0, 128, 129, 5,
		116, 0, 0, 129, 130, 5, 121, 0, 0, 130, 131, 5, 112, 0, 0, 131, 132, 5,
		101, 0, 0, 132, 28, 1, 0, 0, 0, 133, 134, 5, 61, 0, 0, 134, 30, 1, 0, 0,
		0, 135, 139, 7, 0, 0, 0, 136, 138, 7, 1, 0, 0, 137, 136, 1, 0, 0, 0, 138,
		141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 32, 1,
		0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 143, 5, 47, 0, 0, 143, 144, 5, 47,
		0, 0, 144, 148, 1, 0, 0, 0, 145, 147, 8, 2, 0, 0, 146, 145, 1, 0, 0, 0,
		147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		151, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 152, 6, 16, 0, 0, 152, 34,
		1, 0, 0, 0, 153, 154, 5, 47, 0, 0, 154, 155, 5, 42, 0, 0, 155, 159, 1,
		0, 0, 0, 156, 158, 9, 0, 0, 0, 157, 156, 1, 0, 0, 0, 158, 161, 1, 0, 0,
		0, 159, 160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 162, 1, 0, 0, 0, 161,
		159, 1, 0, 0, 0, 162, 163, 5, 42, 0, 0, 163, 164, 5, 47, 0, 0, 164, 165,
		1, 0, 0, 0, 165, 166, 6, 17, 0, 0, 166, 36, 1, 0, 0, 0, 167, 173, 5, 34,
		0, 0, 168, 169, 5, 92, 0, 0, 169, 172, 9, 0, 0, 0, 170, 172, 8, 3, 0, 0,
		171, 168, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173,
		1, 0, 0, 0, 176, 188, 5, 34, 0, 0, 177, 183, 5, 96, 0, 0, 178, 179, 5,
		92, 0, 0, 179, 182, 9, 0, 0, 0, 180, 182, 8, 4, 0, 0, 181, 178, 1, 0, 0,
		0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183,
		184, 1, 0, 0, 0, 184, 186, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 188,
		5, 96, 0, 0, 187, 167, 1, 0, 0, 0, 187, 177, 1, 0, 0, 0, 188, 38, 1, 0,
		0, 0, 189, 191, 7, 5, 0, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0,
		192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194,
		195, 6, 19, 0, 0, 195, 40, 1, 0, 0, 0, 10, 0, 139, 148, 159, 171, 173,
		181, 183, 187, 192, 1, 0, 1, 0,
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
	MyDSLLexerT__13             = 14
	MyDSLLexerEQUAL_SIGN        = 15
	MyDSLLexerIDENTIFIER        = 16
	MyDSLLexerCOMMENT           = 17
	MyDSLLexerMULTILINE_COMMENT = 18
	MyDSLLexerSTRING_LITERAL    = 19
	MyDSLLexerWS                = 20
)
