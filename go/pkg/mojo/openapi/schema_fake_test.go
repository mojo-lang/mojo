package openapi

import (
	"net"
	"net/mail"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFakeNumericBounds(t *testing.T) {
	for _, bits := range []int{8, 16, 32, 64} {
		schema := &Schema{Minimum: core.NewIntValue(10), Maximum: core.NewIntValue(10), MultipleOf: 5}
		assert.Equal(t, 10, FakeInt(schema, bits))
		assert.Equal(t, uint(10), FakeUInt(schema, bits))
		assert.Zero(t, FakeInt(&Schema{}, bits))
		assert.Zero(t, FakeUInt(&Schema{}, bits))
	}
	schema := &Schema{Minimum: core.NewFloat64Value(1.25), Maximum: core.NewFloat64Value(1.25)}
	assert.Equal(t, float32(1.25), FakeFloat32(schema))
	assert.Equal(t, 1.25, FakeFloat64(schema))
	assert.Zero(t, FakeFloat32(&Schema{}))
	assert.Zero(t, FakeFloat64(&Schema{}))
}

func TestFakeStringFormats(t *testing.T) {
	for _, tt := range []struct{ format, layout string }{
		{"date-time", time.RFC3339}, {"Date", time.DateOnly}, {"Time", time.TimeOnly},
	} {
		_, err := time.Parse(tt.layout, FakeString(&Schema{Format: tt.format}))
		assert.NoError(t, err, tt.format)
	}
	for _, format := range []string{"Timestamp", "Duration", "Id", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64"} {
		_, err := strconv.ParseInt(FakeString(&Schema{Format: format}), 10, 64)
		assert.NoError(t, err, format)
	}
	for _, format := range []string{"float32", "float64", "double", "float"} {
		_, err := strconv.ParseFloat(FakeString(&Schema{Format: format}), 64)
		assert.NoError(t, err, format)
	}
	for _, format := range []string{"ipv4", "ipv6", "IPAddress"} {
		assert.NotNil(t, net.ParseIP(FakeString(&Schema{Format: format})), format)
	}
	_, err := net.ParseMAC(FakeString(&Schema{Format: "MacAddress"}))
	assert.NoError(t, err)
	_, err = mail.ParseAddress(FakeString(&Schema{Format: "Email"}))
	assert.NoError(t, err)
	address, err := url.Parse(FakeString(&Schema{Format: "Url"}))
	require.NoError(t, err)
	assert.NotEmpty(t, address.Host)
	assert.Regexp(t, `^[0-9a-f-]{36}$`, FakeString(&Schema{Format: "uuid"}))
	assert.Regexp(t, `^1[0-9]{10}$`, FakeString(&Schema{Format: "PhoneNumber"}))
	assert.Regexp(t, `^[0-9]+$`, FakeString(&Schema{Format: "Digital"}))
	assert.Regexp(t, `^[0-9A-F]{8}$`, FakeString(&Schema{Format: "Hex", MinLength: 8, MaxLength: 8}))
	for _, format := range []string{"Name", "Password", "Path", "unknown"} {
		assert.NotEmpty(t, FakeString(&Schema{Format: format}))
	}
	assert.Equal(t, "chosen", FakeString(&Schema{Enum: []*core.Value{core.NewStringValue("chosen")}}))
}
