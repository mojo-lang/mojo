package syntax

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAttributeDeclarationVisitor_VisitAttributeAliasDeclaration(t *testing.T) {
    const attributeAliasDecl = `@foo('bar') attribute mailbox = pkg1.pkg2.mail<String>`

    decl := parseAttributeAliasDecl(t, attributeAliasDecl)

    assert.Equal(t, "mailbox", decl.Name)
    assert.Equal(t, "pkg1.pkg2", decl.Attribute.PackageName)
    assert.Equal(t, "mail", decl.Attribute.Name)
    assert.Equal(t, 1, len(decl.Attribute.GenericArguments))
    assert.Equal(t, "String", decl.Attribute.GenericArguments[0].Name)
    assert.Equal(t, 1, len(decl.Attributes))
    assert.Equal(t, "foo", decl.Attributes[0].Name)
}

func getAttributeAliasDecl(file *lang.SourceFile) *lang.AttributeAliasDecl {
    if len(file.Statements) > 0 {
        statement := file.Statements[0]
        if decl := statement.GetDeclaration(); decl != nil {
            return decl.GetAttributeAliasDecl()
        }
    }
    return nil
}

func parseAttributeAliasDecl(t *testing.T, str string) *lang.AttributeAliasDecl {
    parser := &Parser{}
    file, err := parser.ParseString(str)
    assert.NoError(t, err)

    decl := getAttributeAliasDecl(file)
    assert.NotNil(t, decl)
    return decl
}
