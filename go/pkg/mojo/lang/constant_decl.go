package lang

func (x *ConstantDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *ConstantDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *ConstantDecl) GetInitialValue() *Expression {
	if x != nil {
		return x.GetInitializer().GetValue()
	}
	return nil
}
