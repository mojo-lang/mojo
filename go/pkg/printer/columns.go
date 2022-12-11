package printer

type Columns []int

func (c Columns) Column(index int) int {
	if len(c) > 0 {
		if index < 0 {
			return c[0]
		}
		if index >= len(c) {
			return c[len(c)-1]
		}
		return c[index]
	}
	return 0
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
