package printer

// A Config node controls the output.
type Config struct {
	IndentWidth          int // default: 4
	Indent               int // default: 0
	MaxCharactersPerLine int // default: 125 https://hilton.org.uk/blog/source-code-line-length
}
