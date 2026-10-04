package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCodeBlock(t *testing.T) {
	for _, tt := range []struct {
		name  string
		lines []string
		want  []string
		code  string
	}{
		{"no lines", nil, nil, ""},
		{"empty line", []string{""}, []string{""}, ""},
		{"separate lines", []string{"first", "second"}, []string{"first", "second"}, "first\nsecond"},
		{"embedded newlines", []string{"first\n\nthird\n", "last"}, []string{"first", "", "third", "", "last"}, "first\n\nthird\n\nlast"},
		{"unicode and indentation", []string{"  你好\n\tworld"}, []string{"  你好", "\tworld"}, "  你好\n\tworld"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			block := NewCodeBlock("go", tt.lines...)
			assert.Equal(t, "go", block.Language)
			require.Len(t, block.Lines, len(tt.want))
			for i, want := range tt.want {
				require.Len(t, block.Lines[i].Vals, 1)
				assert.Equal(t, want, block.Lines[i].Vals[0].GetText().GetVal())
			}
			assert.Equal(t, tt.code, string(block.GetCode()))
		})
	}
}

func TestCodeBlockAppendLine(t *testing.T) {
	var block CodeBlock
	block.AppendLine("first\nsecond")
	block.AppendLine("")
	block.AppendLine("last\n")
	assert.Equal(t, "first\nsecond\n\nlast\n", string(block.GetCode()))
	assert.Len(t, block.Lines, 5)
	assert.NotPanics(t, func() { (*CodeBlock)(nil).AppendLine("ignored") })
	assert.Empty(t, (*CodeBlock)(nil).GetCode())
}

func TestCodeBlockGetCodeMixedInlines(t *testing.T) {
	block := &CodeBlock{Lines: []*Line{
		NewLine(NewTextInline("hello"), nil, NewCodeInlineFrom("ignored"), NewTextInline(" world")),
		{},
		NewLine(NewTextInline("end")),
	}}
	assert.Equal(t, "hello world\n\nend", string(block.GetCode()))
	assert.Equal(t, "hello world\n\nend", string(block.GetCode()), "reading must preserve content")
	block.Lines = []*Line{NewLine(NewTextInline("before")), nil, NewLine(NewTextInline("after"))}
	require.NotPanics(t, func() { assert.Equal(t, "before\n\nafter", string(block.GetCode())) })
}
