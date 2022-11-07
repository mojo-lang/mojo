package printer

import (
	"fmt"
	"io"
)

type Printer struct {
	Config
	Writer io.Writer

	Error  error
	Cursor Cursor

	indent lineIndent
}

func New(config Config, writer io.Writer) *Printer {
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
