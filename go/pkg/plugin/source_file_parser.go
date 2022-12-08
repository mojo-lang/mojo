package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type SourceFileParser interface {
	ParseSourceFile(ctx context.Context, file *lang.SourceFile) error
}
