package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.com/dvictor/api-design/parser"
)

func main() {
	dsl := `
service MyService {
	method MyMethod {
		payload {
			field MyField T {
				description = ""
				required
			}
		}
	}
}
`
	src := "met"
	cursorPosition := strings.Index(dsl, src) + len(src)

	//find word start
	wStart := findWordStart(dsl, cursorPosition)
	fragment := dsl[wStart:cursorPosition]
	cursorPosition = wStart

	//consume whitespace
	rest := dsl[cursorPosition:]
	wsLen := len(rest) - len(strings.TrimLeft(rest, " \t\r\n"))
	cursorPosition += wsLen

	getFieldTypesForCompletion(dsl, cursorPosition)
	println("fragment", fragment)
}

func findWordStart(s string, index int) int {
	if index == 0 {
		return 0
	}
	for i := index - 1; i >= 0; i-- {
		if unicode.IsSpace(rune(s[i])) || unicode.IsPunct(rune(s[i])) {
			return i
		}
	}
	return index
}

func getFieldTypesForCompletion(dsl string, cursorPosition int) {
	input := antlr.NewInputStream(dsl)
	lexer := parser.NewMyDSLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMyDSLParser(stream)

	// Implement a custom error listener to suppress syntax errors
	// (since the DSL might be incomplete while typing)
	errorListener := &errorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	nextTokensByRule := make(map[string][]string)

	atn := p.GetATN()
	for _, state := range atn.DecisionToState {
		next := atn.NextTokensNoContext(state).GetIntervals()
		//fmt.Println(intervalsToTokens(next, p.GetLiteralNames()))
		ruleName := p.RuleNames[state.GetRuleIndex()]
		if strings.HasSuffix(ruleName, "Body") {
			tokens := intervalsToTokens(next, p.GetLiteralNames())
			if len(tokens) > 0 {
				nextTokensByRule[p.RuleNames[state.GetRuleIndex()]] = tokens
				fmt.Println(ruleName, tokens)
			}
		}
		//println()
	}

	tree := p.ApiBody()

	//lookup the result from the rule with the smallest characters span
	span := len(dsl)
	var result []string

	listener := &completionListener{
		callback: func(ctx antlr.ParserRuleContext) {
			start, end := ctx.GetStart().GetStart(), ctx.GetStop().GetStop()
			if cursorPosition < start || cursorPosition > end {
				return
			}
			rule := p.RuleNames[ctx.GetRuleIndex()]
			tokens := nextTokensByRule[rule]
			if len(tokens) == 0 {
				return
			}
			crtSpan := end - start
			if crtSpan >= span {
				return
			}
			span = crtSpan
			result = tokens

			println("rule:", rule)
			fmt.Println(tokens)
			println(ctx.GetStart().GetText(), "...", ctx.GetStop().GetText())
			println(cursorPosition, start, "->", end)

			println()
		},
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	fmt.Println("result", result)
}

func intervalsToTokens(is []*antlr.Interval, literalNames []string) []string {
	var names []string
	for _, v := range is {
		for j := v.Start; j < v.Stop; j++ {
			if j > 0 && j < len(literalNames) {
				names = append(names, literalNames[j])
			}
		}
	}
	return names
}

type completionListener struct {
	*parser.BaseMyDSLListener
	callback func(ctx antlr.ParserRuleContext)
}

func (l *completionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	l.callback(ctx)
}

//func (l *completionListener) EnterFieldName(ctx *parser.FieldNameContext) {
//	println("enter:", ctx.GetStop().GetText())
//	//println("cursor:", l.cursorPosition)
//	//println("token:", ctx.GetStop().GetStop())
//
//	p := ctx.GetParser()
//	state := ctx.GetInvokingState()
//	p.SetState(state)
//	println(p.GetParserRuleContext())
//	println(ctx)
//	msg := p.GetExpectedTokens().StringVerbose(p.GetLiteralNames(), p.GetSymbolicNames(), false)
//	println(msg)
//
//	// Check if the fieldType is missing (incomplete DSL)
//	//if ctx.FieldType() == nil {
//	// Get the token position of the equal sign
//	/*	equalSignToken := ctx.GetToken(parser.MyDSLLexerEQUAL_SIGN, 0).GetSymbol()
//		equalSignPosition := equalSignToken.GetStop()
//		println(equalSignPosition, l.cursorPosition)
//
//		// Check if the cursor is at the equal sign position
//		if l.cursorPosition == equalSignPosition+1 {
//			l.completions = []string{"String", "Int", "Bool", "SoftwareTitle"}
//		}
//		//}*/
//}

type errorListener struct {
	antlr.ErrorListener
}

func (el *errorListener) SyntaxError(
	recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException) {
	// Suppress syntax errors
	fmt.Printf("line %d col %d: %s\n", line, column, msg)
}
