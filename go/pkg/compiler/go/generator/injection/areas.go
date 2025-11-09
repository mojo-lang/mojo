package injection

import "sort"

type Areas []Area

func (a Areas) Len() int {
	return len(a)
}

func (a Areas) Less(i int, j int) bool {
	return a[i].GetStart() < a[j].GetStart()
}

func (a Areas) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func Sort(areas []Area) {
	a := Areas(areas)
	sort.Sort(a)
}
