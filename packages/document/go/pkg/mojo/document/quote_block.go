package document

func NewQuoteBlock(blocks ...*Block) *QuoteBlock {
	return &QuoteBlock{Blocks: blocks}
}
