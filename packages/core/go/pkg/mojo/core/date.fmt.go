package core

import (
	"fmt"
	"github.com/araddon/dateparse"
	"time"
)

func (x *Date) Parse(str string) error {
	if x != nil {
		t, err := dateparse.ParseIn(str, time.UTC)
		if err != nil {
			return err
		}
		y, m, d := t.Date()
		x.Year = int32(y)
		x.Month = int32(m)
		x.Day = int32(d)
	}
	return nil
}

func (x *Date) Format() string {
	if x != nil {
		return fmt.Sprintf("%d-%02d-%02d", x.Year, x.Month, x.Day)
	}
	return ""
}
