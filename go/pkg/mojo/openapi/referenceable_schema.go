package openapi

import (
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/markdown"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/document"
)

const SchemaReferenceRoot = "/components/schemas/"

func NewReferenceableSchema(schema *Schema) *ReferenceableSchema {
	s := &ReferenceableSchema{}
	s.SetSchema(schema)
	return s
}

func NewReferencedSchema(reference *Reference) *ReferenceableSchema {
	s := &ReferenceableSchema{}
	s.SetReference(reference)
	return s
}

func (x *ReferenceableSchema) SetSchema(schema *Schema) *ReferenceableSchema {
	if x != nil && schema != nil {
		x.ReferenceableSchema = &ReferenceableSchema_Schema{
			Schema: schema,
		}
	}
	return x
}

func (x *ReferenceableSchema) SetReference(reference *Reference) *ReferenceableSchema {
	if x != nil && reference != nil {
		x.ReferenceableSchema = &ReferenceableSchema_Reference{
			Reference: reference,
		}
	}
	return x
}

func (x *ReferenceableSchema) SetReferenceUrl(reference string) error {
	if x != nil {
		url, err := core.NewUrl(reference)
		if err != nil {
			return err
		}
		x.ReferenceableSchema = &ReferenceableSchema_Reference{
			Reference: &Reference{
				Ref: url,
			},
		}
	}
	return nil
}

func (x *ReferenceableSchema) GetReferenceUrl() *core.Url {
	return x.GetReference().GetRef()
}

func (x *ReferenceableSchema) GetSchemaName() string {
	reference := x.GetReference()
	if reference != nil {
		return reference.GetSchemaName()
	}

	schema := x.GetSchema()
	if schema != nil {
		return schema.Title
	}

	return ""
}

func (x *ReferenceableSchema) GetSchemaOf(index map[string]*Schema) *Schema {
	if x != nil {
		reference := x.GetReference()
		if reference != nil && index != nil {
			fragment := reference.GetRef().GetFragment()
			key := strings.TrimPrefix(fragment, SchemaReferenceRoot)
			if s, ok := index[key]; ok {
				return s
			}
			return nil
		}
		return x.GetSchema()
	}
	return nil
}

func (x *ReferenceableSchema) GetTypeName(index map[string]*Schema) string {
	s := x.GetSchemaOf(index)
	if s != nil {
		return s.GetTypeName(index)
	}
	return ""
}

func (x *ReferenceableSchema) GetFormat(index map[string]*Schema) string {
	s := x.GetSchemaOf(index)
	if s != nil {
		return s.Format
	}
	return ""
}

func (x *ReferenceableSchema) GetSummary(index map[string]*Schema) string {
	if summary := x.GetReference().GetSummary(); len(summary) > 0 {
		return summary
	}

	s := x.GetSchemaOf(index)
	if s != nil {
		doc := s.GetDescription().GetDocument()
		if doc != nil && len(doc.Blocks) > 0 {
			str, _ := markdown.RenderToString(&document.Document{Blocks: doc.Blocks[0:1]})
			return str
		}
	}

	return ""
}

func (x *ReferenceableSchema) GetTitle(index map[string]*Schema) string {
	s := x.GetSchemaOf(index)
	if s != nil {
		return s.Title
	}
	return ""
}

func (x *ReferenceableSchema) GetDescription(index map[string]*Schema) *document.Document {
	if doc := x.GetReference().GetDescription(); doc != nil {
		return doc.GetDocument()
	}

	s := x.GetSchemaOf(index)
	if s != nil {
		return s.GetDescription().GetDocument()
	}

	return nil
}

func (x *ReferenceableSchema) SetDescription(doc *CachedDocument) {
	if x != nil {
		reference := x.GetReference()
		if reference != nil {
			reference.Description = doc
		}

		schema := x.GetSchema()
		if schema != nil {
			schema.Description = doc
		}
	}
}

func (m *Reference) GetSchemaName() string {
	fragment := m.GetRef().GetFragment()
	return strings.TrimPrefix(fragment, SchemaReferenceRoot)
}
