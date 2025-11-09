package data

type HttpParameters []*HTTPParameter

func (p HttpParameters) Len() int {
	return len(p)
}
func (p HttpParameters) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p HttpParameters) Less(i, j int) bool {
	return p[i].FullName < p[j].FullName
}
