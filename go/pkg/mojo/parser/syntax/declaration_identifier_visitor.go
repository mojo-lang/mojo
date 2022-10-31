package syntax

type DeclarationIdentifierVisitor struct {
	*BaseMojoParserVisitor
}

func NewDeclarationIdentifierVisitor() *DeclarationIdentifierVisitor {
	visitor := &DeclarationIdentifierVisitor{}
	return visitor
}

func (d *DeclarationIdentifierVisitor) VisitDeclarationIdentifier(ctx *DeclarationIdentifierContext) interface{} {
	return ctx.GetText()
}

func GetDeclarationIdentifier(ctx IDeclarationIdentifierContext) string {
	if ctx != nil {
		visitor := &DeclarationIdentifierVisitor{}
		if str, ok := ctx.Accept(visitor).(string); ok {
			return str
		}
	}

	return ""
}
