package lang

import "errors"

func NewFunctionSignature(result *NominalType, params ...*ValueDecl) *FunctionSignature {
	return &FunctionSignature{
		Parameter: NewFunctionParameters(params...),
		Result:    NewFunctionResult(result),
	}
}

func NewFunctionParameters(params ...*ValueDecl) *FunctionSignature_Parameter {
	if len(params) > 0 {
		return &FunctionSignature_Parameter{Decls: params}
	}
	return nil
}

func NewFunctionResult(result *NominalType) *FunctionSignature_Result {
	if result != nil {
		return &FunctionSignature_Result{Type: result}
	}
	return nil
}

func (x *FunctionSignature) SetParameters(decls ...*ValueDecl) *FunctionSignature {
	if x != nil && len(decls) > 0 {
		if x.Parameter != nil {
			x.Parameter = &FunctionSignature_Parameter{}
		}
		x.Parameter.Decls = decls
	}
	return x
}

func (x *FunctionSignature) AppendParameter(decl *ValueDecl) *FunctionSignature {
	if x != nil && decl != nil {
		if x.Parameter == nil {
			x.Parameter = &FunctionSignature_Parameter{}
		}
		x.Parameter.Decls = append(x.Parameter.Decls, decl)
	}
	return x
}

func (x *FunctionSignature) GetParameters() []*ValueDecl {
	return x.GetParameter().GetDecls()
}

func (x *FunctionSignature) ParameterDecl(index int) *ValueDecl {
	if decls := x.GetParameters(); len(decls) > 0 && index >= 0 && index < len(decls) {
		return decls[index]
	}
	return nil
}

func (x *FunctionSignature) ParameterDeclByName(name string) *ValueDecl {
	if decls := x.GetParameters(); len(decls) > 0 && len(name) > 0 {
		for _, decl := range decls {
			if decl.Name == name {
				return decl
			}
		}
	}
	return nil
}

func (x *FunctionSignature) GetResultType() *NominalType {
	return x.GetResult().GetType()
}

func (x *FunctionSignature) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *FunctionSignature) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *FunctionSignature) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil FunctionSignature")
}

func (x *FunctionSignature) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		if x.Parameter != nil {
			includingParen := x.GetStartPosition().Compare(x.Parameter.GetStartPosition()) == 0
			if includingParen {
				terms = append(terms,
					&Term{StartPosition: x.StartPosition,
						EndPosition: x.StartPosition,
						Type:        TermTypeStart,
						Value:       "(",
					})
			}

			for _, parameter := range x.GetParameters() {
				terms = append(terms, parameter)
			}

			if includingParen {
				terms = append(terms,
					&Term{StartPosition: x.StartPosition,
						EndPosition: x.StartPosition,
						Type:        TermTypeEnd,
						Value:       ")",
					})
			}
		}

		if x.Result != nil && !x.Result.Implicit {
			terms = append(terms,
				&Term{StartPosition: x.Result.GetStartPosition(),
					EndPosition: x.Result.GetStartPosition(),
					Type:        TermTypeSymbol,
					Value:       "->",
				})

			terms = append(terms, x.GetResultType())
		}

		return terms
	}

	return nil
}

func (x *FunctionSignature_Parameter) HasFollowingDocument() bool {
	for _, param := range x.Decls {
		if param.Document.Following {
			return true
		}
	}
	return false
}

func (x *FunctionSignature_Result) HasFollowingDocument() bool {
	return x != nil && x.Type != nil && x.Type.Document.Following
}
