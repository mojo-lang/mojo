package plugin

import (
	"reflect"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

type Plugin interface {
	GetName() string
	GetPriority() int

	GetGroup() string
	GetGroupPriority() int

	Create(options core.Options) Plugin
}

func ContainsAnyMethod(p interface{}, methods ...string) bool {
	v := reflect.ValueOf(p)
	for _, method := range methods {
		m := v.MethodByName(method)
		if m.IsValid() {
			return true
		}
	}
	return false
}
