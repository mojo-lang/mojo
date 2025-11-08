package db

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniRecordData_Value(t *testing.T) {
	data := &UniRecordData{}
	data.SetValue("key", core.NewStringValue("value"))

	val, err := data.Value()
	assert.NoError(t, err)
	assert.Equal(t, "{\"key\":\"value\"}", val.(string))
}

func TestUniRecordData_Scan(t *testing.T) {
	const Val = "{\"key\":\"value\"}"
	data := &UniRecordData{}
	err := data.Scan(Val)
	assert.NoError(t, err)
	assert.Equal(t, "value", data.GetString("key"))
}
