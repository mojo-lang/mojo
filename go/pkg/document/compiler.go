package document

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/document/compiler"
	"strings"
)
import "github.com/mojo-lang/mojo/go/pkg/openapi"

type Compiler struct {
	Documents Documents

	Path     string
	Package  *lang.Package
	OpenAPIs *openapi.OpenAPIs
}

// compile openapi to document of markdown style
func NewCompiler(path string, pkg *lang.Package, apis *openapi.OpenAPIs) *Compiler {
	return &Compiler{
		Path:      path,
		Package:   pkg,
		OpenAPIs:  apis,
		Documents: make(Documents),
	}
}

func (c *Compiler) Compile() (Documents, error) {
	apiCompiler := &compiler.ApiCompiler{}
	for name, api := range c.OpenAPIs.APIs {
		if !strings.HasPrefix(name, c.Package.FullName) {
			continue
		}

		document, err := apiCompiler.Compile(c.Package, api)
		if err != nil {
			return nil, err
		}
		c.Documents[lang.TypeNameToFileName(name)] = document
	}

	schemaCompiler := &compiler.SchemaCompiler{
		Components: c.OpenAPIs.Components,
	}
	for name, schema := range c.OpenAPIs.Components.Schemas {
		if schema == nil {
			//FIXME make sure scheme here should NOT be nil
			continue
		}

		if !strings.HasPrefix(name, c.Package.FullName) {
			continue
		}

		decl := c.Package.Scope.GetIdentifier(name).GetDeclaration()
		document, err := schemaCompiler.Compile(decl, schema)
		if err != nil {
			return nil, err
		}

		name = lang.TypeNameToFileName(name)
		c.Documents[name] = document
	}

	return c.Documents, nil
}
