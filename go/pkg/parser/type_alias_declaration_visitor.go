package parser

type TypeAliasDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeAliasDeclarationVisitor() *TypeAliasDeclarationVisitor {
	visitor := &TypeAliasDeclarationVisitor{}
	return visitor
}
