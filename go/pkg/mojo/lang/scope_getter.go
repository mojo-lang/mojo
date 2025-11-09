package lang

type ScopeGetter interface {
	GetScope() *Scope
}

func GetScope(v interface{}) *Scope {
	if getter, ok := v.(ScopeGetter); ok {
		return getter.GetScope()
	}
	return nil
}
