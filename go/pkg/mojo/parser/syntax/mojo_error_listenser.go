package syntax

import (
    "fmt"
    "github.com/antlr/antlr4/runtime/Go/antlr"
    "github.com/mojo-lang/core/go/pkg/logs"
)

type MojoErrorListener struct {
    *antlr.DiagnosticErrorListener

    Error      error
    FileName   string
    Diagnosing bool
}

func NewMojoErrorListener() *MojoErrorListener {
    return &MojoErrorListener{
        DiagnosticErrorListener: antlr.NewDiagnosticErrorListener(true),
    }
}

func (m *MojoErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
    m.Error = fmt.Errorf("file: %s, line: %d, column: %d, message: %s %w", m.FileName, line, column, msg, e)
    logs.Errorw(msg, "file", m.FileName, "line", line, "column", column)
}

func (m *MojoErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
    if m.Diagnosing {
        m.DiagnosticErrorListener.ReportAmbiguity(recognizer, dfa, startIndex, startIndex, exact, ambigAlts, configs)
    }
}

func (m *MojoErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
    if m.Diagnosing {
        m.DiagnosticErrorListener.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
    }
}

func (m *MojoErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
    if m.Diagnosing {
        m.DiagnosticErrorListener.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
    }
}
