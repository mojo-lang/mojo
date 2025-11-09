package document

func NewHeader(level int64, texts ...*Inline) *Header {
	return &Header{
		Attribute: NewAttribute(),
		Level:     level,
		Text:      texts,
	}
}
