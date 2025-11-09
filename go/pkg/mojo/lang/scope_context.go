package lang

// ScopeContext
// A symbol table, which is organized as a stack of scopes, maintained by [[decaf.frontend.typecheck.Namer]].
//
// A typical full scope stack looks like the following:
// {{{
//     LocalScope   --- stack top (current scope)
//     ...          --- many nested local scopes
//     LocalScope
//     FormalScope
//     ClassScope
//     ...          --- many parent class scopes
//     ClassScope
//     GlobalScope  --- stack bottom
// }}}
// Make sure the global scope is always at the bottom, and NO class scope appears in neither formal nor local scope.
//
// @param global        the global scope at stack bottom
// @param scopes        a list of scopes above the bottom (first is the top)
// @param currentScope  the current scope
// @param currentClass  the current class symbol, i.e. the owner of the latest class scope
// @param currentMethod the current method symbol, i.e. owner of the latest formal scope
// @see [[Scope]]
type ScopeContext struct {
	Scopes []*Scope
}

func NewScopeContext() *ScopeContext {
	return &ScopeContext{}
}

// Open a new scope.
//
// @param scope scope
// @return a new scope context after opening `scope`
func (c *ScopeContext) Open(scope *Scope) *ScopeContext {
	c.Scopes = append(c.Scopes, scope)
	return c
}

func (c *ScopeContext) Close() {
	if len(c.Scopes) > 0 {
		c.Scopes = c.Scopes[:len(c.Scopes)-1]
	}
}

func (c *ScopeContext) Lookup(key string) *Identifier {
	for i := len(c.Scopes) - 1; i >= 0; i-- {
		identifier := c.Scopes[i].Identifiers[key]
		if identifier != nil {
			return identifier
		}
	}
	return nil
}

func (c *ScopeContext) Declare(decl *Declaration) *Identifier {
	return c.CurrentScope().Declare(decl)
}

func (c *ScopeContext) CurrentScope() *Scope {
	if len(c.Scopes) == 0 {
		return nil
	}

	return c.Scopes[len(c.Scopes)-1]
}

func (c *ScopeContext) GlobalScope() *Scope {
	if len(c.Scopes) == 0 {
		return nil
	}

	return c.Scopes[0]
}
