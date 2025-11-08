package document

func NewTextTableHeader(texts ...string) *Table_Header {
	header := &Table_Header{}
	for _, text := range texts {
		header.Vals = append(header.Vals, NewTextTableCell(text))
	}
	return header
}
