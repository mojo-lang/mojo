package printer

const maxSpace = 128

var spaceIndents []byte

func init() {
	spaceIndents = make([]byte, 0, maxSpace)
	for i := 0; i < maxSpace; i++ {
		spaceIndents = append(spaceIndents, ' ')
	}
}

type lineIndent struct {
	Width int
	Count int
}

func (l *lineIndent) Indent() *lineIndent {
	if l != nil {
		l.Count++
	}
	return l
}

func (l *lineIndent) Outdent() *lineIndent {
	if l != nil && l.Count > 0 {
		l.Count--
	}
	return l
}

func (l *lineIndent) IndentSpace() []byte {
	if l != nil {
		return spaceIndents[0 : l.Count*l.Width]
	}
	return nil
}
