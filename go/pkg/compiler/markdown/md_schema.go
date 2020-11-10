package markdown

type MdSchema struct {
	Id string
	Name string
	Description string

	Fields []*MdSchemaField
}

// MdSchemaField name is field
type MdSchemaField struct {
	Name string
	Type string

	Required bool
}
