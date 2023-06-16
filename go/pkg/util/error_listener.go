package util

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/core/go/pkg/logs"
)

type ErrorListener struct {
	*antlr.DiagnosticErrorListener

	Errors ParseErrors

	FileName   string
	Diagnosing bool
}

func NewErrorListener(fileName string, diagnosing bool) *ErrorListener {
	return &ErrorListener{
		DiagnosticErrorListener: antlr.NewDiagnosticErrorListener(true),
		FileName:                fileName,
		Diagnosing:              diagnosing,
	}
}

func (m *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	_ = recognizer
	_ = offendingSymbol
	_ = e
	m.Errors = append(m.Errors, NewParseError(m.FileName, int64(line), int64(column), msg))
	logs.Errorw(msg, "file", m.FileName, "line", line, "column", column)
}

func (m *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	_ = stopIndex
	if m.Diagnosing {
		m.DiagnosticErrorListener.ReportAmbiguity(recognizer, dfa, startIndex, startIndex, exact, ambigAlts, configs)
	}
}

func (m *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	if m.Diagnosing {
		m.DiagnosticErrorListener.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (m *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	if m.Diagnosing {
		m.DiagnosticErrorListener.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
	}
}
