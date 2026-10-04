package fake

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentify(t *testing.T) {
	id := Identify()

	assert.NotEmpty(t, id)
	assert.Len(t, id, 18)
}

func TestIdentifyFormatAcrossSeeds(t *testing.T) {
	gofakeit.Seed(1234)
	t.Cleanup(func() { gofakeit.Seed(0) })
	for i := 0; i < 100; i++ {
		var id string
		require.NotPanics(t, func() { id = Identify() })
		require.Regexp(t, `^[0-9]{17}[0-9X]$`, id)
		_, err := time.Parse("20060102", id[6:14])
		assert.NoError(t, err, "date in %s", id)
	}
}
