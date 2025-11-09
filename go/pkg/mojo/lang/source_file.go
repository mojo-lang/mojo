package lang

import "strings"

func (x *SourceFile) IsGenericInstantiated() bool {
	return strings.Contains(x.GetName(), "<")
}

func (x *SourceFile) SetScope(scope *Scope) {
	if x != nil {
		x.Scope = scope
	}
}

func (x *SourceFile) GetTypeDeclarations() []*TypeDeclaration {
	var decls []*TypeDeclaration
	for _, statement := range x.GetStatements() {
		if decl := statement.GetDeclaration(); decl != nil {
			switch decl.Declaration.(type) {
			case *Declaration_StructDecl, *Declaration_EnumDecl, *Declaration_InterfaceDecl, *Declaration_TypeAliasDecl:
				decls = append(decls, NewTypeDeclarationFromDeclaration(decl))
			}
		}
	}
	return decls
}

func (x *SourceFile) GetStructDecls() []*StructDecl {
	var decls []*StructDecl
	for _, statement := range x.GetStatements() {
		if decl := statement.GetDeclaration().GetStructDecl(); decl != nil {
			decls = append(decls, decl)
		}
	}
	return decls
}

func (x *SourceFile) GetTypeAliasDecls() []*TypeAliasDecl {
	var decls []*TypeAliasDecl
	for _, statement := range x.GetStatements() {
		if decl := statement.GetDeclaration().GetTypeAliasDecl(); decl != nil {
			decls = append(decls, decl)
		}
	}
	return decls
}
