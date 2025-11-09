package printer

import (
	"sync"
)

type OnceLineBreaker struct {
	once sync.Once
}

func (b *OnceLineBreaker) Break(p *Printer) *Printer {
	if p != nil {
		b.once.Do(func() {
			if p.P.Cursor.Line > 0 {
				p.BreakLine()
			}
		})
	}
	return p
}
