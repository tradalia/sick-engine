//=============================================================================
/*
Copyright © 2025 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tradalia/sick-engine/core/ast"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================

func main() {
	filename := os.Args[1]
	res := ParseFile(filename)
	if res.Errors != nil {
		for _, err := range res.Errors {
			fmt.Println(err)
		}
		return
	}

	println(res.Script)
}

//=============================================================================
//===
//=== Parsing functions
//===
//=============================================================================

func ParseFile(filename string) *ParseResult {
	stream,err := antlr.NewFileStream(filename)
	if err != nil {
		pr := &ParseResult{}
		pe := NewParseError(-1, -1, err.Error())
		pr.AddError(pe)
		return pr
	}

	return parse(stream)
}

//=============================================================================

func ParseString(input string) *ParseResult {
	return parse(antlr.NewInputStream(input))
}

//=============================================================================

func ParseReader(r io.Reader) *ParseResult {
	return parse(antlr.NewIoStream(r))
}

//=============================================================================

func parse(input antlr.CharStream) *ParseResult {
	lexer  := parser.NewTslLexer(input)
	stream := antlr .NewCommonTokenStream(lexer,0)
	p      := parser.NewTslParser(stream)

	res := &ParseResult{}
	lis := NewParseErrorListener(res)

	//--- Add error collection listener
	p.RemoveErrorListeners()
	p.AddErrorListener(lis)

	//--- Ask the parser to report all ambiguities
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
	tree := p.Script()

	script := ast.Build(tree)
	res.Script = script
	return res
}

//=============================================================================
//===
//=== Parse result
//===
//=============================================================================

type ParseResult struct {
	Script *ast.Script
	Errors []*ParseError
}

//=============================================================================

func (pr *ParseResult) AddError(pe *ParseError) {
	pr.Errors = append(pr.Errors, pe)
}

//=============================================================================

type ParseError struct {
	Line    int
	Column  int
	Error string
}

//=============================================================================

func NewParseError(line int, column int, err string) *ParseError {
	return &ParseError{line, column, err}
}

//=============================================================================

func (pe *ParseError) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Line:%d, Col:%d, Error: %s", pe.Line, pe.Column, pe.Error))
	return sb.String()
}

//=============================================================================
//===
//=== Error listener
//===
//=============================================================================

type ParseErrorListener struct {
	result *ParseResult
}

//=============================================================================

func NewParseErrorListener(result *ParseResult) *ParseErrorListener {
	return &ParseErrorListener{
		result: result,
	}
}

//=============================================================================

func (l *ParseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	pe := NewParseError(line,column, msg)
	l.result.AddError(pe)
}

//=============================================================================

func (l *ParseErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	println("Ambiguity")
}

//=============================================================================

func (l *ParseErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	println("Attempting full context")
}

//=============================================================================

func (l *ParseErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	println("Context sensitivity")
}

//=============================================================================
