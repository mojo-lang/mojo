package compiler

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

const irregularCaseRuleName = "compiler.irregular-case-rule"

func init() {
	plugin.RegisterPlugin(NewIrregularCaseRuleCompiler(nil))
}

type IrregularCaseRuleCompiler struct {
	plugin.BasicPlugin

	Options core.Options
}

func NewIrregularCaseRuleCompiler(options core.Options) *IrregularCaseRuleCompiler {
	return &IrregularCaseRuleCompiler{
		BasicPlugin: plugin.BasicPlugin{
			Name:          irregularCaseRuleName,
			Group:         "compiler",
			GroupPriority: 10,
			Priority:      2,
			Creator: func(options core.Options) plugin.Plugin {
				return NewIrregularCaseRuleCompiler(options)
			},
			MarkedPackages: make(map[string]bool),
		},
		Options: options,
	}
}

func (c *IrregularCaseRuleCompiler) CompileSourceFile(ctx context.Context, sourceFile *lang.SourceFile) error {
	pkg := context.Package(ctx)
	pkgFullName := pkg.GetFullName()

	if !c.MarkedPackages[pkgFullName] {
		c.MarkedPackages[pkgFullName] = true
		logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkgFullName)
	}

	for _, statement := range sourceFile.Statements {
		if decl := statement.GetDeclaration(); decl != nil {
			err := c.ApplyAttribute(ctx, lang.GetAttributes(decl).GetAttribute(core.IrregularCaseRuleAttributeName))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *IrregularCaseRuleCompiler) ApplyAttribute(ctx context.Context, attribute *lang.Attribute) (err error) {
	_ = ctx
	if attribute != nil && len(attribute.Arguments) > 0 {
		if len(attribute.Arguments) == 2 {
			rule, replacement := "", ""
			if rule, err = attribute.Arguments[0].GetString(); err != nil {
				return err
			}
			if replacement, err = attribute.Arguments[1].GetString(); err != nil {
				return err
			}
			err = core.ApplyIrregularCaseRuleAttribute(&core.Regex{Expression: rule}, replacement)
		}
	}
	return
}
