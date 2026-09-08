package lang

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
)

func TestPackageGoModule(t *testing.T) {
	for _, tc := range []struct{ name, repository, module, importPath string }{
		{"mojo.core", "github.com/mojo-lang/mojo/packages/core", MojoGoModule, MojoGoModule + "/pkg/mojo/core"},
		{"mojo.db.sql", "https://github.com/mojo-lang/mojo/packages/db/", MojoGoModule, MojoGoModule + "/pkg/mojo/db/sql"},
		{"example.api", "https://github.com/example/api", "github.com/example/api/go", "github.com/example/api/go/pkg/example/api"},
		{"mojo.custom", "github.com/example/custom", "github.com/example/custom/go", "github.com/example/custom/go/pkg/mojo/custom"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			repository, err := core.NewUrl(tc.repository)
			require.NoError(t, err)
			pkg := &Package{FullName: tc.name, Repository: repository}
			require.Equal(t, tc.module, pkg.GoModName())
			require.Equal(t, tc.importPath, pkg.GoFullPackageName())
			require.Equal(t, tc.importPath, pkg.GetGoPackageImport())
		})
	}
}
