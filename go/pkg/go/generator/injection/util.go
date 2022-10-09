package injection

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

func GetStructField(pkg *lang.Package, structName string, fieldName string) *lang.ValueDecl {
	if id := pkg.GetIdentifier(structName); id != nil {
		if decl := id.GetDeclaration().GetStructDecl(); decl != nil {
			return decl.GetField(fieldName)
		}
	}
	return nil
}
