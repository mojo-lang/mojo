package lang

import "errors"

func (x *ClosureExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *ClosureExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *ClosureExpr) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil StructDecl")
}

func (x *ClosureExpr) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		if x.Signature != nil {
			if x.Signature.Parameter != nil {
				terms = append(terms, x.Signature.Parameter)
			}

			if x.Signature.Result != nil /*&& !x.Signature.Result.Implicit*/ {
				terms = append(terms, x.Signature.Result)
			}
		}

		if x.Body != nil {
			terms = append(terms, x.Body)
		}

		return terms
	}
	return nil
}
