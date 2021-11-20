package render

import "io"

type HandlerTemplate struct {
	Interface string
	Methods   string
	Extension string
}

type Config struct {
	Repository    string // the repository for the generated gokit service
	MixedInAPI    bool
	ApiRepository string
	Version       string
	VersionDate   string
	Output        string
	PreviousFiles map[string]io.Reader

	ExtensionData map[string]interface{}

	HandlerTemplate *HandlerTemplate
}
