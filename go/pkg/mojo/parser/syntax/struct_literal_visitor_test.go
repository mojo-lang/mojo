package syntax

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestStructLiteralVisitor_VisitStructLiteral(t *testing.T) {
    const structLiteral = `Image{type: "png", link: "https://company.com/path/to/image.png"}`
    expr := parseExpression(t, structLiteral)

    structLiteralExpr := expr.GetStructLiteralExpr()
    assert.NotNil(t, structLiteralExpr)

    assert.Equal(t, "Image", structLiteralExpr.Callee.GetIdentifierExpr().Identifier.Name)
    assert.Equal(t, 2, len(structLiteralExpr.Value.Fields))
    assert.Equal(t, "type", structLiteralExpr.Value.Fields[0].Name)
}

func TestStructLiteralVisitor_VisitStructLiteral_Enclosing(t *testing.T) {
    const structLiteral = `Wrap.Image{type: "png", link: "https://company.com/path/to/image.png"}`
    expr := parseExpression(t, structLiteral)

    structLiteralExpr := expr.GetStructLiteralExpr()
    assert.NotNil(t, structLiteralExpr)

    assert.Equal(t, "Image", structLiteralExpr.Callee.GetIdentifierExpr().Identifier.Name)
    assert.Equal(t, 1, len(structLiteralExpr.Callee.GetIdentifierExpr().Identifier.EnclosingTypeNames))
    assert.Equal(t, "Wrap", structLiteralExpr.Callee.GetIdentifierExpr().Identifier.EnclosingTypeNames[0])
    assert.Equal(t, 2, len(structLiteralExpr.Value.Fields))
    assert.Equal(t, "type", structLiteralExpr.Value.Fields[0].Name)
}

func TestStructLiteralVisitor_VisitStructLiteral_PackageName(t *testing.T) {
    const structLiteral = `mojo.core.inner.Image{type: "png", link: "https://company.com/path/to/image.png"}`
    expr := parseExpression(t, structLiteral)

    structLiteralExpr := expr.GetStructLiteralExpr()
    assert.NotNil(t, structLiteralExpr)

    assert.Equal(t, "Image", structLiteralExpr.Callee.GetIdentifierExpr().Identifier.Name)
    assert.Equal(t, "mojo.core.inner", structLiteralExpr.Callee.GetIdentifierExpr().Identifier.PackageName)
    assert.Equal(t, 2, len(structLiteralExpr.Value.Fields))
    assert.Equal(t, "type", structLiteralExpr.Value.Fields[0].Name)
}
