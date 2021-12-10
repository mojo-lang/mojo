package context

import (
	"context"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

const ScopeKey = "@scope"

func WithScope(ctx context.Context) context.Context {
	return WithValues(ctx, ScopeKey, lang.NewScope())
}

func WithScopeType(ctx context.Context, value interface{}) context.Context {
	if scope := lang.GetScope(value); scope != nil {
		return WithValues(WithType(ctx, value), ScopeKey, scope)
	}
	return WithScope(WithType(ctx, value))
}

func Scope(ctx context.Context) *lang.Scope {
	if scope, ok := ctx.Value(ScopeKey).(*lang.Scope); ok {
		return scope
	}
	return nil
}

func GlobalScope(ctx context.Context) *lang.Scope {
	scopes := ScopeValues(ctx)
	if len(scopes) == 0 {
		return nil
	}
	return scopes[len(scopes)-1].(*lang.Scope)
}

func LookupIdentifier(ctx context.Context, key string) *lang.Identifier {
	scopes := ScopeValues(ctx)
	for _, scope := range scopes {
		identifier := scope.(*lang.Scope).Identifiers[key]
		if identifier != nil {
			return identifier
		}
	}
	return nil
}

func ScopeValue(ctx context.Context) interface{} {
	return ctx.Value(ScopeKey)
}

func ScopeValues(ctx context.Context) []interface{} {
	return Values(ctx, ScopeKey)
}

func PreviousScopeValue(ctx context.Context, index int) interface{} {
	return PreviousValue(ctx, ScopeKey, index)
}
