package lang

import "errors"

func (x *BlockStmt) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *BlockStmt) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *BlockStmt) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil BlockStmt")
}

func (x *BlockStmt) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		terms = append(terms, NewSymbolTerm(x.StartPosition, TermTypeStart, ""))

		for _, statement := range x.Statements {
			terms = append(terms, statement)
		}

		terms = append(terms, NewSymbolTerm(x.EndPosition, TermTypeEnd, ""))
		return terms
	}
	return nil
}
