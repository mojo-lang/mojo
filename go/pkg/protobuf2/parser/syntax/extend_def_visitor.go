package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type ExtendDefVisitor struct {
	*BaseProtobuf2Visitor
}

func NewExtendDefVisitor() *ExtendDefVisitor {
	visitor := &ExtendDefVisitor{}
	return visitor
}

func (v *ExtendDefVisitor) VisitExtendDef(ctx *ExtendDefContext) interface{} {
	if messageType := ctx.MessageType(); messageType != nil {
		if typ, ok := messageType.Accept(NewIdentifierVisitor()).(*lang.NominalType); ok {
			decl := &lang.StructDecl{
				Type: &lang.StructType{},
			}
			decl.Type.Inherits = append(decl.Type.Inherits, typ)

			allElements := ctx.AllExtendElement()
			for _, element := range allElements {
				if value, ok := element.Accept(v).(*lang.ValueDecl); ok {
					decl.Type.Fields = append(decl.Type.Fields, value)
				}
			}

			return decl
		}
	}

	return nil
}

func (v *ExtendDefVisitor) VisitExtendElement(ctx *ExtendElementContext) interface{} {
	if field := ctx.Field(); field != nil {
		return field.Accept(NewMessageDefVisitor())
	}
	return nil
}
