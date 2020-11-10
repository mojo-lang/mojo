package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type InterfaceDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewInterfaceDeclarationVisitor() *InterfaceDeclarationVisitor {
	decl := &InterfaceDeclarationVisitor{}
	return decl
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	decl := &lang.InterfaceDecl{}
	decl.Type = &lang.InterfaceType{}

	decl.Document = GetDocument(ctx.Document())
	decl.Attributes = GetAttributes(ctx.Attributes())

	interfaceName := ctx.InterfaceName()
	if interfaceName != nil {
		decl.Name = interfaceName.Accept(i).(string)
	} else {
		// error happened
	}

	decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())
	decl.Type.Inherits = GetTypeInheritances(ctx.TypeInheritanceClause())

	bodyCtx := ctx.InterfaceBody()
	if bodyCtx != nil {
		decl.Type.Methods = bodyCtx.Accept(i).([]*lang.FuncDecl)
	}

	return decl
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceName(ctx *InterfaceNameContext) interface{} {
	return ctx.GetText()
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceBody(ctx *InterfaceBodyContext) interface{} {
	membersCtx := ctx.InterfaceMembers()
	if membersCtx != nil {
		return membersCtx.Accept(i)
	}

	return []*lang.FuncDecl{}
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceMembers(ctx *InterfaceMembersContext) interface{} {
	memberCtxes := ctx.AllInterfaceMember()
	var methods []*lang.FuncDecl
	for _, memberCtx := range memberCtxes {
		funcDecl := memberCtx.Accept(i).(*lang.FuncDecl)
		if funcDecl != nil {
			methods = append(methods, funcDecl)
		}
	}

	return methods
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceMember(ctx *InterfaceMemberContext) interface{} {
	methodCtx := ctx.InterfaceMethodDeclaration()
	if methodCtx != nil {
		visitor := NewFuncDeclarationVisitor()
		return methodCtx.Accept(visitor)
	}
	return nil
}
