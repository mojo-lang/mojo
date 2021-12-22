package plugin

type SkipError struct {
}

func (SkipError) Error() string {
	return "Skip"
}
