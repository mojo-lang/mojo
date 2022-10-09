package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type ServiceDefVisitor struct {
	*BaseProtobuf3Visitor
}

func NewServiceDefVisitor() *ServiceDefVisitor {
	visitor := &ServiceDefVisitor{}
	return visitor
}

func (v *ServiceDefVisitor) VisitServiceDef(ctx *ServiceDefContext) interface{} {
	if name := ctx.ServiceName(); name != nil {
		decl := &lang.InterfaceDecl{
			Name: name.GetText(),
			Type: &lang.InterfaceType{},
		}

		elements := ctx.AllServiceElement()
		for _, element := range elements {
			ele := element.Accept(v)
			if funcDecl, ok := ele.(*lang.FunctionDecl); ok {
				decl.Type.Methods = append(decl.Type.Methods, funcDecl)
			}
		}
		return decl
	}
	return nil
}

func (v *ServiceDefVisitor) VisitServiceElement(ctx *ServiceElementContext) interface{} {
	if rpc := ctx.Rpc(); rpc != nil {
		return rpc.Accept(v)
	}
	return nil
}

func (v *ServiceDefVisitor) VisitRpc(ctx *RpcContext) interface{} {
	if name := ctx.RpcName(); name != nil {
		decl := &lang.FunctionDecl{
			Name:      name.GetText(),
			Signature: &lang.FunctionSignature{},
		}

		allMessages := ctx.AllMessageType()
		if req := allMessages[0]; req != nil {
			if typ, ok := req.Accept(NewIdentifierVisitor()).(*lang.NominalType); ok {
				valueDecl := &lang.ValueDecl{
					Type: typ,
				}
				if ctx.RequestStream() != nil {
					valueDecl.Type.Attributes = append(valueDecl.Type.Attributes, lang.NewBoolAttribute("protobuf", "stream"))
				}
				decl.Signature.AppendParameter(valueDecl)
			}
		}
		if resp := allMessages[1]; resp != nil {
			if typ, ok := resp.Accept(NewIdentifierVisitor()).(*lang.NominalType); ok {
				if ctx.ResponseStream() != nil {
					typ.Attributes = append(typ.Attributes, lang.NewBoolAttribute("protobuf", "stream"))
				}
				decl.Signature.Result = lang.NewFunctionResult(typ)
			}
		}

		allOptions := ctx.AllOptionStatement()
		for _, option := range allOptions {
			if attr, ok := option.Accept(NewOptionStatementVisitor()).(*lang.Attribute); ok {
				decl.Attributes = append(decl.Attributes, attr)
			}
		}

		return decl
	}

	return nil
}
