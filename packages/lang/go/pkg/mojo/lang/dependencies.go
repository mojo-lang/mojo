package lang

func MergeDependencies(dependencies ...[]*Identifier) []*Identifier {
	index := make(map[string]*Identifier)

	for _, deps := range dependencies {
		for _, dep := range deps {
			index[dep.FullName] = dep
		}
	}

	identifiers := make([]*Identifier, 0, len(index))
	for _, value := range index {
		identifiers = append(identifiers, value)
	}

	return identifiers
}
