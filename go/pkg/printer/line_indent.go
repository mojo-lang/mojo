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

func (l *lineIndent) Indent() {
	l.Count++
}

func (l *lineIndent) Outdent() {
	if l.Count > 0 {
		l.Count--
	}
}

func (l *lineIndent) IndentSpace() []byte {
	return spaceIndents[0 : l.Count*l.Width]
}
