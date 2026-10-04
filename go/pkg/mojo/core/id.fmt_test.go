package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdParseFormat(t *testing.T) {
	for _, tt := range []struct{ input, kind string }{
		{"0", "integer"}, {"42", "integer"}, {"18446744073709551615", "integer"},
		{"invoice-42", "string"}, {"", "string"}, {"-1", "string"},
		{"18446744073709551616", "string"},
		{"00112233-4455-4677-8899-aabbccddeeff", "uuid"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			id, err := ParseId(tt.input)
			require.NoError(t, err)
			require.NotNil(t, id)
			switch tt.kind {
			case "integer":
				assert.IsType(t, &Id_Uint64Val{}, id.Id)
			case "string":
				assert.IsType(t, &Id_StringVal{}, id.Id)
			case "uuid":
				assert.IsType(t, &Id_UuidVal{}, id.Id)
			}
			assert.Equal(t, tt.input, id.Format())
		})
	}
	assert.Empty(t, (*Id)(nil).Format())
	assert.Empty(t, (&Id{}).Format())
}
