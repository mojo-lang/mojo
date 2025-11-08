package lang

import "errors"

func (x *InterfaceDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *InterfaceDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *InterfaceDecl) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil InterfaceDecl")
}

func (x *InterfaceDecl) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		if x.Document != nil {
			terms = append(terms, x.Document)
		}
		for _, attribute := range x.Attributes {
			terms = append(terms, attribute)
		}

		terms = append(terms,
			NewSymbolTerm(x.KeywordPosition, TermTypeKeyword, KeywordInterface),
			NewSymbolTerm(x.NamePosition, TermTypeName, x.Name))

		for _, parameter := range x.GenericParameters {
			terms = append(terms, parameter)
		}

		if x.Type != nil {
			if len(x.Type.Inherits) > 0 {
				terms = append(terms, NewSymbolTerm(x.Type.InheritePosition, TermTypeSymbol, ":"))

				for _, inherit := range x.Type.Inherits {
					terms = append(terms, inherit)
				}
			}

			terms = append(terms, NewSymbolTerm(x.Type.StartPosition, TermTypeStart, ""))

			for _, method := range x.Type.Methods {
				terms = append(terms, method)
			}

			terms = append(terms, NewSymbolTerm(x.Type.EndPosition, TermTypeEnd, ""))
		}

		return terms
	}
	return nil
}

func (x *InterfaceDecl) GetFullName() string {
	if x != nil {
		return GetFullName(x.PackageName, nil, x.Name)
	}
	return ""
}

func (x *InterfaceDecl) SetScope(scope *Scope) {
	if x != nil {
		x.Scope = scope
	}
}

func (x *InterfaceDecl) GetInheritMethods() []*FunctionDecl {
	if x != nil && x.Type != nil {
		return x.Type.GetInheritMethods()
	}
	return nil
}

func (x *InterfaceDecl) GetMethodGroups() map[string][]*FunctionDecl {
	if x != nil && x.Type != nil {
		return x.Type.GetMethodGroups()
	}
	return nil
}

func (x *InterfaceDecl) IsBodyEmpty() bool {
	if x != nil {
		return len(x.TypeAliasDecls) == 0 && x.GetMethodCount() == 0
	}
	return true
}

func (x *InterfaceDecl) GetMethodCount() int {
	return len(x.GetType().GetMethods())
}
