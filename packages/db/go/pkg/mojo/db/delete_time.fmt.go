package db

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

func (x *DeleteTime) Format() string {
	return x.Val.Format()
}

func ParseDeleteTime(value string) (*DeleteTime, error) {
	deleteTime := &DeleteTime{}
	if err := deleteTime.Parse(value); err != nil {
		return nil, err
	}
	return deleteTime, nil
}

func (x *DeleteTime) Parse(value string) error {
	if x != nil {
		x.Val = &core.Timestamp{}
		return x.Val.Parse(value)
	}
	return nil
}
