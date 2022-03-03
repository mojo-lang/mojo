package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getStructDecl(file *lang.SourceFile) *lang.StructDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetStructDecl()
		}
	}
	return nil
}

func parseStructDecl(t *testing.T, decl string) *lang.StructDecl {
	parser := &Parser{}
	file, err := parser.ParseString(decl)
	assert.NoError(t, err)

	structDecl := getStructDecl(file)
	assert.NotNil(t, structDecl)

	return structDecl
}

func TestStructDeclarationVisitor_VisitStructDeclaration_FreeFloatingDocument(t *testing.T) {
	const typeDecl = `
@foo('bar')
type Mailbox {
	/// free floating document
}
`
	decl := parseStructDecl(t, typeDecl)
	assert.Equal(t, "Mailbox", decl.Name)
	assert.NotEmpty(t, decl.Type.EndPosition.LeadingComments)
}

func TestStructDeclarationVisitor_VisitStructDeclaration_Enclosing(t *testing.T) {
	const typeDecl = `
@foo('bar')
type Mailbox {
	type Mail {
		type Box {
 			type Inner {
			}
		}
		type Box2 {
		}
	}
}
`
	decl := parseStructDecl(t, typeDecl)

	assert.Equal(t, "Mailbox", decl.Name)
	assert.Equal(t, "Mailbox", decl.StructDecls[0].EnclosingType.Name)

	assert.Equal(t, "Mail", decl.StructDecls[0].StructDecls[0].EnclosingType.Name)
	assert.Equal(t, "Mail", decl.StructDecls[0].StructDecls[1].EnclosingType.Name)

	assert.Equal(t, "Inner", decl.StructDecls[0].StructDecls[0].StructDecls[0].Name)
	assert.Equal(t, "Box", decl.StructDecls[0].StructDecls[0].StructDecls[0].EnclosingType.Name)
}

func TestStructDeclarationVisitor_VisitStructDeclaration_Enclosing2(t *testing.T) {
	const typeDecl = `
/// A Package represents a set of source files
/// collectively building a Mojo package.
type Package {
    ///
    ///
    type Requirement {
        type Version {
            enum Type {
                caret @0
                tilde @1
                wildcard @2
                comparison @3
            }
            type: Type @1
            range: VersionRange @2
        }

        /// ^1.2.3
        /// ~1.2.3
        /// >= 1.2, < 1.5
        version: Version @1
    }

    type Author {
        /// the people that are considered the "author" of the package.
        author: String @1
    }

    /// package name
    name : String @1

    ///
    full_name: String @2
}
`

	decl := parseStructDecl(t, typeDecl)

	assert.Equal(t, "Package", decl.Name)
	assert.Equal(t, "Requirement", decl.StructDecls[0].StructDecls[0].EnclosingType.Name)
	assert.Equal(t, "Package", decl.StructDecls[0].StructDecls[0].EnumDecls[0].EnclosingType.EnclosingType.EnclosingType.Name)
}
