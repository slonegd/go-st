package ast

import (
	"fmt"

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
	x.err = fmt.Errorf("some ReportAttemptingFullContext (TODO): %+v", dfa)
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
	err := &Error{
		Message: msg,
		Snippet: x.source.Snippet(line, column, 2),
	}
	x.err = err
}

func (x visitor) CheckError(ctx antlr.ParserRuleContext, err error) error {
	if err == nil {
		return nil
	}
	x.SyntaxError(nil, nil, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), nil)
	panic(err) // прекращаем парсинг, паника словится
}

type Error struct {
	Message string
	Snippet string
}

type Source []rune

func (e *Error) Error() string {
	if e.Snippet == "" {
		return e.Message
	}
	return fmt.Sprintf(e.Snippet, e.Message)
}
