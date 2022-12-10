package printer

import (
	"bytes"

	"github.com/mojo-lang/mojo/go/pkg/printer"
)

// Printer is the type whose methods generate the output, stored in the associated response structure.
type Printer struct {
	*printer.Printer
	Buffer *bytes.Buffer
}

// New creates a new printer for print protobuf.
func New(config *printer.Config) *Printer {
	p := new(Printer)
	p.Buffer = new(bytes.Buffer)
	p.Printer = printer.New(config, p.Buffer)
	return p
}

func (p *Printer) Reset() {
	if p != nil {
		p.Buffer = new(bytes.Buffer)
		p.Printer.Reset(p.Buffer)
	}
}

// SetError reports a problem, including an error, and exits the program.
func (p *Printer) SetError(err error, msgs ...string) {
	// s := fmt.Errorf() strings.Join(msgs, " ") + ":" + err.Error()
	// logs.Warnw("protobuf print failed", "error", s)
	// p.Error =
}

// deprecationComment is the standard comment added to deprecated
// messages, fields, enums, and enum values.
var deprecationComment = "// Deprecated: Do not use."

// PrintComments prints any comments from the source .proto file.
// The path is a comma-separated list of integers.
// It returns an indication of whether any comments were printed.
// See descriptor.proto for its format.
func (p *Printer) PrintComments(path string) bool {
	if c, ok := p.makeComments(path); ok {
		p.PrintLine(c)
		return true
	}
	return false
}

// makeComments generates the comment string for the field, no "\n" at the end
func (p *Printer) makeComments(path string) (string, bool) {
	// loc, ok := p.file.Comments[path]
	// if !ok {
	// return "", false
	// }
	w := new(bytes.Buffer)
	// nl := ""
	// for _, line := range strings.Split(strings.TrimSuffix(loc.GetLeadingComments(), "\n"), "\n") {
	//    fmt.Fprintf(w, "%s//%s", nl, line)
	//    nl = "\n"
	// }
	return w.String(), true
}
