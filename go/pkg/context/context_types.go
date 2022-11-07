package context

import (
	"context"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

const (
	TypeKey              = "@type"
	PackageTypeKey       = "@type/Package"
	SourceFileTypeKey    = "@type/SourceFile"
	InterfaceDeclTypeKey = "@type/Interface"
	StructDeclTypeKey    = "@type/Struct"
	EnumDeclTypeKey      = "@type/Enum"
	TypeAliasDeclTypeKey = "@type/TypeAlias"
	ValueDeclTypeKey     = "@type/ValueDecl"
	FunctionDeclTypeKey  = "@type/FunctionDecl"
	NominalTypeKey       = "@type/NominalType"
)

func WithType(ctx context.Context, value interface{}) context.Context {
	key := ""
	switch value.(type) {
	case *lang.Package:
		key = PackageTypeKey
	case *lang.SourceFile:
		key = SourceFileTypeKey
	case *lang.StructDecl:
		key = StructDeclTypeKey
	case *lang.EnumDecl:
		key = EnumDeclTypeKey
	case *lang.TypeAliasDecl:
		key = TypeAliasDeclTypeKey
	case *lang.InterfaceDecl:
		key = InterfaceDeclTypeKey
	case *lang.ValueDecl:
		key = ValueDeclTypeKey
	case *lang.FunctionDecl:
		key = FunctionDeclTypeKey
	case *lang.NominalType:
		key = NominalTypeKey
	}
	return WithValues(ctx, TypeKey, value, key, value)
}

func Package(ctx context.Context) *lang.Package {
	if pkg, ok := ctx.Value(PackageTypeKey).(*lang.Package); ok {
		return pkg
	}
	return nil
}

func SourceFile(ctx context.Context) *lang.SourceFile {
	if file, ok := ctx.Value(SourceFileTypeKey).(*lang.SourceFile); ok {
		return file
	}
	return nil
}

func StructDecl(ctx context.Context) *lang.StructDecl {
	if decl, ok := ctx.Value(StructDeclTypeKey).(*lang.StructDecl); ok {
		return decl
	}
	return nil
}

func EnumDecl(ctx context.Context) *lang.EnumDecl {
	if decl, ok := ctx.Value(EnumDeclTypeKey).(*lang.EnumDecl); ok {
		return decl
	}
	return nil
}

func TypeAliasDecl(ctx context.Context) *lang.TypeAliasDecl {
	if decl, ok := ctx.Value(TypeAliasDeclTypeKey).(*lang.TypeAliasDecl); ok {
		return decl
	}
	return nil
}

func InterfaceDecl(ctx context.Context) *lang.InterfaceDecl {
	if decl, ok := ctx.Value(InterfaceDeclTypeKey).(*lang.InterfaceDecl); ok {
		return decl
	}
	return nil
}

func ValueDecl(ctx context.Context) *lang.ValueDecl {
	if decl, ok := ctx.Value(ValueDeclTypeKey).(*lang.ValueDecl); ok {
		return decl
	}
	return nil
}

func FunctionDecl(ctx context.Context) *lang.FunctionDecl {
	if decl, ok := ctx.Value(FunctionDeclTypeKey).(*lang.FunctionDecl); ok {
		return decl
	}
	return nil
}

func NominalType(ctx context.Context) *lang.NominalType {
	if nominal, ok := ctx.Value(NominalTypeKey).(*lang.NominalType); ok {
		return nominal
	}
	return nil
}

func TypeValue(ctx context.Context) interface{} {
	return ctx.Value(TypeKey)
}

func TypeValues(ctx context.Context) []interface{} {
	return Values(ctx, TypeKey)
}

func PreviousTypeValue(ctx context.Context, index int) interface{} {
	return PreviousValue(ctx, TypeKey, index)
}

func IsTypeEnclosed(ctx context.Context) bool {
	_, ok := PreviousTypeValue(ctx, 1).(lang.EnclosingTypeDecl)
	return ok
}
