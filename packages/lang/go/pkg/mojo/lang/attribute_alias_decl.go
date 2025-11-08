package lang

import "errors"

func (x *AttributeAliasDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *AttributeAliasDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *AttributeAliasDecl) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil AttributeAliasDecl")
}

func (x *AttributeAliasDecl) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		for _, attribute := range x.Attributes {
			terms = append(terms, attribute)
		}

		if x.Group == nil {
			terms = append(terms,
				&Term{StartPosition: x.KeywordPosition,
					EndPosition: &Position{
						Line:   x.KeywordPosition.Line,
						Column: x.KeywordPosition.Column + int64(len(KeywordAttribute)),
					},
					Type:  "Keyword",
					Value: KeywordAttribute,
				},
			)
		}

		terms = append(terms,
			&Term{StartPosition: x.NamePosition,
				EndPosition: &Position{
					Line:   x.NamePosition.Line,
					Column: x.NamePosition.Column + int64(len(x.Name)),
				},
				Type:  "Name",
				Value: x.Name,
			})

		for _, parameter := range x.GenericParameters {
			terms = append(terms, parameter)
		}

		terms = append(terms, x.Attribute)

		return terms
	}
	return nil
}
