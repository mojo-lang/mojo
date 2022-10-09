package injection

import (
	"github.com/stretchr/testify/assert"
	"go/token"
	"testing"
)

func TestSort(t *testing.T) {
	areas := []Area{
		&TextArea{
			Start:   100,
			End:     105,
			Content: "test-data-2",
		},
		&TextArea{
			Start:   90,
			End:     95,
			Content: "test-data-1",
		},
	}

	Sort(areas)
	assert.Equal(t, token.Pos(90), areas[0].GetStart())
}
