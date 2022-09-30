package format

// FirstLineNoneBreaker first line of the scope do not break, following will break
type FirstLineNoneBreaker struct {
    Cursor
    Printer *Printer
}

func NewFirstLineNoneBreaker(printer *Printer) *FirstLineNoneBreaker {
    return &FirstLineNoneBreaker{Cursor: printer.Cursor, Printer: printer}
}

func (f *FirstLineNoneBreaker) Break() *Printer {
    if f.Line > f.Cursor.Line {
        f.Printer.BreakLine()
    }
    return f.Printer
}