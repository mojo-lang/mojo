package template

import (
	_ "embed"
)

//go:embed api_header.md.tmpl
var apiDocumentHeader string
