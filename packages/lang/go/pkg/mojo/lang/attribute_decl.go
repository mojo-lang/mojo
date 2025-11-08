package lang

import (
	"errors"
)

func (x *AttributeDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *AttributeDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *AttributeDecl) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil AttributeDecl")
}

func (x *AttributeDecl) GetTerms() []interface{} {
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

		if nominalType := x.GetNominalType(); nominalType != nil {
			terms = append(terms, nominalType)
		} else if structType := x.GetStructType(); structType != nil {
			terms = append(terms,
				&Term{StartPosition: structType.StartPosition,
					EndPosition: structType.StartPosition,
					Type:        TermTypeStart,
				})

			for _, field := range structType.Fields {
				terms = append(terms, field)
			}

			terms = append(terms,
				&Term{StartPosition: structType.EndPosition,
					EndPosition: structType.EndPosition,
					Type:        TermTypeEnd,
				})
		}

		if x.DefaultValue != nil {
			terms = append(terms, x.DefaultValue)
		}

		return terms
	}
	return nil
}
