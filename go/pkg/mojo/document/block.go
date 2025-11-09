package document

func NewPainBlock(inlines ...*Inline) *Block {
	return &Block{Block: &Block_Plain{Plain: &Plain{Inlines: inlines}}}
}

func NewHeaderBlock(header *Header) *Block {
	return &Block{Block: &Block_Header{Header: header}}
}

func NewQuoteBlockBlock(quotes ...*Block) *Block {
	return &Block{Block: &Block_QuoteBlock{QuoteBlock: NewQuoteBlock(quotes...)}}
}

func NewCodeBlockBlock(block *CodeBlock) *Block {
	return &Block{Block: &Block_CodeBlock{CodeBlock: block}}
}

func NewOrderedListBlock(block *OrderedList) *Block {
	return &Block{Block: &Block_OrderedList{OrderedList: block}}
}

func NewBulletListBlock(block *BulletList) *Block {
	return &Block{Block: &Block_BulletList{BulletList: block}}
}

func NewParagraphBlock(inlines ...*Inline) *Block {
	return &Block{Block: &Block_Paragraph{Paragraph: &Paragraph{Inlines: inlines}}}
}

func NewLineBlockBlock(lines ...*Line) *Block {
	return &Block{Block: &Block_LineBlock{LineBlock: &LineBlock{Lines: lines}}}
}

func NewDivisionBlock(division *Division) *Block {
	return &Block{Block: &Block_Division{Division: division}}
}

func NewTableBlock(table *Table) *Block {
	return &Block{Block: &Block_Table{Table: table}}
}

func NewTextPlainBlock(text string) *Block {
	return NewPainBlock(NewTextInline(text))
}
