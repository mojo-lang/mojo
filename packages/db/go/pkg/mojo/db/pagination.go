package db

type Pagination struct {
	Skip int32

	PageSize int32

	PageToken string
}

func (p Pagination) Offset() int32 {
	return 0
}

// Limit
// Maximum number of records to return. Default limit_clause.mojo is 25.
// Maximum limit_clause.mojo is 1000. Anything higher will return an error.
func (p Pagination) Limit() int32 {
	return 0
}
