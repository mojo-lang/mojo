package format

type VerticalLine int
type VerticalLines []int

func (l VerticalLines) GetColumn(index int) int {
    if index >= len(l) || index < 0 {
        return -1
    }
    return l[index]
}

func (l VerticalLines) PrintTo(index int, printer *Printer) *Printer {
    column := l.GetColumn(index) + printer.GetIndent()
    if column > 0 {
        printer.PrintTo(Cursor{
            Line:   printer.Cursor.Line,
            Column: column,
        })
    }
    return printer
}
