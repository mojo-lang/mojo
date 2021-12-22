package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type InterfaceDeclarationVisitor struct {
	*BaseMojoParserVisitor

	Interface *lang.InterfaceDecl
}

func NewInterfaceDeclarationVisitor() *InterfaceDeclarationVisitor {
	decl := &InterfaceDeclarationVisitor{}
	return decl
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	i.Interface = &lang.InterfaceDecl{}
	i.Interface.Type = &lang.InterfaceType{}

	interfaceName := ctx.InterfaceName()
	if interfaceName != nil {
		i.Interface.Name = interfaceName.Accept(i).(string)
	} else {
		// error happened
	}

	i.Interface.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())
	i.Interface.Type.Inherits = GetTypeInheritances(ctx.TypeInheritanceClause())

	bodyCtx := ctx.InterfaceBody()
	if bodyCtx != nil {
		i.Interface.Type.Methods = bodyCtx.Accept(i).([]*lang.FunctionDecl)
	}

	return i.Interface
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceName(ctx *InterfaceNameContext) interface{} {
	return ctx.GetText()
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceBody(ctx *InterfaceBodyContext) interface{} {
	membersCtx := ctx.InterfaceMembers()
	if membersCtx != nil {
		return membersCtx.Accept(i)
	}

	return []*lang.FunctionDecl{}
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceMembers(ctx *InterfaceMembersContext) interface{} {
	memberCtxes := ctx.AllInterfaceMember()
	var methods []*lang.FunctionDecl
	for _, memberCtx := range memberCtxes {
		funcDecl := memberCtx.Accept(i).(*lang.FunctionDecl)
		if funcDecl != nil {
			methods = append(methods, funcDecl)
		}
	}

	return methods
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceMember(ctx *InterfaceMemberContext) interface{} {
	document := GetDocument(ctx.Document())
	attributes := GetAttributes(ctx.Attributes())

	methodCtx := ctx.InterfaceMethodDeclaration()
	if methodCtx != nil {
		visitor := NewFuncDeclarationVisitor()
		if funcDecl, ok := methodCtx.Accept(visitor).(*lang.FunctionDecl); ok {
			funcDecl.Document = document
			funcDecl.Attributes = attributes
			return funcDecl
		} else {
			logs.Errorw("failed to parse the method in the interface", "interface", i.Interface.Name)
		}
	}
	return nil
}
