package lang

func NewIdentifierExpr(names ...string) *IdentifierExpr {
	return &IdentifierExpr{
		Identifier: NewIdentifier(names...),
	}
}

func NewIdentifierExprFromType(nominal *NominalType) *IdentifierExpr {
	if nominal != nil {
		return &IdentifierExpr{
			StartPosition: nominal.StartPosition,
			EndPosition:   nominal.EndPosition,
			Kind:          0,
			Implicit:      false,
			Identifier: &Identifier{
				StartPosition: nil,
				EndPosition:   nil,
				Kind:          Identifier_KIND_STRUCT,
				PackageName:   nominal.PackageName,
				Name:          nominal.Name,
				Declaration:   NewDeclarationFromTypeDeclaration(nominal.TypeDeclaration),
				FullName:      nominal.GetFullName(),
				Enclosing:     nominal.Enclosing.ToIdentifier(),
			},
			GenericArguments: nominal.GenericArguments,
		}
	}
	return nil
}

func (x *IdentifierExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *IdentifierExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *IdentifierExpr) GetName() string {
	if x != nil && x.Identifier != nil {
		return x.Identifier.Name
	}
	return ""
}

func (x *IdentifierExpr) GetFullName() string {
	if x != nil && x.Identifier != nil {
		return x.Identifier.FullName
	}
	return ""
}

func (x *IdentifierExpr) GetPackageName() string {
	if x != nil && x.Identifier != nil {
		return x.Identifier.PackageName
	}
	return ""
}
