package openapi

type Schemas []*Schema

func (s Schemas) Len() int {
	return len(s)
}
func (s Schemas) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Schemas) Less(i, j int) bool {
	return s[i].Title < s[j].Title
}
