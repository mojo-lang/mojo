package format

import (
    "fmt"
    "io"
)

const needIndent = "need-indent"

// A Config node controls the output.
type Config struct {
    IndentWidth          int // default: 4
    Indent               int // default: 0
    MaxCharactersPerLine int // default: 125 https://hilton.org.uk/blog/source-code-line-length
}

const maxSpace = 128

var spaceIndents []byte

func init() {
    spaceIndents = make([]byte, 0, maxSpace)
    for i := 0; i < maxSpace; i++ {
        spaceIndents = append(spaceIndents, ' ')
    }
}

type lineIndent struct {
    Width int
    Count int
}

func (l *lineIndent) Indent() {
    l.Count++
}

func (l *lineIndent) Outdent() {
    if l.Count > 0 {
        l.Count--
    }
}

func (l *lineIndent) IndentSpace() []byte {
    return spaceIndents[0 : l.Count*l.Width]
}

type Cursor struct {
    Line   int
    Column int
}

type Printer struct {
    Config
    Writer io.Writer

    Error  error
    Cursor Cursor
    //Offset int

    indent  lineIndent
    columns []int
}

func NewPrinter(config Config, writer io.Writer) *Printer {
    printer := &Printer{
        Config: config,
        Writer: writer,
        indent: lineIndent{
            Width: config.IndentWidth,
        },
    }

    if printer.Config.IndentWidth == 0 {
        printer.Config.IndentWidth = 4
        printer.indent.Width = 4
    }

    return printer
}

func (p *Printer) GetIndent() int {
    return p.indent.Width * p.indent.Count
}

func (p *Printer) Indent() *Printer {
    if p != nil {
        p.indent.Indent()
    }
    return p
}

func (p *Printer) Outdent() *Printer {
    if p != nil {
        p.indent.Outdent()
    }
    return p
}

func (p *Printer) BreakLine() *Printer {
    if p.Error == nil {
        p.PrintRaw("\n")

        p.columns = append(p.columns, p.Cursor.Column)
        p.Cursor.Line++
        p.Cursor.Column = 0
    }
    return p
}

func (p *Printer) IsNewLine() bool {
    if p.Error == nil {
        return p.Cursor.Column == 0
    }
    return false
}

func (p *Printer) PrintRaw(values ...interface{}) *Printer {
    if n, err := p.Writer.Write([]byte(fmt.Sprint(values...))); err != nil {
        p.Error = err
    } else {
        p.Cursor.Column += n
    }
    return p
}

func (p *Printer) PrintIndent() *Printer {
    if p == nil || p.Error != nil {
        return p
    }

    if n, err := p.Writer.Write(p.indent.IndentSpace()); err != nil {
        p.Error = err
        return p
    } else {
        p.Cursor.Column += n
    }
    return p
}

func (p *Printer) PrintLine(values ...interface{}) *Printer {
    if p.Cursor.Column > 0 {
        p.BreakLine()
    }

    p.PrintIndent()

    if n, err := p.Writer.Write([]byte(fmt.Sprint(values...))); err != nil {
        p.Error = err
    } else {
        p.Cursor.Column += n
    }

    return p
}

func (p *Printer) PrintTo(cursor Cursor) *Printer {
    if p.Error != nil {
        return p
    }

    dLine := cursor.Line - p.Cursor.Line
    if dLine == 0 {
        dColumn := cursor.Column - p.Cursor.Column
        if dColumn > 0 {
            p.PrintRaw(string(spaceIndents[0:dColumn]))
        }
    } else if dLine > 0 {
        for i := 0; i < dLine; i++ {
            p.BreakLine()
        }
        p.PrintRaw(string(spaceIndents[0:cursor.Column]))
    }

    return p
}
