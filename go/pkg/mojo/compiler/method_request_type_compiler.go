package compiler

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

const methodRequestTypeName = "compiler.method-request-type"

func init() {
	plugin.RegisterPlugin(NewMethodRequestTypeCompiler(nil))
}

type MethodRequestTypeCompiler struct {
	plugin.BasicPlugin
}

func NewMethodRequestTypeCompiler(options core.Options) *MethodRequestTypeCompiler {
	return &MethodRequestTypeCompiler{
		BasicPlugin: plugin.BasicPlugin{
			Name:          methodRequestTypeName,
			Group:         "compiler",
			GroupPriority: 10,
			Priority:      55,
			Creator: func(options core.Options) plugin.Plugin {
				return NewMethodRequestTypeCompiler(options)
			},
		},
	}
}

func (c *MethodRequestTypeCompiler) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	pkg := context.Package(ctx)
	logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompileInterface", "pkg", pkg.FullName)

	for _, method := range decl.GetType().GetMethods() {
		if err := c.CompileMethod(ctx, method); err != nil {
			return err
		}
	}
	return nil
}

func (c *MethodRequestTypeCompiler) CompileMethod(ctx context.Context, method *lang.FunctionDecl) error {
	wrapRequestName := strcase.ToCamel(method.Name) + "Request"

	// special case for the request, do NOT generate new request type
	if parameters := method.Signature.GetParameters(); len(parameters) == 1 {
		param := parameters[0]
		if len(param.Name) == 0 || param.Type.Name == wrapRequestName {
			logs.Infow("the only one argument is the request message type in the method, will use it", "method", method.FullName)
			param.SetImplicitBoolAttribute(protobuf.MethodRequestTypeAttributeName, true)
		}
	}

	return nil
}
