package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const md1 = `
# Header

This is a paragraph.
`

func TestMarkdown_Parse(t *testing.T) {
	doc, err := ParseMarkdown([]byte(md1))
	assert.NoError(t, err)

	assert.Equal(t, 2, len(doc.Blocks))

	header := doc.Blocks[0].GetHeader()
	assert.Equal(t, int64(1), header.Level)
	assert.Equal(t, 1, len(header.Text))

	headerContent, _ := FormatInlines(header.Text...)
	assert.Equal(t, "Header", headerContent)

	paragraph, _ := FormatBlock(doc.Blocks[1])
	assert.Equal(t, "This is a paragraph.", paragraph)
}
