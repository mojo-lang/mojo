package document

import "github.com/mojo-lang/mojo/go/pkg/mojo/core"

const (
	StrongTypeFullName = "mojo.document.Strong"
)

func NewTextInline(text string) *Inline {
	return &Inline{Inline: &Inline_Text{Text: &Text{Val: text}}}
}

func NewCodeInline(code *Code) *Inline {
	return &Inline{Inline: &Inline_Code{Code: code}}
}

func NewCodeInlineFrom(code string) *Inline {
	return NewCodeInline(&Code{
		Attribute: NewAttribute(),
		Content:   code,
	})
}

func NewLinkInline(link *Link) *Inline {
	return &Inline{Inline: &Inline_Link{Link: link}}
}

func NewLinkInlineFrom(target *core.Url, description ...*Inline) *Inline {
	return NewLinkInline(&Link{
		Attribute:   NewAttribute(),
		Description: description,
		Target:      &Target{Url: target},
	})
}

func NewImageInline(image *Image) *Inline {
	return &Inline{Inline: &Inline_Image{Image: image}}
}

func NewImageInlineFrom(target *core.Url, description ...*Inline) *Inline {
	return NewImageInline(&Image{
		Attribute:   NewAttribute(),
		Description: description,
		Target:      &Target{Url: target},
	})
}

func NewEmphasizedInline(values ...*Inline) *Inline {
	return &Inline{Inline: &Inline_Emphasized{Emphasized: &Emphasized{Vals: values}}}
}

func NewStrongInline(values ...*Inline) *Inline {
	return &Inline{Inline: &Inline_Strong{Strong: &Strong{Vals: values}}}
}

func (x *Inline) NewInline() *Inline {
	return &Inline{}
}

func (x *Inline) SetText(text string) *Inline {
	if x != nil {
		x.Inline = &Inline_Text{NewText(text)}
	}
	return x
}
