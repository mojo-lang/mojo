package parser

import (
	"errors"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type Parser struct {
	SrcPath string
	Term    *Term
}

func (p *Parser) ParseJSON(json []byte) (*lang.SourceFile, error) {
	j := jsoniter.Config{
		IndentionStep:                 0,
		MarshalFloatWith6Digits:       false,
		EscapeHTML:                    true,
		SortMapKeys:                   false,
		UseNumber:                     false,
		DisallowUnknownFields:         false,
		OnlyTaggedField:               false,
		ValidateJsonRawMessage:        false,
		ObjectFieldMustBeSimpleString: false,
		CaseSensitive:                 false,
	}.Froze()

	term := &Term{}
	if err := j.Unmarshal(json, term); err != nil {
		return nil, err
	}
	return p.ParseRootTerm(term)
}

func (p *Parser) ParseRootTerm(term *Term) (*lang.SourceFile, error) {
	if p != nil && term != nil {
		p.Term = term

		if term.Kind != "TranslationUnitDecl" {
			return nil, errors.New("not complete term tree, need TranslationUnitDecl for the root term")
		}

		file := &lang.SourceFile{}
		for _, t := range term.Inner {
			switch t.Kind {
			case "FunctionDecl":
				decl, err := p.ParseFunctionDecl(t)
				if err != nil {
					return nil, err
				}
				if decl != nil && !decl.Implicit {
					file.Statements = append(file.Statements, lang.NewFunctionDeclStatement(decl))
					if decl.Name == "main" {
						file.Attributes = append(file.Attributes, lang.NewBoolAttribute("", "main"))
					}
				}
			case "TypedefDecl":
				decl, err := p.ParseTypedefDecl(t)
				if err != nil {
					return nil, err
				}
				if decl != nil {
					file.Statements = append(file.Statements, lang.NewTypeAliasDeclStatement(decl))
				}
			case "RecordDecl":
				decl, err := p.ParseRecordDecl(t)
				if err != nil {
					return nil, err
				}
				if decl != nil && !decl.Implicit && len(decl.Name) > 0 {
					file.Statements = append(file.Statements, lang.NewStructDeclStatement(decl))
				}
			default:
			}
		}
		return file, nil
	}
	return nil, nil
}

func (p *Parser) ParseFunctionDecl(term *Term) (*lang.FunctionDecl, error) {
	if p == nil || term == nil {
		return nil, errors.New("null term")
	}

	decl := &lang.FunctionDecl{
		Name:      term.Name,
		Signature: &lang.FunctionSignature{},
	}

	switch term.StorageClass {
	case "static":
		decl.Attributes = append(decl.Attributes, lang.NewBoolAttribute("", "c_static"))
	case "extern":
		decl.Attributes = append(decl.Attributes, lang.NewBoolAttribute("", "c_extern"))
	}

	if term.Range != nil && term.Range.Begin != nil && term.Range.End != nil {
		decl.StartPosition = &lang.Position{
			Filename: term.Range.Begin.File,
			Offset:   int64(term.Range.Begin.Offset),
			Line:     int64(term.Range.Begin.Line),
			Column:   int64(term.Range.Begin.Col),
		}
		decl.EndPosition = &lang.Position{
			Filename: term.Range.End.File,
			Offset:   int64(term.Range.End.Offset),
			Line:     int64(term.Range.End.Line),
			Column:   int64(term.Range.End.Col),
		}

		if len(decl.StartPosition.Filename) == 0 && term.Loc != nil {
			decl.StartPosition.Filename = term.Loc.File
			decl.EndPosition.Filename = term.Loc.File
		}
	}

	decl.Implicit = term.IsImplicit

	sig := term.Type.QualType
	pos := strings.Index(sig, "(")
	if pos < 0 {
		sig = term.Type.DesugaredQualType
		pos = strings.Index(sig, "(")
	}
	if pos > 0 {
		decl.Signature.Result = lang.NewFunctionResult(&lang.NominalType{
			Name: strings.TrimSpace(sig[0:pos]),
		})
	}

	var args []*lang.ValueDecl
	for _, t := range term.Inner {
		switch t.Kind {
		case "ParmVarDecl":
			args = append(args, &lang.ValueDecl{
				Name: t.Name,
				Type: &lang.NominalType{Name: t.Type.QualType},
			})
		case "CompoundStmt":
			decl.Body = &lang.BlockStmt{}
		}
	}

	if len(args) > 0 {
		decl.Signature.Parameter = lang.NewFunctionParameters(args...)
	}

	if decl.Body != nil {
		return decl, nil
	}
	return nil, nil
}

func (p *Parser) ParseTypedefDecl(term *Term) (*lang.TypeAliasDecl, error) {
	if p == nil || term == nil {
		return nil, errors.New("null term")
	}

	decl := &lang.TypeAliasDecl{
		Name:     term.Name,
		Type:     &lang.NominalType{},
		Implicit: term.IsImplicit,
	}

	if term.Range != nil && term.Range.Begin != nil && term.Range.End != nil {
		decl.StartPosition = &lang.Position{
			Filename: term.Range.Begin.File,
			Offset:   int64(term.Range.Begin.Offset),
			Line:     int64(term.Range.Begin.Line),
			Column:   int64(term.Range.Begin.Col),
		}
		decl.EndPosition = &lang.Position{
			Filename: term.Range.End.File,
			Offset:   int64(term.Range.End.Offset),
			Line:     int64(term.Range.End.Line),
			Column:   int64(term.Range.End.Col),
		}

		if len(decl.StartPosition.Filename) == 0 && term.Loc != nil {
			decl.StartPosition.Filename = term.Loc.File
			decl.EndPosition.Filename = term.Loc.File
		}
	}

	if term.Type != nil {
		decl.Type.Name = term.Type.QualType
	}

	// for _, t := range term.Inner {
	// }

	return decl, nil
}

func (p *Parser) ParseRecordDecl(term *Term) (*lang.StructDecl, error) {
	if p == nil || term == nil {
		return nil, errors.New("null term")
	}

	decl := &lang.StructDecl{
		Name:     term.Name,
		Type:     &lang.StructType{},
		Implicit: term.IsImplicit,
	}

	if term.Range != nil && term.Range.Begin != nil && term.Range.End != nil {
		decl.StartPosition = &lang.Position{
			Filename: term.Range.Begin.File,
			Offset:   int64(term.Range.Begin.Offset),
			Line:     int64(term.Range.Begin.Line),
			Column:   int64(term.Range.Begin.Col),
		}
		decl.EndPosition = &lang.Position{
			Filename: term.Range.End.File,
			Offset:   int64(term.Range.End.Offset),
			Line:     int64(term.Range.End.Line),
			Column:   int64(term.Range.End.Col),
		}

		if len(decl.StartPosition.Filename) == 0 && term.Loc != nil {
			decl.StartPosition.Filename = term.Loc.File
			decl.EndPosition.Filename = term.Loc.File
		}
	}

	for _, t := range term.Inner {
		if t.Kind != "FieldDecl" {
			continue
		}

		field := &lang.ValueDecl{
			Name: t.Name,
			Type: &lang.NominalType{Name: t.Type.QualType},
		}

		decl.Type.Fields = append(decl.Type.Fields, field)
	}

	return decl, nil
}

func (p *Parser) ParseEnumDecl(term *Term) (*lang.EnumDecl, error) {
	return nil, nil
}

func (p *Parser) ParseUnionDecl(term *Term) (*lang.NominalType, error) {
	return nil, nil
}
