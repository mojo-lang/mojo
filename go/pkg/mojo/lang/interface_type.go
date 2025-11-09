package lang

import (
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
)

// type name pluralize
// generic type name
// union label

var pluralizing *pluralize.Client

func init() {
	pluralizing = pluralize.NewClient()
}

type MethodName struct {
	Method   string
	Resource string
}

const (
	GetMethod    = "get"
	ListMethod   = "list"
	CreateMethod = "create"
	UpdateMethod = "update"
	DeleteMethod = "delete"

	BatchMethod       = "batch"
	BatchGetMethod    = "batch_get"
	BatchCreateMethod = "batch_create"
	BatchUpdateMethod = "batch_update"
	BatchDeleteMethod = "batch_delete"
)

func (m *InterfaceType) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *InterfaceType) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *InterfaceType) GetInheritMethods() []*FunctionDecl {
	var decls []*FunctionDecl
	for _, inherit := range m.GetInherits() {
		if typ := inherit.GetTypeDeclaration().GetInterfaceDecl().GetType(); typ != nil {
			if methods := typ.GetInheritMethods(); len(methods) > 0 {
				decls = append(decls, methods...)
			}

			for _, method := range typ.GetMethods() {
				decls = append(decls, method)
			}
		}
	}

	return decls
}

func (m *InterfaceType) GetMethodGroups() map[string][]*FunctionDecl {
	groups := make(map[string][]*FunctionDecl)
	if m == nil {
		return groups
	}

	var unprocessedFuncs []*FunctionDecl
	for _, method := range m.Methods {
		if name := method.parseStandardMethod(); name != nil {
			// set attribution
			method.SetImplicitStringAttribute(core.MethodAttributeName, name.Method)
			method.SetImplicitStringAttribute(core.ResourceAttributeName, name.Resource)

			groups[name.Resource] = append(groups[name.Resource], method)
		} else {
			unprocessedFuncs = append(unprocessedFuncs, method)
		}
	}

	for _, function := range unprocessedFuncs {
		name := pluralizing.Singular(function.Name)
		for group := range groups {
			gn := "_" + group
			if strings.HasSuffix(name, gn) {
				function.SetImplicitStringAttribute(core.MethodAttributeName, strings.TrimSuffix(function.Name, gn))
				function.SetImplicitStringAttribute(core.ResourceAttributeName, group)
				groups[group] = append(groups[group], function)
			}
		}
	}

	return groups
}

func (x *FunctionDecl) parseStandardMethod() *MethodName {
	if x != nil {
		getMethodName := func(name string) *MethodName {
			return &MethodName{Method: name, Resource: strcase.ToCamel(pluralizing.Singular(x.Name[len(name)+1:]))}
		}

		segments := strings.Split(x.Name, "_")
		if len(segments) > 0 {
			switch segments[0] {
			case GetMethod:
				return getMethodName(GetMethod)
			case ListMethod:
				return getMethodName(ListMethod)
			case CreateMethod:
				return getMethodName(CreateMethod)
			case UpdateMethod:
				return getMethodName(UpdateMethod)
			case DeleteMethod:
				return getMethodName(DeleteMethod)
			case BatchMethod:
				if len(segments) > 1 {
					switch segments[1] {
					case GetMethod:
						return getMethodName(BatchGetMethod)
					case CreateMethod:
						return getMethodName(BatchCreateMethod)
					case UpdateMethod:
						return getMethodName(BatchUpdateMethod)
					case DeleteMethod:
						return getMethodName(BatchDeleteMethod)
					}
				}
			}
		}
	}
	return nil
}
