package lang

func MergeUnresolvedIdentifiers(identifiers ...[]*Identifier) []*Identifier {
	index := make(map[string]*Identifier)

	for _, deps := range identifiers {
		for _, dep := range deps {
			index[dep.FullName] = dep
		}
	}

	ids := make([]*Identifier, 0, len(index))
	for _, value := range index {
		ids = append(ids, value)
	}

	return ids
}
