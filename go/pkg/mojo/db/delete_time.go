package db

import (
	"time"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

const DeleteTimeTypeFullName = "mojo.db.DeleteTime"

func (x *DeleteTime) ToTime() time.Time {
	return x.Val.ToTime()
}

func (x *DeleteTime) FromTime(t time.Time) {
	x.Val = &core.Timestamp{}
	x.Val.FromTime(t)
}
