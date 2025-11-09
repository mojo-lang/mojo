package lang

func NewVariableDeclFromValueDecl(decl *ValueDecl) *VariableDecl {
	return (&VariableDecl{}).FromValueDecl(decl)
}

func (x *VariableDecl) ToValueDecl() *ValueDecl {
	if x != nil {
		return &ValueDecl{
			StartPosition:   x.StartPosition,
			EndPosition:     x.EndPosition,
			Implicit:        x.Implicit,
			Document:        x.Document,
			PackageName:     x.PackageName,
			SourceFileName:  x.SourceFileName,
			KeywordPosition: x.KeywordPosition,
			Name:            x.Name,
			Attributes:      x.Attributes,
			Group:           x.Group,
			NamePosition:    x.NamePosition,
			Type:            x.Type,
			Initializer:     x.Initializer,
		}
	}
	return nil
}

func (x *VariableDecl) FromValueDecl(decl *ValueDecl) *VariableDecl {
	if x != nil && decl != nil {
		x.StartPosition = decl.StartPosition
		x.EndPosition = decl.EndPosition
		x.Implicit = decl.Implicit
		x.Document = decl.Document
		x.PackageName = decl.PackageName
		x.SourceFileName = decl.SourceFileName
		x.KeywordPosition = decl.KeywordPosition
		x.Name = decl.Name
		x.Attributes = decl.Attributes
		x.Group = decl.Group
		x.Type = decl.Type
		x.Initializer = decl.Initializer
	}
	return x
}

func (x *VariableDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *VariableDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *VariableDecl) GetInitialValue() *Expression {
	if x != nil {
		return x.GetInitializer().GetValue()
	}
	return nil
}
