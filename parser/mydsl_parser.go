// Code generated from MyDSL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // MyDSL

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MyDSLParser struct {
	*antlr.BaseParser
}

var mydslParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mydslParserInit() {
	staticData := &mydslParserStaticData
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
		"apiBody", "importDirective", "service", "serviceName", "serviceBody",
		"method", "methodBody", "response", "payload", "description", "fieldsBody",
		"field", "fieldName", "fieldType", "fieldBody", "required", "validation",
		"type", "typeBody",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 144, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 4, 0, 42, 8,
		0, 11, 0, 12, 0, 43, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 4, 4, 61, 8, 4, 11, 4, 12, 4, 62,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 4, 6, 74, 8, 6, 11,
		6, 12, 6, 75, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 85, 8, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 94, 8, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 4, 10, 101, 8, 10, 11, 10, 12, 10, 102, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 112, 8, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 4, 14, 121, 8, 14, 11, 14, 12, 14, 122, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 18, 4, 18, 140, 8, 18, 11, 18, 12, 18, 141, 1,
		18, 0, 0, 19, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 0, 1, 1, 0, 11, 12, 140, 0, 41, 1, 0, 0, 0, 2, 47, 1, 0, 0,
		0, 4, 50, 1, 0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 60, 1, 0, 0, 0, 10, 64, 1,
		0, 0, 0, 12, 73, 1, 0, 0, 0, 14, 77, 1, 0, 0, 0, 16, 86, 1, 0, 0, 0, 18,
		95, 1, 0, 0, 0, 20, 100, 1, 0, 0, 0, 22, 104, 1, 0, 0, 0, 24, 113, 1, 0,
		0, 0, 26, 115, 1, 0, 0, 0, 28, 120, 1, 0, 0, 0, 30, 124, 1, 0, 0, 0, 32,
		128, 1, 0, 0, 0, 34, 132, 1, 0, 0, 0, 36, 139, 1, 0, 0, 0, 38, 42, 3, 2,
		1, 0, 39, 42, 3, 4, 2, 0, 40, 42, 3, 34, 17, 0, 41, 38, 1, 0, 0, 0, 41,
		39, 1, 0, 0, 0, 41, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 41, 1, 0, 0,
		0, 43, 44, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 5, 0, 0, 1, 46, 1, 1,
		0, 0, 0, 47, 48, 5, 1, 0, 0, 48, 49, 5, 19, 0, 0, 49, 3, 1, 0, 0, 0, 50,
		51, 5, 2, 0, 0, 51, 52, 3, 6, 3, 0, 52, 53, 5, 3, 0, 0, 53, 54, 3, 8, 4,
		0, 54, 55, 5, 4, 0, 0, 55, 5, 1, 0, 0, 0, 56, 57, 5, 16, 0, 0, 57, 7, 1,
		0, 0, 0, 58, 61, 3, 18, 9, 0, 59, 61, 3, 10, 5, 0, 60, 58, 1, 0, 0, 0,
		60, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1,
		0, 0, 0, 63, 9, 1, 0, 0, 0, 64, 65, 5, 5, 0, 0, 65, 66, 5, 16, 0, 0, 66,
		67, 5, 3, 0, 0, 67, 68, 3, 12, 6, 0, 68, 69, 5, 4, 0, 0, 69, 11, 1, 0,
		0, 0, 70, 74, 3, 18, 9, 0, 71, 74, 3, 16, 8, 0, 72, 74, 3, 14, 7, 0, 73,
		70, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0,
		0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 13, 1, 0, 0, 0, 77, 84,
		5, 6, 0, 0, 78, 79, 5, 3, 0, 0, 79, 80, 3, 20, 10, 0, 80, 81, 5, 4, 0,
		0, 81, 85, 1, 0, 0, 0, 82, 83, 5, 15, 0, 0, 83, 85, 5, 16, 0, 0, 84, 78,
		1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 15, 1, 0, 0, 0, 86, 93, 5, 7, 0, 0,
		87, 88, 5, 3, 0, 0, 88, 89, 3, 20, 10, 0, 89, 90, 5, 4, 0, 0, 90, 94, 1,
		0, 0, 0, 91, 92, 5, 15, 0, 0, 92, 94, 5, 16, 0, 0, 93, 87, 1, 0, 0, 0,
		93, 91, 1, 0, 0, 0, 94, 17, 1, 0, 0, 0, 95, 96, 5, 8, 0, 0, 96, 97, 5,
		15, 0, 0, 97, 98, 5, 19, 0, 0, 98, 19, 1, 0, 0, 0, 99, 101, 3, 22, 11,
		0, 100, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102,
		103, 1, 0, 0, 0, 103, 21, 1, 0, 0, 0, 104, 105, 5, 9, 0, 0, 105, 106, 3,
		24, 12, 0, 106, 111, 3, 26, 13, 0, 107, 108, 5, 3, 0, 0, 108, 109, 3, 28,
		14, 0, 109, 110, 5, 4, 0, 0, 110, 112, 1, 0, 0, 0, 111, 107, 1, 0, 0, 0,
		111, 112, 1, 0, 0, 0, 112, 23, 1, 0, 0, 0, 113, 114, 5, 16, 0, 0, 114,
		25, 1, 0, 0, 0, 115, 116, 5, 16, 0, 0, 116, 27, 1, 0, 0, 0, 117, 121, 3,
		18, 9, 0, 118, 121, 3, 32, 16, 0, 119, 121, 3, 30, 15, 0, 120, 117, 1,
		0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0,
		0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 29, 1, 0, 0, 0, 124,
		125, 5, 10, 0, 0, 125, 126, 5, 15, 0, 0, 126, 127, 7, 0, 0, 0, 127, 31,
		1, 0, 0, 0, 128, 129, 5, 13, 0, 0, 129, 130, 5, 15, 0, 0, 130, 131, 5,
		19, 0, 0, 131, 33, 1, 0, 0, 0, 132, 133, 5, 14, 0, 0, 133, 134, 5, 16,
		0, 0, 134, 135, 5, 3, 0, 0, 135, 136, 3, 36, 18, 0, 136, 137, 5, 4, 0,
		0, 137, 35, 1, 0, 0, 0, 138, 140, 3, 22, 11, 0, 139, 138, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 37, 1,
		0, 0, 0, 13, 41, 43, 60, 62, 73, 75, 84, 93, 102, 111, 120, 122, 141,
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

// MyDSLParserInit initializes any static state used to implement MyDSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMyDSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyDSLParserInit() {
	staticData := &mydslParserStaticData
	staticData.once.Do(mydslParserInit)
}

// NewMyDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMyDSLParser(input antlr.TokenStream) *MyDSLParser {
	MyDSLParserInit()
	this := new(MyDSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mydslParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MyDSL.g4"

	return this
}

// MyDSLParser tokens.
const (
	MyDSLParserEOF               = antlr.TokenEOF
	MyDSLParserT__0              = 1
	MyDSLParserT__1              = 2
	MyDSLParserT__2              = 3
	MyDSLParserT__3              = 4
	MyDSLParserT__4              = 5
	MyDSLParserT__5              = 6
	MyDSLParserT__6              = 7
	MyDSLParserT__7              = 8
	MyDSLParserT__8              = 9
	MyDSLParserT__9              = 10
	MyDSLParserT__10             = 11
	MyDSLParserT__11             = 12
	MyDSLParserT__12             = 13
	MyDSLParserT__13             = 14
	MyDSLParserEQUAL_SIGN        = 15
	MyDSLParserIDENTIFIER        = 16
	MyDSLParserCOMMENT           = 17
	MyDSLParserMULTILINE_COMMENT = 18
	MyDSLParserSTRING_LITERAL    = 19
	MyDSLParserWS                = 20
)

// MyDSLParser rules.
const (
	MyDSLParserRULE_apiBody         = 0
	MyDSLParserRULE_importDirective = 1
	MyDSLParserRULE_service         = 2
	MyDSLParserRULE_serviceName     = 3
	MyDSLParserRULE_serviceBody     = 4
	MyDSLParserRULE_method          = 5
	MyDSLParserRULE_methodBody      = 6
	MyDSLParserRULE_response        = 7
	MyDSLParserRULE_payload         = 8
	MyDSLParserRULE_description     = 9
	MyDSLParserRULE_fieldsBody      = 10
	MyDSLParserRULE_field           = 11
	MyDSLParserRULE_fieldName       = 12
	MyDSLParserRULE_fieldType       = 13
	MyDSLParserRULE_fieldBody       = 14
	MyDSLParserRULE_required        = 15
	MyDSLParserRULE_validation      = 16
	MyDSLParserRULE_type            = 17
	MyDSLParserRULE_typeBody        = 18
)

// IApiBodyContext is an interface to support dynamic dispatch.
type IApiBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllImportDirective() []IImportDirectiveContext
	ImportDirective(i int) IImportDirectiveContext
	AllService() []IServiceContext
	Service(i int) IServiceContext
	AllType_() []ITypeContext
	Type_(i int) ITypeContext

	// IsApiBodyContext differentiates from other interfaces.
	IsApiBodyContext()
}

type ApiBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiBodyContext() *ApiBodyContext {
	var p = new(ApiBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_apiBody
	return p
}

func (*ApiBodyContext) IsApiBodyContext() {}

func NewApiBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiBodyContext {
	var p = new(ApiBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_apiBody

	return p
}

func (s *ApiBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiBodyContext) EOF() antlr.TerminalNode {
	return s.GetToken(MyDSLParserEOF, 0)
}

func (s *ApiBodyContext) AllImportDirective() []IImportDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IImportDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportDirectiveContext); ok {
			tst[i] = t.(IImportDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *ApiBodyContext) ImportDirective(i int) IImportDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportDirectiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportDirectiveContext)
}

func (s *ApiBodyContext) AllService() []IServiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceContext); ok {
			len++
		}
	}

	tst := make([]IServiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceContext); ok {
			tst[i] = t.(IServiceContext)
			i++
		}
	}

	return tst
}

func (s *ApiBodyContext) Service(i int) IServiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *ApiBodyContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ApiBodyContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ApiBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterApiBody(s)
	}
}

func (s *ApiBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitApiBody(s)
	}
}

func (p *MyDSLParser) ApiBody() (localctx IApiBodyContext) {
	this := p
	_ = this

	localctx = NewApiBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MyDSLParserRULE_apiBody)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16390) != 0) {
		p.SetState(41)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MyDSLParserT__0:
			{
				p.SetState(38)
				p.ImportDirective()
			}

		case MyDSLParserT__1:
			{
				p.SetState(39)
				p.Service()
			}

		case MyDSLParserT__13:
			{
				p.SetState(40)
				p.Type_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(45)
		p.Match(MyDSLParserEOF)
	}

	return localctx
}

// IImportDirectiveContext is an interface to support dynamic dispatch.
type IImportDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsImportDirectiveContext differentiates from other interfaces.
	IsImportDirectiveContext()
}

type ImportDirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDirectiveContext() *ImportDirectiveContext {
	var p = new(ImportDirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_importDirective
	return p
}

func (*ImportDirectiveContext) IsImportDirectiveContext() {}

func NewImportDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDirectiveContext {
	var p = new(ImportDirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_importDirective

	return p
}

func (s *ImportDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDirectiveContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MyDSLParserSTRING_LITERAL, 0)
}

func (s *ImportDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterImportDirective(s)
	}
}

func (s *ImportDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitImportDirective(s)
	}
}

func (p *MyDSLParser) ImportDirective() (localctx IImportDirectiveContext) {
	this := p
	_ = this

	localctx = NewImportDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MyDSLParserRULE_importDirective)

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
		p.SetState(47)
		p.Match(MyDSLParserT__0)
	}
	{
		p.SetState(48)
		p.Match(MyDSLParserSTRING_LITERAL)
	}

	return localctx
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ServiceName() IServiceNameContext
	ServiceBody() IServiceBodyContext

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) ServiceName() IServiceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceContext) ServiceBody() IServiceBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceBodyContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitService(s)
	}
}

func (p *MyDSLParser) Service() (localctx IServiceContext) {
	this := p
	_ = this

	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MyDSLParserRULE_service)

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
		p.SetState(50)
		p.Match(MyDSLParserT__1)
	}
	{
		p.SetState(51)
		p.ServiceName()
	}
	{
		p.SetState(52)
		p.Match(MyDSLParserT__2)
	}
	{
		p.SetState(53)
		p.ServiceBody()
	}
	{
		p.SetState(54)
		p.Match(MyDSLParserT__3)
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyDSLParserIDENTIFIER, 0)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (p *MyDSLParser) ServiceName() (localctx IServiceNameContext) {
	this := p
	_ = this

	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MyDSLParserRULE_serviceName)

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
		p.SetState(56)
		p.Match(MyDSLParserIDENTIFIER)
	}

	return localctx
}

// IServiceBodyContext is an interface to support dynamic dispatch.
type IServiceBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDescription() []IDescriptionContext
	Description(i int) IDescriptionContext
	AllMethod() []IMethodContext
	Method(i int) IMethodContext

	// IsServiceBodyContext differentiates from other interfaces.
	IsServiceBodyContext()
}

type ServiceBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceBodyContext() *ServiceBodyContext {
	var p = new(ServiceBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_serviceBody
	return p
}

func (*ServiceBodyContext) IsServiceBodyContext() {}

func NewServiceBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceBodyContext {
	var p = new(ServiceBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_serviceBody

	return p
}

func (s *ServiceBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceBodyContext) AllDescription() []IDescriptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDescriptionContext); ok {
			len++
		}
	}

	tst := make([]IDescriptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDescriptionContext); ok {
			tst[i] = t.(IDescriptionContext)
			i++
		}
	}

	return tst
}

func (s *ServiceBodyContext) Description(i int) IDescriptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ServiceBodyContext) AllMethod() []IMethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodContext); ok {
			len++
		}
	}

	tst := make([]IMethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodContext); ok {
			tst[i] = t.(IMethodContext)
			i++
		}
	}

	return tst
}

func (s *ServiceBodyContext) Method(i int) IMethodContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *ServiceBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterServiceBody(s)
	}
}

func (s *ServiceBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitServiceBody(s)
	}
}

func (p *MyDSLParser) ServiceBody() (localctx IServiceBodyContext) {
	this := p
	_ = this

	localctx = NewServiceBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MyDSLParserRULE_serviceBody)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MyDSLParserT__4 || _la == MyDSLParserT__7 {
		p.SetState(60)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MyDSLParserT__7:
			{
				p.SetState(58)
				p.Description()
			}

		case MyDSLParserT__4:
			{
				p.SetState(59)
				p.Method()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	MethodBody() IMethodBodyContext

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyDSLParserIDENTIFIER, 0)
}

func (s *MethodContext) MethodBody() IMethodBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodBodyContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *MyDSLParser) Method() (localctx IMethodContext) {
	this := p
	_ = this

	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MyDSLParserRULE_method)

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
		p.SetState(64)
		p.Match(MyDSLParserT__4)
	}
	{
		p.SetState(65)
		p.Match(MyDSLParserIDENTIFIER)
	}
	{
		p.SetState(66)
		p.Match(MyDSLParserT__2)
	}
	{
		p.SetState(67)
		p.MethodBody()
	}
	{
		p.SetState(68)
		p.Match(MyDSLParserT__3)
	}

	return localctx
}

// IMethodBodyContext is an interface to support dynamic dispatch.
type IMethodBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDescription() []IDescriptionContext
	Description(i int) IDescriptionContext
	AllPayload() []IPayloadContext
	Payload(i int) IPayloadContext
	AllResponse() []IResponseContext
	Response(i int) IResponseContext

	// IsMethodBodyContext differentiates from other interfaces.
	IsMethodBodyContext()
}

type MethodBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodBodyContext() *MethodBodyContext {
	var p = new(MethodBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_methodBody
	return p
}

func (*MethodBodyContext) IsMethodBodyContext() {}

func NewMethodBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodBodyContext {
	var p = new(MethodBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_methodBody

	return p
}

func (s *MethodBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodBodyContext) AllDescription() []IDescriptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDescriptionContext); ok {
			len++
		}
	}

	tst := make([]IDescriptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDescriptionContext); ok {
			tst[i] = t.(IDescriptionContext)
			i++
		}
	}

	return tst
}

func (s *MethodBodyContext) Description(i int) IDescriptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *MethodBodyContext) AllPayload() []IPayloadContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPayloadContext); ok {
			len++
		}
	}

	tst := make([]IPayloadContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPayloadContext); ok {
			tst[i] = t.(IPayloadContext)
			i++
		}
	}

	return tst
}

func (s *MethodBodyContext) Payload(i int) IPayloadContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPayloadContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPayloadContext)
}

func (s *MethodBodyContext) AllResponse() []IResponseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResponseContext); ok {
			len++
		}
	}

	tst := make([]IResponseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResponseContext); ok {
			tst[i] = t.(IResponseContext)
			i++
		}
	}

	return tst
}

func (s *MethodBodyContext) Response(i int) IResponseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResponseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResponseContext)
}

func (s *MethodBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterMethodBody(s)
	}
}

func (s *MethodBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitMethodBody(s)
	}
}

func (p *MyDSLParser) MethodBody() (localctx IMethodBodyContext) {
	this := p
	_ = this

	localctx = NewMethodBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MyDSLParserRULE_methodBody)
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
		p.SetState(73)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MyDSLParserT__7:
			{
				p.SetState(70)
				p.Description()
			}

		case MyDSLParserT__6:
			{
				p.SetState(71)
				p.Payload()
			}

		case MyDSLParserT__5:
			{
				p.SetState(72)
				p.Response()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IResponseContext is an interface to support dynamic dispatch.
type IResponseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL_SIGN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	FieldsBody() IFieldsBodyContext

	// IsResponseContext differentiates from other interfaces.
	IsResponseContext()
}

type ResponseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResponseContext() *ResponseContext {
	var p = new(ResponseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_response
	return p
}

func (*ResponseContext) IsResponseContext() {}

func NewResponseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponseContext {
	var p = new(ResponseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_response

	return p
}

func (s *ResponseContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponseContext) EQUAL_SIGN() antlr.TerminalNode {
	return s.GetToken(MyDSLParserEQUAL_SIGN, 0)
}

func (s *ResponseContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyDSLParserIDENTIFIER, 0)
}

func (s *ResponseContext) FieldsBody() IFieldsBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsBodyContext)
}

func (s *ResponseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterResponse(s)
	}
}

func (s *ResponseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitResponse(s)
	}
}

func (p *MyDSLParser) Response() (localctx IResponseContext) {
	this := p
	_ = this

	localctx = NewResponseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MyDSLParserRULE_response)

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
		p.SetState(77)
		p.Match(MyDSLParserT__5)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MyDSLParserT__2:
		{
			p.SetState(78)
			p.Match(MyDSLParserT__2)
		}
		{
			p.SetState(79)
			p.FieldsBody()
		}
		{
			p.SetState(80)
			p.Match(MyDSLParserT__3)
		}

	case MyDSLParserEQUAL_SIGN:
		{
			p.SetState(82)
			p.Match(MyDSLParserEQUAL_SIGN)
		}
		{
			p.SetState(83)
			p.Match(MyDSLParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPayloadContext is an interface to support dynamic dispatch.
type IPayloadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL_SIGN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	FieldsBody() IFieldsBodyContext

	// IsPayloadContext differentiates from other interfaces.
	IsPayloadContext()
}

type PayloadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPayloadContext() *PayloadContext {
	var p = new(PayloadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_payload
	return p
}

func (*PayloadContext) IsPayloadContext() {}

func NewPayloadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PayloadContext {
	var p = new(PayloadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_payload

	return p
}

func (s *PayloadContext) GetParser() antlr.Parser { return s.parser }

func (s *PayloadContext) EQUAL_SIGN() antlr.TerminalNode {
	return s.GetToken(MyDSLParserEQUAL_SIGN, 0)
}

func (s *PayloadContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyDSLParserIDENTIFIER, 0)
}

func (s *PayloadContext) FieldsBody() IFieldsBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsBodyContext)
}

func (s *PayloadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PayloadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PayloadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterPayload(s)
	}
}

func (s *PayloadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitPayload(s)
	}
}

func (p *MyDSLParser) Payload() (localctx IPayloadContext) {
	this := p
	_ = this

	localctx = NewPayloadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MyDSLParserRULE_payload)

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
		p.SetState(86)
		p.Match(MyDSLParserT__6)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MyDSLParserT__2:
		{
			p.SetState(87)
			p.Match(MyDSLParserT__2)
		}
		{
			p.SetState(88)
			p.FieldsBody()
		}
		{
			p.SetState(89)
			p.Match(MyDSLParserT__3)
		}

	case MyDSLParserEQUAL_SIGN:
		{
			p.SetState(91)
			p.Match(MyDSLParserEQUAL_SIGN)
		}
		{
			p.SetState(92)
			p.Match(MyDSLParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL_SIGN() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) EQUAL_SIGN() antlr.TerminalNode {
	return s.GetToken(MyDSLParserEQUAL_SIGN, 0)
}

func (s *DescriptionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MyDSLParserSTRING_LITERAL, 0)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (p *MyDSLParser) Description() (localctx IDescriptionContext) {
	this := p
	_ = this

	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MyDSLParserRULE_description)

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
		p.SetState(95)
		p.Match(MyDSLParserT__7)
	}
	{
		p.SetState(96)
		p.Match(MyDSLParserEQUAL_SIGN)
	}
	{
		p.SetState(97)
		p.Match(MyDSLParserSTRING_LITERAL)
	}

	return localctx
}

// IFieldsBodyContext is an interface to support dynamic dispatch.
type IFieldsBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsFieldsBodyContext differentiates from other interfaces.
	IsFieldsBodyContext()
}

type FieldsBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsBodyContext() *FieldsBodyContext {
	var p = new(FieldsBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_fieldsBody
	return p
}

func (*FieldsBodyContext) IsFieldsBodyContext() {}

func NewFieldsBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsBodyContext {
	var p = new(FieldsBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_fieldsBody

	return p
}

func (s *FieldsBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsBodyContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldsBodyContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldsBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterFieldsBody(s)
	}
}

func (s *FieldsBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitFieldsBody(s)
	}
}

func (p *MyDSLParser) FieldsBody() (localctx IFieldsBodyContext) {
	this := p
	_ = this

	localctx = NewFieldsBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MyDSLParserRULE_fieldsBody)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MyDSLParserT__8 {
		{
			p.SetState(99)
			p.Field()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	FieldType() IFieldTypeContext
	FieldBody() IFieldBodyContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) FieldType() IFieldTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldContext) FieldBody() IFieldBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldBodyContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *MyDSLParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MyDSLParserRULE_field)
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
		p.SetState(104)
		p.Match(MyDSLParserT__8)
	}
	{
		p.SetState(105)
		p.FieldName()
	}
	{
		p.SetState(106)
		p.FieldType()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MyDSLParserT__2 {
		{
			p.SetState(107)
			p.Match(MyDSLParserT__2)
		}
		{
			p.SetState(108)
			p.FieldBody()
		}
		{
			p.SetState(109)
			p.Match(MyDSLParserT__3)
		}

	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyDSLParserIDENTIFIER, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *MyDSLParser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MyDSLParserRULE_fieldName)

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
		p.SetState(113)
		p.Match(MyDSLParserIDENTIFIER)
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyDSLParserIDENTIFIER, 0)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *MyDSLParser) FieldType() (localctx IFieldTypeContext) {
	this := p
	_ = this

	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MyDSLParserRULE_fieldType)

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
		p.SetState(115)
		p.Match(MyDSLParserIDENTIFIER)
	}

	return localctx
}

// IFieldBodyContext is an interface to support dynamic dispatch.
type IFieldBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDescription() []IDescriptionContext
	Description(i int) IDescriptionContext
	AllValidation() []IValidationContext
	Validation(i int) IValidationContext
	AllRequired() []IRequiredContext
	Required(i int) IRequiredContext

	// IsFieldBodyContext differentiates from other interfaces.
	IsFieldBodyContext()
}

type FieldBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldBodyContext() *FieldBodyContext {
	var p = new(FieldBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_fieldBody
	return p
}

func (*FieldBodyContext) IsFieldBodyContext() {}

func NewFieldBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldBodyContext {
	var p = new(FieldBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_fieldBody

	return p
}

func (s *FieldBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldBodyContext) AllDescription() []IDescriptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDescriptionContext); ok {
			len++
		}
	}

	tst := make([]IDescriptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDescriptionContext); ok {
			tst[i] = t.(IDescriptionContext)
			i++
		}
	}

	return tst
}

func (s *FieldBodyContext) Description(i int) IDescriptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *FieldBodyContext) AllValidation() []IValidationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValidationContext); ok {
			len++
		}
	}

	tst := make([]IValidationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValidationContext); ok {
			tst[i] = t.(IValidationContext)
			i++
		}
	}

	return tst
}

func (s *FieldBodyContext) Validation(i int) IValidationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValidationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValidationContext)
}

func (s *FieldBodyContext) AllRequired() []IRequiredContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRequiredContext); ok {
			len++
		}
	}

	tst := make([]IRequiredContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRequiredContext); ok {
			tst[i] = t.(IRequiredContext)
			i++
		}
	}

	return tst
}

func (s *FieldBodyContext) Required(i int) IRequiredContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequiredContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequiredContext)
}

func (s *FieldBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterFieldBody(s)
	}
}

func (s *FieldBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitFieldBody(s)
	}
}

func (p *MyDSLParser) FieldBody() (localctx IFieldBodyContext) {
	this := p
	_ = this

	localctx = NewFieldBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MyDSLParserRULE_fieldBody)
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
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9472) != 0) {
		p.SetState(120)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MyDSLParserT__7:
			{
				p.SetState(117)
				p.Description()
			}

		case MyDSLParserT__12:
			{
				p.SetState(118)
				p.Validation()
			}

		case MyDSLParserT__9:
			{
				p.SetState(119)
				p.Required()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRequiredContext is an interface to support dynamic dispatch.
type IRequiredContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL_SIGN() antlr.TerminalNode

	// IsRequiredContext differentiates from other interfaces.
	IsRequiredContext()
}

type RequiredContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequiredContext() *RequiredContext {
	var p = new(RequiredContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_required
	return p
}

func (*RequiredContext) IsRequiredContext() {}

func NewRequiredContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequiredContext {
	var p = new(RequiredContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_required

	return p
}

func (s *RequiredContext) GetParser() antlr.Parser { return s.parser }

func (s *RequiredContext) EQUAL_SIGN() antlr.TerminalNode {
	return s.GetToken(MyDSLParserEQUAL_SIGN, 0)
}

func (s *RequiredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequiredContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequiredContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterRequired(s)
	}
}

func (s *RequiredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitRequired(s)
	}
}

func (p *MyDSLParser) Required() (localctx IRequiredContext) {
	this := p
	_ = this

	localctx = NewRequiredContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MyDSLParserRULE_required)
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
		p.SetState(124)
		p.Match(MyDSLParserT__9)
	}
	{
		p.SetState(125)
		p.Match(MyDSLParserEQUAL_SIGN)
	}
	{
		p.SetState(126)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MyDSLParserT__10 || _la == MyDSLParserT__11) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValidationContext is an interface to support dynamic dispatch.
type IValidationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL_SIGN() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode

	// IsValidationContext differentiates from other interfaces.
	IsValidationContext()
}

type ValidationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidationContext() *ValidationContext {
	var p = new(ValidationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_validation
	return p
}

func (*ValidationContext) IsValidationContext() {}

func NewValidationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidationContext {
	var p = new(ValidationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_validation

	return p
}

func (s *ValidationContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidationContext) EQUAL_SIGN() antlr.TerminalNode {
	return s.GetToken(MyDSLParserEQUAL_SIGN, 0)
}

func (s *ValidationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MyDSLParserSTRING_LITERAL, 0)
}

func (s *ValidationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterValidation(s)
	}
}

func (s *ValidationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitValidation(s)
	}
}

func (p *MyDSLParser) Validation() (localctx IValidationContext) {
	this := p
	_ = this

	localctx = NewValidationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MyDSLParserRULE_validation)

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
		p.SetState(128)
		p.Match(MyDSLParserT__12)
	}
	{
		p.SetState(129)
		p.Match(MyDSLParserEQUAL_SIGN)
	}
	{
		p.SetState(130)
		p.Match(MyDSLParserSTRING_LITERAL)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	TypeBody() ITypeBodyContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyDSLParserIDENTIFIER, 0)
}

func (s *TypeContext) TypeBody() ITypeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeBodyContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *MyDSLParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MyDSLParserRULE_type)

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
		p.SetState(132)
		p.Match(MyDSLParserT__13)
	}
	{
		p.SetState(133)
		p.Match(MyDSLParserIDENTIFIER)
	}
	{
		p.SetState(134)
		p.Match(MyDSLParserT__2)
	}
	{
		p.SetState(135)
		p.TypeBody()
	}
	{
		p.SetState(136)
		p.Match(MyDSLParserT__3)
	}

	return localctx
}

// ITypeBodyContext is an interface to support dynamic dispatch.
type ITypeBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsTypeBodyContext differentiates from other interfaces.
	IsTypeBodyContext()
}

type TypeBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBodyContext() *TypeBodyContext {
	var p = new(TypeBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyDSLParserRULE_typeBody
	return p
}

func (*TypeBodyContext) IsTypeBodyContext() {}

func NewTypeBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBodyContext {
	var p = new(TypeBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyDSLParserRULE_typeBody

	return p
}

func (s *TypeBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBodyContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeBodyContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TypeBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.EnterTypeBody(s)
	}
}

func (s *TypeBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyDSLListener); ok {
		listenerT.ExitTypeBody(s)
	}
}

func (p *MyDSLParser) TypeBody() (localctx ITypeBodyContext) {
	this := p
	_ = this

	localctx = NewTypeBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MyDSLParserRULE_typeBody)
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
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MyDSLParserT__8 {
		{
			p.SetState(138)
			p.Field()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
