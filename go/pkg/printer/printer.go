package printer

import (
	"fmt"
	"io"
)

type Printer struct {
	*Config
	Writer io.Writer

	Error  error
	Cursor Cursor

	indent lineIndent
}

func New(config *Config, writer io.Writer) *Printer {
	if config == nil {
		config = &Config{
			IndentWidth:          4,
			Indent:               0,
			MaxCharactersPerLine: 125,
		}
	}

	if config.IndentWidth == 0 {
		config.IndentWidth = 4
	}
	if config.MaxCharactersPerLine == 0 {
		config.MaxCharactersPerLine = 125
	}

	printer := &Printer{
		Config: config,
		Writer: writer,
		indent: lineIndent{
			Width: config.IndentWidth,
			Count: config.Indent,
		},
	}

	return printer
}

func (p *Printer) Reset(writer io.Writer) *Printer {
	if p != nil {
		p.Writer = writer
		p.Error = nil
		p.Cursor = Cursor{}
		p.indent = lineIndent{}
	}
	return p
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
	if p == nil || p.Error != nil {
		return p
	}

	p.PrintRaw("\n")
	p.Cursor.Line++
	p.Cursor.Column = 0

	return p
}

func (p *Printer) PrintBlankLine() *Printer {
	if p == nil || p.Error != nil {
		return p
	}

	if p.Cursor.Column > 0 {
		p.BreakLine()
	}

	return p.BreakLine()
}

func (p *Printer) IsNewLine() bool {
	if p.Error == nil {
		return p.Cursor.Column == 0
	}
	return false
}

func (p *Printer) PrintRaw(values ...interface{}) *Printer {
	if p == nil || p.Error != nil {
		return p
	}

	if n, err := p.Writer.Write([]byte(fmt.Sprint(p.normalize(values...)...))); err != nil {
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
	if p == nil || p.Error != nil {
		return p
	}

	if p.Cursor.Column > 0 {
		p.BreakLine()
	}

	p.PrintIndent()

	if n, err := p.Writer.Write([]byte(fmt.Sprint(p.normalize(values...)...))); err != nil {
		p.Error = err
	} else {
		p.Cursor.Column += n
	}

	return p
}

func (p *Printer) PrintTo(cursor Cursor) *Printer {
	if p == nil || p.Error != nil {
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

func (p *Printer) normalize(values ...interface{}) []interface{} {
	var vs []interface{}
	for _, v := range values {
		switch val := v.(type) {
		case *string:
			vs = append(vs, *val)
		case *bool:
			vs = append(vs, *val)
		case *int:
			vs = append(vs, *val)
		case *uint:
			vs = append(vs, *val)
		case *int8:
			vs = append(vs, *val)
		case *uint8:
			vs = append(vs, *val)
		case *int16:
			vs = append(vs, *val)
		case *uint16:
			vs = append(vs, *val)
		case *int32:
			vs = append(vs, *val)
		case *uint32:
			vs = append(vs, *val)
		case *int64:
			vs = append(vs, *val)
		case *uint64:
			vs = append(vs, *val)
		case *float32:
			vs = append(vs, *val)
		case *float64:
			vs = append(vs, *val)
		default:
			vs = append(vs, val)
		}
	}
	return vs
}
