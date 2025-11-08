package plugin

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type packageParser struct{}

func (p *packageParser) ParsePackage(ctx context.Context, pkg *lang.Package) error {
	return nil
}

type otherParser struct{}

func (p *otherParser) Parse(ctx context.Context, pkg *lang.Package) error {
	return nil
}

func TestContainsAnyMethod(t *testing.T) {
	assert.True(t, ContainsAnyMethod(&packageParser{}, parsePackageMethod, parseInterfaceMethod))
	assert.False(t, ContainsAnyMethod(&otherParser{}, parsePackageMethod, parseInterfaceMethod))
}
