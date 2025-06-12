package ast

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func (x visitor) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	if x.err != nil {
		return
	}
	x.err = fmt.Errorf("some ReportAmbiguity (TODO)")
}

func (x visitor) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	if x.err != nil {
		return
	}
	x.err = fmt.Errorf("some ReportAttemptingFullContext (TODO)")
}

func (x visitor) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	if x.err != nil {
		return
	}
	x.err = fmt.Errorf("some ReportContextSensitivity (TODO)")
}

func (x visitor) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line int, column int, msg string, e antlr.RecognitionException) {
	if x.err != nil {
		return
	}
	err := &Error{Line: line, Column: column, Message: msg}
	err.AddSnippet(x.source)
	x.err = err
}

type Error struct {
	Line, Column int
	Message      string
	Snippet      string
}

type Source []rune

func (e *Error) Error() string {
	return e.format()
}

func (e *Error) AddSnippet(source Source) *Error {
	lines, messageLine, ok := source.FindLines(e.Line)
	if !ok {
		return e
	}
	var snippet string
	for i, line := range lines {
		if i == 0 {
			snippet = "\n"
		}
		snippet += line
		if i != messageLine {
			snippet += "\n"
			continue
		}
		for range e.Column { // TODO учесть табуляцую
			snippet += "."
		}
		snippet += "^    %s\n" // форматер для сообщения
	}

	e.Snippet = snippet
	return e
}

func (e *Error) format() string {
	if e.Snippet == "" {
		return e.Message
	}
	return fmt.Sprintf(e.Snippet, e.Message)
}

// поиск строк с несколькими строками до и после для контекста
// а также вставка пустой строки для сообщения
// вставка номера строки
func (s Source) FindLines(line int) ([]string, int, bool) {
	if s == nil {
		return nil, 0, false
	}
	line -= 1 // массивы с нуля
	lines := strings.Split(string(s), "\n")
	startLine := 0
	if line > 2 {
		startLine = line - 2 // захватить 2 строки до для контекста
	}
	endLine := len(lines)
	if line < endLine-2 {
		endLine = line + 2 // захватить 2 строки после для контекста
	}
	messageLine := line - startLine + 1
	r := slices.Clone(lines[startLine : line+1])
	r = append(r, "")
	r = append(r, lines[line+1:endLine]...)
	positionFormatter := "%0" + strconv.Itoa(powerOf10(endLine)) + "d|"
	for i := range r {
		switch {
		case i < messageLine:
			r[i] = fmt.Sprintf(positionFormatter, i+startLine+1) + r[i]
		case i == messageLine:
			r[i] = fmt.Sprintf(positionFormatter, 0) + r[i]
		case i > messageLine:
			r[i] = fmt.Sprintf(positionFormatter, i+startLine) + r[i]
		}
	}
	return r, messageLine, true
}

func powerOf10(n int) int {
	r := 1
	for n > 0 {
		n /= 10
		r++
	}
	return r
}
