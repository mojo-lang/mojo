package core

func NewStringMap() *StringMap {
	return &StringMap{Vals: make(map[string]string)}
}
