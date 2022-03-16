package document

import (
    "github.com/mojo-lang/document/go/pkg/markdown"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
)

type Generator struct {
    Documents Documents
    Files     []*util.GeneratedFile
}

func NewGenerator(documents Documents) *Generator {
    return &Generator{
        Documents: documents,
    }
}

func (g *Generator) Generate(output string) error {
    if err := g.generateMarkdown(); err != nil {
        return err
    }

    if err := g.writeFiles(output); err != nil {
        return err
    }

    return nil
}

func (g *Generator) generateMarkdown() error {
    md := markdown.New()
    for name, document := range g.Documents {
        m, err := md.RenderToString(document)
        if err != nil {
            return err
        }
        g.Files = append(g.Files, &util.GeneratedFile{
            Name:    name + ".md",
            Content: m,
        })
    }
    return nil
}

func (g *Generator) writeFiles(dir string) error {
    guard := &util.PathGuard{
        Suffixes: []string{".md"},
    }

    for _, file := range g.Files {
        if err := file.WriteTo(dir, guard); err != nil {
            return err
        }
    }

    return nil
}
