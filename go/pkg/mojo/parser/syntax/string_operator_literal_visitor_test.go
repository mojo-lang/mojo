package syntax

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestStringOperatorLiteralVisitor_VisitStringOperatorLiteral(t *testing.T) {
    const StringOperatorExpression = `r"^ab$"`
    expr := parseExpression(t, StringOperatorExpression)

    stringExpr := expr.GetStringPrefixLiteralExpr()
    assert.NotNil(t, stringExpr)
    assert.Equal(t, "r", stringExpr.GetOperator().Symbol)
    assert.Equal(t, "^ab$", stringExpr.Argument.GetStringLiteralExpr().Value)
}

func TestStringOperatorLiteralVisitor_VisitStringOperatorLiteral2(t *testing.T) {
    const StringOperatorExpression = `r"^ab$"i`
    expr := parseExpression(t, StringOperatorExpression)

    suffixExpr := expr.GetStringSuffixLiteralExpr()
    if assert.NotNil(t, suffixExpr) {
        assert.Equal(t, "i", suffixExpr.GetOperator().Symbol)

        prefixExpr := suffixExpr.Argument.GetStringPrefixLiteralExpr()
        if assert.NotNil(t, prefixExpr) {
            assert.Equal(t, "r", prefixExpr.GetOperator().Symbol)
            assert.Equal(t, "^ab$", prefixExpr.Argument.GetStringLiteralExpr().Value)
        }
    }
}

func TestStringOperatorLiteralVisitor_VisitStringOperatorLiteral3(t *testing.T) {
    const StringOperatorExpression = `r "^ab$"i`
    parser := &Parser{}
    _, err := parser.ParseString(StringOperatorExpression)
    assert.Error(t, err)
}

func TestStringOperatorLiteralVisitor_VisitStringOperatorLiteral4(t *testing.T) {
    const StringOperatorExpression = `r"^ab$" i`
    parser := &Parser{}
    _, err := parser.ParseString(StringOperatorExpression)
    assert.Error(t, err)
}
