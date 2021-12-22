package syntax

import "fmt"

type PathIdentifierVisitor struct {
	*BaseMojoParserVisitor
}

func GetPathIdentifier(ctx IPathIdentifierContext) []string {
	if ctx != nil {
		if ids, ok := ctx.Accept(NewPathIdentifierVisitor()).([]string); ok {
			return ids
		} else {
			fmt.Print("===> error")
		}
	}
	return nil
}

func NewPathIdentifierVisitor() *PathIdentifierVisitor {
	visitor := &PathIdentifierVisitor{}
	return visitor
}

func (t *PathIdentifierVisitor) VisitPathIdentifier(ctx *PathIdentifierContext) interface{} {
	identifiers := ctx.AllDeclarationIdentifier()
	var ids []string
	for _, i := range identifiers {
		ids = append(ids, GetDeclarationIdentifier(i))
	}

	return ids
}
