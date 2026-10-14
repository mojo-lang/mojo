package transformer

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestGenericTypeNamer(t *testing.T) {
	for _, tt := range []struct {
		name      string
		arguments []string
		want      string
	}{
		{"User", nil, "User"},
		{"Array", []string{"User"}, "UserArray"},
		{"Map", []string{"String", "User"}, "StringUserMap"},
		{"Cached", []string{"User"}, "CachedUser"},
		{"Comparable", []string{"User", "Admin"}, "ComparableUserAdmin"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			typ := &lang.NominalType{Name: tt.name, PackageName: "sample"}
			for _, name := range tt.arguments {
				typ.GenericArguments = append(typ.GenericArguments, &lang.NominalType{Name: name})
			}
			require.Equal(t, tt.want, (GenericTypeNamer{}).Name(typ))
		})
	}
}

func TestGenericTypeNamerNameFrom(t *testing.T) {
	namer := GenericTypeNamer{}
	require.Equal(t, "StringUserMap", namer.NameFrom("Map<String,User>"))
	require.Equal(t, "CachedUser", namer.NameFrom("Cached<User>"))
	require.Equal(t, "User", namer.NameFrom("sample.User"))
	require.Equal(t, "Map<", namer.NameFrom("Map<"))
}
