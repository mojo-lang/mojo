package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type ParameterVisitor struct {
	*BaseCVisitor
}

func NewParameterVisitor() *ParameterVisitor {
	visitor := &ParameterVisitor{}
	return visitor
}

func (v *ParameterVisitor) VisitParameterTypeList(ctx *ParameterTypeListContext) interface{} {
	var parameters []*lang.ValueDecl
	if parameterList := ctx.ParameterList(); parameterList != nil {
		if params, ok := parameterList.Accept(v).([]*lang.ValueDecl); ok {
			parameters = params
		}
	}

	if ellipsis := ctx.Ellipsis(); ellipsis != nil {
		parameters = append(parameters, &lang.ValueDecl{Name: "..."})
	}
	return parameters
}

func (v *ParameterVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	var parameters []*lang.ValueDecl
	if params := ctx.AllParameterDeclaration(); len(params) > 0 {
		for _, param := range params {
			if decl, ok := param.Accept(v).(*lang.ValueDecl); ok {
				parameters = append(parameters, decl)
			}
		}
	}
	return parameters
}

func (v *ParameterVisitor) VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{} {
	if specifiers := ctx.DeclarationSpecifiers(); specifiers != nil {
		if typ, ok := specifiers.Accept(NewDeclarationVisitor()).(*lang.NominalType); ok {
			if declarator := ctx.Declarator(); declarator != nil {
				if name, ok := declarator.Accept(NewDeclaratorVisitor()).(*lang.ValueDecl); ok {
					return CombineValueDecl(name, typ)
				}
			}
		}
	}
	if specifiers := ctx.DeclarationSpecifiers2(); specifiers != nil {
		if typ, ok := specifiers.Accept(NewDeclarationVisitor()).(*lang.NominalType); ok {
			if declarator := ctx.AbstractDeclarator(); declarator != nil {
				if name, ok := declarator.Accept(NewDeclaratorVisitor()).(*lang.ValueDecl); ok {
					return CombineValueDecl(name, typ)
				}
			} else {
				return &lang.ValueDecl{Type: typ}
			}
		}
	}
	return nil
}
