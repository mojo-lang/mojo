package _go

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/go/compiler"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Compiler struct {
	Path    string
	Package *lang.Package
	Files   []*descriptor.FileDescriptor
}

func (c *Compiler) Compile() (util.CodeGeneratedFiles, error) {
	var files util.CodeGeneratedFiles

	fs, err := compiler.ProtocGo(c.Path, c.Package, c.Files)
	if err != nil {
		return nil, err
	}

	files = append(files, fs...)

	// other compilers
	return files, nil
}
