package lang

func (x *EmptyStmt) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *EmptyStmt) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}
