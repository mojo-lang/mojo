package document

import (
	"bytes"
	"strings"
)

func NewCodeBlock(lang string, lines ...string) *CodeBlock {
	codeBlock := &CodeBlock{Language: lang}
	for _, line := range lines {
		ls := strings.Split(line, "\n")
		for _, l := range ls {
			codeBlock.Lines = append(codeBlock.Lines, &Line{
				Vals: []*Inline{NewTextInline(l)},
			})
		}
	}
	return codeBlock
}

func (x *CodeBlock) AppendLine(line string) {
	if x != nil {
		ls := strings.Split(line, "\n")
		for _, l := range ls {
			x.Lines = append(x.Lines, &Line{
				Vals: []*Inline{NewTextInline(l)},
			})
		}
	}
}

func (x *CodeBlock) GetCode() []byte {
	if x == nil {
		return []byte{}
	}

	buffer := bytes.Buffer{}
	for i, line := range x.Lines {
		if i > 0 {
			buffer.WriteByte('\n')
		}

		for _, inline := range line.Vals {
			if text := inline.GetText(); text != nil {
				buffer.WriteString(text.Val)
			}
		}
	}

	return buffer.Bytes()
}
