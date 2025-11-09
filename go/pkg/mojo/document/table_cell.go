package document

func NewTableCell(blocks ...*Block) *Table_Cell {
	return &Table_Cell{Vals: blocks}
}

func NewTextTableCell(text string) *Table_Cell {
	return &Table_Cell{Vals: []*Block{NewTextPlainBlock(text)}}
}
