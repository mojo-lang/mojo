package document

func NewTocDivision() *Division {
	return &Division{
		Attribute: &Attribute{
			Identifier: "toc",
		},
		Content: []*Block{NewTextPlainBlock("[TOC]")},
	}
}
