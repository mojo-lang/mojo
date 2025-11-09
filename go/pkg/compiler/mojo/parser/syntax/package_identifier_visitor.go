package syntax

import (
	"fmt"
	"strings"
)

type PackageIdentifierVisitor struct {
	*BaseMojoParserVisitor
}

func GetPackageIdentifier(ctx IPackageIdentifierContext) string {
	if ctx != nil {
		visitor := NewPackageIdentifierVisitor()
		if str, ok := ctx.Accept(visitor).(string); ok {
			return str
		} else {
			fmt.Print("===> error")
		}
	}
	return ""
}

func NewPackageIdentifierVisitor() *PackageIdentifierVisitor {
	visitor := &PackageIdentifierVisitor{}
	return visitor
}

func (t *PackageIdentifierVisitor) VisitPackageIdentifier(ctx *PackageIdentifierContext) interface{} {
	identifiers := ctx.AllPackageName()

	var names []string
	for _, identifier := range identifiers {
		names = append(names, identifier.GetText())
	}

	return strings.Join(names, ".")
}
