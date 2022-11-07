package printer

type Columns []int

func (c Columns) Column(index int) int {
	if index >= len(c) || index < 0 {
		return -1
	}
	return c[index]
}

func (c Columns) PrintTo(index int, printer *Printer) *Printer {
	column := c.Column(index) + printer.GetIndent()
	if column > 0 {
		printer.PrintTo(Cursor{
			Line:   printer.Cursor.Line,
			Column: column,
		})
	}
	return printer
}
