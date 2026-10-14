package db

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUniRecordData_Value(t *testing.T) {
	data := &UniRecordData{}
	data.Init()
	data.SetValue("key", core.NewStringValue("value"))

	val, err := data.Value()
	require.NoError(t, err)
	require.Equal(t, "{\"key\":\"value\"}", val)
}

func TestUniRecordData_Scan(t *testing.T) {
	const Val = "{\"key\":\"value\"}"
	data := &UniRecordData{}
	err := data.Scan(Val)
	assert.NoError(t, err)
	assert.Equal(t, "value", data.GetString("key"))
}
