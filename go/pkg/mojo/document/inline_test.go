package document

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInline_SetText(t *testing.T) {
	foo := NewTextInline("").SetText("foo")
	assert.Equal(t, "foo", foo.GetText().Val)
	link := NewLinkInline(&Link{})
	assert.Same(t, link, link.SetText("replacement"))
	assert.Nil(t, link.GetLink())
	assert.Equal(t, "replacement", link.GetText().GetVal())
	assert.Equal(t, "", link.SetText("").GetText().GetVal())
	assert.Nil(t, (*Inline)(nil).SetText("ignored"))
}

func TestNewTextInline(t *testing.T) {
	for _, text := range []string{"", "hello", "你好\nworld"} {
		inline := NewTextInline(text)
		require.NotNil(t, inline.GetText())
		assert.Equal(t, text, inline.GetText().Val)
	}
}

func TestNewCodeInline(t *testing.T) {
	code := &Code{Content: "fmt.Println(1)"}
	assert.Same(t, code, NewCodeInline(code).GetCode())
	assert.IsType(t, &Inline_Code{}, NewCodeInline(nil).Inline)
	for _, content := range []string{"", "first\nsecond"} {
		inline := NewCodeInlineFrom(content)
		require.NotNil(t, inline.GetCode())
		assert.Equal(t, content, inline.GetCode().Content)
		require.NotNil(t, inline.GetCode().Attribute)
		assert.NotNil(t, inline.GetCode().Attribute.Properties)
	}
}

func TestNewLinkAndImageInline(t *testing.T) {
	url, err := core.NewUrl("https://example.com/docs")
	require.NoError(t, err)
	description := []*Inline{NewTextInline("first"), NewStrongInline(NewTextInline("second"))}
	link := &Link{Description: description, Target: &Target{Url: url}}
	image := &Image{Description: description, Target: &Target{Url: url}}
	assert.Same(t, link, NewLinkInline(link).GetLink())
	assert.Same(t, image, NewImageInline(image).GetImage())
	assert.IsType(t, &Inline_Link{}, NewLinkInline(nil).Inline)
	assert.IsType(t, &Inline_Image{}, NewImageInline(nil).Inline)
	createdLink := NewLinkInlineFrom(url, description...).GetLink()
	createdImage := NewImageInlineFrom(url, description...).GetImage()
	for _, attribute := range []*Attribute{createdLink.Attribute, createdImage.Attribute} {
		require.NotNil(t, attribute)
		assert.NotNil(t, attribute.Properties)
	}
	assert.Equal(t, description, createdLink.Description)
	assert.Equal(t, description, createdImage.Description)
	assert.Same(t, url, createdLink.Target.Url)
	assert.Same(t, url, createdImage.Target.Url)
	assert.Empty(t, NewLinkInlineFrom(nil).GetLink().Description)
	assert.Nil(t, NewLinkInlineFrom(nil).GetLink().Target.Url)
	assert.Empty(t, NewImageInlineFrom(nil).GetImage().Description)
	assert.Nil(t, NewImageInlineFrom(nil).GetImage().Target.Url)
}

func TestNewFormattingInline(t *testing.T) {
	values := []*Inline{NewTextInline("first"), NewCodeInlineFrom("second")}
	assert.Equal(t, values, NewEmphasizedInline(values...).GetEmphasized().Vals)
	assert.Equal(t, values, NewStrongInline(values...).GetStrong().Vals)
	assert.Empty(t, NewEmphasizedInline().GetEmphasized().Vals)
	assert.Empty(t, NewStrongInline().GetStrong().Vals)
}

func TestInlineNewInline(t *testing.T) {
	original := NewTextInline("preserved")
	created := original.NewInline()
	require.NotNil(t, created)
	assert.NotSame(t, original, created)
	assert.Nil(t, created.Inline)
	assert.Equal(t, "preserved", original.GetText().Val)
	require.NotNil(t, (*Inline)(nil).NewInline())
}
