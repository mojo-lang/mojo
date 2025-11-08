package lang

type Packages []*Package

func (p Packages) Len() int           { return len(p) }
func (p Packages) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Packages) Less(i, j int) bool { return p[i].FullName < p[j].FullName }
