package lang

func (x *PackageDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *PackageDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}
