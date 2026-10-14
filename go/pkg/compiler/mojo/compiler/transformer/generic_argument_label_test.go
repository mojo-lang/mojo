package transformer

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestGenericArgumentLabel(t *testing.T) {
	for _, tt := range []struct{ name, pkg, label, format, want string }{
		{"UserAccount", "sample", "", "", "user_account"},
		{"String", "mojo.core", "", "", "string_val"},
		{"String", "mojo.core", "", "{}", "string"},
		{"String", "mojo.core", "display_name", "{}", "display_name"},
		{"String", "mojo.core", "", "unsupported", "string_val"},
	} {
		t.Run(tt.want+"/"+tt.format, func(t *testing.T) {
			typ := &lang.NominalType{Name: tt.name, PackageName: tt.pkg}
			typ.TypeDeclaration = lang.NewStructTypeDeclaration(&lang.StructDecl{Name: tt.name, PackageName: tt.pkg})
			if tt.label != "" {
				typ.Attributes = lang.SetStringAttribute(typ.Attributes, core.LabelAttributeName, tt.label)
			}
			if tt.format != "" {
				typ.Attributes = lang.SetStringAttribute(typ.Attributes, core.LabelFormatAttributeName, tt.format)
			}
			require.Equal(t, tt.want, GenericArgumentLabel(typ))
		})
	}
}
