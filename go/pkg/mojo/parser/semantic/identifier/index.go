package identifier

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

// global identifier index
type Index struct {
	Packages    map[string][]*lang.Package
	Identifiers map[string][]*lang.Identifier
}

func (i *Index) Find(key string) (*lang.Identifier, error) {
	return nil, nil
}

func (i *Index) Add(key string, identifier *lang.Identifier) {
}

var identifierIndex *Index

func getIdentifierIndex() *Index {
	if identifierIndex == nil {
		identifierIndex = &Index{
			Identifiers: make(map[string][]*lang.Identifier),
		}
	}

	return identifierIndex
}

func AddIdentifier(key string, identifier *lang.Identifier) {
	getIdentifierIndex().Add(key, identifier)
}

func FindIdentifier(ctx context.Context, key string) (*lang.Identifier, error) {
	return getIdentifierIndex().Find(key)
}
