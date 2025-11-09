package lang

const (
	TermTypeKeyword = "keyword"
	TermTypeName    = "name"
	TermTypeStart   = "start"
	TermTypeEnd     = "end"
	TermTypeSymbol  = "symbol"
)

func (x *Term) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *Term) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func NewSymbolTerm(position *Position, termType string, value string) *Term {
	if position != nil {
		endPosition := position
		if len(value) > 0 {
			endPosition = &Position{
				Line:   position.Line,
				Column: position.Column + int64(len(value)),
			}
		}
		return &Term{StartPosition: position,
			EndPosition: endPosition,
			Type:        termType,
			Value:       value,
		}
	} else if len(termType) > 0 && len(value) > 0 {
		return &Term{Type: termType, Value: value}
	}

	return nil
}
