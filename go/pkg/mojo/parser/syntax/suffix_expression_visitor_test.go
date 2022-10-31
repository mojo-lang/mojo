package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffixExpressionVisitor_VisitFunctionCallSuffix(t *testing.T) {
	const funcCall = "a(b, c, d)"
	expr := parseExpression(t, funcCall)

	funcCallExpr := expr.GetFunctionCallExpr()
	assert.NotNil(t, funcCallExpr)

	assert.Equal(t, "a", funcCallExpr.GetExpression().GetIdentifierExpr().Identifier.Name)
	assert.Equal(t, 3, len(funcCallExpr.Arguments))
	assert.Equal(t, "b", funcCallExpr.Arguments[0].Value.GetIdentifierExpr().Identifier.Name)
}

func TestSuffixExpressionVisitor_VisitSubscriptSuffix(t *testing.T) {
	const subscript = "a[b]"
	expr := parseExpression(t, subscript)

	subscriptExpr := expr.GetSubscriptExpr()
	assert.NotNil(t, subscriptExpr)
	assert.Equal(t, "a", subscriptExpr.GetExpression().GetIdentifierExpr().Identifier.Name)
	assert.Equal(t, 1, len(subscriptExpr.Arguments))
	assert.Equal(t, "b", subscriptExpr.Arguments[0].Value.GetIdentifierExpr().Identifier.Name)
}

func TestSuffixExpressionVisitor_VisitExplicitMemberSuffix(t *testing.T) {
	const explicitMember = "a.b.c"
	expr := parseExpression(t, explicitMember)

	explicitMemberExpr := expr.GetExplicitMemberExpr()
	assert.NotNil(t, explicitMemberExpr)

	callee := explicitMemberExpr.GetExpression().GetExplicitMemberExpr()

	assert.Equal(t, "a", callee.GetExpression().GetIdentifierExpr().Identifier.Name)
	assert.Equal(t, "b", callee.Member)
	assert.Equal(t, "c", explicitMemberExpr.Member)
}
