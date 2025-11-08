package document

func NewLine(inlines ...*Inline) *Line {
	return &Line{Vals: inlines}
}
