package converter

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestConvertFileOptionsImportsAndGenericSkipping(t *testing.T) {
	converter := New()
	pkg := &lang.Package{FullName: "sample"}
	source := &lang.SourceFile{Name: "order_item.mojo", FullName: "sample/order_item.mojo", PackageName: "sample",
		Statements: []*lang.Statement{
			lang.NewStructDeclStatement(&lang.StructDecl{Name: "Item"}),
			lang.NewStructDeclStatement(&lang.StructDecl{Name: "Generic", GenericParameters: []*lang.GenericParameter{{Name: "T"}}}),
		},
		ResolvedIdentifiers: []*lang.Identifier{
			{SourceFileName: "external/value.mojo"}, {SourceFileName: "external/value.mojo"},
			{SourceFileName: "sample/order_item.mojo"}, {SourceFileName: "mojo/core/string.mojo"},
		},
	}
	file, err := converter.compileFile(context.WithType(context.Empty(), pkg), source)
	require.NoError(t, err)
	require.Equal(t, "sample/order_item.proto", file.GetName())
	require.Equal(t, "sample", file.GetPackageName())
	require.Equal(t, "proto3", file.Proto.GetSyntax())
	require.Equal(t, "sample", file.GetOptions().GetJavaPackage())
	require.Equal(t, "OrderItemProto", file.GetOptions().GetJavaOuterClassname())
	require.True(t, file.GetOptions().GetJavaMultipleFiles())
	require.Len(t, file.Messages, 1)
	require.Equal(t, "Item", file.Messages[0].GetName())
	require.Equal(t, []string{"external/value.proto"}, file.GetDependencies())
	source.Name = "Generic<String>.mojo"
	file, err = converter.compileFile(context.Empty(), source)
	require.NoError(t, err)
	require.Nil(t, file)
}

func TestConvertFileReportsMissingContextAndDeclaration(t *testing.T) {
	converter := New()
	source := &lang.SourceFile{Name: "item.mojo"}
	file, err := converter.CompileFile(source)
	require.ErrorContains(t, err, "has no package found")
	require.Nil(t, file)
	source.Statements = []*lang.Statement{{Statement: &lang.Statement_Declaration{}}}
	ctx := context.WithType(context.Empty(), &lang.Package{FullName: "sample"})
	file, err = converter.compileFile(ctx, source)
	require.ErrorContains(t, err, "declaration statement is nil")
	require.Nil(t, file)
}
