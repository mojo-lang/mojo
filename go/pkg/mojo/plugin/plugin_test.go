package plugin

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type PackageParser struct{}

func (p *PackageParser) ParsePackage(ctx context.Context, pkg *lang.Package) error {
	return nil
}

type OtherParser struct{}

func (p *OtherParser) Parse(ctx context.Context, pkg *lang.Package) error {
	return nil
}

func TestContainsAnyMethod(t *testing.T) {
	assert.True(t, ContainsAnyMethod(&PackageParser{}, parsePackageMethod, parseInterfaceMethod))
	assert.False(t, ContainsAnyMethod(&OtherParser{}, parsePackageMethod, parseInterfaceMethod))
}
