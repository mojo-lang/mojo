package lang

type ScopeSetter interface {
	SetScope(scope *Scope)
}

func SetScope(object interface{}, scope *Scope) {
	if setter, ok := object.(ScopeSetter); ok {
		setter.SetScope(scope)
	}
}
