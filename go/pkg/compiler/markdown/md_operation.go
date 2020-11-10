package markdown

type MdOperation struct {
	Summary string
	Description string
	Deprecated bool

	PathParameters []*MdPathParameter
	QueryParameters []*MdQueryParameter
}

type MdQueryParameter struct {
}

type MdPathParameter struct {
}

type MdRequestBody struct {
}

type MdResponse struct {
}
