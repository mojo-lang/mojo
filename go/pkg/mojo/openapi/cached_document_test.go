package openapi

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/document"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCachedDocumentConversion(t *testing.T) {
	assert.Nil(t, (*CachedDocument)(nil).GetDocument())
	assert.Empty(t, (*CachedDocument)(nil).GetStringDocument())
	cached := &CachedDocument{Cache: "# Title\n\nBody\n"}
	assert.Equal(t, cached.Cache, cached.GetStringDocument())
	doc := cached.GetDocument()
	require.NotNil(t, doc)
	assert.Same(t, doc, cached.GetDocument())
	require.Len(t, doc.Blocks, 2)
	assert.Equal(t, "Title", doc.Blocks[0].GetHeader().GetText()[0].GetText().GetVal())
	doc.AppendParagraph(document.NewTextInline("Added"))
	rendered := cached.GetStringDocument()
	assert.Contains(t, rendered, "Added")
	assert.Equal(t, rendered, cached.Cache)
	assert.Same(t, doc, cached.GetDocument())
	assert.Empty(t, (&CachedDocument{}).GetStringDocument())
	require.NotNil(t, (&CachedDocument{}).GetDocument())
}
