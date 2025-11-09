package lang

import (
	"errors"
	"reflect"
)

const (
	ReturnStmtName   = "ReturnStmt"
	BreakStmtName    = "BreakStmt"
	ContinueStmtName = "ContinueStmt"
	MatchStmtName    = "MatchStmt"
	IfStmtName       = "IfStmt"
	ForStmtName      = "ForStmt"
	WhileStmtName    = "WhileStmt"
	RepeatStmtName   = "RepeatStmt"
	BlockStmtName    = "BlockStmt"
)

func NewExpressionStatement(expression *Expression) *Statement {
	return &Statement{
		Statement: &Statement_Expression{Expression: expression},
	}
}

func NewDeclarationStatement(decl *Declaration) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: decl,
		},
	}
}

func NewPackageDeclStatement(decl *PackageDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_PackageDecl{
					PackageDecl: decl,
				},
			},
		},
	}
}

func NewImportDeclStatement(decl *ImportDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_ImportDecl{
					ImportDecl: decl,
				},
			},
		},
	}
}

func NewEnumDeclStatement(decl *EnumDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_EnumDecl{
					EnumDecl: decl,
				},
			},
		},
	}
}

func NewStructDeclStatement(decl *StructDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_StructDecl{
					StructDecl: decl,
				},
			},
		},
	}
}

func NewInterfaceDeclStatement(decl *InterfaceDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_InterfaceDecl{
					InterfaceDecl: decl,
				},
			},
		},
	}
}

func NewTypeAliasDeclStatement(decl *TypeAliasDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_TypeAliasDecl{
					TypeAliasDecl: decl,
				},
			},
		},
	}
}

func NewFunctionDeclStatement(decl *FunctionDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_FunctionDecl{
					FunctionDecl: decl,
				},
			},
		},
	}
}

func NewConstructorDeclStatement(decl *ConstructorDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_ConstructorDecl{
					ConstructorDecl: decl,
				},
			},
		},
	}
}

func NewVariableDeclStatement(decl *VariableDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_VariableDecl{
					VariableDecl: decl,
				},
			},
		},
	}
}

func NewConstantDeclStatement(decl *ConstantDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_ConstantDecl{
					ConstantDecl: decl,
				},
			},
		},
	}
}

func NewAttributeDeclStatement(decl *AttributeDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_AttributeDecl{
					AttributeDecl: decl,
				},
			},
		},
	}
}

func NewAttributeAliasDeclStatement(decl *AttributeAliasDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_AttributeAliasDecl{
					AttributeAliasDecl: decl,
				},
			},
		},
	}
}

func (x *Statement) IsUnion() {
}

func (x *Statement) GetStartPosition() *Position {
	return GetUnionPosition(x, StartPositionFieldName)
}

func (x *Statement) GetEndPosition() *Position {
	return GetUnionPosition(x, EndPositionFieldName)
}

func (x *Statement) SetStartPosition(position *Position) {
	SetUnionPosition(x, StartPositionFieldName, position)
}

func (x *Statement) SetEndPosition(position *Position) {
	SetUnionPosition(x, EndPositionFieldName, position)
}

func (x *Statement) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		value := reflect.ValueOf(x.Statement)
		value = reflect.Indirect(value).Field(0)
		if merger, ok := value.Interface().(CommentMerger); ok {
			return merger.MergeComment(comment)
		} else {
			// error
		}
	}

	return false, errors.New("nil Declaration")
}
