package lang

import "errors"

func (x *EnumDecl) GetFullName() string {
	if x != nil {
		return GetFullName(x.PackageName, GetEnclosingNames(x.Enclosing), x.Name)
	}
	return ""
}

func (x *EnumDecl) SetScope(scope *Scope) {
	if x != nil {
		x.Scope = scope
	}
}

func (x *EnumDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *EnumDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *EnumDecl) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil EnumDecl")
}

func (x *EnumDecl) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		for _, attribute := range x.Attributes {
			terms = append(terms, attribute)
		}

		terms = append(terms,
			NewSymbolTerm(x.KeywordPosition, TermTypeKeyword, KeywordEnum),
			NewSymbolTerm(x.NamePosition, TermTypeName, x.Name))

		for _, parameter := range x.GenericParameters {
			terms = append(terms, parameter)
		}

		if x.Type != nil {
			if x.Type.UnderlyingType != nil {
				terms = append(terms,
					NewSymbolTerm(x.Type.UnderlyingTypePosition, TermTypeSymbol, ":"),
					x.Type.UnderlyingType)
			}

			terms = append(terms, NewSymbolTerm(x.Type.StartPosition, TermTypeStart, ""))

			for _, enumerator := range x.Type.Enumerators {
				terms = append(terms, enumerator)
			}

			terms = append(terms, NewSymbolTerm(x.Type.EndPosition, TermTypeEnd, ""))
		}
	}

	return nil
}
