package transformer

import "github.com/gertd/go-pluralize"

// type name pluralize
// generic type name
// union label

var pluralizing *pluralize.Client

var irregularPlurals = map[string]string{}

func init() {
	pluralizing = pluralize.NewClient()
}

func RegisterIrregularPluralize(singular string, plural string) {
	irregularPlurals[singular] = plural
}

func UnregisterIrregularPluralize(singular string, plural string) {
}

// Plural -- Pluralize a word.
func Plural(word string) string {
	return pluralizing.Plural(word)
}

// IsPlural -- Check if a word is plural.
func IsPlural(word string) bool {
	return pluralizing.IsPlural(word)
}

// Singular -- Singularize a word.
func Singular(word string) string {
	return pluralizing.Singular(word)
}

// IsSingular -- Check if a word is singular.
func IsSingular(word string) bool {
	return pluralizing.IsSingular(word)
}
