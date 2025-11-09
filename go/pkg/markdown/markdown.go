package markdown

import (
	"bytes"
	"context"
	"io"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/stephenafamo/goldmark-pdf"

	"github.com/mojo-lang/mojo/go/pkg/mojo/document"
)

func RenderToString(doc *document.Document) (string, error) {
	md := New()
	str, err := md.RenderToString(doc)
	if err != nil {
		logs.Warnw("failed to render markdown to string document", "error", err)
		return "", err
	}
	return str, nil
}

type Markdown struct {
	*Parser
	*Renderer
}

func New() *Markdown {
	return &Markdown{
		Parser:   NewParser(),
		Renderer: &Renderer{},
	}
}

func (m *Markdown) ConvertToHtml(source []byte, w io.Writer) error {
	return m.MdEngine.Convert(source, w)
}

func (m *Markdown) ConvertToPdf(source []byte, w io.Writer) error {
	renderer := m.MdEngine.Renderer()
	m.MdEngine.SetRenderer(pdf.New(
		// pdf.WithTraceWriter(os.Stdout),
		pdf.WithContext(context.Background()),
		// pdf.WithImageFS(os.DirFS(".")),
		// pdf.WithLinkColor("cc4578"),
		pdf.WithHeadingFont(pdf.GetTextFont("IBM Plex Serif", pdf.FontLora)),
		pdf.WithBodyFont(pdf.GetTextFont("Open Sans", pdf.FontRoboto)),
		pdf.WithCodeFont(pdf.GetCodeFont("Inconsolata", pdf.FontRobotoMono)),
	))
	m.MdEngine.Convert(source, w)
	m.MdEngine.SetRenderer(renderer)
	return nil
}

func (m *Markdown) Parse(content string) (*document.Document, error) {
	return m.ParseBytes([]byte(content))
}

func (m *Markdown) Render(writer io.Writer, doc *document.Document) error {
	return m.RenderDocument(doc, writer)
}

func (m *Markdown) RenderBlocks(writer io.Writer, blocks ...*document.Block) error {
	doc := &document.Document{Blocks: blocks}
	return m.Render(writer, doc)
}

func (m *Markdown) RenderInlines(writer io.Writer, inlines ...*document.Inline) error {
	return m.RenderBlocks(writer, document.NewPainBlock(inlines...))
}

func (m *Markdown) RenderToString(doc *document.Document) (string, error) {
	buffer := &bytes.Buffer{}
	err := m.Render(buffer, doc)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (m *Markdown) RenderBlocksToString(blocks ...*document.Block) (string, error) {
	buffer := &bytes.Buffer{}
	err := m.RenderBlocks(buffer, blocks...)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (m *Markdown) RenderInlinesToString(inlines ...*document.Inline) (string, error) {
	buffer := &bytes.Buffer{}
	err := m.RenderInlines(buffer, inlines...)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (m *Markdown) RenderHtml(doc *document.Document, writer io.Writer) error {
	buffer := &bytes.Buffer{}
	err := m.Render(buffer, doc)
	if err != nil {
		return err
	}

	return m.ConvertToHtml(buffer.Bytes(), writer)
}

func (m *Markdown) RenderPdf(doc *document.Document, writer io.Writer) error {
	buffer := &bytes.Buffer{}
	err := m.Render(buffer, doc)
	if err != nil {
		return err
	}

	return m.ConvertToPdf(buffer.Bytes(), writer)
}
