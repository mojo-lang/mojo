package document

func (x *Document) AppendBlock(block *Block) {
	if x != nil {
		x.Blocks = append(x.Blocks, block)
	}
}

func (x *Document) AppendBlocks(blocks ...*Block) {
	if x != nil {
		x.Blocks = append(x.Blocks, blocks...)
	}
}

func (x *Document) AppendPain(inlines ...*Inline) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewPainBlock(inlines...))
	}
}

func (x *Document) AppendHeader(header *Header) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewHeaderBlock(header))
	}
}

func (x *Document) AppendHeaderFrom(level int64, inlines ...*Inline) {
	x.AppendHeader(&Header{
		Level: level,
		Text:  inlines,
	})
}

func (x *Document) AppendHeaderFromText(level int64, header string) {
	x.AppendHeader(&Header{
		Level: level,
		Text:  []*Inline{NewTextInline(header)},
	})
}

func (x *Document) AppendBlockQuote(blocks ...*Block) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewQuoteBlockBlock(blocks...))
	}
}

func (x *Document) AppendCodeBlock(codeBlock *CodeBlock) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewCodeBlockBlock(codeBlock))
	}
}

func (x *Document) AppendCodeBlockFrom(lang string, lines ...string) {
	x.AppendCodeBlock(NewCodeBlock(lang, lines...))
}

func (x *Document) AppendOrderedList(list *OrderedList) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewOrderedListBlock(list))
	}
}

func (x *Document) AppendBulletList(list *BulletList) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewBulletListBlock(list))
	}
}

func (x *Document) AppendParagraph(inlines ...*Inline) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewParagraphBlock(inlines...))
	}
}

func (x *Document) AppendLineBlock(lines ...*Line) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewLineBlockBlock(lines...))
	}
}

func (x *Document) AppendDivision(division *Division) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewDivisionBlock(division))
	}
}

func (x *Document) AppendTable(table *Table) {
	if x != nil {
		x.Blocks = append(x.Blocks, NewTableBlock(table))
	}
}
